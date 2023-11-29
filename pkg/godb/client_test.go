package godb

import (
	"log"
	"testing"
	"time"

	"github.com/matthew-james-laidlaw/godb/pkg/assert"
)

func TestClient_EndToEnd(t *testing.T) {

	go func() {
		s := NewServer()
		err := s.Listen(":8000")
		if err != nil {
			log.Fatalln(err)
		}
	}()

	time.Sleep(1 * time.Second)

	client, err := NewClient(":8000")
	if err != nil {
		t.Error(err)
	}

	res, err := client.Get("non-existent-key")
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(res.Result, "", t)

	res, err = client.Set("key", "value")
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(res.Result, "1", t)

	res, err = client.Get("key")
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(res.Result, "value", t)

	res, err = client.Del("key")
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(res.Result, "1", t)

	res, err = client.Get("key")
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(res.Result, "", t)
}
