package bitlib_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/EliCDavis/bitlib"
	"github.com/stretchr/testify/assert"
)

func TestReadWrite(t *testing.T) {
	buf := bytes.Buffer{}
	var f64Val float64 = 1.23456789
	var f32Val float32 = 1.23456789
	var i64Val int64 = -12334567890000
	var i32Val int32 = -123345678
	var i16Val int16 = -12334
	var u64Val uint64 = 12334567890000
	var u32Val uint32 = 123345678
	var u16Val uint16 = 12334
	var bVal byte = 14
	var bArr []byte = []byte{1, 2, 3, 4, 5}
	var f64Arr []float64 = []float64{1, 2, 3, 4, 5}
	var f32Arr []float32 = []float32{1, 2, 3, 4, 5}

	// Write
	writer := bitlib.NewWriter(&buf, binary.LittleEndian)
	writer.Float64(f64Val)
	writer.Float32(f32Val)
	writer.Int64(i64Val)
	writer.Int32(i32Val)
	writer.Int16(i16Val)
	writer.UInt64(u64Val)
	writer.UInt32(u32Val)
	writer.UInt16(u16Val)
	writer.VarInt(i64Val)
	writer.UVarInt(u64Val)
	writer.Byte(bVal)
	writer.ByteArray(bArr)
	writer.Float64Array(f64Arr)
	writer.Float32Array(f32Arr)

	// Read
	reader := bitlib.NewReader(bytes.NewBuffer(buf.Bytes()), binary.LittleEndian)
	readf64Val := reader.Float64()
	readf32Val := reader.Float32()
	readi64Val := reader.Int64()
	readi32Val := reader.Int32()
	readi16Val := reader.Int16()
	readu64Val := reader.UInt64()
	readu32Val := reader.UInt32()
	readu16Val := reader.UInt16()
	readivar64Val := reader.VarInt()
	readuvar64Val := reader.UVarInt()
	readbVal := reader.Byte()
	readbArrVal := reader.ByteArray(len(bArr))
	readf64ArrVal := reader.Float64Array(len(bArr))
	readf32ArrVal := reader.Float32Array(len(bArr))

	assert.NoError(t, writer.Error())
	assert.NoError(t, reader.Error())
	assert.Equal(t, f64Val, readf64Val)
	assert.Equal(t, f32Val, readf32Val)
	assert.Equal(t, i64Val, readi64Val)
	assert.Equal(t, i32Val, readi32Val)
	assert.Equal(t, i16Val, readi16Val)
	assert.Equal(t, u64Val, readu64Val)
	assert.Equal(t, u32Val, readu32Val)
	assert.Equal(t, u16Val, readu16Val)
	assert.Equal(t, i64Val, readivar64Val)
	assert.Equal(t, u64Val, readuvar64Val)
	assert.Equal(t, bVal, readbVal)
	assert.Equal(t, bArr, readbArrVal)
	assert.Equal(t, f64Arr, readf64ArrVal)
	assert.Equal(t, f32Arr, readf32ArrVal)
}
