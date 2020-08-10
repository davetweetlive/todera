package models

import "time"

type Blogs struct {
	Blog
}

type Blog struct {
	Title      string
	Content    string
	Author     User
	Thumbnail  []byte
	PublishOn  time.Time
	ModifiedOn time.Time
	Tags       []Tag
}

type Comment struct {
	Content     string
	CommentTime time.Time
	ModifiedOn  time.Time
	User
	Blog
}
type Tag struct {
	tag string
}

func PostBlog() {

}

func GetAllBlog() {

}

func GetBlogById() {

}
