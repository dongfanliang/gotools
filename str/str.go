package str

import "unsafe"

func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
