package api

type Film struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
}