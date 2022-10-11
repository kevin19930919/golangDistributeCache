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
