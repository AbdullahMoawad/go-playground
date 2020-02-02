package ops

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"real-estate/models"
	"real-estate/server"

	"os"
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
			server.Conn().AutoMigrate(&models.User{})
		case "estate":
			server.Conn().AutoMigrate(&models.Property{})
		case "session":
			server.Conn().AutoMigrate(&models.Session{})
		case "category":
			server.Conn().AutoMigrate(&models.Category{})
		default:
			fmt.Println("This model hasn't created yet :(")
		}
		log.Println("Successfully Migrated :)")
	},
}

func init() { command.AddCommand(pgMigrate) }

func Execute() {
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
