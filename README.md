# ISBN Book API

This is a simple API to manage books using their ISBN numbers. It allows you to add, retrieve, and delete books from a collection. It will log all requests and responses to a file (app.log) and stdout.

## Table of Contents
- [How to run locally](#how-to-run-locally)
- [How to run with Docker](#how-to-run-with-docker)
- [How to test](#how-to-test)
- [API Endpoints](#api-endpoints)
  - [Create a book](#create-a-book)
  - [Get book by ISBN](#get-book-by-isbn)
  - [Get list of books](#get-list-of-books)
  - [Update a book by ISBN](#update-a-book-by-isbn)
  - [Delete a book by ISBN](#delete-a-book-by-isbn)
  - [Analytics Endpoints](#analytics-endpoints)

## How to run locally
Run the following command to start the API locally:

```bash
make dev
```

## How to run with Docker
To run the API using Docker, use the following command:

```bash
docker build -t isbn-book-api .
docker run -p 8080:8080 isbn-book-api
```

## How to test
To run the tests, use the following command:

```bash
make test
```

You can also run the tests with actual API calls using the command below. But make sure to have the API running locally or in Docker.

```bash
make test_curl
```

## API Endpoints
### Create a book

```mermaid
sequenceDiagram
    accTitle: Create a book

    User->>Handler: Send payload
    Handler->>Service: Validate payload
    Service-->>Handler: If not valid, return error
    Handler-->>User: If not valid, return error
    Service->>Repository: Create book
    Repository->>Database: Check if ISBN existed
    Database->>Repository: 
    Repository-->>Service: If existed, return error
    Service-->>Handler: If existed, return error
    Handler-->>User: If existed, return error
    Repository->>Database: If not existed, create new record in the DB
    Database->>Repository: 
    Repository->>Service: Book created
    Service->>Handler: Construct response
    Handler->>User: Book created
```

This endpoint allows you to create a new book entry. This endpoint expects a JSON body with the book's details. The request payload will be validated, and if successful, the book will be added to the collection.

POST `/books`

```json
{
  "isbn": "1234567890123",
  "title": "Example Book",
  "author": "John Doe",
  "release_date": "2023-10-01"
}
```
#### Response
```json
{
    "isbn": "1234567890123",
    "title": "Example Book",
    "author": "John Doe",
    "release_date": "2023-10-01"
}
```

### Get book by ISBN
```mermaid
sequenceDiagram
    accTitle: Get book by ISBN

    User->>Handler: Send request with ISBN
    Handler->>Service: Validate ISBN
    Service-->>Handler: If not valid, return error
    Handler-->>User: If not valid, return error
    Service->>Repository: Get book by ISBN
    Repository->>Database: Check if ISBN existed
    Database-->>Repository: If not existed, return error
    Repository-->>Service: If not existed, return error
    Service-->>Handler: If not existed, return error
    Handler-->>User: If not existed, return error
    Repository->>Database: If existed, get book details
    Database->>Repository: 
    Repository->>Service: Book found
    Service->>Handler: Construct response
    Handler->>User: Book details
```
This endpoint allows you to retrieve a book's details using its ISBN. If the book exists, its details will be returned.

GET `/books/{isbn}`

#### Response
```json
{
    "isbn": "1234567890123",
    "title": "Example Book",
    "author": "John Doe",
    "release_date": "2023-10-01"
}
```
 
### Get list of books
```mermaid
sequenceDiagram
    accTitle: Get list of books

    User->>Handler: Send request
    Handler->>Service: Get all books
    Service->>Repository: Get all books
    Repository->>Database: Get all records
    Database->>Repository: 
    Repository->>Service: List of books
    Service->>Handler: Construct response
    Handler->>User: List of books
```
This endpoint allows you to retrieve a list of all books in the collection. The response will include all book details. You can also paginate the results by providing `page` and `limit` query parameters.

GET `/books?page={page}&limit={limit}`

#### Response
```json
{
    "data": [
        {
            "isbn": "1234567890123",
            "title": "Example Book",
            "author": "John Doe",
            "release_date": "2023-10-01"
        },
        {
            "isbn": "9876543210987",
            "title": "Another Book",
            "author": "Jane Smith",
            "release_date": "2023-09-15"
        }
    ],
    "total": 2,
    "total_pages": 1,
    "page": 1,
    "limit": 10
}
```


### Update a book by ISBN

```mermaid
sequenceDiagram
    accTitle: Update a book by ISBN

    User->>Handler: Send payload with ISBN
    Handler->>Service: Validate payload
    Service-->>Handler: If not valid, return error
    Handler-->>User: If not valid, return error
    Service->>Repository: Update book by ISBN
    Repository->>Database: Check if ISBN existed
    Database-->>Repository: If not existed, return error
    Repository-->>Service: If not existed, return error
    Service-->>Handler: If not existed, return error
    Handler-->>User: If not existed, return error
    Repository->>Database: If existed, update record in the DB
    Database->>Repository: 
    Repository->>Service: Book updated
    Service->>Handler: Construct response
    Handler->>User: Book updated
```
This endpoint allows you to update the details of an existing book using its ISBN. The request payload should include the updated book details.


PUT `/books/{isbn}`

```json
{
  "title": "Updated Book Title",
  "author": "Updated Author",
  "release_date": "2023-11-01"
}
```
#### Response
```json
{
    "isbn": "1234567890123",
    "title": "Updated Book Title",
    "author": "Updated Author",
    "release_date": "2023-11-01"
}
```

### Delete a book by ISBN
```mermaid
sequenceDiagram
    accTitle: Delete a book by ISBN

    User->>Handler: Send request with ISBN
    Handler->>Service: Validate ISBN
    Service-->>Handler: If not valid, return error
    Handler-->>User: If not valid, return error
    Service->>Repository: Delete book by ISBN
    Repository->>Database: Check if ISBN existed
    Database-->>Repository: If not existed, return error
    Repository-->>Service: If not existed, return error
    Service-->>Handler: If not existed, return error
    Handler-->>User: If not existed, return error
    Repository->>Database: If existed, delete record in the DB
    Database->>Repository: 
    Repository->>Service: Book deleted
    Service->>Handler: Construct response
    Handler->>User: Book deleted
```
This endpoint allows you to delete a book from the collection using its ISBN. If the book exists, it will be removed from the collection.

DELETE `/books/{isbn}`

#### Response
```json
{
    "message": "Book deleted successfully"
}
```

### Analytics Endpoints
To simulate the usage of goroutines, I added two endpoints to trigger the analytics process. These endpoints will run in the background and will not block the main thread. The analytics handler will trigger the analytics service, which will perform the analytics tasks concurrently using goroutines. Inside this service, multiple goroutines will be used to perform different analytics tasks, such as counting books, finding the oldest and newest release dates, and identifying the most productive author. The results of these tasks will be collected and stored in the database. After the analytics process is complete, the user can retrieve the analytics results using get analytics endpoint.



```mermaid
sequenceDiagram
    accTitle: Analytics pipeline

    User->>Handler: Trigger analytics
    Handler->>User: Analytics pipeline started
    Handler-->>Service: Start analytics
    Service-->>Repository: Count all books
    Repository-->>Database: Count all records
    Database-->>Repository:
    Repository-->>Service: Count all books result
    Service-->>Repository: Count all author
    Repository-->>Database: Count all authors
    Database-->>Repository:
    Repository-->>Service: Count all authors result
    Service-->>Repository: Find oldest release date
    Repository-->>Database: Find oldest release date
    Database-->>Repository:
    Repository-->>Service: Find oldest release date result
    Service-->>Repository: Find newest release date
    Repository-->>Database: Find newest release date
    Database-->>Repository:
    Repository-->>Service: Find newest release date result
    Service-->>Repository: Find most productive author
    Repository-->>Database: Find most productive author
    Database-->>Repository:
    Repository-->>Service: Find most productive author result
    Service-->>Repository: Find longest book title
    Repository-->>Database: Find longest book title
    Database-->>Repository:
    Repository-->>Service: Find longest book title result
    Service-->>Repository: Find shortest book title
    Repository-->>Database: Find shortest book title
    Database-->>Repository:
    Repository-->>Service: Find shortest book title result
    Service-->>Repository: Update analytics
    Repository-->>Database: Update analytics
    Database-->>Repository:
    Repository-->>Service: Update analytics result
    User-->>Handler: Get analytics
    Handler-->>Service: Get analytics
    Service-->>Repository: Get analytics
    Repository-->>Database: Get analytics
    Database-->>Repository:
    Repository-->>Service: Return analytics result
    Service-->>Handler: Return analytics result
    Handler-->>User: Return analytics result
```