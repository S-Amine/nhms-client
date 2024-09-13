package main

import (
	"log"
	"time"

	chaincodeclient "github.com/S-Amine/nhms-client/chaincodeClient"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func main() {
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

	// Call functions
	chaincodeclient.PublishPatient(contract, "111223333", "Alice", "Johnson", "1995-05-05", "F", "987654321", "123456789", "None", "None", "None", "")
	chaincodeclient.GetPatient(contract, "111223333")
	chaincodeclient.GetAllPatients(contract)
}
