# Bitlib

![Coverage](https://img.shields.io/badge/Coverage-82.3%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/bitlib)](https://goreportcard.com/report/github.com/EliCDavis/bitlib)

Utilities for reading and writing binary data that for some reason I keep re-writing over and over.

Reader Implements:

* io.Reader
* io.ByteReader

Writer Implements:

* io.Writer
* io.ByteWriter
* io.StringWriter

## Example

```go
package bitlib_test

import (
	"bytes"
	"encoding/binary"
	"log"
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
	log.Println(reader.Float64())
	log.Println(reader.Float32())
	log.Println(reader.Int64())
	log.Println(reader.Int32())
	log.Println(reader.Int16())
	log.Println(reader.UInt64())
	log.Println(reader.UInt32())
	log.Println(reader.UInt16())
	log.Println(reader.VarInt())
	log.Println(reader.UVarInt())
	log.Println(reader.Byte())
	log.Println(reader.ByteArray(len(bArr)))
	log.Println(reader.Float64Array(len(f64Arr)))
	log.Println(reader.Float32Array(len(f32Arr)))
}
```
