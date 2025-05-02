# CMS API

This is a Content Management System API with user authentication using JWT and CRUD operations for categories.

## Project Structure

The project follows a clean architecture approach with the following components:

- **Handler**: HTTP handlers to process requests and responses
- **Usecase**: Application business logic
- **Repository**: Data access layer using GORM
- **Model**: Data structures with GORM annotations
- **Cmd**: CLI commands using Cobra

## Features

- JWT-based authentication
- GORM for ORM with MySQL
- Viper for configuration management
- Cobra for CLI commands
- RESTful API for categories

## Setup

1. Clone the repository
2. Copy `config.example.yaml` to `config.yaml` and update with your configuration
3. Set up the database
   ```
   mysql -u root -p
   CREATE DATABASE cms_db;
   exit;
   ```
4. Run database migrations
   ```
   go run main.go migrate --up
   ```
5. Create a user
   ```
   go run main.go user create --username admin --email admin@example.com --password password
   ```
6. Run the application
   ```
   go run main.go serve
   ```

## Environment Variables

You can configure the application using environment variables instead of config file:

```
CMS_SERVER_HOST=localhost
CMS_SERVER_PORT=8080
CMS_DATABASE_HOST=localhost
CMS_DATABASE_PORT=3306
CMS_DATABASE_USERNAME=root
CMS_DATABASE_PASSWORD=password
CMS_DATABASE_NAME=cms_db
CMS_JWT_SECRET_KEY=your-secret-key-here
CMS_JWT_EXPIRATION_MINUTES=60
```

## Available Commands

```
  help        Help about any command
  migrate     Run database migrations
  serve       Start the API server
  user        Manage users
```

### Migration Commands

```
go run main.go migrate --up    # Apply migrations
go run main.go migrate --down  # Roll back migrations
```

### User Commands

```
go run main.go user create --username admin --email admin@example.com --password password
```

## API Endpoints

### Authentication

- **POST** `/api/login`: Log in with username and password

### Categories (JWT protected)

- **POST** `/api/categories`: Create a new category
- **GET** `/api/categories`: Get all categories
- **GET** `/api/categories/:id`: Get a category by ID
- **PUT** `/api/categories/:id`: Update a category
- **DELETE** `/api/categories/:id`: Delete a category

## Example Requests

### Login

```
POST /api/login
Content-Type: application/json

{
  "username": "admin",
  "password": "password"
}
```

### Create Category

```
POST /api/categories
Content-Type: application/json
Authorization: Bearer <jwt-token>

{
  "name": "Technology",
  "description": "Technology related content"
}
```
