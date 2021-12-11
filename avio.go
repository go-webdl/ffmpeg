package ffmpeg

/*
#cgo pkg-config: libavformat libavutil

#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <string.h>

#include "libavutil/avutil.h"
#include "libavformat/avio.h"
#include "libavformat/avformat.h"

extern int		_golang_ffmpeg_avio_read_callback	(void *opaque, uint8_t *buf, 	int buf_size);
extern int		_golang_ffmpeg_avio_write_callback	(void *opaque, uint8_t *buf, 	int buf_size);
extern int64_t	_golang_ffmpeg_avio_seek_callback	(void *opaque, int64_t offset, 	int whence);

*/
import "C"

import (
	"errors"
	"fmt"
	"io"
	"runtime/cgo"
	"unsafe"
)

// Functions prototypes for custom IO. Implement necessary prototypes and pass instance pointer to NewAVIOContext.
//
// E.g.:
// 	var reader io.Reader
//
//	avioctx := NewAVIOContext(ctx, &AVIOHandlers{Read: reader.Read})
type AVIOHandlers struct {
	Read  func(p []byte) (n int, err error)
	Write func(p []byte) (n int, err error)
	Seek  func(offset int64, whence int) (int64, error)
}

type AVIOContext struct {
	AVIOContext *C.AVIOContext
	Handlers    AVIOHandlers
	CgoHandle   cgo.Handle
}

var _ AVClass = (*AVIOContext)(nil)

func AVIOOpen(url string, flags int) (avioctx *AVIOContext, err error) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ctx := &AVIOContext{}
	if err = GetAVError(C.avio_open((**C.AVIOContext)(unsafe.Pointer(&ctx.AVIOContext)), curl, C.int(flags))); err != nil {
		return
	}
	avioctx = ctx
	return
}

// AVIOContext constructor. Use it only if You need custom IO behaviour!
func NewAVIOContext(handlers *AVIOHandlers) (avioctx *AVIOContext, err error) {
	if handlers == nil {
		err = fmt.Errorf("handlers for NewAVIOContext cannot be nil: %w", ErrInvalidParam)
		return
	}

	buffer := C.av_malloc(C.size_t(IO_BUFFER_SIZE))
	if buffer == nil {
		err = fmt.Errorf("failed to alloc buffer for NewAVIOContext: %w", ErrMemoryFailure)
		return
	}
	defer func() {
		if buffer != nil {
			C.av_free(buffer)
		}
	}()

	ctx := &AVIOContext{
		Handlers: *handlers,
	}

	// we have to explicitly set it to nil, to force library using default handlers
	var ptrRead, ptrWrite, ptrSeek *[0]byte = nil, nil, nil
	writeFlag := C.int(0)

	if handlers.Read != nil {
		ptrRead = (*[0]byte)(C._golang_ffmpeg_avio_read_callback)
	}

	if handlers.Write != nil {
		ptrWrite = (*[0]byte)(C._golang_ffmpeg_avio_write_callback)
		writeFlag = C.int(1)
	}

	if handlers.Seek != nil {
		ptrSeek = (*[0]byte)(C._golang_ffmpeg_avio_seek_callback)
	}

	ctx.CgoHandle = cgo.NewHandle(ctx)

	if ctx.AVIOContext = C.avio_alloc_context((*C.uchar)(buffer), C.int(IO_BUFFER_SIZE), writeFlag, unsafe.Pointer(ctx.CgoHandle), ptrRead, ptrWrite, ptrSeek); ctx.AVIOContext == nil {
		err = errors.New("unable to initialize avio context")
		return
	}

	buffer = nil
	avioctx = ctx
	return
}

func (c *AVIOContext) AVClass() AVClassObject {
	return AVClassObject{AVClass: unsafe.Pointer(c.AVIOContext)}
}

func (c *AVIOContext) Close() (err error) {
	if c.AVIOContext != nil {
		if c.AVIOContext.buffer != nil {
			C.av_free(unsafe.Pointer(c.AVIOContext.buffer))
			c.AVIOContext.buffer = nil
		}
		C.av_free(unsafe.Pointer(c.AVIOContext))
		c.AVIOContext = nil
		c.CgoHandle.Delete()
	}
	return
}

func (c *AVIOContext) Flush() {
	C.avio_flush(c.AVIOContext)
}

//export _golang_ffmpeg_avio_read_callback
func _golang_ffmpeg_avio_read_callback(opaque unsafe.Pointer, buf *C.uint8_t, buf_size C.int) (ret C.int) {
	handlers := cgo.Handle(opaque).Value().(*AVIOContext).Handlers
	ret = C.int(0)
	if n, err := handlers.Read(unsafe.Slice((*byte)(unsafe.Pointer(buf)), buf_size)); err != nil {
		if err != io.EOF {
			ret = C.AVERROR_EXTERNAL
		} else if ret == 0 {
			ret = C.AVERROR_EOF
		}
	} else {
		ret += C.int(n)
	}
	return
}

//export _golang_ffmpeg_avio_write_callback
func _golang_ffmpeg_avio_write_callback(opaque unsafe.Pointer, buf *C.uint8_t, buf_size C.int) (ret C.int) {
	handlers := cgo.Handle(opaque).Value().(*AVIOContext).Handlers
	ret = C.int(0)
	if n, err := handlers.Write(unsafe.Slice((*byte)(unsafe.Pointer(buf)), buf_size)); err != nil {
		ret = C.AVERROR_EXTERNAL
	} else {
		ret += C.int(n)
	}
	return
}

//export _golang_ffmpeg_avio_seek_callback
func _golang_ffmpeg_avio_seek_callback(opaque unsafe.Pointer, offset C.int64_t, whence C.int) (ret C.int64_t) {
	handlers := cgo.Handle(opaque).Value().(*AVIOContext).Handlers
	var w int
	if whence == C.SEEK_SET {
		w = io.SeekStart
	} else if whence == C.SEEK_CUR {
		w = io.SeekCurrent
	} else if whence == C.SEEK_END {
		w = io.SeekEnd
	} else {
		ret = C.int64_t(C.AVERROR_EXTERNAL)
		return
	}

	n, err := handlers.Seek(int64(offset), w)
	if err != nil {
		ret = C.int64_t(C.AVERROR_EXTERNAL)
		return
	}

	ret = C.int64_t(n)
	return
}
