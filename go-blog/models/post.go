package models

import "gorm.io/gorm"

// post model -> blog post
type Post struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	AuthorID uint `gorm:"not null" json:"author_id"`
	Author User `gorm:"foreignKey:AuthorID" json:"author"`
	Tags []Tag `gorm:"many2many:post_tags;" json:"tags"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"`
	Likes []Like `gorm:"foreignKey:PostID" json:"likes"`
	Shares uint `gorm:"default:0" json:"shares"`
}

// tag model
type Tag struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;" json:"posts"`
}