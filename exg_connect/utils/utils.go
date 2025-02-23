package utils

import "bytes"

func Int8SliceToString(arr []int8) string {
	b := make([]byte, len(arr))
	for i, v := range arr {
		b[i] = byte(v)
	}
	b = bytes.Trim(b, "\x00")
	return string(b)
}
