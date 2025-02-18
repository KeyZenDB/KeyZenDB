package internal

import (
	"testing"
	"time"
)

func TestDb_SetAndGet(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	if db.Get("key1") != "value1" {
		t.Errorf("expected value1, got %v", db.Get("key1"))
	}
}

func TestDb_Del(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	db.Del("key1")
	if db.Get("key1") != nil {
		t.Errorf("expected nil, got %v", db.Get("key1"))
	}
}

func TestDb_Clear(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	db.Set("key2", "value2")
	db.Clear()
	if db.Len() != 0 {
		t.Errorf("expected length 0, got %d", db.Len())
	}
}

func TestDb_Keys(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	db.Set("key2", "value2")
	keys := db.Keys()
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}
}

func TestDb_Exists(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	if !db.Exists("key1") {
		t.Errorf("expected key1 to exist")
	}
	if db.Exists("key2") {
		t.Errorf("expected key2 to not exist")
	}
}

func TestDb_TTL(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1", 2*time.Second)
	if db.Get("key1") != "value1" {
		t.Errorf("expected value1 before TTL expires, got %v", db.Get("key1"))
	}
	time.Sleep(3 * time.Second)
	if db.Get("key1") != nil {
		t.Errorf("expected nil after TTL expires, got %v", db.Get("key1"))
	}
}

func TestDb_IncrDecr(t *testing.T) {
	db := NewDb()
	db.Set("counter", 0)
	if db.Incr("counter") != 1 {
		t.Errorf("expected 1 after increment, got %d", db.Get("counter"))
	}
	if db.Decr("counter") != 0 {
		t.Errorf("expected 0 after decrement, got %d", db.Get("counter"))
	}
}

func TestDb_Len(t *testing.T) {
	db := NewDb()
	db.Set("key1", "value1")
	db.Set("key2", "value2")
	if db.Len() != 2 {
		t.Errorf("expected length 2, got %d", db.Len())
	}
}
