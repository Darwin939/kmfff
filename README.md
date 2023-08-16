# HTTP Proxy Server

This project is an HTTP proxy server that proxies HTTP requests to 3rd-party services. It listens for incoming HTTP requests and expects the request's body to be in JSON format. The server then forms a valid HTTP request to the 3rd-party service using the data from the client's message and responds to the client with a JSON object containing relevant information.

## Getting Started

1. Run the following command to build and start the server using Docker Compose:

   ```bash
   docker compose up --build

2. check it out in localhost:8000/proxy
