package rueidisprob

import "github.com/twmb/murmur3"

// indexes returns a list of hash indexes representing a data item.
func indexes(data []byte, iteration, maxSize uint) []uint64 {
	indices := make([]uint64, iteration)

	h1, h2 := murmur3.Sum128(data)
	size := uint64(maxSize)
	for i := uint(0); i < iteration; i++ {
		indices[i] = index(h1, h2, i, size)
	}

	return indices
}

func index(h1, h2 uint64, i uint, maxSize uint64) uint64 {
	offset := h1 + uint64(i)*h2
	return offset % maxSize
}
