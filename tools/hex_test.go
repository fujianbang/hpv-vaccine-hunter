package tools

import (
	"testing"
	"time"
)

func TestMd5Hex(t *testing.T) {
	s := "fheoaf8eiaj"
	r := Md5Hex(s)
	t.Log(r)

	t.Log(time.Now().UnixMilli())
}
