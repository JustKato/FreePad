package post

import (
	"errors"

	"github.com/JustKato/FreePad/helper"
	"github.com/JustKato/FreePad/models/database"
)

var postList []*Post = []*Post{}

var postMap map[string]Post = make(map[string]Post)

func GetPostList() []*Post {
	return postList
}

func Retrieve(name string) (*Post, error) {

	if len(name) < 1 {
		return nil, errors.New("the name of the post must contain at least 1 character")
	}

	if len(name) > 256 {
		return nil, errors.New("the name of the post must not exceed 256 characters")
	}

	// Check if we have the post cached
	if val, ok := postMap[name]; ok {
		return &val, nil
	}

}

func Create(name string, content string) (*Post, error) {

	if len(name) < 1 {
		return nil, errors.New("the name of the post must contain at least 1 character")
	}

	if len(name) > 256 {
		return nil, errors.New("the name of the post must not exceed 256 characters")
	}

	if len(content) > 16777200 {
		return nil, errors.New("provided content is too long, please do not exceed ")
	}

	// Initialize the post
	myPost := Post{
		Name:    name,
		Content: content,
	}

	// Check if we can cache this element
	if len(postMap) < helper.GetCacheMapLimit() {
		// Set the post by name
		postMap[name] = myPost
	}

	// Add the post to the database
	db, err := database.GetConn()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := `REPLACE INTO freepad.t_posts (name, content) VALUES (?, ?)`
	s, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	_, err = s.Exec(myPost.Name, myPost.Content)
	if err != nil {
		return nil, err
	}

	// Return the post
	return &myPost, nil
}
