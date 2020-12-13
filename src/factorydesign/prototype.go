///bin/true; exec /usr/bin/env go run "" "$@"
package main

import "fmt"

// Post struct
type Post struct {
	Content string
	Website string
	Author  string
}

// Website string
type Website string

// const
const (
	Dev        Website = "Dev.to"
	Medium     Website = "Medium.com"
	HackerNews Website = "Hackernews.com"
)

// NewPost struct
func NewPost(website Website) *Post {
	author := "Tomassirio"
	switch website {
	case Dev:
		return &Post{"", string(website), author}
	case Medium:
		return &Post{"", string(website), author}
	case HackerNews:
		return &Post{"", string(website), author}
	default:
		return &Post{"", "", ""}
	}
}

func main() {
	content := "This is what I'm Cross-Posting today!"

	devPost := NewPost(Dev)
	mediumPost := NewPost(Medium)
	hackerNewsPost := NewPost(HackerNews)

	devPost.Content = content
	mediumPost.Content = content
	hackerNewsPost.Content = content

	fmt.Println(devPost)
	fmt.Println(mediumPost)
	fmt.Println(hackerNewsPost)
}
