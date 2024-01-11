package utils

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
)

// GenerateQRCode generates a QR code and returns the image buffer
func GenerateQRCode(data string) (*bytes.Buffer, error) {
	qrCode, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrCode, err = barcode.Scale(qrCode, 300, 300)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = png.Encode(&buffer, qrCode)
	if err != nil {
		return nil, err
	}

	return &buffer, nil
}
