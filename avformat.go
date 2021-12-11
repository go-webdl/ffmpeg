package ffmpeg

/*
#cgo pkg-config: libavformat libavutil

#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <string.h>

#include "libavutil/avutil.h"
#include "libavformat/avformat.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

type AVFormatContext struct {
	avfmtctx *C.AVFormatContext
}

var _ AVClass = (*AVFormatContext)(nil)

func NewAVFormatContext() (avfmtctx *AVFormatContext, err error) {
	ctx := C.avformat_alloc_context()
	if ctx == nil {
		err = fmt.Errorf("avformat_alloc_context failed: %w", ErrMemoryFailure)
		return
	}
	avfmtctx = &AVFormatContext{
		avfmtctx: ctx,
	}
	return
}

func AVFormatOpenInput(c *AVFormatContext, avioctx *AVIOContext, url string, inputFormat *AVInputFormat, options *AVDictionary) (ctx *AVFormatContext, err error) {
	defer func() {
		if c != nil {
			if e := c.Close(); err == nil {
				err = e
			}
		}
	}()

	if c == nil {
		if c, err = NewAVFormatContext(); err != nil {
			return
		}
	}

	if avioctx != nil {
		c.SetAVIOContext(avioctx)
	}

	var curl *C.char
	if len(url) > 0 {
		curl = C.CString(url)
	}
	defer func() {
		if curl != nil {
			C.free(unsafe.Pointer(curl))
		}
	}()

	var cfmt *C.AVInputFormat
	if inputFormat != nil && inputFormat.AVInputFormat != nil {
		cfmt = inputFormat.AVInputFormat
	}

	var coptions **C.AVDictionary
	if options != nil && options.AVDictionary != nil {
		coptions = (**C.AVDictionary)(unsafe.Pointer(&options.AVDictionary))
	}

	if err = GetAVError(C.avformat_open_input((**C.AVFormatContext)(unsafe.Pointer(&c.avfmtctx)), curl, cfmt, coptions)); err != nil {
		return
	}

	ctx = c
	c = nil
	return
}

func AVFormatOpenOutput(c *AVFormatContext, avioctx *AVIOContext, url string, outputFormat *AVOutputFormat, options *AVDictionary) (ctx *AVFormatContext, err error) {
	defer func() {
		if c != nil {
			if e := c.Close(); err == nil {
				err = e
			}
		}
	}()

	if c == nil {
		if c, err = NewAVFormatContext(); err != nil {
			return
		}
	}

	if avioctx != nil {
		c.SetAVIOContext(avioctx)
	}

	if options != nil {
		c.AVClass().AVOptSetDict(options, 0)
	}

	if outputFormat == nil {
		outputFormat = AVFindOutputFormat("", url, "")
	}

	if outputFormat == nil {
		err = fmt.Errorf("cannot determine output format for output url %s: %w", url, ErrInvalidParam)
		return
	}

	if err = c.SetAVOutputFormat(outputFormat); err != nil {
		return
	}

	if err = c.SetURL(url); err != nil {
		return
	}

	ctx = c
	c = nil
	return
}

func (c *AVFormatContext) Close() (err error) {
	if c.avfmtctx != nil {
		C.avformat_free_context(c.avfmtctx)
		c.avfmtctx = nil
	}
	return
}

func (c *AVFormatContext) AVClass() AVClassObject {
	return AVClassObject{AVClass: unsafe.Pointer(c.avfmtctx)}
}

func (c *AVFormatContext) Flags() int {
	return int(c.avfmtctx.flags)
}

func (c *AVFormatContext) SetFlag(flag int) {
	c.avfmtctx.flags |= C.int(flag)
}

func (c *AVFormatContext) UnsetFlag(flag int) {
	c.avfmtctx.flags &^= C.int(flag)
}

func (c *AVFormatContext) AVStreams() (streams []*AVStream) {
	count := uint(c.avfmtctx.nb_streams)
	for i := uint(0); i < count; i++ {
		streams = append(streams, c.AVStreamAtIndex(int(i)))
	}
	return
}

func (c *AVFormatContext) AVStreamAtIndex(index int) (stream *AVStream) {
	if c == nil || c.avfmtctx == nil || c.avfmtctx.streams == nil || index < 0 || uint(index) >= uint(c.avfmtctx.nb_streams) {
		return
	}
	stream = &AVStream{
		AVStream: *(**C.AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(c.avfmtctx.streams)) + unsafe.Sizeof((*C.AVStream)(unsafe.Pointer(nil)))*uintptr(index))),
	}
	return
}

func (c *AVFormatContext) NewAVStream(codec *AVCodec) (stream *AVStream, err error) {
	var ccodec *C.AVCodec
	if codec != nil && codec.AVCodec != nil {
		ccodec = codec.AVCodec
	}
	s := &AVStream{
		AVStream: C.avformat_new_stream(c.avfmtctx, ccodec),
	}
	if s.AVStream == nil {
		err = errors.New("avformat_new_stream failed")
		return
	}
	stream = s
	return
}

func (c *AVFormatContext) SetAVIOContext(avioctx *AVIOContext) {
	if avioctx == nil || avioctx.AVIOContext == nil {
		c.avfmtctx.pb = nil
	} else {
		c.avfmtctx.pb = avioctx.AVIOContext
	}
}

func (c *AVFormatContext) SetAVInputFormat(avfmt *AVInputFormat) (err error) {
	if avfmt == nil || avfmt.AVInputFormat == nil {
		c.avfmtctx.iformat = nil
	} else {
		c.avfmtctx.iformat = avfmt.AVInputFormat
	}
	return
}

func (c *AVFormatContext) SetAVOutputFormat(avfmt *AVOutputFormat) (err error) {
	if avfmt == nil || avfmt.AVOutputFormat == nil {
		c.avfmtctx.oformat = nil
	} else {
		c.avfmtctx.oformat = avfmt.AVOutputFormat
		if c.avfmtctx.oformat.priv_data_size > 0 {
			if c.avfmtctx.priv_data = C.av_mallocz(C.size_t(c.avfmtctx.oformat.priv_data_size)); c.avfmtctx.priv_data == nil {
				err = fmt.Errorf("unable to alloc AVFormatContext output format's priv_data: %w", ErrMemoryFailure)
				return
			}
			if c.avfmtctx.oformat.priv_class != nil {
				avcl := AVClassObject{AVClass: unsafe.Pointer(c.avfmtctx.priv_data)}
				avcl.SetAVClass(c.avfmtctx.oformat.priv_class)
				avcl.AVOptSetDefaults()
			}
		}
	}
	return
}

func (c *AVFormatContext) SetURL(url string) (err error) {
	var curl *C.char
	if len(url) > 0 {
		curl = C.CString(url)
	}
	defer func() {
		if curl != nil {
			C.free(unsafe.Pointer(curl))
		}
	}()
	if curl != nil {
		c.avfmtctx.url = C.av_strdup(curl)
		if c.avfmtctx.url == nil {
			err = fmt.Errorf("unable to av_strdup url: %w", ErrMemoryFailure)
			return
		}
	} else {
		if c.avfmtctx.url != nil {
			C.av_free(unsafe.Pointer(c.avfmtctx.url))
		}
		c.avfmtctx.url = nil
	}
	return
}

func (c *AVFormatContext) AVFindStreamInfo() (err error) {
	if err = GetAVError(C.avformat_find_stream_info(c.avfmtctx, nil)); err != nil {
		return
	}
	return
}

func (c *AVFormatContext) AVReadFrame() (packet *AVPacket, err error) {
	var pkt *AVPacket
	if pkt, err = NewAVPacket(); err != nil {
		return
	}
	defer func() {
		if pkt != nil {
			if e := pkt.Close(); err == nil && e != nil {
				err = e
			}
		}
	}()
	if err = GetAVError(C.av_read_frame(c.avfmtctx, pkt.AVPacket)); err != nil {
		return
	}
	packet, pkt = pkt, nil
	return
}

func (c *AVFormatContext) AVFormatWriteHeader() (err error) {
	return GetAVError(C.avformat_write_header(c.avfmtctx, nil))
}

func (c *AVFormatContext) AVInterleavedWriteFrame(packet *AVPacket) (err error) {
	return GetAVError(C.av_interleaved_write_frame(c.avfmtctx, packet.AVPacket))
}

func (c *AVFormatContext) AVWriteTrailer() (err error) {
	return GetAVError(C.av_write_trailer(c.avfmtctx))
}

func (c *AVFormatContext) DumpFormat(index int, url string, isOutput bool) {
	cindex := C.int(index)
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	cisOutput := C.int(0)
	if isOutput {
		cisOutput = C.int(1)
	}
	C.av_dump_format(c.avfmtctx, cindex, curl, cisOutput)
}

type AVInputFormat struct {
	AVInputFormat *C.AVInputFormat
}

func AVFindInputFormat(shortName string) *AVInputFormat {
	cname := C.CString(shortName)
	defer C.free(unsafe.Pointer(cname))
	avfmt := C.av_find_input_format(cname)
	if avfmt == nil {
		return nil
	}
	return &AVInputFormat{
		AVInputFormat: avfmt,
	}
}

func (avfmt AVInputFormat) ShortNames() []string {
	if avfmt.AVInputFormat == nil || avfmt.AVInputFormat.name == nil {
		return nil
	}
	return strings.Split(C.GoString(avfmt.AVInputFormat.name), ",")
}

func (avfmt AVInputFormat) LongName() string {
	if avfmt.AVInputFormat == nil || avfmt.AVInputFormat.long_name == nil {
		return ""
	}
	return C.GoString(avfmt.AVInputFormat.long_name)
}

type AVOutputFormat struct {
	AVOutputFormat *C.AVOutputFormat
}

func AVFindOutputFormat(shortName, filename, mimeType string) *AVOutputFormat {
	var cshortName, cfilename, cmimeType *C.char

	if len(shortName) > 0 {
		cshortName = C.CString(shortName)
	}
	if len(filename) > 0 {
		cfilename = C.CString(filename)
	}
	if len(mimeType) > 0 {
		cmimeType = C.CString(mimeType)
	}

	defer func() {
		if cshortName != nil {
			C.free(unsafe.Pointer(cshortName))
		}
		if cfilename != nil {
			C.free(unsafe.Pointer(cfilename))
		}
		if cmimeType != nil {
			C.free(unsafe.Pointer(cmimeType))
		}
	}()

	avfmt := C.av_guess_format(cshortName, cfilename, cmimeType)
	if avfmt == nil {
		return nil
	}
	return &AVOutputFormat{
		AVOutputFormat: avfmt,
	}
}

func (avfmt AVOutputFormat) ShortNames() []string {
	if avfmt.AVOutputFormat == nil || avfmt.AVOutputFormat.name == nil {
		return nil
	}
	return strings.Split(C.GoString(avfmt.AVOutputFormat.name), ",")
}

func (avfmt AVOutputFormat) LongName() string {
	if avfmt.AVOutputFormat == nil || avfmt.AVOutputFormat.long_name == nil {
		return ""
	}
	return C.GoString(avfmt.AVOutputFormat.long_name)
}
