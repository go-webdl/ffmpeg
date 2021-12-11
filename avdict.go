package ffmpeg

/*

#cgo pkg-config: libavutil

#include <stdlib.h>
#include "libavutil/dict.h"

*/
import "C"
import (
	"unsafe"
)

type AVDictionary struct {
	GoMap        map[string]string
	AVDictionary *C.struct_AVDictionary
}

func NewDict(values map[string]string) (avdict *AVDictionary, err error) {
	ret := &AVDictionary{
		GoMap:        make(map[string]string),
		AVDictionary: nil,
	}

	for key, value := range values {
		if err = avdict.Set(key, value); err != nil {
			return
		}
	}

	avdict = ret
	return
}

func (avdict *AVDictionary) Set(key, value string) (err error) {
	ckey := C.CString(key)
	cval := C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(ckey))
		C.free(unsafe.Pointer(cval))
	}()

	if err = GetAVError(C.av_dict_set(&avdict.AVDictionary, ckey, cval, 0)); err != nil {
		return
	}

	avdict.GoMap[key] = value
	return
}

func (avdict *AVDictionary) Del(key string) (err error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	if err = GetAVError(C.av_dict_set(&avdict.AVDictionary, ckey, nil, 0)); err != nil {
		return
	}

	delete(avdict.GoMap, key)
	return
}
