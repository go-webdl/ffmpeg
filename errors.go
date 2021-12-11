package ffmpeg

/*
#cgo pkg-config: libavutil

#include "libavutil/avutil.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"io"
	"unsafe"
)

var (
	ErrInvalidParam  = errors.New("invalid parameter")
	ErrMemoryFailure = errors.New("memory failure")
)

type AVError struct {
	Code    int
	message string
}

var _ error = (*AVError)(nil)

func GetAVError(cerr C.int) (err error) {
	if cerr >= 0 {
		return
	}

	if cerr == C.AVERROR_EOF {
		err = io.EOF
		return
	}

	return &AVError{
		Code: int(cerr),
	}
}

func (err AVError) Error() string {
	if len(err.message) == 0 {
		var b [0x400]byte
		C.av_strerror(C.int(err.Code), (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)))
		err.message = string(b[:])
		if len(err.message) == 0 {
			err.message = fmt.Sprintf("ffmpeg error code %d", err.Code)
		}
	}
	return err.message
}
