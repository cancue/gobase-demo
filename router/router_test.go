package router

import (
	"encoding/json"
	"testing"

	"github.com/cancue/gobase"
	"github.com/cancue/gobase/router"
	"github.com/stretchr/testify/assert"
)

type API struct {
	Method string
	Path   string
}

type APIs []API

func TestApply(t *testing.T) {
	expect := &APIs{
		{"POST", "/posts"},
		{"GET", "/posts/:postID"},
	}

	gb := gobase.Server{Router: Apply}
	gb.Set(0, "..")

	data, err := router.Routes(gb.Echo)
	assert.Nil(t, err, "gobase.Controllers error")

	result := new(APIs)
	err = json.Unmarshal(data, result)

	assert.Nil(t, err)
	assert.ElementsMatch(t, *expect, *result, "routes changed.")
}
