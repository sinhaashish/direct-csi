// // This file is part of MinIO DirectPV
// // Copyright (c) 2021, 2022 MinIO, Inc.
// //
// // This program is free software: you can redistribute it and/or modify
// // it under the terms of the GNU Affero General Public License as published by
// // the Free Software Foundation, either version 3 of the License, or
// // (at your option) any later version.
// //
// // This program is distributed in the hope that it will be useful,
// // but WITHOUT ANY WARRANTY; without even the implied warranty of
// // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// // GNU Affero General Public License for more details.
// //
// // You should have received a copy of the GNU Affero General Public License
// // along with this program.  If not, see <http://www.gnu.org/licenses/>.

package installer

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _config_crd_direct_csi_min_io_directcsidrives_yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6f\x73\xdb\x36\xd2\x7f\xef\x4f\xb1\xa3\xe7\x99\x89\x9d\x47\xa2\xe2\xe4\x99\x5c\xab\x99\x4c\x26\xe7\x9c\x6f\x3c\x6d\x1a\x4f\xec\xf6\xc5\x59\xbe\x2b\x44\xae\x24\xd4\x24\xc0\x02\xa0\x6c\xf5\xe6\xbe\xfb\xcd\x02\xa4\x44\x49\xa4\xfe\xb8\x49\xd3\x36\x8b\x57\x12\x01\x2c\x16\x8b\xdd\xc5\x62\x7f\x7b\xd4\xeb\xf5\x8e\x44\x2e\x7f\x40\x63\xa5\x56\x03\x10\xb9\xc4\x07\x87\x8a\xfe\xd9\xe8\xee\x2b\x1b\x49\xdd\x9f\x9d\x1e\xdd\x49\x95\x0c\xe0\xac\xb0\x4e\x67\x1f\xd0\xea\xc2\xc4\xf8\x16\xc7\x52\x49\x27\xb5\x3a\xca\xd0\x89\x44\x38\x31\x38\x02\x10\x4a\x69\x27\xe8\xb3\xa5\xbf\x00\xb1\x56\xce\xe8\x34\x45\xd3\x9b\xa0\x8a\xee\x8a\x11\x8e\x0a\x99\x26\x68\x3c\xf1\x6a\xe9\xd9\xb3\xe8\x65\x74\x7a\x04\x10\x1b\xf4\xd3\xaf\x65\x86\xd6\x89\x2c\x1f\x80\x2a\xd2\xf4\x08\x40\x89\x0c\x07\x90\x48\x83\xb1\x8b\xad\x4c\x8c\x9c\xa1\x8d\xc2\xff\x28\xb6\x32\xca\xa4\x8a\xa4\x3e\xb2\x39\xc6\xb4\xf6\xc4\xe8\x22\xaf\x26\xd4\x07\x04\x52\x25\x7f\x61\x6f\x6f\xfd\xa0\xb3\xab\x8b\xb7\x44\xd5\x77\xa4\xd2\xba\x6f\x1a\x3a\xbf\x95\xd6\xf9\x01\x79\x5a\x18\x91\x6e\x70\xe4\xfb\xac\x54\x93\x22\x15\x66\xbd\xf7\x08\xc0\xc6\x3a\xc7\x01\x9c\xa5\x85\x75\x68\x8e\x00\x4a\x19\x78\x7e\x7a\xe5\x2e\x67\xa7\x22\xcd\xa7\xe2\x34\x10\x8b\xa7\x98\x89\xc0\x2e\x80\xce\x51\xbd\xb9\xbc\xf8\xe1\xc5\xd5\xca\x67\x80\x04\x6d\x6c\x64\xee\xbc\x3c\x57\x79\x86\x04\x95\x76\x68\xc1\x33\x01\x67\x1f\xde\x82\x1e\xfd\x44\x62\x59\xcc\xce\x8d\xce\xd1\x38\x59\xc9\x25\xb4\x9a\x76\xd4\xbe\xae\xad\xf5\x84\xd8\x09\xa3\x20\x21\xb5\x40\x0b\x6e\x8a\xd5\xc6\x30\x29\x77\x00\x7a\x0c\x6e\x2a\x2d\x18\xcc\x0d\x5a\x54\x41\x51\x56\x08\x03\x0d\x12\xaa\x62\x0f\xae\xd0\x10\x19\xb0\x53\x5d\xa4\x09\x69\xd3\x0c\x8d\x03\x83\xb1\x9e\x28\xf9\xcb\x82\xb6\x05\xa7\xfd\xa2\xa9\x70\x58\x1e\xd0\xb2\x49\xe5\xd0\x28\x91\xc2\x4c\xa4\x05\x76\x41\xa8\x04\x32\x31\x07\x83\xb4\x0a\x14\xaa\x46\xcf\x0f\xb1\x11\xbc\xd3\x06\x41\xaa\xb1\x1e\xc0\xd4\xb9\xdc\x0e\xfa\xfd\x89\x74\x95\x55\xc4\x3a\xcb\x0a\x25\xdd\xbc\xef\x15\x5c\x8e\x0a\xa7\x8d\xed\x27\x38\xc3\xb4\x6f\xe5\xa4\x27\x4c\x3c\x95\x0e\x63\x57\x18\xec\x8b\x5c\xf6\x3c\xeb\xca\x5b\x46\x94\x25\xff\x63\x4a\x3b\xb2\x4f\x56\x78\x75\x73\x52\x0e\xeb\x8c\x54\x93\x5a\x87\xd7\xd2\x2d\x27\x40\x8a\x0a\xd2\x82\x28\xa7\x86\x5d\x2c\x05\x4d\x9f\x48\x3a\x1f\xfe\x76\x75\x0d\xd5\xd2\xfe\x30\xd6\xa5\xef\xe5\xbe\x9c\x68\x97\x47\x40\x02\x93\x6a\x8c\x26\x1c\xe2\xd8\xe8\xcc\xd3\x44\x95\xe4\x5a\x2a\xe7\xff\xc4\xa9\x44\xb5\x2e\x7e\x5b\x8c\x32\xe9\xe8\xdc\x7f\x2e\xd0\x3a\x3a\xab\x08\xce\xbc\xab\x80\x11\x42\x91\x27\xc2\x61\x12\xc1\x85\x82\x33\x91\x61\x7a\x26\x2c\x7e\xf2\x03\x20\x49\xdb\x1e\x09\x76\xbf\x23\xa8\x7b\xb9\xf5\xc1\x41\x6a\xb5\x8e\xca\x07\xb5\x9c\xd7\xaa\x75\x5e\xe5\x18\xaf\x59\x28\xcd\x97\x63\x19\x7b\x03\x89\x56\x08\x35\x1b\xaa\x5f\xa2\xa2\xfa\xfe\x5e\x61\xb2\xde\xbb\xc6\x02\x9d\x85\x34\x98\x6c\x8c\x0a\x3b\x1a\x69\x9d\xa2\x58\xb7\x4d\xcf\xdc\xb5\x90\xca\x6d\x52\x17\x49\xe2\xaf\x03\x91\x5e\xb6\x72\xb8\x45\xbc\x5b\xc5\x49\xad\x54\x1e\x4c\xce\xb5\xc9\x44\x03\x03\x2b\xdb\xfb\xb0\x3a\x7a\x4d\xbc\xe3\xf0\xb1\x24\xe9\x95\x8c\x3e\x6c\xc8\x7a\xbb\xbc\xa9\x8d\x65\x8a\x76\x6e\x1d\x66\x4d\xbd\x3b\x76\x0b\xc4\x48\x8c\xdb\x66\x36\x9f\x03\xb5\x4c\x17\xca\xbd\xcf\x6b\x57\xed\x7a\x93\x0e\xb3\x96\xae\x9d\x8c\x55\x03\x84\x31\x62\xde\xd8\xff\xd0\xa3\xbb\xdc\x28\x74\x68\x7b\x74\x59\xf6\xca\x19\x4e\x67\x32\x6e\x63\xd8\x7b\x8a\x47\x89\x2a\x2f\xcc\xe4\x51\xa2\x6a\xd5\xa9\xca\x04\x56\x89\xf6\xd6\xec\x68\x2f\x73\x77\xc2\x15\x76\x7f\x83\xf7\xc3\xd7\x74\xb2\x55\x09\xdb\x15\x50\xa4\xa9\x8e\xc9\x75\x9e\x89\x5c\xc4\xd2\xcd\x37\xc5\x13\x68\x0e\xe8\x06\x7c\xf9\xff\x2d\xa2\xa1\xdb\x71\xe2\x43\x91\x7a\x8b\xb5\x0a\x06\xdd\xa0\x42\xad\x9a\xb5\xb2\xe9\xce\x59\x45\xc2\x47\x81\x42\x2a\xda\xb3\x13\x32\xb5\xc4\x17\x68\x85\x20\xc8\xd3\xb9\x10\x19\x20\xc4\x85\x31\x9b\xd7\xc7\x52\xc6\xb8\x08\x21\xde\x5c\x5e\x40\x15\x8a\x46\xd0\xeb\xf5\xe0\x9a\x3e\x5b\x67\x8a\xd8\xd1\x4d\x48\x9b\x52\x09\x26\x7e\xa5\x70\xa2\x8d\x64\x0b\x4b\x4c\x50\xc8\xe1\x55\x1d\x44\xb8\xc7\xc6\x12\xd3\x04\x72\xe1\xa6\x10\x85\xd3\x8d\x96\x02\x89\x00\xce\xb5\x01\x7c\x10\x59\x9e\x62\xb7\x55\x27\xe1\x5c\xeb\xf2\xac\x03\x63\xff\xf6\x5d\xfd\x3e\x7c\x58\xdc\xaf\x7e\x35\x3d\xb2\x68\x66\x21\x6c\xf6\x01\x50\x23\xc9\xb1\xd6\x4f\x6c\x25\xa3\x20\x8f\xa8\x22\xf8\x8d\xd2\xf7\xaa\x89\x55\xcf\x87\x30\x2d\x96\x33\xec\xbc\x99\x09\x99\x8a\x51\x8a\xc3\x4e\x17\x86\x9d\x4b\xa3\x27\x06\x2d\xc5\xaf\xf4\x81\x02\xa5\x61\xe7\x2d\x4e\x8c\x48\x30\x19\x76\xaa\xe5\xfe\x2f\x17\x2e\x9e\xbe\x43\x33\xc1\x6f\x70\xfe\x8a\x16\x69\xa6\xbf\x32\xfe\xca\x19\xe1\x70\x32\x7f\x95\xd1\xc4\x05\x2d\x72\x1e\xd7\xf3\x1c\x5f\x65\x22\x5f\xf9\xf8\x4e\xe4\xbb\xa9\x2f\x94\xcc\xc2\xcd\x2d\x5d\xd2\xb3\xd3\x68\xa9\x78\x3f\xfe\x64\xb5\x1a\x0c\x3b\x4b\x89\x74\x75\x46\xea\x9b\xbb\xf9\xb0\xd3\x48\x75\x85\xd5\xc1\xb0\xe3\x99\x1d\x76\x60\x65\xcb\x83\x61\x87\xd8\xa2\xcf\x46\x3b\x3d\x2a\xc6\x83\x61\x67\x34\x77\x68\xbb\xa7\x5d\x83\x79\x97\xe2\xf8\x57\xcb\x55\x87\x9d\x1f\x9b\xb7\xa0\xaa\x1d\x6b\x37\x45\x13\xf4\xce\xc2\x7f\x9a\x58\xdb\x7e\x13\x01\xa4\xc2\xba\x6b\x23\x94\x95\xd5\x03\xaa\xcd\xf9\xaf\x98\xe9\xe6\x34\xb2\x9f\x10\x4b\x5b\x07\x8e\x3e\x78\xe3\xac\x36\xd3\x42\x14\xc0\x2d\xa8\x90\xdd\x51\x7c\x48\x26\x1e\x74\x92\xe2\x73\xa1\xfc\x26\xa3\xd2\x56\x43\x48\x3f\x42\xb8\x9f\xe2\x16\xa2\x53\x84\x42\x25\x68\xd2\x39\x45\xb1\xf1\xd2\xa7\x4c\x85\x9a\x50\xd8\x08\x17\xe4\x14\x84\x37\x7b\x0a\x29\xef\xc8\x16\xba\x34\xb1\x9d\x6a\x61\xab\x90\xd8\xef\x8f\x38\xf0\xff\xc8\xaf\x04\xdb\x2f\xc9\xfb\xa8\x3a\x8e\x31\x77\x64\x24\x9b\xe1\x41\x68\x95\x9b\xa5\x40\xb6\x47\x14\x1f\x7b\xeb\x66\x68\xad\x68\xbb\xe7\xd6\x0e\xae\x1c\x1b\xe2\xfe\x69\x91\x09\x05\x06\x45\x42\x7c\x2e\xfb\x54\xe2\xa3\xc8\x96\xe5\x02\xcd\xe0\x92\xc5\x48\x17\xc1\xf9\x2d\xcf\xb1\x3c\x2a\x0a\xfd\x47\x48\x4e\xd2\x1b\x4e\xb9\x81\x36\x61\x64\xe2\xe1\x5b\x54\x13\x37\x1d\xc0\x8b\xe7\x7f\x79\xf9\xd5\x63\x65\x11\xbc\x22\x26\x7f\x47\x85\xc6\x3b\xc7\xbd\xc4\xb2\x39\xad\xf6\x9c\xf1\xfb\x8b\xaa\x58\x3e\x9a\x2c\xc6\x6c\xd1\xbf\xf2\x4a\x58\x6a\xde\xbd\xb0\x60\xd1\xc1\x48\x58\x4c\xa0\xc8\x49\x4e\x74\x21\x48\x65\x9d\x50\x31\x76\x41\x8e\x0f\x5b\x44\x2e\xfc\x7a\x3a\x87\xd3\xe7\x5d\x18\x95\x47\xb1\xe9\xd1\x6f\x1e\x6e\xa3\xcd\x2d\x6e\xa3\xfc\x75\x77\x8d\x7f\x69\x81\x8e\x5a\x8f\xbd\xbe\xc2\xbd\x74\x53\x7a\x14\xfa\x9b\xb8\x7c\x46\x6f\xbb\x89\x61\xf5\x36\xc6\xc5\xbe\x77\x59\x47\x73\x10\x12\x5a\x26\x95\xcc\x8a\x6c\x00\xcf\xb6\xaa\x4b\x73\xac\x12\x9a\x41\x61\xf7\xd4\x91\x30\x74\x19\x96\x08\x72\xae\x13\x23\x32\x0a\xc0\x62\x90\x09\x3d\x14\xc7\x12\xcd\x3e\x06\x44\x22\x28\x09\x52\xb0\xb1\x22\xeb\x27\xb6\xf4\xa2\x35\x93\xba\x34\x3a\x29\x62\x34\xeb\x6f\xef\x65\xd3\xe3\xc5\x0b\xb0\x76\x6c\xfe\xc5\xea\x6d\x31\x64\x59\x00\x1f\xe8\xc8\x16\x39\x0b\xba\xad\x5b\x49\x66\x28\x94\x54\x13\x5b\xb2\x48\x0f\x78\x72\x73\xe1\x8a\xbf\x9f\xa2\xbf\x7d\x7c\xd6\xa6\xa4\x65\xfc\x2e\xac\x4c\xb0\xe9\x95\x58\x35\x01\x93\x42\x18\xa1\x1c\x62\x42\xce\x93\x1c\x46\x49\xa3\xe6\xe0\xc5\xf2\x5d\xbf\xc3\x77\x40\x70\x38\xc1\x05\xd3\x56\xcb\x1c\x81\xf7\x3b\x7b\x38\x9c\xd3\x67\xcf\xb7\x68\xd8\x62\x54\xcb\x90\x5c\x38\x87\x46\x0d\xe0\x9f\x37\x6f\x7a\xff\x10\xbd\x5f\x6e\x8f\xcb\x1f\xcf\x7a\x5f\xff\xab\x3b\xb8\x7d\x5a\xfb\x7b\x7b\xf2\xfa\x7f\x1f\xeb\xda\x9a\x1e\x0c\xcb\xb6\xa2\xaa\xe5\xf5\x59\x45\xc8\x95\x36\x74\xfd\xdd\xaa\xc7\x70\x6d\x0a\xec\xc2\xb9\x48\x2d\x76\xe1\x7b\xe5\x2f\xbf\x36\x41\xa1\x2a\x5a\xde\xa9\xf4\xee\xe9\x10\xa9\xe6\x98\xc8\x77\xfb\x35\xda\xfb\xcb\xb5\x7f\xd5\x7b\x73\x1f\x81\xf8\x88\x56\x8f\xeb\xfe\xac\x96\x37\x02\xef\x87\x29\x56\x8e\xca\xf8\x3c\x8a\x75\xd6\x5f\xe6\x95\x5a\x15\x8f\x1e\x11\xef\x84\x9a\xc3\xd2\xd9\x86\xe8\x79\xdd\x22\xac\xa3\xf8\x5b\xc4\x46\x5b\xbb\x48\xa6\xb5\x1b\x73\x2a\xef\x10\x16\x61\x76\x70\xed\x23\x8c\x85\x7f\x79\x98\x91\x74\x46\x98\x79\xed\xb9\x05\xb1\x50\x3e\x2d\x66\x71\x5c\xa4\xad\x64\x8f\x2d\x22\x44\x4a\x27\xb8\x79\x47\x9c\x04\x8f\x2f\x46\x32\x95\x6e\x4e\x3e\x3d\xc1\x58\xab\x71\x2a\xfd\xe3\xa8\xfd\xb2\xc8\x72\x6d\x9c\x50\x2e\x98\xb1\xc1\x09\x3e\x80\x74\x90\x51\xe8\x8b\x96\x2e\x8e\xe3\x44\xd9\xd3\xd3\xe7\x2f\xae\x8a\x51\xa2\x33\x21\xd5\x79\xe6\xfa\x27\xaf\x8f\x7f\x2e\x44\x4a\x1e\x33\xf9\x4e\x64\x78\x9e\xb9\x93\x3d\x82\x83\xd3\x97\x3b\xed\xf0\xf8\x26\x58\xdb\xed\xf1\x4d\xaf\xfc\xf5\xb4\xfa\x74\xf2\xfa\x78\x18\x6d\xed\x3f\x79\x4a\xac\xd5\x6c\xf8\xf6\xa6\xb7\x34\xe0\xe8\xf6\xe9\xc9\xeb\x5a\xdf\xc9\x23\xcd\xb9\x39\x8f\x10\x5a\xaf\x21\xbc\x6e\x1c\x56\x06\x6c\x8d\x7d\xe1\x72\x69\xec\x0a\x47\xdf\xd8\xd5\xf2\x6c\xda\x92\x62\xdb\x9e\xf4\xd9\x4c\xf8\x64\x22\xef\xdd\xe1\xbc\xc1\x8f\xb5\xac\xde\x96\x33\xca\x44\xde\x94\x69\xbc\x6a\xf1\x92\xab\xa9\x95\xd6\x8c\x4a\x69\x16\x2d\x9b\x6c\x3c\xce\x6d\xe9\xbc\x6d\xd3\x0c\xe2\xa7\xc8\xc1\xa4\x7a\x22\x63\x91\xfe\x35\xd5\xf1\xdd\x95\xfc\xa5\xc1\x3f\x3e\x9e\x76\xa6\x13\x4c\xbf\x2b\xb2\x11\x9a\x83\xf6\xba\x3d\xef\xd8\x9a\x19\xda\x23\xed\xbb\xaf\xda\x6d\xc9\x33\x6e\xcb\x31\x6e\xe1\x80\xbc\x28\xf9\xad\x83\x26\xe5\xc2\x38\x6f\xd3\xdf\x35\x5d\xaa\xdb\x44\x9f\x0b\x37\x3d\x6c\xa9\xe9\xdc\x7e\x32\x45\x30\x5a\xbb\xcb\x6a\x2f\x07\xb1\x65\xd1\x48\xf1\x18\x1d\x72\x3a\xd7\xa9\x9e\x34\xd8\xca\xa7\x46\x11\x9c\x76\x22\xfd\xf8\xa6\xda\x96\x4a\xa6\x93\xde\x9d\x40\xde\x9c\xdd\x5b\xc0\x4d\xb5\x4f\xf4\x24\x38\x6a\x25\x14\x5e\x84\x03\x70\xa6\x08\x8e\xd7\x3a\x6d\xc4\x04\x07\x30\xa6\xb8\x6d\x05\x5c\x1e\xa1\x63\x6c\x99\xb1\xe5\xd0\x18\x5b\x66\x6c\xb9\x6c\x8c\x2d\x33\xb6\x0c\x8c\x2d\x6f\xce\xfc\x02\xb1\xe5\x38\x46\x6b\xaf\x65\x53\x64\xb7\xb2\xfc\x9b\xc5\xc0\xc5\xa2\x61\x2e\x38\x89\xe6\xa0\xd7\x17\xe3\xd9\x8c\x67\x03\xe3\xd9\x8c\x67\x87\xc6\x78\x36\xe3\xd9\x8c\x67\x33\x9e\xbd\x4e\x99\xf1\x6c\x60\x3c\x9b\xf1\x6c\xc6\xb3\x19\xcf\x66\x3c\x9b\xf1\x6c\xc6\xb3\x19\xcf\x66\x3c\x9b\xf1\xec\x5a\x63\x3c\x9b\xf1\x6c\xdf\xbe\x20\x3c\xfb\x79\x18\xc4\x78\x36\xe3\xd9\x8c\x67\x33\x9e\xed\x1b\xe3\xd9\x8c\x67\x03\xe3\xd9\x9b\x33\x19\xcf\x6e\x5d\x9e\xf1\x6c\xc6\xb3\x19\xcf\x66\x3c\xbb\xdc\x02\xe3\xd9\x8c\x67\xd7\x1a\xe3\xd9\xcb\xc6\x78\x36\xe3\xd9\xb5\x73\x60\x3c\x9b\xf1\xec\x96\xbd\x31\x9e\xcd\x78\x36\xe3\xd9\x8c\x67\x33\x9e\x5d\xeb\x62\x3c\xfb\xb7\xc3\xb3\x17\xd3\xbe\xff\xfe\xe2\xed\x9f\x1f\x0a\x17\x3f\x69\xd3\x06\x63\xd6\xc8\xbe\x78\x7e\x18\x59\xa9\x3e\x09\x59\x06\xee\x17\xed\x37\x07\xee\xcb\x99\x07\x9b\x05\x43\xfe\x0c\xf9\x7f\x76\xc8\xff\x45\x18\xc4\x90\x3f\x43\xfe\x0c\xf9\x33\xe4\xef\x1b\x43\xfe\x0c\xf9\x03\x43\xfe\x9b\x33\x19\xf2\x6f\x5d\x9e\x21\x7f\x86\xfc\x19\xf2\x67\xc8\xbf\xdc\x02\x43\xfe\x0c\xf9\xd7\x1a\x43\xfe\xcb\xc6\x90\x3f\x43\xfe\xb5\x73\x60\xc8\x9f\x21\xff\x96\xbd\x31\xe4\xcf\x90\x3f\x43\xfe\x0c\xf9\x33\xe4\x5f\xeb\xfa\x63\x43\xfe\xd9\xc1\xc8\x64\x72\x38\xde\xce\x85\x05\xa1\xf7\x8b\x2a\x2c\x10\xd6\x1d\x0a\xfe\x27\x07\x0b\x9c\xcb\x17\x7e\xbf\xe5\x0b\xd7\x74\x0f\x5f\x37\x86\x1b\xfb\xcc\x7c\x44\xf9\xc2\xe7\x28\x99\x28\x67\x36\x5d\x4b\xdb\x72\xf4\xbf\xb3\x5a\x0b\x14\xc9\x7b\x95\x36\xf8\xac\x6d\x7b\xf8\x0c\x15\x1a\xf6\x5e\xe4\xef\x5b\xd7\x6a\x66\xf3\x4f\x57\xd5\x01\x50\xe0\x0c\x95\x3b\xbf\x3a\x58\x5f\xc3\xc4\x2b\x2f\xfe\x83\x26\xce\x50\x25\xfa\xb0\xb3\x9a\x49\xe3\x8a\xf6\x65\x9a\x0f\xeb\xfe\x5e\xb6\x5a\x52\xc3\x2a\x7f\xc0\x02\x97\x70\xd2\x5c\xe0\xc2\x05\x2e\x5c\xe0\xc2\x05\x2e\xa1\x71\x81\x0b\x17\xb8\x00\x17\xb8\x6c\xce\xe4\x02\x97\xd6\xe5\xb9\xc0\x85\x0b\x5c\xb8\xc0\x85\x0b\x5c\xca\x2d\x70\x81\x0b\x17\xb8\xd4\x1a\x17\xb8\x2c\x1b\x17\xb8\x70\x81\x4b\xed\x1c\xb8\xc0\x85\x0b\x5c\x5a\xf6\xc6\x05\x2e\x5c\xe0\xc2\x05\x2e\x5c\xe0\xc2\x05\x2e\xb5\x2e\x2e\x70\xd9\x3d\x85\x0b\x5c\x80\x0b\x5c\xea\x73\xbe\xc8\x02\x97\xb5\x6c\xbd\x48\x4a\x44\xa9\x3f\x13\xa6\x9f\xca\x51\x3f\xe4\xcf\x7a\xb1\x95\xfb\xa7\x02\x3f\x7f\x71\xcc\xaf\xd9\xd6\xc7\x2e\xac\xf1\x89\x8f\x77\xc4\xac\xbd\x50\x63\x7d\x40\x4a\x75\x57\xd6\x67\x37\xf6\xb0\x03\x7d\xd8\x23\x36\xdb\x8d\x40\xec\x02\x14\x76\x2e\xf2\x11\xaf\xa5\x2d\xea\xc2\xf5\x4d\x1f\xbd\xbe\x29\x96\x97\x5c\x13\xf5\xa8\x9a\xa8\xda\xc4\x6f\xb5\x9a\x70\x41\x15\x17\x54\x2d\xdb\x9f\xa9\xa0\xca\x7f\x59\xa6\x3c\x42\x3a\x3d\xbc\x14\x4b\x35\xf3\x35\x35\xd0\x09\xc9\x85\x3c\x2d\x8c\x48\xcb\xbf\x35\x18\x12\x6e\x6e\x8f\x02\x55\x4c\xca\x1a\xa7\xf0\xf1\xbf\x01\x00\x00\xff\xff\x19\xda\x2f\x56\x3d\xa7\x00\x00")

