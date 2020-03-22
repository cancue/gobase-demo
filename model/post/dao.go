package post

import (
	"time"

	"github.com/cancue/gobase/db/pg"
	uuid "github.com/satori/go.uuid"
)

// Post is a DAO.
type Post struct {
	PostID    uuid.UUID `json:"postID" pg:",pk"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"-"`
}

// Create interacts with database.
func (post *Post) Create() (err error) {
	post.PostID = uuid.NewV4()

	connection := pg.Get()

	_, err = connection.
		Model(post).Returning("*").Insert()

	return
}

// Read interacts with database.
func (post *Post) Read() error {
	connection := pg.Get()

	return connection.Select(post)
}

// Update interacts with database.
func (post *Post) Update(columns *[]string) (err error) {
	connection := pg.Get()

	_, err = connection.
		Model(post).Column(*columns...).WherePK().Update()

	return
}

// Delete interacts with database.
func (post *Post) Delete() (err error) {
	connection := pg.Get()

	_, err = connection.
		Model(post).WherePK().Delete()

	return
}
