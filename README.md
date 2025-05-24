# ISBN Book API

This is a simple API to manage books using their ISBN numbers. It allows you to add, retrieve, and delete books from a collection.

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
### Response
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

### Response
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

### Response
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
    Repository->>Database: Update record in the DB
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
### Response
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
    Repository->>Database: Delete record in the DB
    Database->>Repository: 
    Repository->>Service: Book deleted
    Service->>Handler: Construct response
    Handler->>User: Book deleted
```
This endpoint allows you to delete a book from the collection using its ISBN. If the book exists, it will be removed from the collection.

DELETE `/books/{isbn}`

### Response
```json
{
    "message": "Book deleted successfully"
}
```