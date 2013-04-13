// +build !devel

package main

import (
	"bytes"
	"compress/gzip"
	"io"
)

// status_html returns raw, uncompressed file data.
func status_html() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0xb4,0x57,
0x5d,0x6f,0xda,0x3c,0x14,0xbe,0x86,0x5f,0x61,0x45,0xaa,0xf4,
0xbe,0xd2,0x5b,0x87,0xf4,0x2d,0xd1,0xd8,0x42,0x6e,0xd6,0xaa,
0xbb,0xa1,0x6a,0x87,0xb6,0x8b,0xdd,0x99,0xf8,0x40,0xa2,0x05,
0x1b,0xd9,0xa6,0x94,0x45,0xfc,0xf7,0xd9,0xce,0x07,0x5f,0x81,
0x51,0x57,0xbb,0x21,0xce,0x63,0x3f,0x9c,0x8f,0xe7,0x9c,0xe4,
0x24,0x4a,0xd5,0x3c,0x8f,0xa3,0x14,0x08,0x8d,0x23,0x95,0xa9,
0x1c,0xe2,0x07,0xe0,0x77,0x8f,0x63,0x54,0x14,0x08,0x7f,0x07,
0x21,0x33,0xce,0xd0,0x66,0x13,0xf9,0xe5,0x66,0x37,0xca,0x33,
0xf6,0x13,0xa5,0x02,0xa6,0x43,0xcf,0xf7,0x19,0x28,0xca,0x08,
0x9e,0x70,0xae,0xa4,0x12,0x64,0x91,0x50,0x86,0x13,0x3e,0xf7,
0xd5,0x2a,0x53,0x0a,0xc4,0x75,0xb3,0xe1,0xdf,0xe0,0xff,0x71,
0xe0,0x27,0x52,0xfa,0x0d,0x76,0xad,0x4f,0x4e,0x32,0x06,0x14,
0xcf,0x33,0x4d,0x93,0xd2,0x43,0x02,0xf2,0xa1,0x27,0xd5,0x3a,
0x07,0x99,0x02,0x28,0xef,0x52,0x7b,0x16,0x58,0x11,0x95,0xa4,
0xb5,0x21,0x10,0xcb,0x1c,0x08,0xdb,0x5a,0x3b,0x6b,0xc4,0xde,
0xc5,0xdd,0x8e,0xa2,0xf8,0x17,0x67,0xc0,0xc8,0x1c,0xfe,0xd3,
0x6b,0x93,0x16,0x10,0xa8,0x40,0x53,0xce,0xd4,0xf5,0x0a,0xb2,
0x59,0xaa,0x3e,0xa2,0x09,0xcf,0xe9,0x27,0xb4,0xe9,0x46,0x7e,
0x45,0x8b,0x26,0x9c,0xae,0xe3,0x6e,0x37,0xa2,0xd9,0x0b,0x4a,
0x72,0x22,0xe5,0xd0,0x4b,0x34,0x83,0xe8,0xe0,0x84,0x67,0x36,
0xd2,0x20,0x7e,0xc8,0xf9,0x84,0xe4,0x91,0xaf,0x97,0xfb,0x27,
0x05,0x5f,0x99,0x33,0x9d,0x5d,0x4c,0x2e,0x08,0xfb,0x60,0xd1,
0x4e,0xa4,0xc8,0x24,0x87,0x7a,0xa3,0xbc,0xb1,0xbf,0x3a,0xb9,
0x42,0x7b,0x07,0xb4,0xba,0xd5,0x16,0x29,0x30,0x09,0x54,0xf3,
0x3a,0x86,0x27,0xec,0x55,0x2f,0x68,0xcd,0xce,0xd8,0x94,0xa3,
0x32,0x28,0xaf,0x72,0x08,0x49,0x45,0x94,0xd4,0xf2,0xd2,0xe6,
0x74,0x1c,0x20,0x9d,0xab,0x7d,0xa8,0x7f,0x0c,0x05,0x2d,0xd8,
0x48,0xa7,0xbc,0x41,0xf4,0x42,0xd8,0x10,0xf6,0x7c,0x89,0x9f,
0x97,0x20,0x32,0x38,0x30,0x59,0x14,0x0b,0x91,0x31,0x35,0x45,
0xde,0x15,0xbe,0x99,0x7a,0x08,0x97,0xde,0xe1,0xea,0x30,0xfe,
0x4a,0x14,0x04,0xb6,0x0c,0xdf,0x48,0xea,0xbb,0x90,0x02,0x27,
0x96,0x09,0x7e,0x87,0x57,0x86,0x6f,0xaf,0x46,0x9f,0xbf,0x2b,
0xe6,0x97,0x4c,0x2a,0x3e,0x13,0x64,0x7e,0x90,0xd7,0xd1,0x59,
0x85,0x2c,0x32,0xe8,0x5d,0x1d,0x00,0x83,0x23,0x00,0x1f,0x42,
0x23,0xf2,0xba,0x0f,0x8c,0x15,0xbd,0x83,0x97,0x3f,0x69,0x1f,
0xf4,0x4c,0xcd,0x2c,0x15,0x48,0x14,0xc9,0x39,0xc9,0xf3,0xf8,
0x9f,0xa2,0xa8,0x73,0xd9,0xc4,0x10,0xf4,0xf0,0x67,0xbe,0x64,
0x6a,0xb3,0xf9,0x57,0xf7,0x98,0x3d,0x76,0x52,0x0d,0xea,0x21,
0x84,0xda,0xfe,0x41,0x47,0x7e,0xa9,0x88,0x7b,0xb4,0x7d,0x11,
0x2f,0xe6,0x3d,0x25,0x6a,0xd0,0x43,0xae,0xcc,0x81,0x3b,0xf3,
0x2c,0xf5,0x74,0x76,0xc8,0xab,0x8b,0xc5,0x52,0x65,0x74,0x54,
0xe5,0x47,0x42,0x87,0x17,0x09,0x1d,0xbe,0x5b,0xe8,0xd0,0x4d,
0xe8,0xd0,0x51,0xe8,0xd0,0x59,0xe8,0xd0,0x59,0xe8,0xd0,0x5d,
0xe8,0xd0,0x4d,0xe8,0xf0,0x0d,0x42,0x07,0x88,0x92,0xf5,0xd9,
0x66,0xbe,0xbd,0x7d,0x7f,0x3b,0x9b,0xff,0x70,0x6a,0x68,0x4b,
0x74,0x6a,0x69,0xc3,0x74,0x6d,0xea,0x9a,0xeb,0xd2,0xd6,0x0d,
0xd7,0xa5,0xb1,0x6d,0xb8,0x2e,0xad,0x6d,0x88,0x67,0x34,0xdf,
0x79,0x87,0x45,0xbe,0x9e,0x53,0xda,0xe6,0x95,0xdb,0x77,0xcd,
0x2b,0x75,0x51,0x9d,0x7a,0xc1,0x7d,0x5b,0xa8,0x6c,0x0e,0xd6,
0xb3,0x32,0x22,0x5c,0x22,0xf8,0x8e,0xac,0xc7,0x4a,0x47,0x37,
0xab,0xfd,0xee,0xb6,0xbe,0x7a,0x6b,0xb7,0x9b,0xab,0x9e,0xc3,
0x7e,0xe8,0x51,0x4f,0x96,0x13,0x99,0x93,0xd7,0xd6,0xe7,0x48,
0xa5,0xa6,0x9c,0xd3,0x72,0x15,0x54,0x8f,0xbd,0x2d,0xd2,0xaf,
0x1f,0x84,0x3b,0x87,0x5a,0x30,0x53,0xa4,0xe8,0xf9,0x69,0x5c,
0x21,0x65,0xe2,0x8b,0x42,0x10,0x36,0x03,0xad,0xb4,0x75,0x55,
0x4b,0x63,0x22,0x51,0x62,0x37,0x43,0x55,0xf6,0x74,0xd6,0x78,
0x6e,0x64,0x18,0xf6,0xeb,0xdd,0x7a,0x92,0xf5,0x4c,0xb6,0x1e,
0xf5,0xa2,0x51,0xb6,0xd6,0xb5,0x4a,0x7a,0xcb,0x54,0xd6,0x5a,
0x35,0x23,0xd0,0x89,0x4e,0xe4,0xa9,0xa1,0xec,0x62,0x4e,0xdf,
0x81,0x13,0xb8,0x90,0xf6,0x3a,0xbf,0x25,0xe8,0x7b,0xf3,0xa1,
0xf3,0x96,0xc8,0xef,0x29,0x93,0x2e,0xd1,0x1f,0xf2,0x2e,0x0d,
0xe6,0xc8,0x9e,0x2b,0xb1,0x2d,0x13,0x45,0x01,0x8c,0x9a,0x92,
0xda,0xb6,0xc9,0xb6,0x4b,0xca,0x4f,0x1a,0xdd,0x1c,0xe6,0xfb,
0xb0,0xfb,0x3b,0x00,0x00,0xff,0xff,0xe0,0x5f,0x4d,0x4f,0x27,
0x0e,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}