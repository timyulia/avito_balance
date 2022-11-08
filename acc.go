package balance

type User struct {
	Id     int    `json:"id" db:"id"`
	Reason string `json:"reason" db:"reason"`
	Amount int    `json:"amount" db:"amount"`
}

type Order struct {
	UserId    int `json:"id" db:"id"`
	ServiceId int `json:"service_id" db:"service_id"`
	OrderId   int `json:"order_id" db:"order_id"`
	Amount    int `json:"amount" db:"amount"`
}

type Report struct {
	ServiceId int    `json:"service_id" db:"service_id"`
	Name      string `json:"name" binding:"required"`
	Amount    int    `json:"amount" db:"amount"`
}

type Service struct {
	ServiceId int    `json:"service_id" db:"service_id"`
	Name      string `json:"name" db:"name"`
}

type History struct {
	Reason string `json:"reason" db:"reason"`
	Amount int    `json:"amount" db:"amount"`
	Date   string `json:"date" db:"date"`
}
