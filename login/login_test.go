package login

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	email    string
	password string
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("please set .env file!!", err)
	}

	email = os.Getenv("email")
	password = os.Getenv("password")
}
func TestHoge(t *testing.T) {
	c := NewClient()
	err := c.login(email, password)
	if err != nil {
		log.Println(err)
	}
}
