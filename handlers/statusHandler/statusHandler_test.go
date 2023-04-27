package statusHandler

import (
	"assignment-2/database"
	"assignment-2/structs"
	"assignment-2/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

// TestMain starts a main sequence for the test file
func TestMain(m *testing.M) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err2 := os.Chdir(workDir + "/../..")
	if err2 != nil {
		return
	}

	database.InitFirestore()
	defer func() {
		err := database.CloseDB()
		if err != nil {
			log.Printf("Error in closing database: %s", err)
		}
	}()
	m.Run()
}

/*
Test function for status endpoint.
*/
func TestStatusEndpoint(t *testing.T) {
	//	webhookAmount, err := database.GetWebhookAmount()
	//	if err != nil {
	//		log.Printf("Error in getting webhook amount: %s", err)
	//	}
	//	fmt.Print(webhookAmount)
}

func TestUptime(t *testing.T) {
	time.Sleep(2 * time.Second)
	assert.Equal(t, 2, utils.Uptime())
}

func TestStatusHandler(t *testing.T) {
	handleStatusRequest()
	status := Status()

	expected := structs.Status{
		CountriesApi:    "200 OK",
		MarkdownHtmlApi: "200 OK",
		NotificationDB:  "200 OK",
		Webhooks:        database.GetWebhookAmount(),
		Version:         "v1",
		Uptime:          utils.Uptime(),
	}

	assert.Equal(t, expected, status)

}
