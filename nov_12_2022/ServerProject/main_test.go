package main

import (
	"net/http"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

// test function

func TestPushHandler(t *testing.T) {

	r, _ := http.NewRequest("GET", "/push", nil)

	w := httptest.NewRecorder()

	PushHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPopHandler(t *testing.T) {

	r, _ := http.NewRequest("GET", "/pop", nil)

	w := httptest.NewRecorder()

	PopHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}
