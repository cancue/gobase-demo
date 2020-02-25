package router

import (
	"github.com/cancue/gobase"

	"gobase_demo/controller/post"
)

// Apply registers routes with controllers to the server.
// This should be called once before starting the server.
func Apply(s *gobase.Server) {
	// post
	s.POST("/posts", gobase.Controller(new(post.Create)))
	s.GET("/posts/:postID", gobase.Controller(new(post.Read)))
}
