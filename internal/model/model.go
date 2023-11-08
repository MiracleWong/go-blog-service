package model

type Model struct {
	// 创建人
	CreatedBy string `json:"created_by"`
	// 创建时间
	CreatedOn uint32 `json:"created_on"`
	// 删除时间
	DeletedOn uint32 `json:"deleted_on"`
	// id
	ID uint32 `json:"id"`
	// 是否删除，0为未删除，1为删除
	IsDel uint8 `json:"is_del"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
	// 修改时间
	ModifiedOn uint32 `json:"modified_on"`
}
