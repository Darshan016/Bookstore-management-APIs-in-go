# Bookstore Backend Application

This is a simple backend application for managing a bookstore. It provides basic CRUD (Create, Read, Update, Delete) operations on the Book struct using Go and MySQL.

## Features

- Add a new book
- Retrieve details of a book
- Update book information
- Delete a book from the inventory
- List all books

## Technologies Used

- Go (Golang)
- MySQL
- `jinzhu/dialects/mysql` for database interactions

## Installation

### Prerequisites

- Go 1.16 or later
- MySQL 5.7 or later

### Steps

1. **Clone the repository:**

    ```sh
    git clone https://github.com/Darshan016/go-bookstore.git
    cd go-bookstore
    ```

2. **Install dependencies:**

    ```sh
    go get -u github.com/jinzhu/gorm
    go get -u github.com/jinzhu/gorm/dialects/mysql
    ```

3. **Set up MySQL:**

    Create a database named `bookstore` and a user with appropriate privileges. For example:

    ```sql
    CREATE DATABASE bookstore;
    CREATE USER 'bookstore_user'@'localhost' IDENTIFIED BY 'password';
    GRANT ALL PRIVILEGES ON bookstore.* TO 'bookstore_user'@'localhost';
    FLUSH PRIVILEGES;
    ```

4. **Configure the application:**

    Update the `main.go` file with your MySQL credentials:

    ```go
    dsn := "bookstore_user:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
    ```

5. **Run the application:**

    ```sh
    go run main.go
    ```

## API Endpoints

### Create a new book

- **URL:** `/book/`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
      "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
    }
    ```

- **Response:**

    ```json
    {
      "id": 1,
     "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
    }
    ```

### Retrieve a book by ID

- **URL:** `/book/{id}`
- **Method:** `GET`
- **Response:**

    ```json
    {
      "id": 1,
    "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
    }
    ```

### Update a book

- **URL:** `/book/{id}`
- **Method:** `PUT`
- **Request Body:**

    ```json
    {
   "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
    }
    ```

- **Response:**

    ```json
    {
      "id": 1,
     "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
    }
    ```

### Delete a book

- **URL:** `/book/{id}`
- **Method:** `DELETE`
- **Response:**

    ```json
    {
      "message": "Book deleted successfully"
    }
    ```

### List all books

- **URL:** `/book/`
- **Method:** `GET`
- **Response:**

    ```json
    [
      {
        "id": 1,
        "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
      },
      {
        "id": 2,
       "name": "Book Title",
      "author": "Author Name",
      "publication": "publication"
      }
    ]
    ```
    

## Acknowledgements

- [Gorm](https://gorm.io/) for the ORM library
- [MySQL](https://www.mysql.com/) for the database
