package config

type ConfigFile struct {
	ShowTerminatedNodes bool                `json:"ShowTerminatedNodes"`
	NodeConfigurations  []NodeConfiguration `json:"NodeConfigurations"`
}

type NodeConfiguration struct {
	ConfigurationName  string   `json:"ConfigurationName"`
	SubnetNameTag      string   `json:"SubnetNameTag"`
	SecurityGroupNames []string `json:"SecurityGroupNames"`
	AMI                string   `json:"AMI"`
	KeyPair            string   `json:"KeyPair"`
	InstanceType       string   `json:"InstanceType"`
	StorageSize        string   `json:"StorageSize"`
}
