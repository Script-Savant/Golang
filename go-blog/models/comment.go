package models

import "gorm.io/gorm"

// comment model
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	PostID uint `gorm:"not null" json:"post_id"`
	Post Post `gorm:"foreignKey:PostID" json:"post"`
	UserID uint `gorm:"not null" json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	Like []Like `gorm:"foreignKey:CommentID" json:"likes"`
}

// like model -> like on either post or comment
type Like struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	PostID *uint `json:"post_id,omitempty"`
	Post *Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CommentID *uint `json:"comment_id,omitempty"`
	Comment *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
	IsLike bool `gorm:"not null" json:"is_like"`
}