package model

type Hello struct {
	Id   uint64 `gorm:"Id"`
	Text string `gorm:"Text"`
}

func (Hello) TableName() string {
	return "hello"
}
