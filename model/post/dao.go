package post

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"gobase_demo/db"
)

type Post struct {
	PostID    uuid.UUID `json:"postID" pg:",pk"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"-"`
}

func (post *Post) Create() (postID uuid.UUID, err error) {
	postID = uuid.NewV4()
	post.PostID = postID

	connection := db.Get()

	_, err = connection.
		Model(post).Returning("*").Insert()

	return
}

func (post *Post) Read() error {
	connection := db.Get()

	return connection.Select(post)
}

func (post *Post) Update(columns *[]string) (err error) {
	connection := db.Get()

	_, err = connection.
		Model(post).Column(*columns...).WherePK().Update()

	return
}

func (post *Post) Delete() (err error) {
	connection := db.Get()

	_, err = connection.
		Model(post).WherePK().Delete()

	return
}
