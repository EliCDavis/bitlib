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
	buf    []byte
}

func NewReader(in io.Reader, byteOrder binary.ByteOrder) *Reader {
	return &Reader{
		in:     in,
		endian: byteOrder,
		buf:    make([]byte, 8),
	}
}

func (r Reader) Error() error {
	return r.err
}

func (r *Reader) Float64() float64 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf)
	return math.Float64frombits(r.endian.Uint64(r.buf))
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
	_, r.err = io.ReadFull(r.in, r.buf[:4])
	return math.Float32frombits(r.endian.Uint32(r.buf))
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
	_, r.err = io.ReadFull(r.in, r.buf)
	i := r.endian.Uint64(r.buf)
	return int64(i)
}

func (r *Reader) Int32() int32 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf[:4])
	i := r.endian.Uint32(r.buf)
	return int32(i)
}

func (r *Reader) Int32Array(len int) []int32 {
	if r.err != nil || len == 0 {
		return nil
	}
	data := make([]byte, 4*len)
	_, r.err = io.ReadFull(r.in, data)
	if r.err != nil {
		return nil
	}

	arr := make([]int32, len)
	for i := 0; i < len; i++ {
		arr[i] = int32(r.endian.Uint32(data[i*4:]))
	}
	return arr
}

func (r *Reader) Uint32Array(len int) []uint32 {
	if r.err != nil || len == 0 {
		return nil
	}
	data := make([]byte, 4*len)
	_, r.err = io.ReadFull(r.in, data)
	if r.err != nil {
		return nil
	}

	arr := make([]uint32, len)
	for i := 0; i < len; i++ {
		arr[i] = r.endian.Uint32(data[i*4:])
	}
	return arr
}

func (r *Reader) Int16() int16 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf[:2])
	i := r.endian.Uint16(r.buf)
	return int16(i)
}

func (r *Reader) UInt64() uint64 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf)
	return r.endian.Uint64(r.buf)
}

func (r *Reader) UInt32() uint32 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf[:4])
	return r.endian.Uint32(r.buf)
}

func (r *Reader) UInt16() uint16 {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf[:2])
	return r.endian.Uint16(r.buf)
}

func (r *Reader) Byte() byte {
	if r.err != nil {
		return 0
	}
	_, r.err = io.ReadFull(r.in, r.buf[:1])
	return r.buf[0]
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

// Reads data and sets the value a passed into this function.
// Function must recieve a pointer to the value to be set.
func (r *Reader) Any(a any) error {
	if r.err != nil {
		return r.err
	}

	r.err = binary.Read(r, r.endian, a)
	return r.err
}

func (r *Reader) String(len int) string {
	if r.err != nil || len == 0 {
		return ""
	}
	data := make([]byte, len)
	_, r.err = io.ReadFull(r.in, data)
	return string(data)
}

// Implementing the io.ByteReader interface
func (r *Reader) ReadByte() (byte, error) {
	return r.Byte(), r.err
}

// Implementing the io.Reader interface
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.err != nil {
		return 0, r.err
	}

	n, err = io.ReadFull(r.in, p)
	r.err = err
	return
}

func Read[T any](reader *Reader) (T, error) {
	var val T
	return val, binary.Read(reader, reader.endian, &val)
}

func ReadArray[T any](reader *Reader, length int) ([]T, error) {
	data := make([]T, length)
	return data, binary.Read(reader, reader.endian, data)
}
