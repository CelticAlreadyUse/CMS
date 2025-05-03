# CMS API

A Content Management System API with JWT authentication, category, news, comment, and custom page management.

## Project Structure

- `internal/model/` — Data models and interfaces
- `internal/repository/` — Database access logic
- `internal/usecase/` — Business logic
- `internal/delivery/http/` — HTTP handlers (Gin)
- `internal/cmd/` — Application entrypoints

## How to Run

1. Configure your database in `config.yaml`.
2. Run migrations:
   ```bash
   go run main.go migrate
   ```
3. Start the server:
   ```bash
   go run main.go httpsrv
   ```

## API Endpoints

### Category
- `GET /v1/categories` — List categories
- `GET /v1/categories/:id` — Get category detail
- `POST /v1/categories` — Create category (auth required)
- `PUT /v1/categories/:id` — Update category (auth required)
- `DELETE /v1/categories/:id` — Delete category (auth required)

### News
- `GET /v1/news` — List news
- `GET /v1/news/:id` — Get news detail
- `POST /v1/news` — Create news (auth required)
- `PUT /v1/news/:id` — Update news (auth required)
- `DELETE /v1/news/:id` — Delete news (auth required)

### Comments
- `POST /v1/comments` — Create comment (auth optional, anonymous allowed)
  - Request: `{ "news_id": int, "comment": string }`
  - Response: `{ "data": { ...comment }, "message": "Comment created successfully" }`
- `DELETE /v1/comments/:id` — Delete comment
  - Response: `{ "message": "Comment deleted successfully" }`

### Custom Pages
- `GET /v1/pages` — List all custom pages
- `GET /v1/pages/:id` — Get custom page by ID
- `GET /v1/pages/url/:url` — Get custom page by URL
- `POST /v1/pages` — Create new custom page
  - Request: `{ "title": string, "url": string, "content": string }`
  - Response: `{ "data": page_id, "message": "Custom page created successfully" }`
- `PUT /v1/pages/:id` — Update custom page
  - Request: `{ "title": string, "url": string, "content": string }`
  - Response: `{ "message": "Custom page updated successfully" }`
- `DELETE /v1/pages/:id` — Delete custom page
  - Response: `{ "message": "Custom page deleted successfully" }`

## Response Format
All endpoints return JSON. Standard response:
```json
{
  "data": ..., // optional
  "message": "...", // optional
  "error": "..." // only on error
}
```

## Error Handling & Typos
- All error messages and response fields are consistently named (`error`, `message`).
- HTTP status codes follow REST conventions (200 OK, 201 Created, 400 Bad Request, 404 Not Found, 500 Internal Server Error).
- All panic sources in comment POST have been fixed with proper error handling and type assertion.

---

For more details, see the code in each respective directory.
