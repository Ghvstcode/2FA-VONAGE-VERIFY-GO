package verify

import (
	"log"
	"net/http"
	"os"
	//"sync"

	"github.com/joho/godotenv"
	"github.com/nexmo-community/nexmo-go"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func createClient() *nexmo.Client{
	Key, _ := os.LookupEnv("API_KEY")
	Secret, _ := os.LookupEnv("API_SECRET")
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(Key, Secret)
	client := nexmo.NewClient(http.DefaultClient, auth)
	return client
}

func VerStart(phoneNumber string) string{
	client := createClient()
	verification, _, err := client.Verify.Start(nexmo.StartVerificationRequest{
		Number: phoneNumber,
		Brand:  "Go-Tut 2FA",
	})
	if err != nil {
		log.Fatal(err)
	}
	return verification.RequestID
}

func VerCheck(reqId, code string) string{
	client := createClient()
	response, _, err := client.Verify.Check(nexmo.CheckVerificationRequest{
		RequestID: reqId,
		Code:      code,
	})
	if err != nil {
		log.Fatal(err)
	}

	return response.Status
}


