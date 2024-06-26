package main

import (
	"io"
	"strconv"
)

type Writer struct {
	writer io.Writer
}

func (w *Writer) Write(v Value) error {
	var bytes = v.Serialize()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (v Value) Serialize() []byte {
	switch v.typ {
	case "array":
		return v.serializeArray()
	case "bulk":
		return v.serializeBulk()
	case "string":
		return v.serializeString()
	case "null":
		return v.serializeNull()
	case "error":
		return v.serializeError()
	default:
		return []byte{}
	}
}

func (v Value) serializeString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) serializeBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) serializeArray() []byte {
	len := len(v.array)
	var bytes []byte
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.array[i].Serialize()...)
	}

	return bytes
}

func (v Value) serializeError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) serializeNull() []byte {
	return []byte("$-1\r\n")
}
