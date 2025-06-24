# GOSimpleCMS
A Simple CMS API built with Go

## Installation

### Running on Docker
Clone this code, open terminal, go to your working directory
#### Run Docker Compose
`docker-compose up --build`
### Manual Installation
Clone this code, open terminal, go to your working directory, and run these following commands:
#### Database Migration
`go run main.go --mode=migrate`
#### Database Seeder (optional)
`go run main.go --mode=seed`
#### Run the application server
`go run main.go` or `go run main.go --mode=serve`

## API Documentation
Open this url: `http://localhost:8080/swagger/index.html` on your browser

## Contact
- Sugiarto
- Phone: 081915523100
- Email: sugiarto.dlingo@example.com