package database

type Book struct {
	ID           int64     `json:"id"`
	Title        string    `json:"tilte"`
	PublishedAt  string    `json:"published_at"`
}
