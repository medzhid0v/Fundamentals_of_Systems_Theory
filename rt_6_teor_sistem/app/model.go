package app

// Структура, описывающая поля данных каждого животного
type Stool struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Higt  int    `json:"higt"`
	Price int    `json:"price"`
}