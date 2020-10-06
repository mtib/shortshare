package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"
)

type (
	// Decoder only provided in case you prefer this calling scheme to the imperative Share()
	Decoder struct {
		reader io.Reader
	}
)

// Share converts the provided data into a sharable (base64 urlencoded) bytes.Buffer, it uses the datas json representation as an intermediary.
func Share(data interface{}) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	cmpress, err := zlib.NewWriterLevel(base64.NewEncoder(base64.URLEncoding, b), zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	enc := json.NewEncoder(cmpress)
	enc.Encode(data)
	cmpress.Close()
	return b, nil
}

// ShareString does the same as Share but converts the bytes.Buffer into a string.
func ShareString(data interface{}) (string, error) {
	r, err := Share(data)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

// Unshare unmarshals the data into the target.
func Unshare(data io.Reader, target interface{}) error {
	uncmpress, err := zlib.NewReader(base64.NewDecoder(base64.URLEncoding, data))
	if err != nil {
		return err
	}
	enc := json.NewDecoder(uncmpress)
	enc.Decode(target)
	return nil
}

// UnshareString unmarshals the data string into the target.
func UnshareString(data string, target interface{}) error {
	return Unshare(strings.NewReader(data), target)
}

// NewDecoder creates a new Decoder which reads from the io.Reader and offers a Decode(target) method.
func NewDecoder(reader io.Reader) Decoder {
	return Decoder{
		reader,
	}
}

// Decode decodes the data the reader this Decoder was created with has into the target.
func (d Decoder) Decode(target interface{}) error {
	return Unshare(d.reader, target)
}
