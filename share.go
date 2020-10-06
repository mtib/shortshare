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
	Decoder struct {
		reader io.Reader
	}
)

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

func ShareString(data interface{}) (string, error) {
	r, err := Share(data)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func Unshare(data io.Reader, target interface{}) error {
	uncmpress, err := zlib.NewReader(base64.NewDecoder(base64.URLEncoding, data))
	if err != nil {
		return err
	}
	enc := json.NewDecoder(uncmpress)
	enc.Decode(target)
	return nil
}

func UnshareString(data string, target interface{}) error {
	return Unshare(strings.NewReader(data), target)
}

func NewDecoder(reader io.Reader) Decoder {
	return Decoder{
		reader,
	}
}

func (d Decoder) Decode(target interface{}) error {
	return Unshare(d.reader, target)
}
