package ops

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"property/models"
	"property/server"
)

var command = &cobra.Command{}

var pgMigrate = &cobra.Command{
	Use: "pg-migrate",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please define which model to migrate ..!")
		}
		model := args[0]
		switch model {
		case "user":
			server.CreatePostgresDbConnection().AutoMigrate(&models.User{})
		case "property":
			server.CreatePostgresDbConnection().AutoMigrate(&models.Property{})
		case "session":
			server.CreatePostgresDbConnection().AutoMigrate(&models.Session{})
		case "category":
			server.CreatePostgresDbConnection().AutoMigrate(&models.Category{})
		default:
			fmt.Println("This model hasn't created yet :(")
		}
		log.Println("Successfully Migrated :)")
	},
}

func init() { command.AddCommand(pgMigrate) }

func Execute() {
	err := command.Execute()
	if err != nil {
		return
	}
}
