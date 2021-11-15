package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	expect "github.com/stretchr/testify/require"
)

func TestEndpoint_WASM(t *testing.T) {
	r := gin.Default()
	setupRoutes(r)

	form := url.Values{
		"snippet": []string{"fmt.Println(\"hello world\")"},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/wasm", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)

	expect.Equal(t, http.StatusOK, w.Code)
	expect.Equal(t, "application/wasm", w.Header().Get("Content-Type"))
}
