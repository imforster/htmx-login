# Login App with HTMX and Go

This is a simple login application that demonstrates how to use HTMX with a Go backend. The application features:

- A login form that submits data as JSON
- A Go backend that processes the login request
- HTMX for handling the response without page reload

## Project Structure

```
login-app/
├── index.html    # Frontend with HTMX
├── main.go       # Go backend server
└── README.md     # This file
```

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the project directory
3. Run the Go server:

```bash
go run main.go
```

4. Open your browser and visit: http://localhost:8080

## How It Works

1. The login form uses HTMX to submit data as JSON to the `/login` endpoint
2. The Go server processes the request and returns an HTML fragment
3. HTMX swaps the response into the page without a full page reload

## Features

- JSON request/response handling
- Form validation
- Dynamic content updates with HTMX
- Clean separation of frontend and backend

## Notes

This is a demonstration application and does not implement actual authentication. In a production environment, you would:

- Use HTTPS
- Implement proper password hashing
- Use sessions or JWT for authentication
- Add CSRF protection
- Store user credentials in a database
