package models

type Order struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	Product   string `json:"product"`
	Amount    int    `json:"amount"`
	CreatedAt string `xorm:"created" json:"created_at"`
}
