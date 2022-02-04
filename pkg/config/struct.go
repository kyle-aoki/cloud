package config

type ConfigFile struct {
	Variables              Variables               `json:"Variables"`
	NodeConfigs []InstanceConfiguration `json:"InstanceConfigurations"`
}

type InstanceConfiguration struct {
	Name           string   `json:"Name"`
	Subnet         string   `json:"Subnet"`
	SecurityGroups []string `json:"SecurityGroups"`
	KeyPair        string   `json:"KeyPair"`
	InstanceType   string   `json:"InstanceType"`
	StorageSize    string   `json:"StorageSize"`
	Ami            string   `json:"AMI"`
	Username       string   `json:"Username"`
}

type Variables struct {
	Subnets        []KeyValue `json:"Subnets"`
	SecurityGroups []KeyValue `json:"SecurityGroups"`
	AMIs           []KeyValue `json:"AMIs"`
}

type KeyValue struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}
