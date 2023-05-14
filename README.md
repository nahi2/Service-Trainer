# Employee Management System API

This API provides endpoints to manage employees and supervisors.

## Requirements

- Go version 1.15 or higher
- Fiber web framework
- godotenv package

## Installation

1. Clone the repository
2. Run `go mod tidy`
3. Run `go run main.go`

## Usage

The API provides the following endpoints:

- `GET /allEmployees` - This endpoint retrieves all employees.
- `GET /allSupervisors` - This endpoint retrieves all supervisors.
- `GET /employee/:id` - This endpoint retrieves an employee by their ID.

## Configuration

To configure the API, set the following environment variables:

- `employee_url`: the URL of the employee management system
