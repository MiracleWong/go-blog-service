package model

type Tag struct {
	*Model
	// Tag名称
	Name  string  `json:"name"`
	// 状态0为禁用，状态1为启用
	State  uint8    `json:"state"`
}

func (model Tag) TableName() string {
	return "blog_tag"
}
