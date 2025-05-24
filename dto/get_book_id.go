package dto

// DTO
type GetBookByIDResponse struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release_date"`
}
