package usermodel

type Filter struct {
	Status int8 `json:"status,omitempty" gorm:"column:status"`
}
