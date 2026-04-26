# Repo Stat

This project is a microservice-based tool for working with GitHub repositories and repository subscriptions.

## Features

- Fetch GitHub repository information by repository URL
- Create, list, and delete repository subscriptions
- Aggregate repository information for all saved subscriptions
- REST API via the `api` service
- gRPC communication between internal services
- PostgreSQL storage for subscriptions
- Swagger web interface for API testing

## Services

- `api` - HTTP API for clients
- `subscriber` - subscription storage and GitHub repository existence validation
- `processor` - gRPC layer between `api` and `collector`
- `collector` - GitHub repository data collection service
- `postgres` - database for subscriptions

## Usage

1. Start the services:

   ```bash
   make up
   ```

2. Open the Swagger UI in your browser:

   `http://localhost:28080/swagger/index.html`

3. Stop and remove containers when you are done:

   ```bash
   make down
   ```
