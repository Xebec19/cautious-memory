package main

import (
	"context"
	"os"

	"github.com/Xebec19/cautious-memory/internal/templates"
)

func main() {
	component := templates.Hello("John")
	component.Render(context.Background(), os.Stdout)
}
