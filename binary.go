package rueidis

import "unsafe"

// BinaryString convert the provided []byte into a string without copy. It does what strings.Builder.String() does.
// Redis Strings are binary safe, this means that it is safe to store any []byte into Redis directly.
// Users can use this BinaryString helper to insert a []byte as the part of redis command. For example:
//   client.B().Set().Key(rueidis.BinaryString([]byte{0})).Value(rueidis.BinaryString([]byte{0})).Build()
// To read back the []byte of the string returned from the Redis, it is recommended to use the RedisMessage.AsReader.
func BinaryString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
