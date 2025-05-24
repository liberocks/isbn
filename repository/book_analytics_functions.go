package repository

import "context"

func (r *Repository) BookCount(ctx context.Context) (int, error) {
	count := 0

	for range bookStore {
		count++
	}

	return count, nil
}

func (r *Repository) BookAuthorCount(ctx context.Context) (int, error) {
	var authorBookCount = make(map[string]int)

	for _, book := range bookStore {
		authorBookCount[book.Author]++
	}

	return len(authorBookCount), nil
}

func (r *Repository) BookGetOldestReleaseDate(ctx context.Context) (string, error) {
	var oldest string
	for _, book := range bookStore {
		if oldest == "" || book.ReleaseDate < oldest {
			oldest = book.ReleaseDate
		}
	}
	return oldest, nil
}

func (r *Repository) BookGetNewestReleaseDate(ctx context.Context) (string, error) {
	var newest string
	for _, book := range bookStore {
		if newest == "" || book.ReleaseDate > newest {
			newest = book.ReleaseDate
		}
	}
	return newest, nil
}

func (r *Repository) BookGetMostProductiveAuthor(ctx context.Context) (string, error) {
	var mostProductive string
	var maxCount int
	authorBookCount := make(map[string]int)

	for _, book := range bookStore {
		authorBookCount[book.Author]++
	}

	for author, count := range authorBookCount {
		if count > maxCount {
			maxCount = count
			mostProductive = author
		}
	}

	return mostProductive, nil
}

func (r *Repository) BookGetLongestTitle(ctx context.Context) (string, error) {
	var longest string
	for _, book := range bookStore {
		if len(book.Title) > len(longest) {
			longest = book.Title
		}
	}
	return longest, nil
}

func (r *Repository) BookGetShortestTitle(ctx context.Context) (string, error) {
	var shortest string
	for _, book := range bookStore {
		if shortest == "" || len(book.Title) < len(shortest) {
			shortest = book.Title
		}
	}
	return shortest, nil
}
