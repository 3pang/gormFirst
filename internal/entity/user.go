package entity

type User struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string
	Age string
}
