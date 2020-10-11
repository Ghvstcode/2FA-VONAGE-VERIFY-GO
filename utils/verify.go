package verify

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func createClient() *vonage.VerifyClient{
	Key, _ := os.LookupEnv("API_KEY")
	Secret, _ := os.LookupEnv("API_SECRET")

	auth := vonage.CreateAuthFromKeySecret(Key, Secret)

	client := vonage.NewVerifyClient(auth)

	return client
}

func VerStart(phoneNumber string) string{
	client := createClient()

	verification, _, err := client.Request(phoneNumber, "Go-Tut 2FA", vonage.VerifyOpts{
		CodeLength: 6,
	})

	if err != nil {
		log.Fatal(err)
	}
	return verification.RequestId
}

func VerCheck(reqId, code string) string{
	client := createClient()
	response, _, err := client.Check(reqId, code)
	if err != nil {
		log.Fatal(err)
	}

	return response.Status
}