func config_crd_direct_csi_min_io_directcsidrives_yaml() ([]byte, error) {
	return bindata_read(
		_config_crd_direct_csi_min_io_directcsidrives_yaml,
		"config/crd/direct.csi.min.io_directcsidrives.yaml",
	)
}

var _config_crd_direct_csi_min_io_directcsivolumes_yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5f\x8f\xdb\xb8\x11\x7f\xf7\xa7\x18\xb8\x05\xb2\x9b\x5a\x72\x9c\x1c\xd2\x3b\x03\x41\x10\x6c\x9a\x22\xc8\xa5\x08\xb2\xdb\x3c\x74\xbd\xed\xd1\xd2\x58\xe6\xad\x44\x2a\x24\xe5\x5d\x5f\xd1\xef\x5e\x0c\x29\x59\xb2\x2d\x39\x9b\xa0\xf7\xd0\x60\xe6\x69\x4d\x52\x3f\x0e\x87\xf3\x8f\xbf\x97\x1d\x45\x51\x34\x12\xa5\xfc\x84\xc6\x4a\xad\xe6\x20\x4a\x89\xf7\x0e\x15\xfd\xb2\xf1\xed\x8f\x36\x96\x7a\xba\x99\x8d\x6e\xa5\x4a\xe7\x70\x51\x59\xa7\x8b\x8f\x68\x75\x65\x12\x7c\x8d\x2b\xa9\xa4\x93\x5a\x8d\x0a\x74\x22\x15\x4e\xcc\x47\x00\x42\x29\xed\x04\x0d\x5b\xfa\x09\x90\x68\xe5\x8c\xce\x73\x34\x51\x86\x2a\xbe\xad\x96\xb8\xac\x64\x9e\xa2\xf1\xe0\xcd\xd6\x9b\x27\xf1\xf3\x78\x36\x02\x48\x0c\xfa\xcf\xaf\x64\x81\xd6\x89\xa2\x9c\x83\xaa\xf2\x7c\x04\xa0\x44\x81\x73\x48\xa5\xc1\xc4\x25\x56\x6e\x74\x5e\x15\x68\xe3\x30\x10\x27\x56\xc6\x85\x54\xb1\xd4\x23\x5b\x62\x42\x9b\x67\x46\x57\x65\xf3\x45\x77\x41\xc0\xaa\x15\x0c\x87\x7b\xed\x17\x5d\x5c\xbe\xfd\xe4\x61\xfd\x4c\x2e\xad\x7b\xd7\x37\xfb\xb3\xb4\xce\xaf\x28\xf3\xca\x88\xfc\x58\x29\x3f\x69\xa5\xca\xaa\x5c\x98\xa3\xe9\x11\x80\x4d\x74\x89\x73\xb8\xc8\x2b\xeb\xd0\x8c\x00\x6a\x43\x78\x9d\xa2\xfa\xa8\x9b\x99\xc8\xcb\xb5\x98\x05\xb4\x64\x8d\x85\x08\x2a\x03\xe8\x12\xd5\xab\x0f\x6f\x3f\x3d\xbb\xdc\x1b\x06\x48\xd1\x26\x46\x96\xce\x1b\xf5\x40\x6d\x48\x51\x69\x87\x16\x82\x1a\x70\xf1\xf1\x35\xe8\xe5\xaf\x64\x9c\xdd\xf7\xa5\xd1\x25\x1a\x27\x1b\xeb\x04\xe9\x38\x49\x67\xf4\x60\xb7\x47\xa4\x50\x58\x05\x29\x79\x07\x5a\x70\x6b\x6c\x8e\x86\x69\x7d\x06\xd0\x2b\x70\x6b\x69\xc1\x60\x69\xd0\xa2\x0a\xfe\xb2\x07\x0c\xb4\x48\xa8\x46\x3d\xb8\x44\x43\x30\x60\xd7\xba\xca\x53\x72\xaa\x0d\x1a\x07\x06\x13\x9d\x29\xf9\xdb\x0e\xdb\x82\xd3\x7e\xd3\x5c\x38\xac\x2f\xa9\x15\xa9\x1c\x1a\x25\x72\xd8\x88\xbc\xc2\x09\x08\x95\x42\x21\xb6\x60\x90\x76\x81\x4a\x75\xf0\xfc\x12\x1b\xc3\x7b\x6d\x10\xa4\x5a\xe9\x39\xac\x9d\x2b\xed\x7c\x3a\xcd\xa4\x6b\x82\x23\xd1\x45\x51\x29\xe9\xb6\x53\xef\xe7\x72\x59\x39\x6d\xec\x34\xc5\x0d\xe6\x53\x2b\xb3\x48\x98\x64\x2d\x1d\x26\xae\x32\x38\x15\xa5\x8c\xbc\xea\xca\x07\x48\x5c\xa4\x7f\x30\x75\x38\xd9\x47\x7b\xba\xba\x2d\xb9\x87\x75\x46\xaa\xac\x33\xe1\x7d\xf5\xc4\x0d\x90\xb7\x82\xb4\x20\xea\x4f\xc3\x29\x5a\x43\xd3\x10\x59\xe7\xe3\x5f\x2e\xaf\xa0\xd9\xda\x5f\xc6\xa1\xf5\xbd\xdd\xdb\x0f\x6d\x7b\x05\x64\x30\xa9\x56\x68\xc2\x25\xae\x8c\x2e\x3c\x26\xaa\xb4\xd4\x52\x39\xff\x23\xc9\x25\xaa\x43\xf3\xdb\x6a\x59\x48\x47\xf7\xfe\xb9\x42\xeb\xe8\xae\x62\xb8\xf0\x19\x03\x96\x08\x55\x99\x0a\x87\x69\x0c\x6f\x15\x5c\x88\x02\xf3\x0b\x61\xf1\x77\xbf\x00\xb2\xb4\x8d\xc8\xb0\x0f\xbb\x82\x6e\xb2\x3b\x5c\x1c\xac\xd6\x99\xb0\x4e\xb8\xca\x9e\xb8\xb1\x83\x08\xbd\xf4\xeb\x0f\xe3\x94\x0e\x6f\x0a\x1f\x24\xf1\x1e\x54\x7f\xb0\x92\x88\x8d\x90\xb9\x58\xe6\x78\x21\x4a\x91\x48\xb7\x3d\x5c\x00\x10\x30\xe7\x14\x14\xcf\x7f\x38\x9a\x0d\x07\xa2\x80\xc9\x7c\x7e\xea\x4a\xa2\x55\x2a\x3b\x29\xbe\x2b\xd2\x61\xd1\x33\x7c\x70\xec\xf1\x45\x03\xe1\xeb\x83\x90\x8a\x0e\xed\x84\xcc\x2d\xe9\x05\x5a\x21\x08\x4a\xe3\x2e\x24\x0b\x84\xa4\x32\xe6\xd8\xa3\x5a\x2b\xe3\x2e\xab\xbc\xfa\xf0\x16\x9a\x22\x15\x43\x14\x45\x70\x45\xc3\xd6\x99\x2a\x71\x14\x1c\x74\x28\x95\x62\xea\x77\x0a\xa9\xb9\x17\xb6\xb2\xa4\x04\x65\x21\x61\x8c\xd8\x82\x08\xae\xbd\x92\x98\xa7\x50\x0a\xb7\x86\x38\xdc\x6f\xdc\x1a\x24\x06\x78\xa3\x0d\xe0\xbd\x28\xca\x1c\x27\xbd\xb8\x64\x5a\x78\xa3\x75\x7d\xd9\x41\xb1\x7f\xfb\xa9\xe9\x14\x3e\xee\x42\xce\xef\xa6\x97\x16\xcd\x26\x14\x54\x9f\x13\x7b\x21\x57\x5a\x3f\xb2\x8d\x8d\x82\x3d\xe2\x06\xf0\x9d\xd2\x77\xaa\x4f\x55\xaf\x87\x30\xd8\x77\x5b\x00\x8b\xf1\xab\xc6\x87\x16\xe3\x09\x2c\xc6\x1f\x8c\xce\x0c\x5a\xaa\x6a\x34\x40\xb9\x73\x31\x7e\x8d\x99\x11\x29\xa6\x8b\x71\xb3\xdd\x9f\x4a\xe1\x92\xf5\x7b\x34\x19\xbe\xc3\xed\x0b\xda\xa4\x1f\x7f\x6f\xfd\xa5\x33\xc2\x61\xb6\x7d\x51\xd0\x87\x3b\x2c\xaa\xc0\x57\xdb\x12\x5f\x14\xa2\xdc\x1b\x7c\x2f\xca\x2f\xa3\xef\x9c\xcc\xc2\xf5\x0d\xc5\xed\x66\x16\xb7\x8e\xf7\xcb\xaf\x56\xab\xf9\x62\xdc\x5a\x64\xa2\x0b\x72\xdf\xd2\x6d\x17\xe3\x5e\xd4\x3d\x55\xe7\x8b\xb1\x57\x76\x31\x86\xbd\x23\xcf\x17\x63\x52\x8b\x86\x8d\x76\x7a\x59\xad\xe6\x8b\xf1\x72\xeb\xd0\x4e\x66\x13\x83\xe5\x84\x8a\xfb\x8b\x76\xd7\xc5\xf8\x97\xfe\x23\xa8\xe6\xc4\xda\xad\xd1\x04\xbf\xb3\xf0\x9f\x3e\xd5\x86\x13\x41\x90\x5c\x58\x77\x65\x84\xb2\xb2\x69\xad\xfa\xd7\x1d\x84\xe9\xf1\x67\x14\x3f\xa1\xbc\x5a\x07\x8e\x06\x7c\x70\x36\x87\x19\x00\x05\x70\x3b\x14\x8a\x3b\x2a\x19\x14\xe2\xc1\x27\xa9\x64\x0b\xe5\x0f\x19\xd7\xb1\x1a\xaa\xfc\x12\xe1\x6e\x8d\x27\x40\xd7\x08\x95\x4a\xd1\xe4\x5b\x2a\x6c\x49\x9b\x53\xd6\x42\x65\x54\x49\xe0\x2d\x25\x05\xe1\xc3\x9e\xaa\xcc\x2d\xc5\xc2\x84\x3e\x1c\x46\xad\x6c\x53\x25\xfd\xf9\x48\x03\xff\x8b\xf2\x4a\x88\xfd\x1a\xde\x17\xda\x24\xc1\xd2\x51\x90\xc4\x03\x80\x4d\x9a\xa5\xda\x16\x11\xe2\xc0\xba\x81\x72\xd3\x4a\x81\xd6\x8a\xec\x61\x17\x57\xaf\x0d\xad\xc0\xba\x2a\x84\x02\x83\x22\x25\x3d\xdb\x39\x95\xca\x44\xb8\xa1\xed\x02\x66\x48\xc9\x62\xa9\xab\x90\xfc\xda\x7b\xac\xaf\x8a\xba\x81\x25\x52\x92\xf4\x81\x53\x1f\x60\xc8\x18\x85\xb8\xff\x19\x55\xe6\xd6\x73\x78\xf6\xf4\xcf\xcf\x7f\xfc\x56\x5b\x84\xac\x88\xe9\x5f\x51\xa1\xf1\xc9\xf1\x41\x66\x39\xfe\xac\xd3\xe1\xf8\xf3\xc5\x4d\x79\x8f\xb3\xdd\x9a\x13\xfe\x57\x97\x84\xd6\xf3\xee\x84\x05\x8b\x0e\x96\xc2\x62\x0a\x55\x49\x76\xa2\x82\x20\x95\x75\x42\x25\x38\x01\xb9\xfa\xba\x4d\xe4\x2e\xaf\xe7\x5b\x98\x3d\x9d\xc0\xb2\xbe\x8a\xe3\x8c\x7e\x7d\x7f\x13\x1f\x1f\xf1\x14\xf2\x4f\x93\x03\xfd\xa5\x05\xba\x6a\xbd\xf2\xfe\x0a\x77\xd2\xad\xa9\x4f\xf4\x95\xb8\xee\xac\x4f\x55\x62\xd8\xaf\xc6\xb8\x3b\xf7\x97\xa2\xa3\xbf\x09\x09\x52\x48\x25\x8b\xaa\x98\xc3\x93\x93\xee\xd2\xdf\xab\x04\x31\x28\xec\x03\x7d\x24\x2c\x6d\xdb\x12\x41\xc9\x35\x33\xa2\xa0\x06\x2c\x01\x99\x52\xef\xb8\x92\x68\x1e\x12\x40\x64\x82\x1a\x90\x9a\x8d\x3d\x5b\x3f\xb2\x75\x16\xed\x84\xd4\x07\xa3\xd3\x2a\x41\x73\xd8\x8e\xb7\xa2\x57\x40\xb7\x21\x57\x32\xe9\x5c\x9b\x6f\x62\x7d\x2c\x86\x87\x17\xe0\x3d\x5d\xd9\xee\x19\x43\xd5\x7a\x10\xb2\x40\xa1\xa4\xca\x6c\xad\x22\xf5\xf4\x94\xe6\x42\x89\xbf\x5b\xa3\xaf\x3e\xfe\x21\x57\x63\x19\x7f\x0a\x2b\x53\x34\x38\x0c\x2b\x20\xab\x84\x11\xca\x21\xa6\x94\x3c\x29\x61\xd4\x18\x9d\x04\x2f\xda\x56\xff\x0b\xb9\x03\x42\xc2\x09\x29\x98\x8e\x5a\x3f\x1b\x7c\xde\x79\x40\xc2\x99\x3d\x79\x7a\xc2\xc3\x76\xab\x06\x96\x94\xc2\xd1\xdb\x71\x0e\xff\xbc\x7e\x15\xfd\x43\x44\xbf\xdd\x9c\xd5\x7f\x3c\x89\x7e\xfa\xd7\x64\x7e\xf3\xb8\xf3\xf3\xe6\xfc\xe5\x1f\xbf\x35\xb5\xf5\x3d\x19\x5a\xd9\x73\xd5\xba\x7c\x36\x1d\x72\xe3\x0d\x13\x5f\x5b\xf5\x0a\xae\x0c\x3d\x72\xdf\x88\xdc\xe2\x04\xfe\xae\x7c\xf1\x1b\x32\x14\xaa\xaa\x18\xda\x34\x82\x31\x41\xf5\xf7\x44\x7e\xda\xef\x31\x3c\x5f\xef\xfd\xad\x26\xf1\x0b\x1e\x62\x10\xdf\xd1\xea\x55\x37\x9f\x75\x9e\x92\xe0\xf3\x30\xf5\xca\x71\xdd\x9f\xc7\x89\x2e\xa6\xed\x53\x73\xd0\xf1\xe8\x11\xf1\x5e\xa8\x2d\xb4\xc9\x36\x74\xcf\x87\x11\x61\x1d\xf5\xdf\x22\x31\xda\xda\xdd\xfb\x7a\x38\x98\x73\x79\x8b\xb0\x6b\xb3\x43\x6a\x5f\x62\x22\xfc\xcb\xc3\x2c\xa5\x33\xc2\x6c\x3b\xcf\x2d\x48\x84\xf2\x2f\x65\x8b\xab\x2a\x1f\x84\x3d\xb3\x88\x10\x2b\x9d\xe2\x71\x8d\x38\x0f\x19\x5f\x2c\x65\x2e\xdd\x96\x72\x7a\x8a\x89\x56\xab\x5c\xfa\xc7\xd1\x70\xb1\x28\x4a\x6d\x9c\x50\x2e\x84\xb1\xc1\x0c\xef\x41\x3a\x28\xa8\xf5\x45\x4b\x85\xe3\x2c\x55\x76\x36\x7b\xfa\xec\xb2\x5a\xa6\xba\x10\x52\xbd\x29\xdc\xf4\xfc\xe5\xd9\xe7\x4a\xe4\x94\x31\xd3\xbf\x89\x02\xdf\x14\xee\xfc\x01\xcd\xc1\xec\xf9\x17\xe3\xf0\xec\x3a\x44\xdb\xcd\xd9\x75\x54\xff\xf5\xb8\x19\x3a\x7f\x79\xb6\x88\x4f\xce\x9f\x3f\x26\xd5\x3a\x31\x7c\x73\x1d\xb5\x01\x1c\xdf\x3c\x3e\x7f\xd9\x99\x3b\xff\xc6\x70\x36\xf8\xb9\x92\x06\xd3\x3e\xef\x8d\x7a\xda\xeb\xde\x65\x75\xc3\xd6\x3b\x17\x8a\x4b\xef\x54\xb8\xfa\xde\xa9\x81\x67\xd3\x00\x89\xd1\x9d\xf4\x2f\xe1\xa3\xb9\xfb\xe8\xb6\x5a\xa2\x51\xe8\xd0\x46\xf4\x3c\x8b\x0a\x51\x46\xb7\xb8\xed\xc9\x63\x03\xbb\x1f\x43\x84\x0d\x0b\x51\x1e\xb3\x0f\x54\x99\xd1\x7c\x10\x6e\x7d\x8c\x7f\xe2\x46\x52\x23\x37\x3d\x89\xe4\xc4\x17\x6b\x6d\xdd\x57\x6f\x43\x81\x47\xae\xfe\x55\x1f\x59\x27\x32\xa9\xb2\xaf\xde\xcc\x69\x27\xf2\xdf\x83\xe4\xa9\x2c\xa6\xff\x7b\xdc\x5e\x17\x3b\x8e\x92\x68\x47\xb3\x8d\x06\xbf\x0c\x7d\xee\x1c\x9c\xa9\x82\x3b\x59\xa7\x0d\x3d\x90\x60\x45\xd5\x68\x8f\x47\x5f\xa2\x63\x1a\x9d\x69\xf4\x46\x98\x46\x67\x1a\xbd\x23\x4c\xa3\xef\xac\xcc\x34\x3a\xd3\xe8\x87\xe8\x4c\xa3\x37\xc2\x34\x3a\xd3\xe8\x4c\xa3\x33\x8d\x7e\x88\xcc\x34\x3a\x30\x8d\xce\x34\x3a\xd3\xe8\x4c\xa3\x33\x8d\xce\x34\x3a\xd3\xe8\x4c\xa3\xef\x09\xd3\xe8\xdf\x0f\x8d\xfe\x94\x69\x74\xa6\xd1\x83\x30\x8d\xce\x34\x7a\x47\x98\x46\xdf\x59\x99\x69\x74\xa6\xd1\x0f\xd1\x99\x46\x6f\x84\x69\x74\xa6\xd1\x99\x46\x67\x1a\xfd\x10\x99\x69\x74\x60\x1a\x9d\x69\x74\xa6\xd1\x99\x46\x67\x1a\x9d\x69\x74\xa6\xd1\x99\x46\xdf\x13\xa6\xd1\xbf\x1f\x1a\xfd\x19\xd3\xe8\x4c\xa3\x07\x61\x1a\x9d\x69\xf4\x8e\x30\x8d\xbe\xb3\x32\xd3\xe8\x4c\xa3\x1f\xa2\x33\x8d\xde\x08\xd3\xe8\x4c\xa3\x33\x8d\xce\x34\xfa\x21\x32\xd3\xe8\xc0\x34\x3a\xd3\xe8\x4c\xa3\x33\x8d\xce\x34\x3a\xd3\xe8\x4c\xa3\x33\x8d\xbe\x27\x4c\xa3\x7f\x3f\x34\xfa\x0f\x4c\xa3\x33\x8d\x1e\x84\x69\x74\xa6\xd1\x3b\xc2\x34\xfa\xce\xca\x4c\xa3\x33\x8d\x7e\x88\xce\x34\x7a\x23\x4c\xa3\x33\x8d\xce\x34\x3a\xd3\xe8\x87\xc8\x4c\xa3\x03\xd3\xe8\x4c\xa3\x33\x8d\xce\x34\x3a\xd3\xe8\x4c\xa3\x33\x8d\xce\x34\xfa\x9e\x30\x8d\xfe\xff\x48\xa3\xfb\x91\xb6\x8e\x86\x37\x5a\x48\x3f\x7b\xff\x0f\x75\x1c\x2a\x56\xf3\x0f\x4e\xfd\xcf\x0e\xb7\x05\xd7\x37\xa3\x80\x8a\xe9\xa7\xe6\x5f\x97\xd2\xe0\x7f\x03\x00\x00\xff\xff\xff\xc8\xfc\x74\x54\x76\x00\x00")

func config_crd_direct_csi_min_io_directcsivolumes_yaml() ([]byte, error) {
	return bindata_read(
		_config_crd_direct_csi_min_io_directcsivolumes_yaml,
		"config/crd/direct.csi.min.io_directcsivolumes.yaml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"config/crd/direct.csi.min.io_directcsidrives.yaml":  config_crd_direct_csi_min_io_directcsidrives_yaml,
	"config/crd/direct.csi.min.io_directcsivolumes.yaml": config_crd_direct_csi_min_io_directcsivolumes_yaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"config": {nil, map[string]*_bintree_t{
		"crd": {nil, map[string]*_bintree_t{
			"direct.csi.min.io_directcsidrives.yaml":  {config_crd_direct_csi_min_io_directcsidrives_yaml, map[string]*_bintree_t{}},
			"direct.csi.min.io_directcsivolumes.yaml": {config_crd_direct_csi_min_io_directcsivolumes_yaml, map[string]*_bintree_t{}},
		}},
	}},
}}
