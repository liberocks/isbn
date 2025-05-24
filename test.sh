#!/bin/bash

# Base URL for the API
BASE_URL="http://localhost:8080"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}=== ISBN Book API Test Script ===${NC}"
echo

# Function to print test results
print_result() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✓ PASS${NC}: $2"
    else
        echo -e "${RED}✗ FAIL${NC}: $2"
    fi
}

# Test 1: Create a new book
echo -e "${YELLOW}Test 1: Create a new book${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X POST "$BASE_URL/books" \
    -H "Content-Type: application/json" \
    -d '{
        "isbn": "9781234567890",
        "title": "Test Book Title",
        "author": "Test Author",
        "release_date": "2023-01-15"
    }')
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 201 ]; then
    print_result 0 "Book created successfully"
else
    print_result 1 "Failed to create book"
fi
echo

# Test 2: Create book with invalid data (missing title)
echo -e "${YELLOW}Test 2: Create book with invalid data (missing title)${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X POST "$BASE_URL/books" \
    -H "Content-Type: application/json" \
    -d '{
        "isbn": "9781234567891",
        "title": "",
        "author": "Test Author",
        "release_date": "2023-01-15"
    }')
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 400 ]; then
    print_result 0 "Validation error handled correctly"
else
    print_result 1 "Validation error not handled"
fi
echo

# Test 3: Get book by ID
echo -e "${YELLOW}Test 3: Get book by ID${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books/9781234567890")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 200 ]; then
    print_result 0 "Book retrieved successfully"
else
    print_result 1 "Failed to retrieve book"
fi
echo

# Test 4: Get non-existent book
echo -e "${YELLOW}Test 4: Get non-existent book${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books/9999999999999")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 500 ] || [ "$HTTP_CODE" -eq 404 ]; then
    print_result 0 "Non-existent book handled correctly"
else
    print_result 1 "Non-existent book not handled properly"
fi
echo

# Test 5: Update book by ID
echo -e "${YELLOW}Test 5: Update book by ID${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$BASE_URL/books/9781234567890" \
    -H "Content-Type: application/json" \
    -d '{
        "title": "Updated Test Book Title",
        "author": "Updated Test Author",
        "release_date": "2023-12-01"
    }')
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 200 ]; then
    print_result 0 "Book updated successfully"
else
    print_result 1 "Failed to update book"
fi
echo

# Test 6: Update book with invalid data
echo -e "${YELLOW}Test 6: Update book with invalid data (short title)${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$BASE_URL/books/9781234567890" \
    -H "Content-Type: application/json" \
    -d '{
        "title": "Hi",
        "author": "Test Author",
        "release_date": "2023-12-01"
    }')
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 400 ]; then
    print_result 0 "Update validation error handled correctly"
else
    print_result 1 "Update validation error not handled"
fi
echo

# Create additional books for list testing
echo -e "${YELLOW}Creating additional books for list testing...${NC}"
curl -s -X POST "$BASE_URL/books" \
    -H "Content-Type: application/json" \
    -d '{
        "isbn": "9781234567892",
        "title": "Second Test Book",
        "author": "Second Author",
        "release_date": "2023-02-15"
    }' > /dev/null

curl -s -X POST "$BASE_URL/books" \
    -H "Content-Type: application/json" \
    -d '{
        "isbn": "9781234567893",
        "title": "Third Test Book",
        "author": "Third Author",
        "release_date": "2023-03-15"
    }' > /dev/null
echo "Additional books created"
echo

# Test 7: Get list of books (default pagination)
echo -e "${YELLOW}Test 7: Get list of books (default pagination)${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 200 ]; then
    print_result 0 "Book list retrieved successfully"
else
    print_result 1 "Failed to retrieve book list"
fi
echo

# Test 8: Get list of books with custom pagination
echo -e "${YELLOW}Test 8: Get list of books with custom pagination${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books?page=1&limit=2")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 200 ]; then
    print_result 0 "Book list with pagination retrieved successfully"
else
    print_result 1 "Failed to retrieve book list with pagination"
fi
echo

# Test 9: Get list with invalid pagination
echo -e "${YELLOW}Test 9: Get list with invalid pagination (page 0)${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books?page=0&limit=10")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 400 ]; then
    print_result 0 "Invalid pagination handled correctly"
else
    print_result 1 "Invalid pagination not handled"
fi
echo

# Test 10: Delete book by ID
echo -e "${YELLOW}Test 10: Delete book by ID${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$BASE_URL/books/9781234567893")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 200 ]; then
    print_result 0 "Book deleted successfully"
else
    print_result 1 "Failed to delete book"
fi
echo

# Test 11: Delete non-existent book
echo -e "${YELLOW}Test 11: Delete non-existent book${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$BASE_URL/books/9999999999999")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 500 ] || [ "$HTTP_CODE" -eq 404 ]; then
    print_result 0 "Delete non-existent book handled correctly"
else
    print_result 1 "Delete non-existent book not handled properly"
fi
echo

# Test 12: Verify deleted book is gone
echo -e "${YELLOW}Test 12: Verify deleted book is gone${NC}"
RESPONSE=$(curl -s -w "%{http_code}" -X GET "$BASE_URL/books/9781234567893")
HTTP_CODE="${RESPONSE: -3}"
BODY="${RESPONSE%???}"
echo "Response: $BODY"
echo "HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 500 ] || [ "$HTTP_CODE" -eq 404 ]; then
    print_result 0 "Deleted book correctly not found"
else
    print_result 1 "Deleted book still exists"
fi
echo

echo -e "${YELLOW}=== Test Script Completed ===${NC}"