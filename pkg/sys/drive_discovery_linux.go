//go:build linux

// This file is part of MinIO DirectPV
// Copyright (c) 2021, 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sys

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/minio/directpv/pkg/fs"
	fserrors "github.com/minio/directpv/pkg/fs/errors"
	"github.com/minio/directpv/pkg/mount"
	"golang.org/x/sys/unix"
)

const (
	defaultBlockSize = 512
)

func getDeviceMajorMinor(device string) (major, minor uint32, err error) {
	stat := syscall.Stat_t{}
	if err = syscall.Stat(device, &stat); err == nil {
		major, minor = uint32(unix.Major(stat.Rdev)), uint32(unix.Minor(stat.Rdev))
	}
	return
}

func readFirstLine(filename string, errorIfNotExist bool) (string, error) {
	getError := func(err error) error {
		if errorIfNotExist {
			return err
		}
		switch {
		case errors.Is(err, os.ErrNotExist), errors.Is(err, os.ErrInvalid):
			return nil
		case strings.Contains(strings.ToLower(err.Error()), "no such device"):
			return nil
		case strings.Contains(strings.ToLower(err.Error()), "invalid argument"):
			return nil
		}
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", getError(err)
	}
	defer file.Close()
	s, err := bufio.NewReader(file).ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", getError(err)
	}
	return strings.TrimSpace(s), nil
}

func readdirnames(dirname string, errorIfNotExist bool) ([]string, error) {
	dir, err := os.Open(dirname)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && !errorIfNotExist {
			err = nil
		}
		return nil, err
	}
	defer dir.Close()
	return dir.Readdirnames(-1)
}

func getSize(name string) (uint64, error) {
	s, err := readFirstLine("/sys/class/block/"+name+"/size", true)
	if err != nil {
		return 0, err
	}
	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ui64 * defaultBlockSize, nil
}

func getRemovable(name string) (bool, error) {
	s, err := readFirstLine("/sys/class/block/"+name+"/removable", false)
	return s != "" && s != "0", err
}

func getReadOnly(name string) (bool, error) {
	s, err := readFirstLine("/sys/class/block/"+name+"/ro", false)
	return s != "" && s != "0", err
}

func getHidden(name string) bool {
	// errors ignored since real devices do not have <sys>/hidden
	// borrow idea from 'lsblk'
	// https://github.com/util-linux/util-linux/commit/c8487d854ba5cf5bfcae78d8e5af5587e7622351
	v, _ := readFirstLine("/sys/class/block/"+name+"/hidden", false)
	return v == "1"
}

func getPartitions(name string) ([]string, error) {
	names, err := readdirnames("/sys/block/"+name, false)
	if err != nil {
		return nil, err
	}

	partitions := []string{}
	for _, n := range names {
		if strings.HasPrefix(n, name) {
			partitions = append(partitions, n)
		}
	}

	return partitions, nil
}

func getHolders(name string) ([]string, error) {
	return readdirnames("/sys/block/"+name+"/holders", false)
}

func probeFS(device *Device) (fs.FS, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	fsInfo, err := fs.Probe(ctx, "/dev/"+device.Name)
	if err != nil && device.Size > 0 {
		switch {
		case errors.Is(err, fserrors.ErrFSNotFound), errors.Is(err, fserrors.ErrCanceled), errors.Is(err, io.ErrUnexpectedEOF):
		default:
			return nil, err
		}
	}
	return fsInfo, nil
}

func updateFSInfo(device *Device) error {
	// Probe only for "xfs" devices
	// UDev may have empty ID_FS_TYPE for xfs devices (ref: https://github.com/minio/directpv/issues/602)
	if device.FSType == "" || strings.EqualFold(device.FSType, "xfs") {
		fsInfo, err := probeFS(device)
		if err != nil {
			return err
		}
		if fsInfo != nil {
			device.FSUUID = fsInfo.ID()
			device.FSType = fsInfo.Type()
			device.TotalCapacity = fsInfo.TotalCapacity()
			device.FreeCapacity = fsInfo.FreeCapacity()
		}
	}
	return nil
}

