package schema

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Daily struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   uint32
	AdjClose float64
}

func (d *Daily) Size() (s uint64) {

	{
		l := uint64(len(d.Date))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 44
	return
}
func (d *Daily) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Date))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Date)
		i += l
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Open)))

		buf[i+0+0] = byte(v >> 0)

		buf[i+1+0] = byte(v >> 8)

		buf[i+2+0] = byte(v >> 16)

		buf[i+3+0] = byte(v >> 24)

		buf[i+4+0] = byte(v >> 32)

		buf[i+5+0] = byte(v >> 40)

		buf[i+6+0] = byte(v >> 48)

		buf[i+7+0] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.High)))

		buf[i+0+8] = byte(v >> 0)

		buf[i+1+8] = byte(v >> 8)

		buf[i+2+8] = byte(v >> 16)

		buf[i+3+8] = byte(v >> 24)

		buf[i+4+8] = byte(v >> 32)

		buf[i+5+8] = byte(v >> 40)

		buf[i+6+8] = byte(v >> 48)

		buf[i+7+8] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Low)))

		buf[i+0+16] = byte(v >> 0)

		buf[i+1+16] = byte(v >> 8)

		buf[i+2+16] = byte(v >> 16)

		buf[i+3+16] = byte(v >> 24)

		buf[i+4+16] = byte(v >> 32)

		buf[i+5+16] = byte(v >> 40)

		buf[i+6+16] = byte(v >> 48)

		buf[i+7+16] = byte(v >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Close)))

		buf[i+0+24] = byte(v >> 0)

		buf[i+1+24] = byte(v >> 8)

		buf[i+2+24] = byte(v >> 16)

		buf[i+3+24] = byte(v >> 24)

		buf[i+4+24] = byte(v >> 32)

		buf[i+5+24] = byte(v >> 40)

		buf[i+6+24] = byte(v >> 48)

		buf[i+7+24] = byte(v >> 56)

	}
	{

		buf[i+0+32] = byte(d.Volume >> 0)

		buf[i+1+32] = byte(d.Volume >> 8)

		buf[i+2+32] = byte(d.Volume >> 16)

		buf[i+3+32] = byte(d.Volume >> 24)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.AdjClose)))

		buf[i+0+36] = byte(v >> 0)

		buf[i+1+36] = byte(v >> 8)

		buf[i+2+36] = byte(v >> 16)

		buf[i+3+36] = byte(v >> 24)

		buf[i+4+36] = byte(v >> 32)

		buf[i+5+36] = byte(v >> 40)

		buf[i+6+36] = byte(v >> 48)

		buf[i+7+36] = byte(v >> 56)

	}
	return buf[:i+44], nil
}

func (d *Daily) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Date = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		v := 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
		d.Open = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)
		d.High = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.Low = *(*float64)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+24]) << 0) | (uint64(buf[i+1+24]) << 8) | (uint64(buf[i+2+24]) << 16) | (uint64(buf[i+3+24]) << 24) | (uint64(buf[i+4+24]) << 32) | (uint64(buf[i+5+24]) << 40) | (uint64(buf[i+6+24]) << 48) | (uint64(buf[i+7+24]) << 56)
		d.Close = *(*float64)(unsafe.Pointer(&v))

	}
	{

		d.Volume = 0 | (uint32(buf[i+0+32]) << 0) | (uint32(buf[i+1+32]) << 8) | (uint32(buf[i+2+32]) << 16) | (uint32(buf[i+3+32]) << 24)

	}
	{

		v := 0 | (uint64(buf[i+0+36]) << 0) | (uint64(buf[i+1+36]) << 8) | (uint64(buf[i+2+36]) << 16) | (uint64(buf[i+3+36]) << 24) | (uint64(buf[i+4+36]) << 32) | (uint64(buf[i+5+36]) << 40) | (uint64(buf[i+6+36]) << 48) | (uint64(buf[i+7+36]) << 56)
		d.AdjClose = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 44, nil
}
