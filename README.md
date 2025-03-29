# GitHub MCP Server in Go

A Go implementation of the GitHub Model Context Protocol (MCP) server. This implementation allows AI assistants to interact with the GitHub API to perform operations such as creating repositories, managing branches, manipulating files, and more.

## Prerequisites

- Go 1.21 or higher
- A GitHub personal access token with appropriate permissions

## Installation

```bash
go get github.com/metoro-k8s/github-mcp-server-go
```

## Authentication

The server supports two methods of authentication:

### Environment Variable Authentication

Set your GitHub personal access token as an environment variable:

```bash
export GITHUB_PERSONAL_ACCESS_TOKEN=your_github_token
```

### HTTP Header Authentication

The server can also extract authentication tokens from HTTP requests. You can pass your GitHub token via the Authorization header:

```
Authorization: Bearer your_github_token
```

or simply:

```
Authorization: your_github_token
```

#### Context Passthrough

For HTTP handlers, the server supports both:

1. The standard `http_request` context value for HTTP requests.
2. The `ginContext` value when using the Gin framework.

This enables seamless integration with different web frameworks while maintaining a consistent authentication mechanism.

## Usage

1. Set your GitHub personal access token (as described in the Authentication section).

2. Run the server:

```bash
go run main.go
```

## Available Tools

The server provides the following tools:

- **search_repositories**: Search for GitHub repositories
- **create_repository**: Create a new GitHub repository in your account
- **fork_repository**: Fork a GitHub repository to your account or specified organization
- **create_branch**: Create a new branch in a GitHub repository
- **get_file_contents**: Get the contents of a file or directory from a GitHub repository
- **create_or_update_file**: Create or update a single file in a GitHub repository
- **push_files**: Push multiple files to a GitHub repository in a single commit
- **create_issue**: Create a new issue in a GitHub repository
- **get_issue**: Get details of a specific issue in a GitHub repository
- **list_issues**: List issues in a GitHub repository with filtering options
- **update_issue**: Update an existing issue in a GitHub repository
- **add_issue_comment**: Add a comment to an existing issue
- **list_commits**: Get list of commits of a branch in a GitHub repository
- **search_code**: Search for code across GitHub repositories
- **search_issues**: Search for issues and pull requests across GitHub repositories
- **search_users**: Search for users on GitHub

## Development

### Project Structure

- `main.go`: Entry point for the application
- `common/`: Common utilities and error handling
- `operations/`: GitHub API operations implementation
- `tools/`: MCP tool definitions and handlers

### Building from Source

```bash
go build -o github-mcp-server
```

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License. 