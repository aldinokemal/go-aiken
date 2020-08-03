package main

import "testing"

func TestReadAiken(t *testing.T)  {
	result, err := handleAiken("aiken-example.txt")
	if err != nil {
		t.Error(err.Error())
	}else{
		t.Log(result)
	}
}
