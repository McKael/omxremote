package main

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

func static_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd4, 0x1a,
		0x6d, 0x6f, 0xdb, 0xb8, 0xf9, 0x7b, 0x7f, 0x05, 0xab, 0x74, 0x88, 0x83,
		0x8b, 0x25, 0xc7, 0x4d, 0x86, 0xd4, 0xb1, 0x3d, 0x6c, 0xdd, 0x0d, 0xdb,
		0x80, 0x5d, 0x0f, 0x68, 0xef, 0xc3, 0x70, 0x18, 0x02, 0x5a, 0xa2, 0x63,
		0x36, 0x94, 0xa8, 0x89, 0x94, 0x9d, 0xdc, 0x5d, 0xfe, 0xfb, 0x1e, 0x92,
		0x12, 0x45, 0x51, 0x92, 0xed, 0x2c, 0xed, 0xed, 0x86, 0x00, 0xb1, 0x4c,
		0x3e, 0x7c, 0xde, 0x5f, 0x29, 0xcf, 0x5f, 0xff, 0xf9, 0xc3, 0xfb, 0x4f,
		0xff, 0xfc, 0xfe, 0x5b, 0xb4, 0x91, 0x29, 0x5b, 0xbe, 0x9a, 0x9b, 0x0f,
		0x84, 0xe6, 0x1b, 0x82, 0x13, 0xf5, 0x00, 0x8f, 0x29, 0x91, 0x18, 0x65,
		0x38, 0x25, 0x8b, 0x60, 0x4b, 0xc9, 0x2e, 0xe7, 0x85, 0x0c, 0x50, 0xcc,
		0x33, 0x49, 0x32, 0xb9, 0x08, 0x76, 0x34, 0x91, 0x9b, 0x45, 0x42, 0xb6,
		0x34, 0x26, 0x63, 0xfd, 0xe5, 0x1c, 0xd1, 0x8c, 0x4a, 0x8a, 0xd9, 0x58,
		0xc4, 0x98, 0x91, 0xc5, 0x45, 0x38, 0x39, 0x47, 0x29, 0x7e, 0xa0, 0x69,
		0x99, 0xba, 0x4b, 0xa5, 0x20, 0x85, 0xfe, 0x8e, 0x57, 0xb0, 0x94, 0xf1,
		0x00, 0x45, 0x5d, 0x8a, 0x38, 0xcf, 0x19, 0x19, 0xa7, 0x7c, 0x45, 0xe1,
		0x63, 0x47, 0x56, 0x63, 0x58, 0x18, 0xc7, 0x38, 0x57, 0x67, 0x1c, 0x2e,
		0x1e, 0x89, 0x68, 0x8e, 0x33, 0x9a, 0xdd, 0xa3, 0x4d, 0x41, 0xd6, 0x8b,
		0x60, 0x23, 0x65, 0x3e, 0x8b, 0x22, 0xa0, 0x1e, 0x27, 0x59, 0xb8, 0xe2,
		0x5c, 0x0a, 0x59, 0xe0, 0x5c, 0x7d, 0x89, 0x79, 0x1a, 0xad, 0xe1, 0xfc,
		0x18, 0xef, 0x88, 0xe0, 0x29, 0x89, 0x2e, 0xc3, 0x69, 0x38, 0x89, 0x62,
		0x21, 0x5a, 0xcb, 0x61, 0x4a, 0x01, 0x56, 0x00, 0xfa, 0x82, 0xb0, 0x45,
		0x20, 0xe4, 0x23, 0x23, 0x62, 0x43, 0x88, 0x6c, 0xe8, 0x89, 0xb8, 0xa0,
		0xb9, 0x44, 0xf2, 0x31, 0x07, 0x86, 0x25, 0x79, 0x90, 0xd1, 0x67, 0xbc,
		0xc5, 0x66, 0x35, 0x40, 0xa2, 0x88, 0x0d, 0x1f, 0x02, 0x18, 0x89, 0x79,
		0x42, 0xc2, 0xcf, 0xff, 0x2e, 0x49, 0xf1, 0xa8, 0x19, 0x30, 0x8f, 0xe3,
		0x69, 0x78, 0x01, 0x7f, 0x8a, 0xd2, 0x67, 0x11, 0x2c, 0xe7, 0x91, 0x39,
		0x5b, 0xa1, 0x97, 0x54, 0x32, 0xb2, 0xe4, 0xe9, 0x43, 0x41, 0x52, 0x2e,
		0xc9, 0x3c, 0x32, 0x0b, 0xca, 0x4a, 0x51, 0x6d, 0xa6, 0xb9, 0xe6, 0xcb,
		0x1c, 0x50, 0x36, 0x44, 0x3f, 0xeb, 0x47, 0xf8, 0x42, 0xe8, 0xdd, 0x46,
		0xce, 0xd0, 0xc5, 0x64, 0xf2, 0xbb, 0x1b, 0xbd, 0xf6, 0xf4, 0x4a, 0x7f,
		0xac, 0x78, 0xf2, 0x68, 0xa1, 0x72, 0x9c, 0x24, 0x34, 0xbb, 0x9b, 0xa1,
		0x49, 0xfe, 0x70, 0x53, 0xad, 0xa5, 0xb8, 0xb8, 0xa3, 0x59, 0x6b, 0xa9,
		0x8b, 0x0c, 0x21, 0xad, 0xab, 0x35, 0x4e, 0x29, 0x7b, 0x9c, 0xa1, 0xbf,
		0x12, 0xb6, 0x25, 0x92, 0xc6, 0xb8, 0xde, 0x5d, 0xe1, 0xf8, 0xfe, 0xae,
		0xe0, 0x65, 0x96, 0xcc, 0xd0, 0xc9, 0x64, 0x32, 0x69, 0xb1, 0x40, 0xc3,
		0x35, 0x1e, 0xe7, 0x0c, 0x77, 0xf8, 0x18, 0x33, 0xb2, 0x06, 0x32, 0xd7,
		0x35, 0xe5, 0x0a, 0x3e, 0x5c, 0x95, 0x52, 0xf2, 0xcc, 0x42, 0xc7, 0x65,
		0x21, 0x78, 0x31, 0x43, 0x39, 0xa7, 0xe0, 0x06, 0x45, 0x4d, 0x33, 0xe6,
		0x4c, 0xad, 0x9e, 0x5c, 0x5e, 0x5e, 0xf6, 0x1d, 0x0f, 0x71, 0x2c, 0xe9,
		0x96, 0x34, 0x58, 0x2a, 0xf0, 0xf5, 0x7a, 0x8d, 0x5e, 0xd3, 0x54, 0xb9,
		0x37, 0xce, 0xa4, 0x77, 0xb2, 0xe0, 0x3b, 0xd1, 0x1c, 0x49, 0xa8, 0x50,
		0x6c, 0xcf, 0x50, 0xc6, 0x33, 0x52, 0x93, 0xfd, 0x69, 0x4c, 0xb3, 0x84,
		0x3c, 0x80, 0x76, 0x7a, 0xcf, 0x96, 0xec, 0x79, 0xda, 0xee, 0x1e, 0x67,
		0xf4, 0xb9, 0xf6, 0x62, 0x54, 0xc8, 0xb1, 0xf6, 0x8c, 0xb1, 0x72, 0xce,
		0x36, 0xbf, 0xcd, 0xa6, 0xbf, 0x9e, 0x91, 0xb1, 0xb5, 0x74, 0x78, 0x45,
		0x52, 0x6b, 0x4c, 0x5e, 0x24, 0x10, 0xb2, 0x2b, 0x0e, 0x6a, 0x4c, 0x61,
		0x2f, 0x7f, 0x40, 0x82, 0x33, 0x9a, 0xa0, 0x93, 0xe9, 0x74, 0x7a, 0xe3,
		0xb3, 0x76, 0x31, 0x85, 0xfd, 0xeb, 0x86, 0x19, 0xed, 0x28, 0x82, 0xfe,
		0x44, 0x14, 0xd6, 0x8b, 0x1a, 0x6b, 0xaf, 0x98, 0xb4, 0x63, 0x9c, 0xab,
		0xab, 0xab, 0xb6, 0x9c, 0xe3, 0xc2, 0x30, 0x78, 0xd9, 0x10, 0xb0, 0x56,
		0xa1, 0x99, 0x96, 0x61, 0xc5, 0x78, 0x7c, 0x5f, 0x6f, 0xea, 0xcc, 0x34,
		0x43, 0x53, 0x47, 0x3b, 0x2a, 0x52, 0xc7, 0x98, 0xd1, 0x3b, 0x50, 0x5a,
		0x4c, 0x1a, 0x0f, 0xea, 0x65, 0x09, 0x5b, 0x96, 0xf4, 0xb1, 0x84, 0xc4,
		0xbc, 0xc0, 0x92, 0xf2, 0xac, 0xad, 0xbc, 0x9a, 0x5f, 0x8c, 0x6d, 0x08,
		0xec, 0x36, 0x54, 0x92, 0xb1, 0xc8, 0x71, 0xac, 0x15, 0xbd, 0x83, 0xf4,
		0x53, 0x6f, 0xf1, 0x2d, 0x29, 0xd6, 0x8c, 0xef, 0x66, 0x68, 0x43, 0x93,
		0x84, 0x64, 0x2d, 0xce, 0x9a, 0x4d, 0xc2, 0x18, 0xcd, 0x05, 0x15, 0x1d,
		0x41, 0x1d, 0x09, 0x6b, 0xa6, 0x55, 0x46, 0x2c, 0x38, 0x13, 0xfb, 0x7d,
		0x35, 0xe7, 0x82, 0x1a, 0xe6, 0x21, 0xa9, 0x61, 0x15, 0x0e, 0x9e, 0x9e,
		0xdc, 0x08, 0xef, 0x8b, 0x7a, 0xeb, 0xec, 0xd3, 0x01, 0xf2, 0x8e, 0xaf,
		0x5a, 0x5a, 0x78, 0x05, 0xee, 0x52, 0x4a, 0x4b, 0x4b, 0xf2, 0x7c, 0x86,
		0xae, 0x00, 0x29, 0x32, 0xf1, 0x7e, 0xd5, 0xe0, 0x57, 0xa9, 0xfe, 0x9e,
		0xca, 0x31, 0x24, 0xeb, 0x4c, 0xac, 0x79, 0x01, 0xee, 0xa6, 0x1f, 0x81,
		0x59, 0x32, 0x1a, 0x03, 0xe0, 0x39, 0x52, 0xff, 0xcf, 0x06, 0xa8, 0x87,
		0x05, 0xd9, 0x01, 0x83, 0x3d, 0x3c, 0xac, 0xe9, 0x03, 0x49, 0x5a, 0x0c,
		0x4c, 0x7c, 0x39, 0xa7, 0x0d, 0x1b, 0x1d, 0x6d, 0x0c, 0x52, 0xf2, 0x73,
		0xd3, 0x41, 0x0d, 0x5f, 0xed, 0x57, 0x30, 0x58, 0x1e, 0xc3, 0x92, 0x52,
		0xcc, 0x21, 0x7f, 0x6d, 0x85, 0xd6, 0xd4, 0x0d, 0xd7, 0x07, 0xb5, 0xa8,
		0x83, 0xd1, 0x86, 0xee, 0xc3, 0x40, 0x30, 0x5f, 0x36, 0xc1, 0x6c, 0xd3,
		0xb4, 0x97, 0xbe, 0x6d, 0x90, 0xef, 0x51, 0x02, 0xc0, 0xef, 0x70, 0x91,
		0x40, 0x41, 0x10, 0xd2, 0xea, 0xa2, 0x22, 0x55, 0x45, 0xec, 0xb4, 0x87,
		0xd2, 0x30, 0x46, 0xb0, 0xfd, 0x20, 0x42, 0xe3, 0x35, 0xc7, 0xe1, 0xdb,
		0x82, 0xe7, 0xa5, 0xe4, 0x90, 0x43, 0xd4, 0xda, 0x78, 0xa1, 0x4f, 0x54,
		0xc4, 0x7e, 0x2b, 0x3e, 0xf1, 0xf6, 0x99, 0x3e, 0xa1, 0xc3, 0xe2, 0x65,
		0x0e, 0x51, 0x6b, 0xc0, 0x7c, 0x8e, 0xcb, 0xfc, 0x0b, 0xd8, 0xae, 0x46,
		0x96, 0xf0, 0x5d, 0xf6, 0x22, 0xdf, 0x4a, 0x31, 0xed, 0x33, 0x4a, 0x37,
		0x35, 0x38, 0x56, 0xaf, 0x3d, 0xe3, 0x79, 0x8e, 0xa0, 0x29, 0x85, 0xed,
		0x16, 0xe7, 0xeb, 0xa7, 0xc3, 0xc6, 0x7b, 0x70, 0x29, 0xb9, 0xc7, 0xee,
		0xa5, 0x2b, 0x93, 0xd1, 0x1d, 0x4e, 0x68, 0x29, 0xda, 0xae, 0x66, 0x9b,
		0x89, 0x16, 0x8a, 0x1e, 0x6f, 0x43, 0x5d, 0x77, 0xbb, 0xf4, 0xdc, 0x6d,
		0x83, 0x13, 0x55, 0xc5, 0x68, 0x26, 0x88, 0x04, 0x84, 0xea, 0xef, 0x2d,
		0xd8, 0xe9, 0xe4, 0xfa, 0xfa, 0xda, 0x2f, 0x9b, 0x76, 0x69, 0x9f, 0x2a,
		0x67, 0x2b, 0x02, 0x2a, 0xe8, 0xe9, 0xc6, 0x5a, 0x05, 0xbf, 0x1a, 0x0d,
		0x66, 0xe8, 0x14, 0x9d, 0x7a, 0xbd, 0x89, 0x71, 0xf0, 0xc3, 0x76, 0xab,
		0xb3, 0xd9, 0xd7, 0xb4, 0xdd, 0xf4, 0xe2, 0xea, 0x19, 0xc6, 0x9b, 0x7e,
		0x75, 0xe3, 0xf9, 0xf5, 0xe3, 0x4b, 0x1a, 0xcf, 0xea, 0xb3, 0xd3, 0xc8,
		0xd6, 0xf1, 0xdb, 0xe9, 0xf8, 0x7d, 0x14, 0x55, 0x39, 0xf8, 0x8a, 0x16,
		0xb9, 0xf8, 0x5f, 0x1a, 0xe4, 0xd7, 0xb5, 0x47, 0xad, 0xcc, 0xae, 0x39,
		0x8c, 0xca, 0x7e, 0x7f, 0xa4, 0x35, 0xaa, 0x80, 0x3c, 0xdf, 0x6f, 0xf5,
		0x5f, 0x2d, 0x6c, 0x45, 0xb9, 0xd2, 0x13, 0xb2, 0x38, 0xd6, 0x4b, 0xf4,
		0x50, 0x80, 0xea, 0x12, 0xd2, 0x33, 0xef, 0xee, 0x37, 0xfc, 0x4b, 0xc2,
		0xad, 0x35, 0x5f, 0xb9, 0x47, 0x34, 0x37, 0xbe, 0x45, 0xdf, 0xbd, 0x7b,
		0x77, 0x40, 0x74, 0x90, 0xe7, 0x79, 0x52, 0x57, 0x75, 0xf8, 0xb7, 0x21,
		0xb4, 0xdb, 0xd9, 0x1c, 0x2d, 0x33, 0xcd, 0xd6, 0xfc, 0x08, 0x99, 0x9b,
		0x02, 0xee, 0xcc, 0xc7, 0xff, 0x37, 0xc2, 0xf7, 0xce, 0x71, 0xb5, 0x46,
		0x52, 0xbe, 0xa5, 0x07, 0x2e, 0x27, 0xbe, 0xdc, 0xc0, 0xf7, 0x76, 0xa0,
		0x19, 0x5c, 0xaf, 0xd7, 0x7d, 0x5c, 0x85, 0x34, 0xc5, 0x77, 0x35, 0x73,
		0x4f, 0xfa, 0x7f, 0xcf, 0x36, 0x4d, 0xef, 0x2c, 0xff, 0x83, 0x6d, 0x55,
		0x75, 0x22, 0x21, 0x12, 0x53, 0x67, 0xc0, 0xb5, 0x17, 0x0e, 0xd3, 0xce,
		0xdd, 0x89, 0x77, 0x22, 0xd4, 0x69, 0xc1, 0x1e, 0x74, 0x3b, 0xe4, 0x89,
		0x77, 0x4f, 0xb1, 0xab, 0x54, 0xb0, 0xe2, 0x2c, 0x69, 0x5b, 0xdd, 0x4e,
		0x4c, 0x9d, 0x5a, 0x55, 0x51, 0xd3, 0x1f, 0xe3, 0x98, 0xf1, 0xbe, 0x0b,
		0xa3, 0x56, 0xae, 0x6b, 0xe9, 0xcf, 0x31, 0xb5, 0x73, 0x23, 0x75, 0x64,
		0xad, 0xd8, 0xc3, 0x6e, 0x13, 0xeb, 0xb5, 0xd3, 0xbb, 0xf7, 0x3b, 0x97,
		0xfe, 0x7d, 0xc8, 0xd0, 0xc5, 0x86, 0x57, 0xdd, 0x2e, 0x5b, 0xc2, 0xcf,
		0x23, 0x7b, 0xfd, 0x38, 0x57, 0xf7, 0x8a, 0xd5, 0xc5, 0x65, 0x42, 0xb7,
		0x28, 0x66, 0x58, 0x88, 0x45, 0x60, 0xee, 0x52, 0x82, 0x65, 0x85, 0x6d,
		0x5e, 0xb2, 0xe5, 0x3c, 0x2a, 0x59, 0x05, 0x18, 0x01, 0xe4, 0xf2, 0x55,
		0xe7, 0x50, 0x1d, 0xe9, 0xcd, 0x31, 0x67, 0xd3, 0x0c, 0x88, 0x76, 0xcb,
		0x23, 0x67, 0x06, 0x2f, 0x75, 0xd3, 0x07, 0x1f, 0xad, 0xa9, 0x34, 0x40,
		0x09, 0x96, 0x78, 0x11, 0x08, 0x42, 0xee, 0x6f, 0xd5, 0xce, 0xad, 0x5e,
		0x5d, 0xce, 0x69, 0x7d, 0x78, 0x8d, 0xd1, 0x1a, 0x6b, 0xd8, 0x71, 0x7d,
		0x52, 0xdd, 0xc9, 0xd2, 0x65, 0xc5, 0xe7, 0x61, 0x82, 0xee, 0xd0, 0xda,
		0xa2, 0x57, 0x6d, 0xec, 0x23, 0x59, 0x81, 0xb8, 0x14, 0x2b, 0x7a, 0xb5,
		0x12, 0x5c, 0x2e, 0x5c, 0x1e, 0x54, 0x3a, 0x1c, 0xd0, 0x47, 0x53, 0x12,
		0x0d, 0x87, 0x96, 0xa9, 0x7a, 0x3d, 0x58, 0x7e, 0xfc, 0xe1, 0x4f, 0x1f,
		0xf7, 0x08, 0xa8, 0x2a, 0x0b, 0x00, 0x7d, 0xfa, 0xf0, 0xfd, 0x1e, 0x20,
		0x95, 0x8a, 0x83, 0xe5, 0xdf, 0xbe, 0xfb, 0xcb, 0x87, 0xa3, 0x55, 0x65,
		0x35, 0xec, 0x9b, 0xa5, 0xab, 0x9e, 0x3e, 0x63, 0x58, 0xad, 0x1c, 0x61,
		0x8e, 0x3e, 0x4b, 0xf4, 0x18, 0xa1, 0xab, 0xff, 0x23, 0x48, 0xa8, 0xe8,
		0xae, 0xf1, 0xe7, 0xb8, 0x04, 0x57, 0x47, 0x34, 0x59, 0x04, 0x95, 0xb6,
		0x3b, 0x44, 0x0c, 0x48, 0x97, 0xc4, 0xa0, 0x6d, 0xcd, 0xe8, 0x7b, 0x94,
		0xb7, 0xeb, 0xf1, 0xd8, 0x19, 0x95, 0x6b, 0xb6, 0xcc, 0xd2, 0xad, 0x5e,
		0xea, 0x30, 0xe4, 0xc2, 0x3f, 0x4f, 0x72, 0x18, 0xed, 0xed, 0x90, 0xef,
		0x91, 0x82, 0x85, 0x21, 0x42, 0x7a, 0x6b, 0x8f, 0xf4, 0x43, 0x29, 0x41,
		0x27, 0xd7, 0xde, 0x7c, 0xa0, 0x0b, 0x89, 0xa2, 0x07, 0xa5, 0x44, 0xbf,
		0x72, 0x51, 0xef, 0x67, 0x06, 0xf5, 0x59, 0x95, 0x84, 0x01, 0x85, 0xea,
		0x98, 0x08, 0xf6, 0xe9, 0x20, 0x21, 0xe6, 0x15, 0x8d, 0xb1, 0xae, 0x07,
		0x87, 0xab, 0x97, 0x4f, 0x27, 0x41, 0x8b, 0x6d, 0x53, 0x13, 0x82, 0xe5,
		0x7b, 0xf5, 0x31, 0x8f, 0xf0, 0x21, 0xb1, 0xf7, 0xbe, 0x56, 0xaa, 0x0f,
		0x6f, 0x71, 0xa1, 0x9d, 0x8f, 0x14, 0x1f, 0x25, 0x96, 0xa5, 0x40, 0x0b,
		0x94, 0x95, 0x8c, 0xdd, 0xbc, 0xaa, 0xcb, 0x43, 0x99, 0x19, 0x3b, 0xc5,
		0x3c, 0x4d, 0x71, 0x96, 0x8c, 0xd4, 0x6b, 0xb5, 0x33, 0x5b, 0x98, 0x10,
		0x7a, 0x13, 0xde, 0x11, 0xf9, 0xf7, 0x8f, 0x1f, 0xbe, 0x1b, 0x05, 0x51,
		0x05, 0x13, 0x05, 0xe8, 0x1b, 0xfd, 0xfa, 0xed, 0x1c, 0xfd, 0xfc, 0x74,
		0x6e, 0x51, 0x8c, 0x0a, 0x22, 0x72, 0xf7, 0xa8, 0xee, 0xd6, 0xa1, 0xcb,
		0x22, 0x21, 0xe3, 0x77, 0x66, 0xf7, 0xc6, 0x6e, 0x3e, 0xd9, 0xe7, 0xa7,
		0x0e, 0x2f, 0x8a, 0xe1, 0x51, 0x8e, 0xe5, 0x66, 0x88, 0x11, 0x1d, 0x4e,
		0x40, 0x1d, 0xad, 0xa9, 0x7a, 0x3d, 0xa1, 0x40, 0xd1, 0x7e, 0x4e, 0xe8,
		0x1a, 0xe9, 0x55, 0xe8, 0xfe, 0xe3, 0x98, 0x08, 0xd1, 0xde, 0x45, 0xe0,
		0xa6, 0xe0, 0x9a, 0xc4, 0xe8, 0x68, 0xe4, 0xb0, 0x59, 0x37, 0x25, 0x87,
		0x58, 0x36, 0x05, 0xac, 0xcb, 0xf4, 0x28, 0x28, 0x59, 0x70, 0x16, 0xaa,
		0x77, 0x6e, 0xa3, 0x20, 0x38, 0xb3, 0x7a, 0x47, 0x86, 0xe9, 0x85, 0xf9,
		0xf8, 0xe5, 0x17, 0x14, 0x04, 0xce, 0x9e, 0x2b, 0x6b, 0x55, 0x1a, 0x95,
		0xb4, 0x0a, 0xf6, 0x28, 0x69, 0xa3, 0x08, 0xfd, 0x31, 0x49, 0xa0, 0x97,
		0x93, 0x65, 0x61, 0x72, 0x28, 0x22, 0x8c, 0xa4, 0xd0, 0x18, 0x28, 0x45,
		0xec, 0x08, 0x82, 0x9e, 0x82, 0x14, 0x08, 0x03, 0xae, 0x42, 0x2d, 0x26,
		0xb4, 0x20, 0xb1, 0xe4, 0xc5, 0xa3, 0xa7, 0x31, 0x4d, 0xea, 0xf5, 0x02,
		0x98, 0xf3, 0xf5, 0x45, 0x25, 0x49, 0x85, 0x7e, 0x32, 0x22, 0x84, 0xd0,
		0xc0, 0x50, 0x09, 0xec, 0x82, 0x8c, 0x66, 0x33, 0xcc, 0x79, 0xde, 0xd6,
		0xa4, 0xe9, 0x68, 0x6e, 0x2b, 0xc1, 0x0d, 0xd0, 0x67, 0x4e, 0x33, 0x73,
		0xea, 0x55, 0x0b, 0x12, 0x14, 0x37, 0x67, 0x10, 0xff, 0x55, 0xac, 0x9c,
		0x9e, 0x9c, 0x9a, 0xdc, 0x71, 0xaa, 0x7c, 0xaf, 0x41, 0xf3, 0x0d, 0x0a,
		0x4e, 0x9b, 0x1c, 0x72, 0x6a, 0x72, 0x08, 0x23, 0x5b, 0xc2, 0x20, 0x85,
		0x9c, 0xea, 0x14, 0xa2, 0xa1, 0x55, 0x38, 0xcd, 0x23, 0x40, 0x08, 0xb6,
		0xc0, 0x79, 0x4e, 0xb2, 0xe4, 0x13, 0x1f, 0x05, 0xcd, 0x2b, 0x9c, 0xc0,
		0x33, 0xb9, 0xf3, 0x05, 0x12, 0x3e, 0x1a, 0x51, 0x98, 0xa8, 0x51, 0x57,
		0xcf, 0x48, 0x3b, 0x20, 0x08, 0xa3, 0xb6, 0x7e, 0xa4, 0xff, 0x6a, 0x4b,
		0xab, 0x42, 0x04, 0xf6, 0x14, 0x48, 0xa8, 0xfe, 0xa9, 0xef, 0x9e, 0x98,
		0xfb, 0x95, 0x6c, 0x51, 0xd4, 0xb2, 0xda, 0xc8, 0x6b, 0x13, 0x7a, 0x6a,
		0x23, 0x05, 0xeb, 0x7f, 0x4b, 0xe5, 0x46, 0x1b, 0x58, 0xf3, 0xc7, 0xd5,
		0x53, 0x9f, 0x8d, 0x91, 0x56, 0xea, 0xad, 0x4a, 0x21, 0x35, 0xa3, 0x16,
		0x0c, 0xfd, 0x01, 0x05, 0xf0, 0x25, 0x40, 0x33, 0x14, 0xa8, 0x9d, 0xa0,
		0x4d, 0x33, 0x16, 0xe2, 0x56, 0x6b, 0xbd, 0xf7, 0xa0, 0x2e, 0x94, 0x0c,
		0x7c, 0xcc, 0x1c, 0x87, 0x6f, 0x94, 0xa5, 0xc1, 0x71, 0x36, 0xd6, 0xef,
		0x22, 0x8d, 0xa1, 0x1b, 0xee, 0x94, 0xa1, 0x1d, 0x0f, 0xd0, 0x7a, 0xe9,
		0x1a, 0x5f, 0x6d, 0x35, 0x8c, 0x99, 0x7d, 0xe5, 0x02, 0x6a, 0xbd, 0x65,
		0x06, 0xb5, 0xf7, 0x5c, 0x9f, 0xb0, 0x4f, 0x7b, 0xd2, 0x40, 0x3b, 0x8d,
		0x0c, 0x64, 0x2f, 0xa1, 0xb7, 0x83, 0x83, 0xd9, 0xd3, 0xcb, 0xdb, 0x0a,
		0xa0, 0xa5, 0x41, 0x9b, 0xd3, 0x8a, 0x32, 0xcb, 0x60, 0xc4, 0xf1, 0xdd,
		0xe7, 0x8d, 0x95, 0x46, 0x65, 0x20, 0x9a, 0x10, 0x3f, 0x1c, 0x15, 0x80,
		0x6d, 0xa2, 0xcf, 0x42, 0xb1, 0xe1, 0xbb, 0x91, 0x1f, 0x88, 0x96, 0x88,
		0x4c, 0x93, 0xd5, 0x2d, 0xce, 0xe9, 0xed, 0x3d, 0x79, 0x1c, 0x72, 0x57,
		0x85, 0x50, 0x37, 0x79, 0x0d, 0x32, 0x77, 0xfb, 0xa9, 0x57, 0x9f, 0x08,
		0x72, 0x93, 0x33, 0x0b, 0xf5, 0xf0, 0xde, 0x87, 0xcb, 0xe3, 0xbd, 0x4f,
		0xbc, 0x2a, 0x27, 0xff, 0x17, 0x66, 0xac, 0xe6, 0xb1, 0x7f, 0xa8, 0xb2,
		0x3c, 0x52, 0x5e, 0xe7, 0x8a, 0xaa, 0x4a, 0xaa, 0xee, 0x26, 0x7e, 0x28,
		0x18, 0x98, 0xc5, 0xfe, 0x76, 0x43, 0xaf, 0x69, 0x35, 0x85, 0xbc, 0xb8,
		0x8b, 0x64, 0x94, 0x47, 0x3b, 0x98, 0x86, 0xa3, 0xda, 0x91, 0x21, 0x21,
		0x0a, 0x18, 0xcc, 0x74, 0xe2, 0x72, 0x73, 0x3d, 0x08, 0x62, 0x46, 0x43,
		0x68, 0x4d, 0x94, 0x2b, 0x4a, 0x59, 0x8c, 0x02, 0x68, 0x51, 0xc0, 0x41,
		0x6a, 0x32, 0x8e, 0x08, 0x0d, 0xb8, 0x19, 0x57, 0xe1, 0x84, 0xaa, 0xfd,
		0x9a, 0x4b, 0xa0, 0x4b, 0x61, 0xa6, 0xc3, 0xec, 0x56, 0x6f, 0xf5, 0x9f,
		0x72, 0x5b, 0x93, 0xd6, 0xd9, 0x2d, 0x29, 0xd4, 0xef, 0x77, 0xce, 0x3c,
		0xde, 0xf6, 0x28, 0xd9, 0x62, 0xed, 0xd8, 0xa8, 0xab, 0x52, 0x70, 0x7f,
		0xa3, 0x4e, 0x9a, 0x9c, 0xa3, 0x94, 0x24, 0xd4, 0x04, 0xb6, 0xaf, 0x58,
		0x28, 0x48, 0x38, 0x55, 0xde, 0xee, 0xba, 0x43, 0xe5, 0x75, 0xb3, 0x56,
		0x44, 0xb4, 0x1c, 0xb2, 0xb1, 0x6a, 0x7f, 0x15, 0xad, 0x6d, 0x04, 0xe0,
		0x21, 0x64, 0x46, 0xcd, 0x74, 0x65, 0xa7, 0xb7, 0xda, 0x40, 0x0d, 0x47,
		0x36, 0xcf, 0x2a, 0x46, 0x0d, 0x3b, 0x7b, 0xa3, 0xb4, 0xe5, 0x2b, 0xcf,
		0x68, 0x72, 0x04, 0xc1, 0x45, 0xbc, 0x31, 0xc7, 0xfc, 0xa6, 0xeb, 0x05,
		0x9a, 0x38, 0x77, 0xa0, 0xf5, 0x8f, 0x87, 0x66, 0x3a, 0x59, 0xbe, 0x48,
		0x41, 0x86, 0xd5, 0x28, 0x2d, 0x99, 0xa4, 0xc1, 0x51, 0x4a, 0x69, 0xb2,
		0x06, 0x97, 0xe0, 0x8e, 0xf0, 0x08, 0x67, 0x05, 0x5a, 0xa2, 0x49, 0x5f,
		0x23, 0x51, 0x25, 0xb7, 0xb0, 0x02, 0xfb, 0x71, 0xe2, 0x15, 0xd2, 0xc6,
		0x77, 0x00, 0x38, 0x54, 0x76, 0xd1, 0x0f, 0x8e, 0x17, 0x1d, 0x1d, 0xe2,
		0x6f, 0x46, 0x96, 0x6b, 0x97, 0x13, 0xbf, 0xfd, 0x6b, 0xb9, 0xb8, 0x9e,
		0x6d, 0xcf, 0x42, 0x38, 0x12, 0xc4, 0x8c, 0xc2, 0xe8, 0xe9, 0x88, 0xee,
		0xf7, 0xbb, 0xa6, 0x89, 0x36, 0xe3, 0x70, 0x8b, 0x2b, 0x75, 0x1b, 0xc1,
		0x77, 0xd0, 0x09, 0xc7, 0xfa, 0xf6, 0x04, 0x64, 0x65, 0x1c, 0x27, 0x23,
		0xcf, 0x59, 0x5a, 0x74, 0xab, 0xe9, 0xf0, 0x59, 0x94, 0xdf, 0x8c, 0xe4,
		0x86, 0x8a, 0x3a, 0x8b, 0xa8, 0xc0, 0x0e, 0xce, 0x5a, 0x7c, 0xd4, 0x00,
		0xea, 0xc7, 0x62, 0x5b, 0xf2, 0x5e, 0x15, 0xca, 0x51, 0x60, 0x7e, 0xf2,
		0xe4, 0x35, 0x61, 0xca, 0x86, 0x35, 0xf4, 0x06, 0x8b, 0x0a, 0x54, 0xb7,
		0xde, 0x67, 0xbe, 0x15, 0x89, 0xca, 0x84, 0x35, 0xf0, 0x9a, 0x2a, 0x15,
		0xd0, 0x4e, 0x4f, 0xa7, 0x10, 0x12, 0xe6, 0xe0, 0xaa, 0x7e, 0xde, 0xd5,
		0x45, 0xa7, 0x10, 0xb6, 0x39, 0xb4, 0xa0, 0x21, 0x4e, 0x12, 0x67, 0x4d,
		0x0f, 0xca, 0x7b, 0xca, 0x4c, 0x6f, 0x71, 0xe9, 0x47, 0x6f, 0x50, 0x79,
		0xf8, 0x35, 0xcd, 0x63, 0xaa, 0x58, 0xc7, 0x7a, 0x55, 0x1d, 0x3c, 0xc2,
		0x76, 0x6e, 0x26, 0x68, 0xc5, 0xb5, 0x4e, 0x0b, 0x7b, 0x1c, 0xc4, 0x1d,
		0x1b, 0x8f, 0xa2, 0xe4, 0x26, 0xec, 0x6e, 0xc5, 0x1c, 0x68, 0x07, 0x1a,
		0xf2, 0x6d, 0xea, 0x66, 0xce, 0xaf, 0x08, 0x4b, 0x5e, 0xc6, 0x1b, 0x68,
		0x6c, 0x0a, 0x89, 0x52, 0x0e, 0x8a, 0xd4, 0x37, 0x05, 0x7b, 0xf8, 0xa8,
		0xbc, 0xd4, 0xea, 0xba, 0xf1, 0xc0, 0x41, 0x69, 0xbb, 0xf4, 0xa0, 0x65,
		0x43, 0xfa, 0x21, 0xc6, 0x59, 0x0c, 0x2e, 0xa8, 0x29, 0x97, 0xf9, 0x61,
		0xba, 0x03, 0xce, 0x3f, 0x44, 0x5a, 0xdd, 0x59, 0x7a, 0x0a, 0x0e, 0xf0,
		0x30, 0x15, 0x93, 0xbd, 0xf5, 0x90, 0xd3, 0x17, 0x8e, 0x37, 0x1e, 0x68,
		0xd5, 0x7d, 0x77, 0x41, 0x75, 0x23, 0xdc, 0x13, 0x94, 0xe6, 0xc0, 0xa2,
		0x6a, 0xcb, 0xfd, 0xd8, 0x69, 0x86, 0xe7, 0xfe, 0xa4, 0xd8, 0x1b, 0x11,
		0xee, 0xf8, 0x7a, 0x30, 0x97, 0x56, 0x4f, 0xee, 0xcf, 0x50, 0xe7, 0x91,
		0xb9, 0xd7, 0x9d, 0x47, 0xfa, 0x17, 0xc2, 0xff, 0x09, 0x00, 0x00, 0xff,
		0xff, 0x5a, 0x72, 0x3c, 0xaa, 0x38, 0x2c, 0x00, 0x00,
	},
		"static/index.html",
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
	"static/index.html": static_index_html,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
	}},
}}
