package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	conn "programa3/database"
	database "programa3/database/models"
)

var email string

func main() {
	// Create connexion to database
	conn.Init()
	// Close connexion database
	defer func() {
		if err := conn.Close(conn.GetConn()); err != nil {
			fmt.Println(err)
		}
	}()
	// CI/CD START
	app := &cli.App{
		Name:  "CRUD Notes APP",
		Usage: "Manage a Notes Database",
		Commands: []*cli.Command{{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Connect to user",
			Action: func(context *cli.Context) error {
				if context.NArg() != 2 {
					fmt.Println("Verify the info")
				}
				email = context.Args().Get(0)
				if err := database.GetUserLogin(email, context.Args().Get(1)); err != nil {
					log.Fatal(err)
				}
				InitAPP()
				return nil
			},
		},
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "Register an user",
				Action: func(context *cli.Context) error {
					fmt.Println(context.Args())
					fmt.Println(context.NArg())
					fmt.Println("Register user")
					return nil
				},
			}},
		Action: func(*cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func InitAPP() {
	var inMenu string
	scanner := bufio.NewScanner(os.Stdin)
	// Main menu
main:
	for {
		fmt.Print("==== Menu Principal ====\n1.List Notes\n2.Create Note\n3.Salir\n$ ")
		scanner.Scan()
		switch scanner.Bytes()[0] {
		case 49: // 49 equals 1 in ASCII
			fmt.Println("Listing...")
			notes, err := database.GetNotes(email)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(notes)
		case 50: // 50 equals 2 in ASCII
			fmt.Println("Creating...")
		case 51: // 51 equals 3 in ASCII
			fmt.Println("Closing...")
		default:
			fmt.Println("Nothing")
		}
		inMenu = scanner.Text()
		fmt.Println(inMenu)
		break main
	}
}
