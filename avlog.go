package ffmpeg

/*

#cgo pkg-config: libavutil

#include <stdlib.h>
#include "libavutil/log.h"
*/
import "C"

func AVLogSetLevel(level AV_LOG_LEVEL) {
	C.av_log_set_level(C.int(level))
}
