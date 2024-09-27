package model

// WhiteList Структура пользователя.
type WhiteList struct {
	ID     int64  `json:"id"`
	Subnet string `json:"subnet" db:"subnet"`
}

// BlackList Структура пользователя.
type BlackList struct {
	ID     int64  `json:"id"`
	Subnet string `json:"subnet" db:"subnet"`
}
