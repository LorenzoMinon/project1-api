# project1-api 

A simple REST API built with Go to manage products. No database — data lives in memory.

## Run

```bash
go run main.go
```

Server starts on `http://localhost:8000`.

## Endpoints

### GET /products
Returns all products.

### GET /products/{id}
Returns a single product by ID.
Returns 404 if not found.

### POST /products
Creates a new product.

**Body:**
```json
{
    "name": "Jeans",
    "price": 100.00,
    "description": "Levis 505",
    "stock": 5
}
```

### PUT /products/{id}
Updates the stock of a product.

**Body:**
```json
{
    "stock": 20
}
```

### DELETE /products/{id}
Deletes a product by ID.
Returns 204 on success, 404 if not found.