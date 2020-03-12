package str

import "unsafe"

func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
