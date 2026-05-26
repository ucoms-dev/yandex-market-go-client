package openapi

import (
	"bytes"
	"mime/multipart"
	"os"
	"strings"
	"testing"
)

func TestAddFileCopiesOpenFileIntoMultipartPart(t *testing.T) {
	tmpFile := t.TempDir() + "/upload.txt"
	const fileBody = "multipart payload"
	if err := os.WriteFile(tmpFile, []byte(fileBody), 0600); err != nil {
		t.Fatalf("write temp upload file: %v", err)
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if err := addFile(writer, "file", tmpFile); err != nil {
		t.Fatalf("addFile returned error: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	if !strings.Contains(body.String(), fileBody) {
		t.Fatalf("multipart body does not contain uploaded file contents: %q", body.String())
	}
}
