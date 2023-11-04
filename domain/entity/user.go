package entity

type User struct {
	Base
	Email      string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	ExternalID string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}
