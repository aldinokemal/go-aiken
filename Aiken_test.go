package Aiken

import (
	aiken "github.com/aldinokemal/go-aiken"
	"testing"
)

func TestReadAiken(t *testing.T) {
	result, err := aiken.ReadAiken("aiken-example.txt")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
	}
}
