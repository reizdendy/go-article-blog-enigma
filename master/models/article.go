package models

// Article is a struct of Model Article
type Article struct {
	ArticleID      string `json:"article_id"`
	ArticleName    string `json:"article_name"`
	ArticleContent string `json:"article_content"`
	UserCreator    string `json:"username"`
	CreatedAt      string `json:"created_at"`
}
