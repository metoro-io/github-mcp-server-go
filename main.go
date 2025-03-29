package main

import (
	"fmt"
	"os"

	mcpgolang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/metoro-k8s/github-mcp-server-go/tools"
)

func main() {
	// Check if the appropriate environment variables are set
	if err := checkEnvVars(); err != nil {
		panic(err)
	}

	done := make(chan struct{})

	mcpServer := mcpgolang.NewServer(stdio.NewStdioServerTransport())

	// Add tools
	for _, tool := range tools.GitHubToolsList {
		err := mcpServer.RegisterTool(tool.Name, tool.Description, tool.Handler)
		if err != nil {
			panic(err)
		}
	}

	err := mcpServer.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}

func checkEnvVars() error {
	if os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN") == "" {
		return fmt.Errorf("GITHUB_PERSONAL_ACCESS_TOKEN environment variable not set")
	}
	return nil
}
