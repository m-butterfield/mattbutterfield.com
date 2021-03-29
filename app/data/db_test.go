package data

import (
	"testing"
)

func TestConnect(t *testing.T) {
	result, err := Connect()
	if err != nil {
		t.Error(err)
	}
	_ = result.(*dbStore)
}
