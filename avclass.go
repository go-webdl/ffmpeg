package ffmpeg

/*

#cgo pkg-config: libavutil
#cgo linux CFLAGS: -Wno-format-security

#include <stdlib.h>
#include "libavutil/log.h"
#include "libavutil/opt.h"

void _golang_ffmpeg_av_log(void *avcl, int level, const char *msg) {
	av_log(avcl, level, msg);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AVClass interface {
	AVClass() AVClassObject
}

type AVClassObject struct {
	AVClass unsafe.Pointer
}

func (avcl AVClassObject) SetAVClass(cls *C.AVClass) {
	*(**C.AVClass)(avcl.AVClass) = cls
}

func (avcl AVClassObject) AVLog(level AV_LOG_LEVEL, format string, a ...interface{}) {
	msg := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(msg))
	C._golang_ffmpeg_av_log(avcl.AVClass, C.int(level), msg)
}

func (avcl AVClassObject) AVOptSetDefaults() {
	C.av_opt_set_defaults(avcl.AVClass)
}

func (avcl AVClassObject) AVOptSetDefaults2(mask, flags int) {
	C.av_opt_set_defaults2(avcl.AVClass, C.int(mask), C.int(flags))
}

func (avcl AVClassObject) AVOptSet(key, value string, searchFlags int) (err error) {
	ckey := C.CString(key)
	cval := C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(ckey))
		C.free(unsafe.Pointer(cval))
	}()

	if err = GetAVError(C.av_opt_set(avcl.AVClass, ckey, cval, C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetInt64(key string, value int64, searchFlags int) (err error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	if err = GetAVError(C.av_opt_set_int(avcl.AVClass, ckey, C.int64_t(value), C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetFloat64(key string, value float64, searchFlags int) (err error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	if err = GetAVError(C.av_opt_set_double(avcl.AVClass, ckey, C.double(value), C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetBytes(key string, value []byte, searchFlags int) (err error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	if err = GetAVError(C.av_opt_set_bin(avcl.AVClass, ckey, (*C.uint8_t)((unsafe.Pointer(&value[0]))), C.int(len(value)), C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetImageSize(key string, width, height int, searchFlags int) (err error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	if err = GetAVError(C.av_opt_set_image_size(avcl.AVClass, ckey, C.int(width), C.int(height), C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetDictVal(key string, avdict *AVDictionary, searchFlags int) (err error) {
	if avdict == nil {
		err = fmt.Errorf("avdict is nil: %w", ErrInvalidParam)
		return
	}

	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	var cval *C.AVDictionary
	if avdict != nil && avdict.AVDictionary != nil {
		cval = avdict.AVDictionary
	}

	if err = GetAVError(C.av_opt_set_dict_val(avcl.AVClass, ckey, cval, C.int(searchFlags))); err != nil {
		return
	}
	return
}

func (avcl AVClassObject) AVOptSetDict(avdict *AVDictionary, searchFlags int) (err error) {
	var cval **C.AVDictionary
	if avdict != nil && avdict.AVDictionary != nil {
		cval = (**C.AVDictionary)(unsafe.Pointer(&avdict.AVDictionary))
	}

	if err = GetAVError(C.av_opt_set_dict2(avcl.AVClass, cval, C.int(searchFlags))); err != nil {
		return
	}
	return
}
