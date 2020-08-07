package models

import "time"

type Blogs struct {
	Blog
}

type Blog struct {
	Title      string
	Content    string
	Author     string
	PublishOn  time.Time
	ModifiedOn time.Time
}

func PostBlog() {

}

func GetAllBlog() {

}

func GetBlogById() {

}
