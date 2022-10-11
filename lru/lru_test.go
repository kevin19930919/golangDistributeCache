package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestNewFunc(t *testing.T) {
	lru := New(int64(0), nil)
	t.Log(lru)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	t.Log(lru.Get("ke1"))
}

func TestPrintKV(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key2", String("1234"))
	if element, ok := lru.cache["key2"]; ok {
		t.Log(reflect.TypeOf(element.Value))
		kv := element.Value.(*entry)
		t.Log("===== type of kv =====: ", reflect.TypeOf(kv))
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	cap := len(k1) + len(k2) + len(v1) + len(v2)

	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Log("remove oldest key1 fail")
	} else {
		t.Log("remove oldest key1 success")
	}
}
