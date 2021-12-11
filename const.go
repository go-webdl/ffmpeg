package ffmpeg

/*
#cgo pkg-config: libavformat libavutil

#include "libavutil/log.h"
#include "libavutil/opt.h"
#include "libavutil/error.h"
#include "libavutil/avutil.h"
#include "libavformat/avio.h"
#include "libavformat/avformat.h"
*/
import "C"

const (
	AVERROR_BSF_NOT_FOUND      = int(C.AVERROR_BSF_NOT_FOUND)
	AVERROR_BUG                = int(C.AVERROR_BUG)
	AVERROR_BUFFER_TOO_SMALL   = int(C.AVERROR_BUFFER_TOO_SMALL)
	AVERROR_DECODER_NOT_FOUND  = int(C.AVERROR_DECODER_NOT_FOUND)
	AVERROR_DEMUXER_NOT_FOUND  = int(C.AVERROR_DEMUXER_NOT_FOUND)
	AVERROR_ENCODER_NOT_FOUND  = int(C.AVERROR_ENCODER_NOT_FOUND)
	AVERROR_EOF                = int(C.AVERROR_EOF)
	AVERROR_EXIT               = int(C.AVERROR_EXIT)
	AVERROR_EXTERNAL           = int(C.AVERROR_EXTERNAL)
	AVERROR_FILTER_NOT_FOUND   = int(C.AVERROR_FILTER_NOT_FOUND)
	AVERROR_INVALIDDATA        = int(C.AVERROR_INVALIDDATA)
	AVERROR_MUXER_NOT_FOUND    = int(C.AVERROR_MUXER_NOT_FOUND)
	AVERROR_OPTION_NOT_FOUND   = int(C.AVERROR_OPTION_NOT_FOUND)
	AVERROR_PATCHWELCOME       = int(C.AVERROR_PATCHWELCOME)
	AVERROR_PROTOCOL_NOT_FOUND = int(C.AVERROR_PROTOCOL_NOT_FOUND)
	AVERROR_STREAM_NOT_FOUND   = int(C.AVERROR_STREAM_NOT_FOUND)
	AVERROR_BUG2               = int(C.AVERROR_BUG2)
	AVERROR_UNKNOWN            = int(C.AVERROR_UNKNOWN)
	AVERROR_EXPERIMENTAL       = int(C.AVERROR_EXPERIMENTAL)
	AVERROR_INPUT_CHANGED      = int(C.AVERROR_INPUT_CHANGED)
	AVERROR_OUTPUT_CHANGED     = int(C.AVERROR_OUTPUT_CHANGED)
	AVERROR_HTTP_BAD_REQUEST   = int(C.AVERROR_HTTP_BAD_REQUEST)
	AVERROR_HTTP_UNAUTHORIZED  = int(C.AVERROR_HTTP_UNAUTHORIZED)
	AVERROR_HTTP_FORBIDDEN     = int(C.AVERROR_HTTP_FORBIDDEN)
	AVERROR_HTTP_NOT_FOUND     = int(C.AVERROR_HTTP_NOT_FOUND)
	AVERROR_HTTP_OTHER_4XX     = int(C.AVERROR_HTTP_OTHER_4XX)
	AVERROR_HTTP_SERVER_ERROR  = int(C.AVERROR_HTTP_SERVER_ERROR)
)

type AV_LOG_LEVEL int

const (
	AV_LOG_QUIET   = AV_LOG_LEVEL(C.AV_LOG_QUIET)
	AV_LOG_PANIC   = AV_LOG_LEVEL(C.AV_LOG_PANIC)
	AV_LOG_FATAL   = AV_LOG_LEVEL(C.AV_LOG_FATAL)
	AV_LOG_ERROR   = AV_LOG_LEVEL(C.AV_LOG_ERROR)
	AV_LOG_WARNING = AV_LOG_LEVEL(C.AV_LOG_WARNING)
	AV_LOG_INFO    = AV_LOG_LEVEL(C.AV_LOG_INFO)
	AV_LOG_VERBOSE = AV_LOG_LEVEL(C.AV_LOG_VERBOSE)
	AV_LOG_DEBUG   = AV_LOG_LEVEL(C.AV_LOG_DEBUG)
	AV_LOG_TRACE   = AV_LOG_LEVEL(C.AV_LOG_TRACE)
)

const IO_BUFFER_SIZE int = 32768 // 32 KB

const (
	AV_OPT_FLAG_ENCODING_PARAM  = int(C.AV_OPT_FLAG_ENCODING_PARAM) ///< a generic parameter which can be set by the user for muxing or encoding
	AV_OPT_FLAG_DECODING_PARAM  = int(C.AV_OPT_FLAG_DECODING_PARAM) ///< a generic parameter which can be set by the user for demuxing or decoding
	AV_OPT_FLAG_AUDIO_PARAM     = int(C.AV_OPT_FLAG_AUDIO_PARAM)
	AV_OPT_FLAG_VIDEO_PARAM     = int(C.AV_OPT_FLAG_VIDEO_PARAM)
	AV_OPT_FLAG_SUBTITLE_PARAM  = int(C.AV_OPT_FLAG_SUBTITLE_PARAM)
	AV_OPT_FLAG_EXPORT          = int(C.AV_OPT_FLAG_EXPORT)
	AV_OPT_FLAG_READONLY        = int(C.AV_OPT_FLAG_READONLY)
	AV_OPT_FLAG_BSF_PARAM       = int(C.AV_OPT_FLAG_BSF_PARAM)       ///< a generic parameter which can be set by the user for bit stream filtering
	AV_OPT_FLAG_RUNTIME_PARAM   = int(C.AV_OPT_FLAG_RUNTIME_PARAM)   ///< a generic parameter which can be set by the user at runtime
	AV_OPT_FLAG_FILTERING_PARAM = int(C.AV_OPT_FLAG_FILTERING_PARAM) ///< a generic parameter which can be set by the user for filtering
	AV_OPT_FLAG_DEPRECATED      = int(C.AV_OPT_FLAG_DEPRECATED)      ///< set if option is deprecated, users should refer to AVOption.help text for more information
	AV_OPT_FLAG_CHILD_CONSTS    = int(C.AV_OPT_FLAG_CHILD_CONSTS)    ///< set if option constants can also reside in child objects
)

