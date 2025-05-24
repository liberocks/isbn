package dto

import "fmt"

// DTO
type UpdateBookByIDRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release_date"`
}

type UpdateBookByIDResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release_date"`
}

// Validator
func (c *UpdateBookByIDRequest) Validate() error {
	// Validating the title
	if c.Title == "" {
		return fmt.Errorf("title is required")
	}

	if len(c.Title) < 3 || len(c.Title) > 100 {
		return fmt.Errorf("title must be between 3 and 100 characters")
	}

	// Validating the author
	if c.Author == "" {
		return fmt.Errorf("author is required")
	}

	if len(c.Author) < 3 || len(c.Author) > 100 {
		return fmt.Errorf("author must be between 3 and 100 characters")
	}

	// Validating the release date
	if c.ReleaseDate == "" {
		return fmt.Errorf("release date is required")
	}
	if len(c.ReleaseDate) != 10 {
		return fmt.Errorf("release date must be in YYYY-MM-DD format")
	}
	if c.ReleaseDate[4] != '-' || c.ReleaseDate[7] != '-' {
		return fmt.Errorf("release date must be in YYYY-MM-DD format")
	}
	return nil
}
