package common

type OrgStatus uint

const (
	_ OrgStatus = iota
	Created
	NotCreated
)

type ConfigType uint

const (
	_ ConfigType = iota
	SetupNetworkConfig
	JoinNetworkConfig
)

const (
	ADD = iota + 1
	DELETE
)

type PublishInfo struct {
	Org         string
	MSPID       string
	OrgFullName string
	Peers       NodeInfo
	Orderers    NodeInfo
	MSPCert     []byte
	Status      OrgStatus
	ConfigType  ConfigType
	Attribute   map[string]string
	Version     int `json:"version"`
}

// NodeInfo ...
type NodeInfo struct {
	// TODO: Cert改为TLSCACrt
	Cert  []byte
	Nodes []*RemoteNode
}

type RemoteNode struct {
	ID           string
	Address      string
	Override     string
	TLSServerCrt []byte
}

// ChangeNodes ...
type ChangeNodes struct {
	OrgName     string
	ReceiveOrgs []string
	Type        int
	Nodes       []string
}
