package bitlib_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/EliCDavis/bitlib"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	A bool
	B int32
}

func TestReadWriteArray(t *testing.T) {
	// ARRANGE ================================================================
	var f64Arr []float64 = []float64{1, 2}
	var f32Arr []float32 = []float32{1, 2}
	buf := bytes.Buffer{}

	// ACT ====================================================================
	writer := bitlib.NewWriter(&buf, binary.LittleEndian)
	assert.NoError(t, bitlib.WriteArray(writer, f64Arr))
	assert.NoError(t, bitlib.WriteArray(writer, f32Arr))

	reader := bitlib.NewReader(bytes.NewBuffer(buf.Bytes()), binary.LittleEndian)
	f64Back, _ := bitlib.ReadArray[float64](reader, len(f64Arr))
	f32Back, _ := bitlib.ReadArray[float32](reader, len(f32Arr))

	// ASSERT =================================================================
	assert.Equal(t, f64Arr, f64Back)
	assert.Equal(t, f32Arr, f32Back)
	assert.NoError(t, reader.Error())
	assert.NoError(t, writer.Error())
}

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
	var bVal2 byte = 222
	var bArr []byte = []byte{1, 2, 3, 4, 5}
	var f64Arr []float64 = []float64{1, 2, 3, 4, 5}
	var f32Arr []float32 = []float32{1, 2, 3, 4, 5}
	var i32Arr []int32 = []int32{-1, -2, -3, -4, -5}
	var u32Arr []uint32 = []uint32{1, 2, 3, 4, 5}
	var bGenArr []byte = []byte{5, 4, 3, 2, 1}
	var structVal TestStruct = TestStruct{
		A: true,
		B: 32,
	}
	var strVal = "test string!!!"

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
	writer.WriteByte(bVal2)
	writer.ByteArray(bArr)
	writer.Float64Array(f64Arr)
	writer.Float32Array(f32Arr)
	writer.Int32Array(i32Arr)
	writer.Uint32Array(u32Arr)
	writer.Any(structVal)
	writer.WriteString(strVal)
	writeCount, writeErr := writer.Write(bGenArr)
	assert.NoError(t, writeErr)

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
	readbVal2, _ := reader.ReadByte()
	readbArrVal := reader.ByteArray(len(bArr))
	readf64ArrVal := reader.Float64Array(len(f64Arr))
	readf32ArrVal := reader.Float32Array(len(f32Arr))
	readi32ArrVal := reader.Int32Array(len(i32Arr))
	readu32ArrVal := reader.Uint32Array(len(u32Arr))
	var readStructVal TestStruct
	reader.Any(&readStructVal)
	readStr := reader.String(len(strVal))
	readByesArr := make([]byte, len(bGenArr))
	readCount, readErr := reader.Read(readByesArr)
	assert.NoError(t, readErr)

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
	assert.Equal(t, bVal2, readbVal2)
	assert.Equal(t, bArr, readbArrVal)
	assert.Equal(t, f64Arr, readf64ArrVal)
	assert.Equal(t, f32Arr, readf32ArrVal)
	assert.Equal(t, i32Arr, readi32ArrVal)
	assert.Equal(t, u32Arr, readu32ArrVal)
	assert.Equal(t, structVal, readStructVal)
	assert.Equal(t, strVal, readStr)
	assert.Equal(t, writeCount, readCount)
	assert.Equal(t, len(bGenArr), readCount)
	assert.Equal(t, readByesArr, bGenArr)

	// Reading arrays with length 0 returns nil
	assert.Equal(t, "", reader.String(0))
	assert.Equal(t, []byte(nil), reader.ByteArray(0))
	assert.Equal(t, []uint32(nil), reader.Uint32Array(0))
	assert.Equal(t, []int32(nil), reader.Int32Array(0))
	assert.Equal(t, []float32(nil), reader.Float32Array(0))
	assert.Equal(t, []float64(nil), reader.Float64Array(0))
}
