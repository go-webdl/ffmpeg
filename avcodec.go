package ffmpeg

/*
#cgo pkg-config: libavcodec

#include "libavcodec/avcodec.h"
*/
import "C"

type AVCodec struct {
	AVCodec *C.AVCodec
}
