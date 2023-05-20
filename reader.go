package bitlib

import (
	"encoding/binary"
	"io"
	"math"
)

type Reader struct {
	in     io.Reader
	err    error
	endian binary.ByteOrder
}

func NewReader(in io.Reader, byteOrder binary.ByteOrder) *Reader {
	return &Reader{
		in:     in,
		endian: byteOrder,
	}
}

func (r Reader) Error() error {
	return r.err
}

func (r *Reader) Float64() float64 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 8)
	_, r.err = io.ReadFull(r.in, data)
	return math.Float64frombits(r.endian.Uint64(data))
}

func (r *Reader) Float64Array(len int) []float64 {
	if r.err != nil || len == 0 {
		return nil
	}

	data := make([]byte, 8*len)
	_, r.err = io.ReadFull(r.in, data)
	if r.err != nil {
		return nil
	}

	arr := make([]float64, len)
	for i := 0; i < len; i++ {
		arr[i] = math.Float64frombits(r.endian.Uint64(data[i*8:]))
	}

	return arr
}

func (r *Reader) Float32() float32 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 4)
	_, r.err = io.ReadFull(r.in, data)
	return math.Float32frombits(r.endian.Uint32(data))
}

func (r *Reader) Float32Array(len int) []float32 {
	if r.err != nil || len == 0 {
		return nil
	}
	data := make([]byte, 4*len)
	_, r.err = io.ReadFull(r.in, data)
	if r.err != nil {
		return nil
	}

	arr := make([]float32, len)
	for i := 0; i < len; i++ {
		arr[i] = math.Float32frombits(r.endian.Uint32(data[i*4:]))
	}
	return arr
}

func (r *Reader) Int64() int64 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 8)
	_, r.err = io.ReadFull(r.in, data)
	i := r.endian.Uint64(data)
	return int64(i)
}

func (r *Reader) Int32() int32 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 4)
	_, r.err = io.ReadFull(r.in, data)
	i := r.endian.Uint32(data)
	return int32(i)
}

func (r *Reader) Int16() int16 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 2)
	_, r.err = io.ReadFull(r.in, data)
	i := r.endian.Uint16(data)
	return int16(i)
}

func (r *Reader) UInt64() uint64 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 8)
	_, r.err = io.ReadFull(r.in, data)
	return r.endian.Uint64(data)
}

func (r *Reader) UInt32() uint32 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 4)
	_, r.err = io.ReadFull(r.in, data)
	return r.endian.Uint32(data)
}

func (r *Reader) UInt16() uint16 {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 2)
	_, r.err = io.ReadFull(r.in, data)
	return r.endian.Uint16(data)
}

func (r *Reader) Byte() byte {
	if r.err != nil {
		return 0
	}
	data := make([]byte, 1)
	_, r.err = io.ReadFull(r.in, data)
	return data[0]
}

func (r *Reader) ByteArray(len int) []byte {
	if r.err != nil || len == 0 {
		return nil
	}
	data := make([]byte, len)
	_, r.err = io.ReadFull(r.in, data)
	return data
}

func (r *Reader) UVarInt() (i uint64) {
	if r.err != nil {
		return 0
	}
	i, r.err = binary.ReadUvarint(r)
	return
}

func (r *Reader) VarInt() (i int64) {
	if r.err != nil {
		return 0
	}
	i, r.err = binary.ReadVarint(r)
	return
}

// Implementing the io.ByteReader interface
func (r *Reader) ReadByte() (byte, error) {
	return r.Byte(), r.err
}
