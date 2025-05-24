package dto

import "fmt"

// DTO
type GetBooksQuery struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetBooksBaseResponse struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release_date"`
}

type GetBooksResponse struct {
	Data       []GetBooksBaseResponse `json:"data"`
	Total      int                    `json:"total"`
	TotalPages int                    `json:"total_pages"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
}

// Validator
func (c *GetBooksQuery) Validate() error {
	if c.Page < 1 {
		return fmt.Errorf("page must be greater than 0")
	}

	if c.Limit < 1 || c.Limit > 100 {
		return fmt.Errorf("limit must be between 1 and 100")
	}

	return nil
}
