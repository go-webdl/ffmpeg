package ffmpeg

/*
#cgo pkg-config: libavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AVPacket struct {
	AVPacket *C.AVPacket
}

func NewAVPacket() (packet *AVPacket, err error) {
	pkt := &AVPacket{
		AVPacket: C.av_packet_alloc(),
	}
	if pkt.AVPacket == nil {
		err = fmt.Errorf("av_packet_alloc failed: %w", ErrMemoryFailure)
		return
	}
	packet = pkt
	return
}

func NewAVPacketWithSize(size int) (packet *AVPacket, err error) {
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
	if err = GetAVError(C.av_new_packet(pkt.AVPacket, C.int(size))); err != nil {
		return
	}
	packet, pkt = pkt, nil
	return
}

func NewAVPacketWithCBytes(data *C.uint8_t, size C.int) (packet *AVPacket, err error) {
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
	if err = GetAVError(C.av_packet_from_data(pkt.AVPacket, data, C.int(size))); err != nil {
		return
	}
	packet, pkt = pkt, nil
	return
}

func (packet *AVPacket) StreamIndex() (index int) {
	if packet == nil || packet.AVPacket == nil {
		index = -1
	} else {
		index = int(packet.AVPacket.stream_index)
	}
	return
}

func (packet *AVPacket) SetStreamIndex(index int) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	packet.AVPacket.stream_index = C.int(index)
}

func (packet *AVPacket) PTS() (pts int64) {
	if packet == nil || packet.AVPacket == nil {
		pts = AV_NOPTS_VALUE
	} else {
		pts = int64(packet.AVPacket.pts)
	}
	return
}

func (packet *AVPacket) SetPTS(pts int64) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	packet.AVPacket.pts = C.int64_t(pts)
}

func (packet *AVPacket) DTS() (dts int64) {
	if packet == nil || packet.AVPacket == nil {
		dts = AV_NOPTS_VALUE
	} else {
		dts = int64(packet.AVPacket.dts)
	}
	return
}

func (packet *AVPacket) SetDTS(dts int64) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	packet.AVPacket.dts = C.int64_t(dts)
}

func (packet *AVPacket) Duration() (duration int64) {
	if packet == nil || packet.AVPacket == nil {
		duration = 0
	} else {
		duration = int64(packet.AVPacket.duration)
	}
	return
}

func (packet *AVPacket) SetDuration(duration int64) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	packet.AVPacket.duration = C.int64_t(duration)
}

func (packet *AVPacket) Pos() (pos int64) {
	if packet == nil || packet.AVPacket == nil {
		pos = -1
	} else {
		pos = int64(packet.AVPacket.pos)
	}
	return
}

func (packet *AVPacket) SetPos(pos int64) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	packet.AVPacket.pos = C.int64_t(pos)
}

func (packet *AVPacket) Shrink(size int) {
	if packet != nil && packet.AVPacket != nil {
		C.av_shrink_packet(packet.AVPacket, C.int(size))
	}
}

func (packet *AVPacket) Grow(growBy int) (err error) {
	if packet == nil || packet.AVPacket == nil {
		return
	}
	return GetAVError(C.av_grow_packet(packet.AVPacket, C.int(growBy)))
}

func (packet *AVPacket) Clone() (ref *AVPacket, err error) {
	pkt := &AVPacket{
		AVPacket: C.av_packet_clone(packet.AVPacket),
	}
	if pkt.AVPacket == nil {
		err = fmt.Errorf("av_packet_clone failed: %w", ErrMemoryFailure)
		return
	}
	ref = pkt
	return
}

func (packet *AVPacket) Ref(dst *AVPacket) (err error) {
	if packet == nil || packet.AVPacket == nil || dst == nil || dst.AVPacket == nil {
		return
	}
	if err = GetAVError(C.av_packet_ref(dst.AVPacket, packet.AVPacket)); err != nil {
		return
	}
	return
}

func (packet *AVPacket) MoveRef(dst *AVPacket) {
	if packet != nil && packet.AVPacket != nil && dst != nil && dst.AVPacket != nil {
		C.av_packet_move_ref(dst.AVPacket, packet.AVPacket)
	}
}

func (packet *AVPacket) CopyProps(dst *AVPacket) (err error) {
	if packet == nil || packet.AVPacket == nil || dst == nil || dst.AVPacket == nil {
		return
	}
	return GetAVError(C.av_packet_copy_props(dst.AVPacket, packet.AVPacket))
}

func (packet *AVPacket) Unref() {
	if packet != nil && packet.AVPacket != nil {
		C.av_packet_unref(packet.AVPacket)
	}
}

func (packet *AVPacket) Close() (err error) {
	if packet != nil && packet.AVPacket != nil {
		C.av_packet_free((**C.AVPacket)(unsafe.Pointer(&packet.AVPacket)))
	}
	return
}
