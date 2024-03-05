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
}

func NewWriter(in io.Writer, byteOrder binary.ByteOrder) *Writer {
	return &Writer{
		out:    in,
		endian: byteOrder,
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
	bytes := make([]byte, 8)
	w.endian.PutUint64(bytes, bits)
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Float32(f float32) error {
	if w.err != nil {
		return w.err
	}

	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	w.endian.PutUint32(bytes, bits)
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Int64(i int64) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 8)
	w.endian.PutUint64(bytes, uint64(i))
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Int32(i int32) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 4)
	w.endian.PutUint32(bytes, uint32(i))
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Int16(i int16) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 2)
	w.endian.PutUint16(bytes, uint16(i))
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) UInt64(i uint64) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 8)
	w.endian.PutUint64(bytes, i)
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) UInt32(i uint32) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 4)
	w.endian.PutUint32(bytes, i)
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) UInt16(i uint16) error {
	if w.err != nil {
		return w.err
	}

	bytes := make([]byte, 2)
	w.endian.PutUint16(bytes, i)
	_, w.err = w.out.Write(bytes)
	return w.err
}

func (w *Writer) Byte(i byte) error {
	if w.err != nil {
		return w.err
	}

	_, w.err = w.out.Write([]byte{i})
	return w.err
}

func (w *Writer) UVarInt(i uint64) error {
	if w.err != nil {
		return w.err
	}
	bytes := make([]byte, 8)
	c := binary.PutUvarint(bytes, i)
	_, w.err = w.out.Write(bytes[:c])
	return w.err
}

func (w *Writer) VarInt(i int64) error {
	if w.err != nil {
		return w.err
	}
	bytes := make([]byte, 8)
	c := binary.PutVarint(bytes, i)
	_, w.err = w.out.Write(bytes[:c])
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
