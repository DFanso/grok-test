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

The server will start on the address specified in `config.json` (default is `localhost:8080`). You should see a message indicating that the server is running.

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

## Configuration for Deployment

To deploy the server on a different address, domain, or with HTTPS, edit the `config.json` file:

```json
{
    "server_address": "yourdomain.com:port",
    "protocol": "https",
    "bind_address": ":8080"
}
```

- Update the `server_address` field with your desired domain and port for URL generation. If no port is specified, the default port based on the protocol (80 for HTTP, 443 for HTTPS) will be used in generated URLs.
- Update the `protocol` field to `https` if you are using SSL/TLS for secure connections. The default is `http`. This will be used when generating shortened URLs.
- Update the `bind_address` field with the address and port the server should listen on. This can be different from `server_address` for deployment scenarios (e.g., behind a proxy). The default is `:8080`.

Note: Setting up HTTPS requires additional configuration for SSL certificates which is not covered in this basic implementation. Ensure your server is configured with valid certificates before changing the protocol to `https`. Also, binding to privileged ports like 80 or 443 may require elevated permissions.

## Notes

- The shortened URLs are stored in memory, so they will be lost if the server restarts. For a production environment, consider using a persistent storage solution like a database.
- The server generates random 6-character short URLs. There is a small chance of collision (generating the same short URL for different long URLs), which is not handled in this basic implementation.

## License

This project is licensed under the MIT License - see the LICENSE file for details (if applicable).
