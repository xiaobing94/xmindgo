package xmindgo

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func readFile(file *zip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, rc); err != nil {
		return nil, err
	}
	if err := rc.Close(); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}


func CheckSum(content []byte) (string, error){
	h := sha256.New()
	if _, err := h.Write(content); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}