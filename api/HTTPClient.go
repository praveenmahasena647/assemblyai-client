package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/praveenmahasena647/client/internal/helpers"
)

func StartClient() error {
	var file, fileErr = helpers.GetFile()

	if fileErr != nil {
		return fileErr
	}

	var HTTPReq, HTTPErr = http.Post("http://localhost:42069", "", bytes.NewReader(file))

	if HTTPErr != nil {
		return HTTPErr
	}

	defer HTTPReq.Body.Close()

	var stringBytes, strErr = io.ReadAll(HTTPReq.Body)

	if strErr != nil {
		return strErr
	}

	println(string(stringBytes))

	return nil
}
