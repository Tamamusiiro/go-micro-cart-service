package model

type Cart struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	UserId    int64 `gorm:"not null;index" json:"user_id"`
	ProductId int64 `gorm:"not null;index" json:"product_id"`
	SizeId    int64 `gorm:"not null;" json:"size_id"`
	Num       int64 `gorm:"not null;" json:"num"`
}
