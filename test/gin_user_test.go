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

	var url = urlBase + "/api/v1/signup"

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
				"estatus": c.status, "astatus": resp.Status,
			}).Info("Add log.Info in order to ensure assert.Equal work well.")
			assert.Equal(t, c.code, w.Code)
			assert.Equal(t, c.status, resp.Status)
		})
	}
}

func Test_Gin_Login(t *testing.T) {
	setupConfig(t)
	r := setupRouter(t)
	defer teardownConfig(t)

	const method = "POST"
	var url = urlBase + "/api/v1/login"
	tcases := []struct {
		name   string
		body   *bytes.Buffer
		code   int
		status string
	}{
		{
			name:   "Gin_Login_1",
			body:   bytes.NewBuffer([]byte(`{"username": "ljg", "password": "gjl"}`)),
			code:   200,
			status: "Username or password wrong!",
		},
		{
			name:   "Gin_Login_2",
			body:   bytes.NewBuffer([]byte(`{"username": "gjl", "password": "gjl"}`)),
			code:   200,
			status: "OK",
		},
		{
			name:   "Gin_Login_3",
			body:   bytes.NewBuffer([]byte(`{"username": "g", "password": "??"}`)),
			code:   200,
			status: "User not existed!",
		},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(method, url, c.body)

			r.ServeHTTP(w, req)

			resp := v1.JSONStatus{}
			jsontool.GetJSON(w.Body, &resp)
			log.WithFields(log.Fields{
				"ecode": c.code, "acode": w.Code,
				"estatus": c.status, "astatus": resp.Status,
			}).Info("Add log.Info in order to ensure assert.Equal work well.")
			assert.Equal(t, c.code, w.Code)
			assert.Equal(t, c.status, resp.Status)
		})
	}
}

func Test_Gin_GetUser(t *testing.T) {
	setupConfig(t)
	r := setupRouter(t)
	defer teardownConfig(t)

	const method = "GET"
	var url = urlBase + "/api/v1/user"
	tcases := []struct {
		name   string
		param  string
		code   int
		status string
	}{
		{
			name:   "Gin_GetUser_1",
			param:  "?query=ljg",
			code:   200,
			status: "OK",
		},
		{
			name:   "Gin_GetUser_2",
			param:  "?query=g",
			code:   200,
			status: "OK",
		},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(method, url+c.param, nil)

			r.ServeHTTP(w, req)

			resp := v1.JSONGetUser{}
			jsontool.GetJSON(w.Body, &resp)
			log.WithFields(log.Fields{
				"ecode": c.code, "acode": w.Code,
				"estatus": c.status, "astatus": resp.Status,
				"user": resp.UserList, "url": url + c.param,
			}).Info("Add log.Info in order to ensure assert.Equal work well.")
			assert.Equal(t, c.code, w.Code)
			assert.Equal(t, c.status, resp.Status)
		})
	}
}
