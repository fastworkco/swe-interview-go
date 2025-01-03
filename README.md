# REST API for Managing Items

This project demonstrates a REST API implementation for managing items, written in Golang. It includes endpoints for CRUD operations and other functionalities for evaluating candidates during interviews.

## Features
- CRUD operations for items
  - RESTful API design
- Get list items with pagination
- Input validation
- Modular and scalable project structure
- Search and filter items (name, price range)
- Sorting items

## Additional Features
- Cache result of each request for 60 seconds
- Integrating with databases like PostgreSQL, MongoDB, etc.

## Endpoints

### Base URL
`http://localhost:3000`

### Endpoints

#### 1. Create an Item
- **POST** `/items`
- **Request Body**:
  ```json
  {
    "name": "string",
    "price": number
    "amount": number
  }
  ```
- **Response**:
  ```json
  {
    "id": "string",
    "name": "string",
    "price": number,
    "amount": number
  }
  ```

#### 2. Get All Items
- **GET** `/items`
- **Response**:
  ```json
  [
    {
      "id": "string",
      "name": "string",
      "price": number,
      "amount": number
    }
  ]
  ```

#### 3. Get an Item by ID
- **GET** `/items/:id`
- **Response**:
  ```json
  {
    "id": "string",
    "name": "string",
    "price": number,
    "amount": number
  }
  ```

#### 4. Update an Item
- **PUT** `/items/:id`
- **Request Body**:
  ```json
  {
    "name": "string",
    "price": number,
    "amount": number
  }
  ```
- **Response**:
  ```json
  {
    "id": "string",
    "name": "string",
    "price": number,
    "amount": number
  }
  ```

#### 5. Delete an Item
- **DELETE** `/items/:id`
- **Response**:
  ```json
  {
    "message": "Item deleted successfully"
  }
  ```

## Project Structure

```
swe-interview-go/
├── main.go
├── items.csv
└── go.mod
```

## Setup and Run

### Prerequisites
Ensure you have the following installed:

- [Golang](https://go.dev/) (>=1.20.x)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/swe-interview-go.git
   cd swe-interview-go
   ```

2. Install dependencies:
   ```bash
   go get
   ```

3. Run the project:
   ```bash
   go run main.go
   ```

4. Access the API at `http://localhost:3000`.
