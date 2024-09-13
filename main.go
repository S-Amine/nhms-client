package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	chaincodeclient "github.com/S-Amine/nhms-client/chaincodeClient"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func main() {
	// Define command-line flags
	action := flag.String("action", "", `Specify the action to perform:
  - 'publish' : Publish a new patient record to the blockchain.
  - 'get'     : Retrieve a specific patient record by ID.
  - 'get-all' : Retrieve all patient records.`)

	patientID := flag.String("patient-id", "", "Unique patient ID (required for 'publish' and 'get').")
	firstName := flag.String("first-name", "", "First name of the patient (required for 'publish').")
	lastName := flag.String("last-name", "", "Last name of the patient (required for 'publish').")
	dob := flag.String("dob", "", "Date of birth in YYYY-MM-DD format (required for 'publish').")
	sex := flag.String("sex", "", "Gender of the patient (required for 'publish').")
	motherNIN := flag.String("mother-nin", "", "National Identification Number of the patient's mother (required for 'publish').")
	fatherNIN := flag.String("father-nin", "", "National Identification Number of the patient's father (required for 'publish').")
	familyHistory := flag.String("family-history", "None", "Family medical history (optional for 'publish').")
	allergies := flag.String("allergies", "None", "Known allergies (optional for 'publish').")
	chronicIllnesses := flag.String("chronic-illnesses", "None", "Chronic illnesses (optional for 'publish').")
	bilans := flag.String("amended-from", "", "Who submitted the doc ?).")

	// Parse the flags
	flag.Parse()

	// Validate the action
	if *action == "" {
		fmt.Println("Error: Action is required.\n")
		fmt.Println("Usage:")
		flag.Usage()
		os.Exit(1)
	}

	// Establish the gRPC client connection
	clientConnection := chaincodeclient.NewGrpcConnection()
	defer clientConnection.Close()

	// Create identity and sign objects for the gateway
	id := chaincodeclient.NewIdentity()
	sign := chaincodeclient.NewSign()

	// Connect to the gateway using the identity and sign objects
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to create gateway connection: %v", err)
	}
	defer gw.Close()

	// Get the network and contract
	network := gw.GetNetwork("mychannel")
	contract := network.GetContract("basic")

	// Perform actions based on the command-line flag
	switch *action {
	case "publish":
		// Validate required flags for 'publish'
		if *patientID == "" || *firstName == "" || *lastName == "" || *dob == "" || *sex == "" || *motherNIN == "" || *fatherNIN == "" {
			fmt.Println("Error: Missing required fields for 'publish'.")
			fmt.Println("Required fields for 'publish': -patient-id, -first-name, -last-name, -dob, -sex, -mother-nin, -father-nin.")
			flag.Usage()
			os.Exit(1)
		}
		// Call the publish function
		chaincodeclient.PublishPatient(contract, *patientID, *firstName, *lastName, *dob, *sex, *motherNIN, *fatherNIN, *familyHistory, *allergies, *chronicIllnesses, *bilans)
		fmt.Printf("Patient with ID %s published successfully.\n", *patientID)

	case "get":
		// Validate required flags for 'get'
		if *patientID == "" {
			fmt.Println("Error: 'patient-id' is required for 'get' action.")
			flag.Usage()
			os.Exit(1)
		}
		// Call the get function
		chaincodeclient.GetPatient(contract, *patientID)

	case "get-all":
		// Call the get all function
		chaincodeclient.GetAllPatients(contract)

	default:
		fmt.Println("Error: Invalid action. Supported actions are 'publish', 'get', and 'get-all'.")
		flag.Usage()
		os.Exit(1)
	}
}
