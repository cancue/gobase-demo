package post

import (
	"github.com/cancue/gobase"
	uuid "github.com/satori/go.uuid"

	"gobase_demo/model/post"
)

// Create is the controller seed.
type Create struct {
	Title   string `validate:"required"`
	Content string `validate:"required"`
}

// Exec is the controller executor.
func (params *Create) Exec(_ gobase.Context) (result interface{}, err error) {
	post := post.Post{
		Title:   params.Title,
		Content: params.Content,
	}

	err = createPost(&post)

	return
}

// Read is the controller seed.
type Read struct {
	PostID uuid.UUID `param:"postID" validate:"required"`
}

// Exec is the controller executor.
func (params *Read) Exec(_ gobase.Context) (result interface{}, err error) {
	post := post.Post{
		PostID: params.PostID,
	}

	err = readPost(&post)
	result = post

	return
}

// For unit tests
var createPost = func(a *post.Post) (err error) {
	_, err = a.Create()

	return
}

var readPost = func(a *post.Post) error {
	return a.Read()
}
