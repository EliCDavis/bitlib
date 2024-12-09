package bitlib

import (
	"encoding/binary"
	"io"
	"math"
)

type Writer struct {
	out    io.Writer
	err    error
	endian binary.ByteOrder
	buf    []byte
}

func NewWriter(in io.Writer, byteOrder binary.ByteOrder) *Writer {
	return &Writer{
		out:    in,
		endian: byteOrder,
		buf:    make([]byte, 8),
	}
}

func (w Writer) Error() error {
	return w.err
}

func (w *Writer) Float64(f float64) error {
	if w.err != nil {
		return w.err
	}

	bits := math.Float64bits(f)
	w.endian.PutUint64(w.buf, bits)
	_, w.err = w.out.Write(w.buf)
	return w.err
}

func (w *Writer) Float32(f float32) error {
	if w.err != nil {
		return w.err
	}

	bits := math.Float32bits(f)
	w.endian.PutUint32(w.buf, bits)
	_, w.err = w.out.Write(w.buf[:4])
	return w.err
}

func (w *Writer) Int64(i int64) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint64(w.buf, uint64(i))
	_, w.err = w.out.Write(w.buf)
	return w.err
}

func (w *Writer) Int32(i int32) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint32(w.buf, uint32(i))
	_, w.err = w.out.Write(w.buf[:4])
	return w.err
}

func (w *Writer) Int16(i int16) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint16(w.buf, uint16(i))
	_, w.err = w.out.Write(w.buf[:2])
	return w.err
}

func (w *Writer) UInt64(i uint64) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint64(w.buf, i)
	_, w.err = w.out.Write(w.buf)
	return w.err
}

func (w *Writer) UInt32(i uint32) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint32(w.buf, i)
	_, w.err = w.out.Write(w.buf[:4])
	return w.err
}

func (w *Writer) UInt16(i uint16) error {
	if w.err != nil {
		return w.err
	}

	w.endian.PutUint16(w.buf, i)
	_, w.err = w.out.Write(w.buf[:2])
	return w.err
}

func (w *Writer) Byte(i byte) error {
	if w.err != nil {
		return w.err
	}

	w.buf[0] = i
	_, w.err = w.out.Write(w.buf[:1])
	return w.err
}

func (w *Writer) UVarInt(i uint64) error {
	if w.err != nil {
		return w.err
	}
	c := binary.PutUvarint(w.buf, i)
	_, w.err = w.out.Write(w.buf[:c])
	return w.err
}

func (w *Writer) VarInt(i int64) error {
	if w.err != nil {
		return w.err
	}
	c := binary.PutVarint(w.buf, i)
	_, w.err = w.out.Write(w.buf[:c])
	return w.err
}

func (w *Writer) ByteArray(b []byte) error {
	if w.err != nil {
		return w.err
	}
	_, w.err = w.out.Write(b)
	return w.err
}

func (w *Writer) Float64Array(fArr []float64) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 8*len(fArr))
	for i := 0; i < len(fArr); i++ {
		bits := math.Float64bits(fArr[i])
		w.endian.PutUint64(bytes[8*i:], bits)
	}

	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Float32Array(fArr []float32) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 4*len(fArr))
	for i := 0; i < len(fArr); i++ {
		bits := math.Float32bits(fArr[i])
		w.endian.PutUint32(bytes[4*i:], bits)
	}

	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Int32Array(iArr []int32) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 4*len(iArr))
	for i := 0; i < len(iArr); i++ {
		w.endian.PutUint32(bytes[4*i:], uint32(iArr[i]))
	}

	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Uint32Array(iArr []uint32) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 4*len(iArr))
	for i := 0; i < len(iArr); i++ {
		w.endian.PutUint32(bytes[4*i:], iArr[i])
	}

	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Any(a any) error {
	if w.err != nil {
		return w.err
	}

	w.err = binary.Write(w, w.endian, a)
	return w.err
}

// Implementing the io.Writer interface
func (w *Writer) Write(p []byte) (n int, err error) {
	if w.err != nil {
		return 0, w.err
	}

	n, err = w.out.Write(p)
	w.err = err
	return n, err
}

// Implementing io.ByteWriter
func (w *Writer) WriteByte(c byte) error {
	if w.err != nil {
		return w.err
	}

	_, w.err = w.out.Write([]byte{c})
	return w.err
}

// Implementing io.StringWriter
func (w *Writer) WriteString(s string) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	var n int
	n, w.err = w.out.Write([]byte(s))
	return n, w.err
}

func WriteArray[T any](writer *Writer, data []T) error {
	return binary.Write(writer, writer.endian, data)
}

func Write[T any](writer *Writer, data T) error {
	return binary.Write(writer, writer.endian, data)
}
