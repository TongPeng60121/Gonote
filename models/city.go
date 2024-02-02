package models

type City struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:TINYTEXT"`
	Mayor string `gorm:"type:TINYTEXT"`
}

type Query struct {
	IDs   []uint   `json:"ids"`
	Names []string `json:"names"`
}
