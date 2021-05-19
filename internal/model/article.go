package model

import "github.com/MiracleWong/go-blog-service/pkg/app"

type Article struct {
	*Model
	// 文章标题
	Title  string  `json:"title"`
	// 文章简述
	Desc   string  `json:"desc"`
	// 封面图片地址
	CoverImageUrl  string  `json:"cover_image_url"`
	// 文章内容
	Content        string  `json:"content"`
	// 状态0为禁用，状态1为启用
	State  uint8    `json:"state"`
}

func (model Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}