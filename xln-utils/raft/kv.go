package raft

import (
	"bytes"
	"encoding/gob"
)

type StringKV struct {

	KV map[string]string

}

func NewStringKV() *StringKV{

	return &StringKV{KV: make(map[string]string)}

}

func (kv *StringKV) Encode() []byte {
 	b := new(bytes.Buffer)
	encoder := gob.NewEncoder(b)

	err := encoder.Encode(kv.KV)
	if err != nil {
		panic("String KV encode failed")
	}
	return b.Bytes()
}

func (kv *StringKV) Decode(byte []byte) {

	kv.KV = make(map[string]string)

	decoder := gob.NewDecoder(bytes.NewBuffer(byte))
	err := decoder.Decode(&kv.KV)
	if err != nil {
		panic("String KV decode failed")
	}
}