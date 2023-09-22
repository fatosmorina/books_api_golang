package models

type Book struct {
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorId string `json:"author_id"`
}