func parseCDROMs(r io.Reader) (map[string]struct{}, error) {
	reader := bufio.NewReader(r)
	names := map[string]struct{}{}
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		if tokens := strings.SplitAfterN(s, "drive name:", 2); len(tokens) == 2 {
			for _, token := range strings.Fields(tokens[1]) {
				if token != "" {
					names[token] = struct{}{}
				}
			}
			break
		}
	}
	return names, nil
}

func getCDROMs() (map[string]struct{}, error) {
	file, err := os.Open("/proc/sys/dev/cdrom/info")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return map[string]struct{}{}, nil
		}
		return nil, err
	}
	defer file.Close()
	return parseCDROMs(file)
}

func getSwaps() (map[string]struct{}, error) {
	file, err := os.Open("/proc/swaps")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return map[string]struct{}{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	filenames := []string{}
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		filenames = append(filenames, strings.Fields(s)[0])
	}

	devices := map[string]struct{}{}
	for _, filename := range filenames[1:] {
		major, minor, err := getDeviceMajorMinor(filename)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}

		devices[fmt.Sprintf("%v:%v", major, minor)] = struct{}{}
	}
	return devices, nil
}

// ProbeSysInfo probes device information from /sys
func (device *Device) ProbeSysInfo() (err error) {
	device.Hidden = getHidden(device.Name)
	if device.Removable, err = getRemovable(device.Name); err != nil {
		return err
	}

	if device.ReadOnly, err = getReadOnly(device.Name); err != nil {
		return err
	}

	if device.Size, err = getSize(device.Name); err != nil {
		return err
	}

	// No partitions for hidden devices.
	if !device.Hidden {
		if device.Partition <= 0 {
			names, err := getPartitions(device.Name)
			if err != nil {
				return err
			}
			device.Partitioned = len(names) > 0
		}
		device.Holders, err = getHolders(device.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProbeDevInfo probes device information from /dev
func (device *Device) ProbeDevInfo() (err error) {
	// No FS information needed for hidden devices
	if !device.Hidden {
		CDROMs, err := getCDROMs()
		if err != nil {
			return err
		}
		if _, found := CDROMs[device.Name]; found {
			device.CDRom = true
		}
		swaps, err := getSwaps()
		if err != nil {
			return err
		}
		majorMinor := fmt.Sprintf("%v:%v", device.Major, device.Minor)
		if _, found := swaps[majorMinor]; found {
			device.SwapOn = true
		}
		if !device.CDRom {
			return updateFSInfo(device)
		}
	}
	return nil
}

// ProbeMountInfo probes mount information from /proc/1/mountinfo
func (device *Device) ProbeMountInfo() (err error) {
	mountInfosAll, err := mount.Probe()
	if err != nil {
		return err
	}
	majorMinor := mount.MajorMinor(device.Major, device.Minor)

	mountInfos, ok := mountInfosAll[majorMinor]
	if !ok {
		// not mounted
		return nil
	}
	for i, mountInfo := range mountInfos {
		if i == 0 {
			device.FirstMountPoint = mountInfo.MountPoint
			device.FirstMountOptions = mountInfo.MountOptions
		} else {
			device.OtherMountsInfo = append(device.OtherMountsInfo, mountInfo)
		}
	}
	return nil
}

func getDeviceName(major, minor uint32) (string, error) {
	filename := fmt.Sprintf("/sys/dev/block/%v:%v/uevent", major, minor)
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		if !strings.HasPrefix(s, "DEVNAME=") {
			continue
		}

		switch tokens := strings.SplitN(s, "=", 2); len(tokens) {
		case 2:
			return strings.TrimSpace(tokens[1]), nil
		default:
			return "", fmt.Errorf("filename %v contains invalid DEVNAME value", filename)
		}
	}
}
