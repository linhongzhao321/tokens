package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BaseAssert(t *testing.T, method string, url string, reqBody io.Reader) (body map[string]interface{}) {
	router := NewRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, reqBody)
	assert.Nil(t, err)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Greater(t, 0, w.Body.Len())
	err = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.NotNil(t, body["code"])
	assert.IsType(t, 0, body["msg"])
	assert.NotNil(t, body["msg"])
	assert.IsType(t, "", body["msg"])
	assert.NotNil(t, body["data"])
	assert.IsType(t, []interface{}{}, body["data"])
	return
}

func TestGetPing(t *testing.T) {
	body := BaseAssert(t, "GET", "/ping", nil)
	assert.Equal(t, PING.code, body["code"])
	assert.Equal(t, PING.msg, body["code"])
}

func BenchmarkGetPing(b *testing.B) {

	router := NewRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, req)
	}
}
