# Cerebro API


## Installation & Setup

1. **Clone or download the project**
   ```bash
   cd cerebro-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the API**
   ```bash
   go run main.go
   ```

4. **The API will be available at:**
   - Base URL: `http://localhost:8080`
   - Health check: `http://localhost:8080/health`

## API Endpoints

### Health Check
- **GET** `/health` - Check if the API is running

### Items CRUD Operations
- **GET** `/api/v1/items` - Get all items
- **GET** `/api/v1/items/:id` - Get a specific item by ID
- **POST** `/api/v1/items` - Create a new item
- **PUT** `/api/v1/items/:id` - Update an existing item
- **DELETE** `/api/v1/items/:id` - Delete an item

## Example Usage

### 1. Check API Health
```bash
curl http://localhost:8080/health
```

### 2. Create a new item
```bash
curl -X POST http://localhost:8080/api/v1/items \
  -H "Content-Type: application/json" \
  -d '{"name": "Sample Item", "description": "This is a sample item"}'
```

### 3. Get all items
```bash
curl http://localhost:8080/api/v1/items
```

### 4. Get a specific item
```bash
curl http://localhost:8080/api/v1/items/1
```

### 5. Update an item
```bash
curl -X PUT http://localhost:8080/api/v1/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Item", "description": "This item has been updated"}'
```

### 6. Delete an item
```bash
curl -X DELETE http://localhost:8080/api/v1/items/1
```

## Data Structure

The API works with `Item` objects that have the following structure:

```json
{
  "id": 1,
  "name": "Item Name",
  "description": "Item Description"
}
```

## Response Format

All API responses follow this format:

**Success Response:**
```json
{
  "data": {...},
  "message": "Operation successful",
  "count": 1
}
```

**Error Response:**
```json
{
  "error": "Error message"
}
```

## Development

### Project Structure
```
cerebro-api/
├── main.go          # Main application file
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksums
└── README.md        # This file
```

### Next Steps for Enhancement

1. **Add Database Integration**
   - Install GORM: `go get gorm.io/gorm gorm.io/driver/sqlite`
   - Replace in-memory storage with a database

2. **Add Authentication**
   - Install JWT library: `go get github.com/golang-jwt/jwt/v5`
   - Implement authentication middleware

3. **Add Configuration Management**
   - Install Viper: `go get github.com/spf13/viper`
   - Create config files for different environments

4. **Add Logging**
   - Install Logrus: `go get github.com/sirupsen/logrus`
   - Implement structured logging

5. **Add Input Validation**
   - Use Gin's built-in validation
   - Add custom validation rules

6. **Add Testing**
   - Write unit tests for handlers
   - Add integration tests

## Useful Go Libraries for APIs

- **Web Framework:** `github.com/gin-gonic/gin`
- **ORM:** `gorm.io/gorm`
- **Database Drivers:** `gorm.io/driver/postgres`, `gorm.io/driver/mysql`
- **Configuration:** `github.com/spf13/viper`
- **Logging:** `github.com/sirupsen/logrus`
- **Authentication:** `github.com/golang-jwt/jwt/v5`
- **Testing:** Built-in `testing` package + `github.com/stretchr/testify`

