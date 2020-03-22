package router

import (
	"github.com/cancue/gobase/router"

	"gobase_demo/controller/post"
)

// Apply registers routes with controllers to the server.
// This should be called once before starting the server.
func Apply(s router.Server) {
	// post
	s.POST("/posts", router.Controller(new(post.Create)))
	s.GET("/posts/:postID", router.Controller(new(post.Read)))
}
