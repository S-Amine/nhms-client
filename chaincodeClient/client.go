package chaincodeclient

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/S-Amine/nhms-client/models"
	"github.com/S-Amine/nhms-client/settings"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GetPatient(contract *client.Contract, nin string) {
	result, err := contract.EvaluateTransaction("ReadPatient", nin)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}

	var patient models.Patient
	err = json.Unmarshal(result, &patient)
	if err != nil {
		log.Fatalf("Failed to unmarshal patient data: %v", err)
	}

	fmt.Printf("Patient Data: %+v\n", patient)
}
func GetAllPatients(contract *client.Contract) {
	result, err := contract.EvaluateTransaction("GetAllPatients")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}

	var patients []models.Patient
	err = json.Unmarshal(result, &patients)
	if err != nil {
		log.Fatalf("Failed to unmarshal patients data: %v", err)
	}

	for _, patient := range patients {
		fmt.Printf("Patient: %+v\n", patient)
	}
}
func PublishPatient(contract *client.Contract, nin, firstName, lastName, dob, sex, motherNIN, fatherNIN, familyHistory, allergy, chronicIllnesses, amendedFrom string) {
	_, err := contract.SubmitTransaction("CreatePatient", nin, firstName, lastName, dob, sex, motherNIN, fatherNIN, familyHistory, allergy, chronicIllnesses, amendedFrom)
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v", err)
	}

	fmt.Println("Patient record published successfully.")
}

func NewGrpcConnection() *grpc.ClientConn {
	certificatePEM, err := os.ReadFile(settings.TlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read TLS certificate file: %v", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		log.Fatalf("Failed to parse certificate: %v", err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, settings.GatewayPeer)

	connection, err := grpc.Dial(settings.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}

	return connection
}

func NewIdentity() *identity.X509Identity {
	certificatePEM, err := os.ReadFile(path.Join(settings.CertPath, "User1@org1.example.com-cert.pem"))
	if err != nil {
		log.Fatalf("Failed to read certificate file: %v", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		log.Fatalf("Failed to parse certificate: %v", err)
	}

	id, err := identity.NewX509Identity(settings.MspID, certificate)
	if err != nil {
		log.Fatalf("Failed to create identity: %v", err)
	}

	return id
}

func NewSign() identity.Sign {
	privateKeyPEM, err := os.ReadFile(path.Join(settings.KeyPath, "priv_sk"))
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		log.Fatalf("Failed to create sign function: %v", err)
	}

	return sign
}