const (
	AV_OPT_SEARCH_CHILDREN       = int(C.AV_OPT_SEARCH_CHILDREN)
	AV_OPT_SEARCH_FAKE_OBJ       = int(C.AV_OPT_SEARCH_FAKE_OBJ)
	AV_OPT_ALLOW_NULL            = int(C.AV_OPT_ALLOW_NULL)
	AV_OPT_MULTI_COMPONENT_RANGE = int(C.AV_OPT_MULTI_COMPONENT_RANGE)
)

const (
	AVFMT_FLAG_GENPTS          = int(C.AVFMT_FLAG_GENPTS)          ///< Generate missing pts even if it requires parsing future frames.
	AVFMT_FLAG_IGNIDX          = int(C.AVFMT_FLAG_IGNIDX)          ///< Ignore index.
	AVFMT_FLAG_NONBLOCK        = int(C.AVFMT_FLAG_NONBLOCK)        ///< Do not block when reading packets from input.
	AVFMT_FLAG_IGNDTS          = int(C.AVFMT_FLAG_IGNDTS)          ///< Ignore DTS on frames that contain both DTS & PTS
	AVFMT_FLAG_NOFILLIN        = int(C.AVFMT_FLAG_NOFILLIN)        ///< Do not infer any values from other values, just return what is stored in the container
	AVFMT_FLAG_NOPARSE         = int(C.AVFMT_FLAG_NOPARSE)         ///< Do not use AVParsers, you also must set AVFMT_FLAG_NOFILLIN as the fillin code works on frames and no parsing -> no frames. Also seeking to frames can not work if parsing to find frame boundaries has been disabled
	AVFMT_FLAG_NOBUFFER        = int(C.AVFMT_FLAG_NOBUFFER)        ///< Do not buffer frames when possible
	AVFMT_FLAG_CUSTOM_IO       = int(C.AVFMT_FLAG_CUSTOM_IO)       ///< The caller has supplied a custom AVIOContext, don't avio_close() it.
	AVFMT_FLAG_DISCARD_CORRUPT = int(C.AVFMT_FLAG_DISCARD_CORRUPT) ///< Discard frames marked corrupted
	AVFMT_FLAG_FLUSH_PACKETS   = int(C.AVFMT_FLAG_FLUSH_PACKETS)   ///< Flush the AVIOContext every packet.
	AVFMT_FLAG_BITEXACT        = int(C.AVFMT_FLAG_BITEXACT)        // When muxing, try to avoid writing any random/volatile data to the output. This includes any random IDs, real-time timestamps/dates, muxer version, etc. This flag is mainly intended for testing.
	AVFMT_FLAG_SORT_DTS        = int(C.AVFMT_FLAG_SORT_DTS)        ///< try to interleave outputted packets by dts (using this flag can slow demuxing down)
	AVFMT_FLAG_FAST_SEEK       = int(C.AVFMT_FLAG_FAST_SEEK)       ///< Enable fast, but inaccurate seeks for some formats
	AVFMT_FLAG_SHORTEST        = int(C.AVFMT_FLAG_SHORTEST)        ///< Stop muxing when the shortest stream stops.
	AVFMT_FLAG_AUTO_BSF        = int(C.AVFMT_FLAG_AUTO_BSF)        ///< Add bitstream filters as requested by the muxer
)

const (
	AVIO_FLAG_READ       = int(C.AVIO_FLAG_READ)
	AVIO_FLAG_WRITE      = int(C.AVIO_FLAG_WRITE)
	AVIO_FLAG_READ_WRITE = int(C.AVIO_FLAG_READ_WRITE)
)

type AVMEDIA_TYPE int

const (
	AVMEDIA_TYPE_UNKNOWN    = AVMEDIA_TYPE(C.AVMEDIA_TYPE_UNKNOWN)
	AVMEDIA_TYPE_VIDEO      = AVMEDIA_TYPE(C.AVMEDIA_TYPE_VIDEO)
	AVMEDIA_TYPE_AUDIO      = AVMEDIA_TYPE(C.AVMEDIA_TYPE_AUDIO)
	AVMEDIA_TYPE_DATA       = AVMEDIA_TYPE(C.AVMEDIA_TYPE_DATA)
	AVMEDIA_TYPE_SUBTITLE   = AVMEDIA_TYPE(C.AVMEDIA_TYPE_SUBTITLE)
	AVMEDIA_TYPE_ATTACHMENT = AVMEDIA_TYPE(C.AVMEDIA_TYPE_ATTACHMENT)
)

const AV_NOPTS_VALUE = int64(C.AV_NOPTS_VALUE)

const (
	AV_ROUND_ZERO        = uint32(C.AV_ROUND_ZERO)
	AV_ROUND_INF         = uint32(C.AV_ROUND_INF)
	AV_ROUND_DOWN        = uint32(C.AV_ROUND_DOWN)
	AV_ROUND_UP          = uint32(C.AV_ROUND_UP)
	AV_ROUND_NEAR_INF    = uint32(C.AV_ROUND_NEAR_INF)
	AV_ROUND_PASS_MINMAX = uint32(C.AV_ROUND_PASS_MINMAX)
)
