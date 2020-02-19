package router

import (
	"github.com/cancue/gobase"

	"gobase_demo/controller/post"
)

func Apply(s *gobase.Server) {
	// post
	s.POST("/posts", gobase.Controller(new(post.Create)))
	s.GET("/posts/:postID", gobase.Controller(new(post.Read)))
}
