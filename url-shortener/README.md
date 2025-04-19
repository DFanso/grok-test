# URL Shortener Server

This is a simple URL shortening server written in Go. It allows you to shorten long URLs into shorter, more manageable links.

## Prerequisites

- Go (version 1.16 or higher) must be installed on your system.

## Installation

1. Clone this repository or copy the code into your project directory.
2. Navigate to the `url-shortener` directory.

## Running the Server

To start the server, run the following command in the terminal:

```bash
go run main.go
```

The server will start on port 8080 by default. You should see a message indicating that the server is running.

## Using the URL Shortener

### Shortening a URL

To shorten a URL, send a POST request to `http://localhost:8080/shorten` with a JSON body containing the URL you want to shorten.

Example using `curl`:

```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://www.example.com/very/long/url"}'
```

Response:

```json
{
  "short_url": "http://localhost:8080/AbCdEf"
}
```

### Accessing the Original URL

To access the original URL, simply visit the shortened URL in your browser or make a GET request to it. For example, if the shortened URL is `http://localhost:8080/AbCdEf`, navigating to this URL will redirect you to the original long URL.

## Notes

- The shortened URLs are stored in memory, so they will be lost if the server restarts. For a production environment, consider using a persistent storage solution like a database.
- The server generates random 6-character short URLs. There is a small chance of collision (generating the same short URL for different long URLs), which is not handled in this basic implementation.

## License

This project is licensed under the MIT License - see the LICENSE file for details (if applicable).
