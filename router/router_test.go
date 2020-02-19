package router

import (
	"encoding/json"
	"testing"

	"github.com/cancue/gobase"
	"github.com/stretchr/testify/assert"
)

type API struct {
	Method string
	Path   string
	Name   string
}

type APIs []API

func TestApply(t *testing.T) {
	server := gobase.NewWithConfig(new(gobase.Config))
	expect := &APIs{
		{"POST", "/posts", "github.com/cancue/gobase.Controller.func1"},
		{"GET", "/posts/:postID", "github.com/cancue/gobase.Controller.func1"},
	}

	Apply(server)

	data, err := server.Routes()
	assert.Nil(t, err, "gobase.Controllers error")

	result := new(APIs)
	err = json.Unmarshal(data, result)

	assert.Nil(t, err)
	assert.ElementsMatch(t, *expect, *result, "routes changed.")
}
