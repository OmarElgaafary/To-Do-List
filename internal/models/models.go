package models

type ToDo struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Finished bool   `json:"finished"`
}
