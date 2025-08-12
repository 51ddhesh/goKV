package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type KV struct {
	data map[string]string
	mu   sync.RWMutex
	file string
}

func NewKVStore(filename string) *KV {
	store := &KV{
		data: make(map[string]string),
		file: filename,
	}

	store.load()
	return store
}

func (kv *KV) Set(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = value
	kv.save()
}

func (kv *KV) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, exists := kv.data[key]
	return val, exists
}

func (kv *KV) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	delete(kv.data, key)
	kv.save()
}

func (kv *KV) save() {
	if kv.file == "" {
		return
	}

	data, err := json.MarshalIndent(kv.data, "", " ")
	if err != nil {
		fmt.Println("Failed to marshal data: ", err)
		return
	}
	if err := os.WriteFile(kv.file, data, 0644); err != nil {
		fmt.Println("Failed to write file: ", err)
	}
}

func (kv *KV) load() {
	if kv.file == "" {
		return
	}

	if _, err := os.Stat(kv.file); os.IsNotExist(err) {
		return
	}

	content, err := os.ReadFile(kv.file)
	if err != nil {
		fmt.Println("Failed to read file: ", err)
		return
	}

	json.Unmarshal(content, &kv.data)
}
