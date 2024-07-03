package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "picturetest"
	pathKey := CASPathTransformFunc(key)
	fmt.Print(pathKey)
	expectedPathname := "ac353/0cf16/67425/e4008/d4dba/6ad90/e55c8/f8d37"
	expectedOriginalKey := "ac3530cf1667425e4008d4dba6ad90e55c8f8d37"
	if pathKey.PathName != expectedPathname {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathname)
	}
	if pathKey.Original != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
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
