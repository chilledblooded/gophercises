package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chilledblooded/gophercises/Exercise_18/primitive"
	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

func TestWelcome(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	r, _ := http.NewRequest("GET", srv.URL, nil)
	res, _ := client.Do(r)
	if res.StatusCode != http.StatusOK {
		t.Error("Expected status ok but got different status")
	}
}

func TestUpload(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/ghoper.jpg")
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := http.DefaultClient.Do(r)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestCreateTempFile(t *testing.T) {
	_, err := createTempFile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestModifyMode(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.jpg?mode=3", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestModifyModeNegative(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.jpg?mode=a", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusOK {
		t.Errorf("Expected Bad request but got different status %v", res.Status)
	}
}

func TestModifyModeNegativeExt(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.txt?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusOK {
		t.Errorf("Expected status internal server error but got different status %v", res.Status)
	}
}

func TestModifyModeShapes(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.jpg?mode=3&n=5", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

// func TestRenderModeChoices(t *testing.T) {
// 	req, err := http.NewRequest("GET", "localhost:8888", nil)
// 	if err != nil {
// 		t.Fatalf("could not create request: %v", err)
// 	}
// 	rs := bytes.NewReader(nil)
// 	rec := httptest.NewRecorder()
// 	renderModeChoices(rec, req, rs, "txt")
// 	res := rec.Result()
// 	if res.StatusCode != http.StatusInternalServerError {
// 		t.Errorf("Expected status Internal server error but got %v", res.Status)
// 	}
// }

func TestGenImage(t *testing.T) {
	rs := bytes.NewReader(nil)
	mode := primitive.ModeCombo
	_, err := genImage(rs, "txt", -1, mode)
	if err == nil {
		t.Error("Expected error but no error")
	}
}
