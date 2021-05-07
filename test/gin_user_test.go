package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/pkg/jsontool"
	v1 "github.com/Arapgp/Arapgp-Server-go/route/api/v1"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Gin_Register(t *testing.T) {
	setupConfig(t)
	r := setupRouter(t)
	defer teardownConfig(t)

	const url = "http://127.0.0.1:3000/api/v1/signup"

	tcases := []struct {
		name   string
		method string
		body   io.Reader
		code   int
		status string
	}{
		{
			name:   "Gin_Register_1",
			method: "POST",
			body:   bytes.NewBuffer([]byte(`{"username": "ljg", "password": "ljg"}`)),
			code:   200,
			status: "OK",
		},
		{
			name:   "Gin_Register_2",
			method: "POST",
			body:   bytes.NewBuffer([]byte(`{"username": "ljg", "password": "ljg"}`)),
			code:   200,
			status: "Username already exists!",
		},
		{
			name:   "Gin_Register_3",
			method: "GET",
			body:   bytes.NewBuffer([]byte(`{"username": "gjl", "password": "gjl"}`)),
			code:   404,
			status: "",
		},
		{
			name:   "Gin_Register_4",
			method: "POST",
			body:   bytes.NewBuffer([]byte(`{"username": "gjl", "password": "gjl"}`)),
			code:   200,
			status: "OK",
		},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(c.method, url, c.body)

			r.ServeHTTP(w, req)

			resp := v1.JSONStatus{}
			jsontool.GetJSON(w.Body, &resp)
			log.WithFields(log.Fields{
				"ecode": c.code, "acode": w.Code,
				"estatus": c.status, "astatus": resp.Status},
			).Info("Add log.Info in order to ensure assert.Equal work well.")
			assert.Equal(t, c.code, w.Code)
			assert.Equal(t, c.status, resp.Status)
		})
	}
}
