package main

import (
	"io"
)

type Serializer struct {
	writer io.Writer
}

func NewSerializer(w io.Writer) *Serializer {
	return &Serializer{writer: w}
}

func (v Value) Serialize() []byte {
	switch v.typ {
	case "array":
		return v.serializeString()
	case "bulk":
		return []byte{}
	case "string":
		return []byte{}
	case "null":
		return []byte{}
	case "error":
		return []byte{}
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
