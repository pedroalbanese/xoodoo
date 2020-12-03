package xoodoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkXoodooPermutation(b *testing.B) {
	newXD, _ := NewXooDoo(12, [48]byte{})
	for n := 0; n < b.N; n++ {
		newXD.Permutation()
	}
}

var perumtationTestTable = []struct {
	inBytes  [48]byte
	rounds   int
	outBytes [48]byte
}{
	{
		inBytes:  [48]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		rounds:   12,
		outBytes: [48]byte{0x8D, 0xD8, 0xD5, 0x89, 0xBF, 0xFC, 0x63, 0xA9, 0x19, 0x2D, 0x23, 0x1B, 0x14, 0xA0, 0xA5, 0xFF, 0x06, 0x81, 0xB1, 0x36, 0xFE, 0xC1, 0xC7, 0xAF, 0xBE, 0x7C, 0xE5, 0xAE, 0xBD, 0x40, 0x75, 0xA7, 0x70, 0xE8, 0x86, 0x2E, 0xC9, 0xB7, 0xF5, 0xFE, 0xF2, 0xAD, 0x4F, 0x8B, 0x62, 0x40, 0x4F, 0x5E},
	},
	{
		inBytes:  [48]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F},
		rounds:   12,
		outBytes: [48]byte{0x76, 0x33, 0xAE, 0xB5, 0x5D, 0xCC, 0xBF, 0x60, 0xD4, 0xA6, 0xDF, 0xD7, 0x50, 0x6D, 0x06, 0xBF, 0xB2, 0xAC, 0x97, 0xAE, 0x97, 0x0D, 0x8A, 0xD3, 0x13, 0x85, 0x11, 0x7B, 0xB7, 0x75, 0xA7, 0x41, 0xB3, 0xB1, 0x54, 0x0B, 0xB5, 0x3B, 0xE9, 0x6F, 0x3B, 0x2B, 0x8F, 0xAF, 0xA6, 0x76, 0xA3, 0xB6},
	},
	{
		inBytes:  [48]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		rounds:   12,
		outBytes: [48]byte{0x64, 0x4E, 0xD8, 0xAC, 0xE6, 0x92, 0x61, 0x9D, 0x30, 0xEB, 0x4F, 0xA0, 0x81, 0xF5, 0x65, 0x54, 0xD0, 0xE9, 0xE2, 0xA5, 0xE1, 0x8D, 0x7B, 0x67, 0x36, 0xAE, 0xA7, 0x32, 0x54, 0xBC, 0x0A, 0xFE, 0x1F, 0x48, 0x47, 0x41, 0x10, 0x9D, 0xBE, 0xA6, 0xAF, 0x0D, 0xC7, 0x83, 0x52, 0x24, 0x2E, 0x83},
	},
	{
		inBytes:  [48]byte{0x71, 0xFA, 0x9C, 0xC2, 0x60, 0xB0, 0x59, 0xF9, 0x42, 0x45, 0xA4, 0x3E, 0x61, 0xC5, 0x39, 0x0D, 0xC4, 0x6A, 0x11, 0x40, 0xC0, 0x65, 0x7B, 0x43, 0x5C, 0x8C, 0x00, 0x06, 0x4B, 0x51, 0x61, 0xF4, 0x9C, 0x27, 0x11, 0xBE, 0x86, 0xD9, 0x32, 0x1F, 0xD1, 0x53, 0x29, 0x08, 0x15, 0x90, 0xDB, 0x18},
		rounds:   12,
		outBytes: [48]byte{0xA4, 0xCA, 0x45, 0x89, 0x33, 0x6C, 0x92, 0x2A, 0x45, 0xF6, 0xC9, 0x99, 0x81, 0x16, 0xF4, 0x52, 0x26, 0xC5, 0x66, 0xF2, 0x9E, 0xB2, 0xE3, 0xD9, 0x98, 0xDB, 0xCE, 0xCB, 0x86, 0x16, 0x81, 0xF5, 0x54, 0xFC, 0xE5, 0x44, 0x4A, 0x62, 0x65, 0xF5, 0x8B, 0x7A, 0x6E, 0xE9, 0xE1, 0x4E, 0xD0, 0x06},
	},
	{
		inBytes:  [48]byte{0xB2, 0x93, 0xDA, 0x74, 0x20, 0x29, 0xBB, 0x4F, 0x6F, 0x84, 0x6E, 0xFB, 0x3B, 0x82, 0xA0, 0x1D, 0xEB, 0xEF, 0x19, 0x90, 0x62, 0x3B, 0xB1, 0x4B, 0xA2, 0xC8, 0x1E, 0xC7, 0xFE, 0x70, 0xA4, 0x8C, 0x05, 0x81, 0x83, 0x64, 0xF8, 0x54, 0x83, 0x5D, 0x0F, 0x7C, 0x5B, 0x98, 0xEA, 0xFE, 0x02, 0xDF},
		rounds:   12,
		outBytes: [48]byte{0x34, 0xF6, 0x32, 0xA4, 0x00, 0xF6, 0x7A, 0x61, 0xE4, 0xF6, 0x4B, 0x54, 0x87, 0x70, 0xCA, 0x95, 0xAE, 0x1A, 0x4D, 0x33, 0x00, 0x11, 0x78, 0xC3, 0xC3, 0x76, 0x02, 0x40, 0x05, 0xFF, 0xB5, 0xBE, 0x06, 0x27, 0x30, 0x40, 0x03, 0x00, 0x09, 0x1B, 0x7E, 0x13, 0x77, 0xC1, 0x09, 0x35, 0x55, 0x4A},
	},
	{
		inBytes:  [48]byte{0xF7, 0x40, 0x86, 0x81, 0x27, 0x57, 0x0F, 0xB2, 0xBB, 0xCF, 0xB4, 0xBE, 0x50, 0x22, 0x9F, 0x27, 0x32, 0x05, 0xB7, 0x42, 0x33, 0x72, 0x3C, 0x57, 0xED, 0xE7, 0xF5, 0x91, 0x24, 0xE3, 0x75, 0x72, 0xE1, 0xE0, 0x51, 0xB8, 0x7C, 0x9D, 0x80, 0x3C, 0xC5, 0x23, 0xED, 0x05, 0x4D, 0x28, 0x54, 0x8A},
		rounds:   12,
		outBytes: [48]byte{0x01, 0x2C, 0x48, 0x7E, 0xDD, 0xAD, 0x45, 0x34, 0x84, 0x2E, 0x31, 0x90, 0x67, 0xCD, 0x04, 0xCF, 0xC7, 0xFE, 0x0B, 0x99, 0xE8, 0xB5, 0xC7, 0x69, 0xB9, 0x5C, 0x4A, 0x6E, 0x3F, 0xB1, 0xBF, 0x1D, 0xAD, 0x8A, 0x2D, 0x09, 0x62, 0x1E, 0x7C, 0xE3, 0x5C, 0x42, 0xE8, 0xF0, 0x5B, 0x94, 0xA0, 0xE4},
	},
	{
		inBytes:  [48]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		rounds:   6,
		outBytes: [48]byte{0xA3, 0xCE, 0xC9, 0x28, 0x60, 0x4F, 0x20, 0xAD, 0xD6, 0xD0, 0xC3, 0x2E, 0xC5, 0xC7, 0x50, 0xF0, 0x25, 0x12, 0xDC, 0x08, 0x04, 0x23, 0x99, 0x61, 0x2D, 0x40, 0x0D, 0x9E, 0x9B, 0x9B, 0xD5, 0x42, 0xFC, 0x14, 0x61, 0x1E, 0x97, 0xB6, 0x6E, 0x18, 0x7F, 0xBC, 0xDB, 0x35, 0x4E, 0x10, 0xF9, 0xA1},
	},
	{
		inBytes:  [48]byte{0x76, 0x33, 0xAE, 0xB5, 0x5D, 0xCC, 0xBF, 0x60, 0xD4, 0xA6, 0xDF, 0xD7, 0x50, 0x6D, 0x06, 0xBF, 0xB2, 0xAC, 0x97, 0xAE, 0x97, 0x0D, 0x8A, 0xD3, 0x13, 0x85, 0x11, 0x7B, 0xB7, 0x75, 0xA7, 0x41, 0xB3, 0xB1, 0x54, 0x0B, 0xB5, 0x3B, 0xE9, 0x6F, 0x3B, 0x2B, 0x8F, 0xAF, 0xA6, 0x76, 0xA3, 0xB6},
		rounds:   6,
		outBytes: [48]byte{0x9F, 0xED, 0x7B, 0xFE, 0xC1, 0x4C, 0x4D, 0x7F, 0x83, 0x2A, 0x15, 0x97, 0x69, 0x22, 0xE1, 0xC0, 0x96, 0x71, 0x89, 0x89, 0x3F, 0x7B, 0x47, 0xE8, 0xC3, 0x84, 0xA4, 0xB9, 0xB9, 0xFF, 0x42, 0x84, 0xDB, 0x79, 0x08, 0x03, 0x90, 0x55, 0x32, 0x08, 0xD9, 0xEB, 0x6C, 0x8D, 0x97, 0xDB, 0x2F, 0xC9},
	},
	{
		inBytes:  [48]byte{0x64, 0x4E, 0xD8, 0xAC, 0xE6, 0x92, 0x61, 0x9D, 0x30, 0xEB, 0x4F, 0xA0, 0x81, 0xF5, 0x65, 0x54, 0xD0, 0xE9, 0xE2, 0xA5, 0xE1, 0x8D, 0x7B, 0x67, 0x36, 0xAE, 0xA7, 0x32, 0x54, 0xBC, 0x0A, 0xFE, 0x1F, 0x48, 0x47, 0x41, 0x10, 0x9D, 0xBE, 0xA6, 0xAF, 0x0D, 0xC7, 0x83, 0x52, 0x24, 0x2E, 0x83},
		rounds:   6,
		outBytes: [48]byte{0x9E, 0x5B, 0x05, 0xCB, 0xD7, 0xDB, 0xDE, 0x94, 0x51, 0x32, 0x84, 0x32, 0x06, 0x6A, 0x09, 0xDD, 0x66, 0xFF, 0x01, 0x55, 0xDA, 0x38, 0x26, 0x4F, 0x9C, 0x59, 0xC5, 0x66, 0x5C, 0xB7, 0xF7, 0xA7, 0x23, 0xC2, 0x0F, 0xE8, 0x70, 0x89, 0x21, 0x63, 0x7B, 0xF3, 0x5F, 0x40, 0x32, 0x4E, 0xE6, 0x49},
	},
	{
		inBytes:  [48]byte{0xA4, 0xCA, 0x45, 0x89, 0x33, 0x6C, 0x92, 0x2A, 0x45, 0xF6, 0xC9, 0x99, 0x81, 0x16, 0xF4, 0x52, 0x26, 0xC5, 0x66, 0xF2, 0x9E, 0xB2, 0xE3, 0xD9, 0x98, 0xDB, 0xCE, 0xCB, 0x86, 0x16, 0x81, 0xF5, 0x54, 0xFC, 0xE5, 0x44, 0x4A, 0x62, 0x65, 0xF5, 0x8B, 0x7A, 0x6E, 0xE9, 0xE1, 0x4E, 0xD0, 0x06},
		rounds:   6,
		outBytes: [48]byte{0x9C, 0x51, 0xB9, 0xCE, 0x01, 0xDC, 0x9B, 0x4B, 0x47, 0xCA, 0x89, 0x91, 0x5D, 0xD0, 0x24, 0x44, 0x9B, 0x09, 0x7B, 0x02, 0xF3, 0x1B, 0x8B, 0x13, 0x2E, 0x05, 0x54, 0xB9, 0xA9, 0xCB, 0xB6, 0xA3, 0x1D, 0x08, 0x76, 0xFB, 0x07, 0x3E, 0xF2, 0x49, 0xE6, 0x44, 0x95, 0xE7, 0x7F, 0xE2, 0xC3, 0x94},
	},
	{
		inBytes:  [48]byte{0x34, 0xF6, 0x32, 0xA4, 0x00, 0xF6, 0x7A, 0x61, 0xE4, 0xF6, 0x4B, 0x54, 0x87, 0x70, 0xCA, 0x95, 0xAE, 0x1A, 0x4D, 0x33, 0x00, 0x11, 0x78, 0xC3, 0xC3, 0x76, 0x02, 0x40, 0x05, 0xFF, 0xB5, 0xBE, 0x06, 0x27, 0x30, 0x40, 0x03, 0x00, 0x09, 0x1B, 0x7E, 0x13, 0x77, 0xC1, 0x09, 0x35, 0x55, 0x4A},
		rounds:   6,
		outBytes: [48]byte{0x75, 0x8A, 0x3A, 0xEA, 0x04, 0x18, 0xF3, 0xF0, 0x1B, 0x84, 0x03, 0xA8, 0xBA, 0xAE, 0x94, 0x96, 0x09, 0xC6, 0x98, 0x18, 0xCC, 0x74, 0x7F, 0x1E, 0x20, 0xA7, 0x15, 0xE0, 0x15, 0x6D, 0x8D, 0x0A, 0xBE, 0x7B, 0xDA, 0xA7, 0xDB, 0x45, 0x70, 0x55, 0x2D, 0xEA, 0xB3, 0x74, 0x4B, 0x21, 0x98, 0x0E},
	},
	{
		inBytes:  [48]byte{0x01, 0x2C, 0x48, 0x7E, 0xDD, 0xAD, 0x45, 0x34, 0x84, 0x2E, 0x31, 0x90, 0x67, 0xCD, 0x04, 0xCF, 0xC7, 0xFE, 0x0B, 0x99, 0xE8, 0xB5, 0xC7, 0x69, 0xB9, 0x5C, 0x4A, 0x6E, 0x3F, 0xB1, 0xBF, 0x1D, 0xAD, 0x8A, 0x2D, 0x09, 0x62, 0x1E, 0x7C, 0xE3, 0x5C, 0x42, 0xE8, 0xF0, 0x5B, 0x94, 0xA0, 0xE4},
		rounds:   6,
		outBytes: [48]byte{0x5E, 0xEB, 0x0A, 0xA8, 0x21, 0x25, 0x0B, 0xB7, 0xDB, 0xCB, 0xFF, 0x9D, 0xA7, 0xBE, 0x15, 0x87, 0x12, 0x5D, 0xB6, 0xA3, 0x03, 0xC0, 0x01, 0x5A, 0xC3, 0xE9, 0xD2, 0x59, 0xF5, 0x41, 0xB3, 0x20, 0x82, 0x1E, 0x30, 0xF1, 0x37, 0xF6, 0x67, 0xDD, 0x29, 0x88, 0x7C, 0xFD, 0x91, 0x97, 0x93, 0x55},
	},
}

func TestPermutation(t *testing.T) {
	for _, tt := range perumtationTestTable {
		newXD, _ := NewXooDoo(tt.rounds, tt.inBytes)
		newXD.Permutation()
		assert.Equal(t, tt.outBytes, newXD.Bytes())
	}
}
