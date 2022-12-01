package guBasic

// # Brief
// ## Functions
//    - pointer(w32 관련)
// ## Packages
//   - basic
// ## Usage
//    -

// # Import
//   - strings : <builtin> Print format
//   - syscall : <builtin> win32 dll
//   - unsafe : <builtin> pointer
//   - time : <builtin> sleep,
//   - golang.org/x/text/encoding/korean : <3rd party> 한글 인코딩
//   - github.com/suapapa/go_hangul/encoding/cp949 : <3rd party> [golang 2byte 한글 처리](http://egloos.zum.com/manwooo/v/1949391)
//   - golang.org/x/text/transform
import (
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/korean"
)

// # Functions(pointer <- data)
// ## UintptrFromBool
//   - bool -> uintptr
func UintptrFromBool(b bool) uintptr {
	if b {
		return uintptr(1)
	} else {
		return uintptr(0)
	}
}

// ## UintptrFromBytes
//   - bytes -> uintptr
func UintptrFromBytes(_bytes []byte) uintptr {
	return uintptr(unsafe.Pointer(&append(_bytes, 0)[0]))
}

// ## UintptrFromStr
//   - str -> uintptr
func UintptrFromStr(str string) uintptr {
	return uintptr(unsafe.Pointer(&append([]byte(str), 0)[0]))
}

// ## UintptrFromUtf
//   - utf -> uintptr
func UintptrFromUtf(str string) uintptr {
	s16, _ := syscall.UTF16PtrFromString(str)
	// return voidptr(unsafe.Pointer(s16))
	return uintptr(unsafe.Pointer(s16))
}

// ## PtrFromStr
//   - str -> ptr(unsafe.Pointer)
func PtrFromStr(str string) unsafe.Pointer {
	return unsafe.Pointer(&append([]byte(str), 0)[0])
}

// # Functions(data <- pointer)
// ## StrFromUintptr
//   - uintptr -> []byte
//   - 문자열 끝(/0)이 나오면 반환, 최대 크기 4096
func BytesFromUintptr(uptr uintptr) []byte {
	bytes := make([]byte, 0)

	for i := 0; i < 4096; i++ { // 4096: max byte
		b := *(*byte)(unsafe.Pointer(uptr + uintptr(i)*unsafe.Sizeof(byte(0))))
		if b == 0 { // NOTE: 문자열 끝(/0)
			return bytes
		}

		bytes = append(bytes, b)
	}
	return bytes
}

// ## StrFromUintptr
//   - uintptr -> str
func StrFromUintptr(ptr uintptr) string {
	return string(BytesFromUintptr(ptr))
}

// ## StrFromUintptr
//   - ptr(unsafe.Pointer) -> str
func StrFromPtr(ptr unsafe.Pointer) string {
	return string(BytesFromUintptr(uintptr(ptr)))
}

// ## BytesFromUintptrWithSize
//   - uintptr -> []byte
//   - size: bytes 크기
func BytesFromPtrWithSize(ptr unsafe.Pointer, size int) []byte {
	bytes := make([]byte, 0)

	for i := 0; i < size; i++ {
		ptr := (*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(byte(0))))
		_byte := *ptr
		bytes = append(bytes, _byte)
	}

	return bytes
}

// # Functions(한글 처리)
// ## KorFromBytes
//   - []byte -> Kor
func KorFromBytes(bytes []byte) string {
	idxNull := strings.Index(string(bytes), "\x00")

	if idxNull >= 0 {
		bytes = bytes[:idxNull]
	}

	bytes_utf8, err := korean.EUCKR.NewDecoder().Bytes(bytes)
	if err != nil {
		if len(bytes) > 0 {
			return KorFromBytes(bytes[:len(bytes)-1])
		}

		return string(bytes)
	}

	return string(bytes_utf8)
}

// ## KorFromUintptr
//   - uintptr -> Kor
func KorFromUintptr(ptr uintptr) string {
	return KorFromBytes(BytesFromUintptr(ptr))
}
