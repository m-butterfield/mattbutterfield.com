package app

import (
	"net/http"
	"testing"
)

func TestAppRuns(t *testing.T) {
	listenAndServeCalled := false
	listenAndServe := func(addr string, handler http.Handler) error {
		listenAndServeCalled = true
		if addr != ":8000" {
			t.Errorf("Addr: \"%s\" != \":8000\"", addr)
		}
		return nil
	}
	if err := Run(listenAndServe, "8000"); err != nil {
		t.Error(err)
	}
	if !listenAndServeCalled {
		t.Error("listenAndServe not called")
	}
}
