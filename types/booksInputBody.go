package types

type BookInputBody struct {
	ID     int              `json:"id"`
	Isbn   string           `json:"isbn"`
	Title  string           `json:"title"`
	Author *AuthorInputBody `json:"author"`
}
