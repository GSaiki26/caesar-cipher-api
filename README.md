# Caesar Cipher API
The Caesar Cipher API is a repository that contains simple APIs for encrypting text using the Caesar Cipher algorithm.

Each API was written in a different programming language, and the following languages were used:
- [Go](/go)
- [Rust](/rust)

All APIs by default, run on port 8080. All projects has same endpoints and has a Dockerfile to build a Docker image.

## Endpoints
### POST /encode
This endpoint receives a JSON with the following structure:
```json
{
  "content": "string"
}
```
And returns a JSON with the following structure:
```json
{
  "shift": "number",
  "text": "string"
}
```
