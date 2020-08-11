package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Articles struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Thumbnail  []byte    `json:"thumbnail"`
	PublishOn  time.Time `json:"release date"`
	ModifiedOn time.Time `json:"update date"`
	Author     User      `json:"author"`
	Tags       []Tag     `json:"tags"`
	Comments   []Comment `json:"comments"`
}

type Comment struct {
	Content     string    `json:"content"`
	CommentedOn time.Time `json:"comment time"`
	ModifiedOn  time.Time `json:"update date"`
	ComentedBy  User      `json:"writer"`
}
type Tag struct {
	Tag string `json:"tag"`
}

func GenerateJson() string {
	articles := Articles{
		[]Article{
			{
				Title:      "Hola I'm writing",
				Content:    "These are all about my skills upgradations",
				Thumbnail:  []byte("someurl"),
				PublishOn:  time.Now(),
				ModifiedOn: time.Now(),
				Author: User{
					username: "davetweetlive",
				},
				Tags: []Tag{
					Tag{Tag: "golang"},
					Tag{Tag: "learning"},
				},
				Comments: []Comment{
					{
						Content:     "Good one",
						CommentedOn: time.Now(),
						ModifiedOn:  time.Now(),
						ComentedBy: User{
							username: "irak",
						},
					},
				},
			},

			{
				Title:      "Another Title",
				Content:    "Well it matters when it runs sucessfully on machine, Some people say it as Feature complete",
				Thumbnail:  []byte("someurl"),
				PublishOn:  time.Now(),
				ModifiedOn: time.Now(),
				Author: User{
					username: "davetweetlive",
				},
				Tags: []Tag{
					Tag{Tag: "golang"},
					Tag{Tag: "learning"},
				},
				Comments: []Comment{
					{
						Content:     "Can be improved",
						CommentedOn: time.Now(),
						ModifiedOn:  time.Now(),
						ComentedBy: User{
							username: "irak",
						},
					},
				},
			},
		},
	}

	jsonData, err := json.MarshalIndent(&articles, " ", "\t")
	if err != nil {
		fmt.Println("Error while marshaling structure")
	}
	return string(jsonData)
}

func PostBlog() {

}

func GetAllBlog() {

}

func GetBlogById() {

}
