package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

func main() {
	data := "This is a simple example of compressing and decompressing data using gzip in Go!"

	// Compress the data
	compressedData, err := compressData(data)
	if err != nil {
		log.Fatal("Error compressing data:", err)
	}

	// Print compressed data (in bytes)
	fmt.Println("Compressed Data:", compressedData)

	// Decompress the data
	decompressedData, err := decompressData(compressedData)
	if err != nil {
		log.Fatal("Error decompressing data:", err)
	}

	// Print decompressed data
	fmt.Println("Decompressed Data:", decompressedData)
}

func compressData(text string) ([]byte, error) {
	var buffer bytes.Buffer

	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write([]byte(text))
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Return the compressed data as a byte slice
	return buffer.Bytes(), nil
}

func decompressData(data []byte) (string, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	text, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(text), nil
}
