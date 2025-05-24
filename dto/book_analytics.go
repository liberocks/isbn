package dto

// DTO
type BookAnalyticsTriggerResponse struct {
	Message string `json:"message"`
}

type BookAnalyticsGetResponse struct {
	TotalBooks            int    `json:"total_books"`
	TotalAuthors          int    `json:"total_authors"`
	OldestBookReleaseDate string `json:"oldest_book_release_date"`
	NewestBookReleaseDate string `json:"newest_book_release_date"`
	MostProductiveAuthor  string `json:"most_productive_author"`
	LongestBookTitle      string `json:"longest_book_title"`
	ShortestBookTitle     string `json:"shortest_book_title"`
}
