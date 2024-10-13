package models

type ToDo struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	Title    string
	Body     string
	Finished bool
}
