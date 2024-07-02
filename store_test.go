package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "picturetest"
	pathName := CASPathTransformFunc(key)
	fmt.Print(pathName)
	expectedPathName := "ac353/0cf16/67425/e4008/d4dba/6ad90/e55c8/f8d37"
	if pathName != expectedPathName {
		t.Errorf("have %s want %s", pathName, expectedPathName)

	}
}

func TestNewStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("picture", data); err != nil {
		t.Error(err)
	}
}
