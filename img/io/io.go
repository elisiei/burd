// modified version by elisiei, original by anthonynsimon.
// this modification targets a more byte-oriented approach.
// Package io provides basic image encoding/decoding from/to bytes.
package io

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

// Encoder encodes the provided image and writes it.
type Encoder func(io.Writer, image.Image) error

func FetchBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Decode decodes an image from raw bytes.
func Decode(data []byte) (image.Image, error) {
	r := bytes.NewReader(data)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// JPEGEncoder returns an encoder to JPEG given the argument 'quality'.
func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

// PNGEncoder returns an encoder to PNG.
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}

// Encode encodes an image into bytes using the provided encoder.
func Encode(img image.Image, encoder Encoder) ([]byte, error) {
	var buf bytes.Buffer
	if err := encoder(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
