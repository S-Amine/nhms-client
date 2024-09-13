package settings

const (
	MspID        = "Org1MSP"
	CryptoPath   = "/home/nhms/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	CertPath     = CryptoPath + "/users/User1@org1.example.com/msp/signcerts"
	KeyPath      = CryptoPath + "/users/User1@org1.example.com/msp/keystore"
	TlsCertPath  = CryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	PeerEndpoint = "localhost:7051"
	GatewayPeer  = "peer0.org1.example.com"
)
