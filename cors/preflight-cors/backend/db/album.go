package db

type Album struct {
	ID     int64   `json:"ID"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
