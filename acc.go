package balance

type User struct {
	Id int `json:"id" db:"id"`
	//Name     string `json:"name" binding:"required"`
	Amount int `json:"amount" db:"amount"`
}

type Order struct {
	ServiceId int `json:"service_id" db:"service_id"`
	OrderId   int `json:"order_id" db:"order_id"`
	Amount    int `json:"amount" db:"amount"`
}