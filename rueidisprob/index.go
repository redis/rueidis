package rueidisprob

import "github.com/twmb/murmur3"

func hash(data []byte) (uint64, uint64) {
	return murmur3.Sum128(data)
}

func index(h1, h2 uint64, i uint, maxSize uint64) uint64 {
	offset := h1 + uint64(i)*h2
	return offset % maxSize
}
