package post

import "errors"

var postList []*Post = []*Post{}

var postMap map[string]Post = make(map[string]Post)

func GetPostList() []*Post {
	return postList
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

	// Set the post by name
	postMap[name] = myPost

	// Add the post to the database

	// Return the post
	return &myPost, nil
}
