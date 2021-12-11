package ffmpeg

/*
#cgo pkg-config: libavformat libavutil

#include "libavutil/avutil.h"
#include "libavutil/timestamp.h"
#include "libavformat/avformat.h"
*/
import "C"
import "unsafe"

type AVRational C.AVRational

func (r AVRational) Numerator() int {
	return int(r.num)
}

func (r AVRational) Denominator() int {
	return int(r.den)
}

func NewAVRational(numerator, demoninator int) AVRational {
	return AVRational(C.av_make_q(C.int(numerator), C.int(demoninator)))
}

func CompareAVRational(a, b AVRational) int {
	return int(C.av_cmp_q(C.AVRational(a), C.AVRational(b)))
}

func AVRescaleQRnd(a int64, bq, cq AVRational, rnd uint32) int64 {
	return int64(C.av_rescale_q_rnd(C.int64_t(a), C.AVRational(bq), C.AVRational(cq), rnd))
}

func AVRescaleQ(a int64, bq, cq AVRational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), C.AVRational(bq), C.AVRational(cq)))
}

func AVTs2Str(ts int64) string {
	var buf [32]byte
	C.av_ts_make_string((*C.char)(unsafe.Pointer(&buf[0])), C.int64_t(ts))
	return string(buf[:])
}

func AVTs2TimeStr(ts int64, timebase AVRational) string {
	var buf [32]byte
	C.av_ts_make_time_string((*C.char)(unsafe.Pointer(&buf[0])), C.int64_t(ts), (*C.AVRational)(unsafe.Pointer(&timebase)))
	return string(buf[:])
}
