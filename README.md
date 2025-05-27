# Login & Registration App with HTMX and Go

This is a simple login and registration application that demonstrates how to use HTMX with a Go backend. The application features:

- A login form that submits data as JSON
- A registration form for new users
- A Go backend that processes both login and registration requests
- HTMX for handling the responses without page reload

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

1. The application provides both login and registration forms
2. Users can toggle between the forms using links
3. Form data is submitted as JSON to the respective endpoints (`/login` or `/register`)
4. The Go server processes the request and returns an HTML fragment
5. The response is displayed on the page without a full page reload

## Features

- JSON request/response handling
- Form validation (including password matching)
- Dynamic content updates with HTMX
- Clean separation of frontend and backend
- Seamless switching between login and registration forms

## Notes

This is a demonstration application and does not implement actual authentication or user storage. In a production environment, you would:

- Use HTTPS
- Implement proper password hashing
- Use sessions or JWT for authentication
- Add CSRF protection
- Store user credentials in a database
- Implement email verification
- Add more robust input validation
