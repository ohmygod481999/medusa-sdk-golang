package schema

import "time"

type ArticleStatus string

const (
	ArticleDraft     ArticleStatus = "draft"
	ArticlePublished ArticleStatus = "published"
)

type Article struct {
	Title             string         `json:"title"`
	Id                string         `json:"id"`
	Handle            string         `json:"handle"`
	Thumbnail         string         `json:"thumbnail"`
	Content           string         `json:"content"`
	Status            ArticleStatus  `json:"status"`
	Metadata          map[string]any `json:"metadata"`
	ArticleCategoryId string         `json:"article_category_id"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         *time.Time     `json:"deleted_at"`
}

type ArticleCategory struct {
	Title     string         `json:"title"`
	Id        string         `json:"id"`
	Handle    string         `json:"handle"`
	Metadata  map[string]any `json:"metadata"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at"`
}
