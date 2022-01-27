package golpe

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/mholt/archiver/v3"
)

// Base64Decode decode a base64 encoded string (to []byte)
func Base64Decode(text string) []byte {
	dec, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		log.Printf("Base64Decode: %v", err)
		return nil
	}
	return dec
}

// ExtractFileFromString base64 decode, and then decompress using bzip2
func ExtractFileFromString(data string) (out []byte, err error) {
	// decode
	decoded := Base64Decode(data)
	if len(decoded) == 0 {
		return nil, fmt.Errorf("ExtractFileFromString: Failed to decode")
	}

	// decompress
	bz2 := &archiver.Bz2{}
	var decompress_buf bytes.Buffer
	reader := bytes.NewReader(decoded)
	err = bz2.Decompress(reader, &decompress_buf)
	if err != nil {
		return nil, fmt.Errorf("ExtractFileFromString: Decompress: %v", err)
	}

	out = decompress_buf.Bytes()
	return
}
