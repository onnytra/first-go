package entity

type Health struct {
	Status string `json:"status" gorm:"type:varchar(255);not null"`
}
