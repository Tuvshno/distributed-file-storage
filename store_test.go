package main

import (
	"bytes"
	"testing"
)

func TestNewStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("picture", data); err != nil {
		t.Error(err)
	}
}
