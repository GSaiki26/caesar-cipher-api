# Go Caesar Cipher
The GO Caesar Cipher is a simple API that receives a text and returns all the possible shifts of the Caesar Cipher.

By default, the API runs on port 3000.

## Endpoints
### POST /encode
This endpoint receives a JSON with the following structure:
```json
{
  "text": "string"
}
```
And returns a JSON with the following structure:
```json
{
  "results": [
    {
      "shift": "int",
      "text": "string"
    }
  ]
}
```

## Running the project
To run the project, you can choose between running the project with Docker or running the project with Go locally.

### Running with Docker
To run the project with Docker, you need to have Docker installed on your machine. After that, you can run the following command:
```bash
docker build -t go-caesar-cipher .
docker run -p 3000:3000 go-caesar-cipher
```

### Running locally
To run the project locally, you need to have Go installed on your machine. After that, you can run the following command:
```bash
go run main.go
```
