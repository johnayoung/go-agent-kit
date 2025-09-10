# go-agent-kit

A brief description of your project.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22 or later

### Installing

1. Clone the repository
2. Run `go mod tidy` to download dependencies
3. Build the project with `go build ./cmd/go-agent-kit`

## Commands

- `init`: Initialize a new project. Sets up necessary files and directories, including the .github folder with .github/copilot-instructions.md and .github/prompts/<prompt_name>.prompt.md

## Prompts

These are to be used in the sidebar of the Github Copilot Chat extension.

- `/specify`: Use the /specify command to describe what you want to build. Focus on the what and why, not the tech stack.
- `/plan`: Use the /plan command to provide your tech stack and architecture choices.
- `/feat`: Use the /feat command to describe new features you want to add.
- `/implement`: Use the /implement command to implement the specified features.
- `/verify`: runs the stability prompt. sample stability prompt included in reference/prompts/stability_prompt.md

## Running the tests

Run `go test ./...` to execute the tests.

## Built With

* [Go](https://golang.org/) - The programming language used

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
