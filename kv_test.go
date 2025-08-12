package main

import (
	"os"
	"testing"
)

func testKV(t *testing.T) {
	tmpFile := "test.json"
	defer os.Remove(tmpFile)

	store := NewKVStore(tmpFile)

	store.Set("foo", "bar")
	val, ok := store.Get("foo")
	if !ok || val != "bar" {
		t.Errorf("expected 'bar', got '%s'", val)
	}

	store.Delete("foo")
	_, ok = store.Get("foo")
	if ok {
		t.Errorf("expected key to be deleted")
	}

	store.Set("hello", "world")
	store2 := NewKVStore(tmpFile)
	val, ok = store2.Get("hello")
	if !ok || val != "world" {
		t.Errorf("expected 'world', got '%s'", val)
	}

}
