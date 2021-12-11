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

type AVStream struct {
	AVStream *C.AVStream
}

func (s *AVStream) AVCodecParameters() *AVCodecParameters {
	return &AVCodecParameters{
		AVCodecParameters: s.AVStream.codecpar,
	}
}

func (s *AVStream) TimeBase() AVRational {
	return AVRational(s.AVStream.time_base)
}

func (s *AVStream) SetTimeBase(timebase AVRational) {
	s.AVStream.time_base = C.AVRational(timebase)
}

type AVCodecParameters struct {
	AVCodecParameters *C.AVCodecParameters
}

func (p *AVCodecParameters) AVCodecType() AVMEDIA_TYPE {
	return AVMEDIA_TYPE(p.AVCodecParameters.codec_type)
}

func (p *AVCodecParameters) SetAVCodecTag(codecTag uint32) {
	p.AVCodecParameters.codec_tag = C.uint32_t(codecTag)
}

func AVCodecParametersCopy(dst, src *AVCodecParameters) (err error) {
	if dst == nil || dst.AVCodecParameters == nil || src == nil || src.AVCodecParameters == nil {
		return
	}
	if err = GetAVError(C.avcodec_parameters_copy(dst.AVCodecParameters, src.AVCodecParameters)); err != nil {
		return
	}
	return
}
