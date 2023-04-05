package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/exepirit/go-graphql/genqlient/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func main() {
	buildApp().Run(os.Args)
}

func buildApp() *cli.App {
	return &cli.App{
		Name: "GraphQL application client",
		Commands: []*cli.Command{
			{
				Name: "todos",
				Description: "get a list of TODO items",
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Description: "create new TODO",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "userId", Required: true, Usage: "TODO item author ID"},
						},
						Action: createNewTodo,
					},
				},
				Action: handleGetTodosQuery,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "server", Value: "http://localhost:8080", Usage: "remote server address"},
		},
		ExitErrHandler: func(cCtx *cli.Context, err error) {
			panic(err)
		},
	}
}

func handleGetTodosQuery(ctx *cli.Context) error {
	// prepare configuration
	serverAddr, err := url.Parse(ctx.String("server"))
	if err != nil {
		return errors.New("invalid server URL")
	}

	// do query
	graphql := client.NewClient(serverAddr)
	todos, err := client.GetTodos(ctx.Context, graphql)
	if err != nil {
		return err
	}

	// print todos
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredGreenWhiteOnBlack)
	t.AppendHeader(table.Row{"ID", "Text", "Done", "Author"})
	for _, todo := range todos.Todos {
		t.AppendRow(table.Row{
			todo.Id, todo.Text, todo.Done, todo.User.Name,
		})
	}
	t.Render()

	return nil
}

func createNewTodo(ctx *cli.Context) error {
	// prepare configuration
	serverAddr, err := url.Parse(ctx.String("server"))
	if err != nil {
		return errors.New("invalid server URL")
	}

	// do query
	graphql := client.NewClient(serverAddr)
	newTodo := client.NewTodo{
		Text:   ctx.Args().Get(0),
		UserId: ctx.String("userId"),
	}

	createdTodo, err := client.CreateTodo(ctx.Context, graphql, newTodo)
	if err != nil {
		return err
	}

	// print result
	fmt.Printf("Created new TODO: %s \"%s\"\n", createdTodo.CreateTodo.Id, createdTodo.CreateTodo.Text)
	return nil
}
