package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/octavore/delta/browser"
)

func getAsset(path string) string {
	a, err := browser.Asset(path)
	if err != nil {
		panic(err)
	}

	gz, err := gzip.NewReader(bytes.NewBuffer(a))
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	io.Copy(buf, gz)

	return buf.String()
}

func md5sum(s string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
