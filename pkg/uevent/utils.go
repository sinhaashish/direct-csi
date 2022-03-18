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

package uevent

import (
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	directcsi "github.com/minio/directpv/pkg/apis/direct.csi.min.io/v1beta4"
	"github.com/minio/directpv/pkg/sys"
)

func getRootBlockPath(devName string) string {
	switch {
	case strings.HasPrefix(devName, sys.HostDevRoot):
		return devName
	case strings.Contains(devName, sys.DirectCSIDevRoot):
		return getRootBlockPath(filepath.Base(devName))
	default:
		name := strings.ReplaceAll(
			strings.Replace(devName, sys.DirectCSIPartitionInfix, "", 1),
			sys.DirectCSIPartitionInfix,
			sys.HostPartitionInfix,
		)
		return filepath.Join(sys.HostDevRoot, name)
	}
}

func validateDevice(device *sys.Device, filteredDrives []*directcsi.DirectCSIDrive) bool {
	if len(filteredDrives) != 1 {
		return false
	}
	directCSIDrive := filteredDrives[0]
	if directCSIDrive.Status.Path != device.DevPath() {
		return false
	}
	if directCSIDrive.Status.MajorNumber != uint32(device.Major) {
		return false
	}
	if directCSIDrive.Status.MinorNumber != uint32(device.Minor) {
		return false
	}
	if directCSIDrive.Status.Virtual != device.Virtual {
		return false
	}
	if directCSIDrive.Status.PartitionNum != device.Partition {
		return false
	}
	if directCSIDrive.Status.WWID != device.WWID {
		return false
	}
	if directCSIDrive.Status.ModelNumber != device.Model {
		return false
	}
	if directCSIDrive.Status.UeventSerial != device.UeventSerial {
		return false
	}
	if directCSIDrive.Status.Vendor != device.Vendor {
		return false
	}
	if directCSIDrive.Status.DMName != device.DMName {
		return false
	}
	if directCSIDrive.Status.DMUUID != device.DMUUID {
		return false
	}
	if directCSIDrive.Status.MDUUID != device.MDUUID {
		return false
	}
	if directCSIDrive.Status.PartTableUUID != device.PTUUID {
		return false
	}
	if directCSIDrive.Status.PartTableType != device.PTType {
		return false
	}
	if directCSIDrive.Status.PartitionUUID != device.PartUUID {
		return false
	}
	if directCSIDrive.Status.Filesystem != device.FSType {
		return false
	}

	if len(device.MountInfos) > 0 {
		if directCSIDrive.Status.Mountpoint != device.MountInfos[0].MountPoint {
			return false
		}
		deviceMountOptions := device.MountInfos[0].MountOptions
		sort.Strings(deviceMountOptions)
		driveMountOptions := directCSIDrive.Status.MountOptions
		sort.Strings(driveMountOptions)
		if !reflect.DeepEqual(deviceMountOptions, driveMountOptions) {
			return false
		}

	}

	return true
}
