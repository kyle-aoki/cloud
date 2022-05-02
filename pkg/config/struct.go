package config

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type ConfigVars struct {
	ShowTerminatedNodes bool `yaml:"ShowTerminatedNodes"`
	Configurations  []Configuration `yaml:"Configurations"`
}

type Configuration struct {
	ConfigName         string   `yaml:"Name"`
	VPCNameTag         string   `yaml:"VPCNameTag"`
	SubnetNameTag      string   `yaml:"SubnetNameTag"`
	SecurityGroupNames []string `yaml:"SecurityGroupNames"`
	AMI                string   `yaml:"AMI"`
	KeyPair            string   `yaml:"KeyPair"`
	InstanceType       string   `yaml:"InstanceType"`
	StorageSize        string   `yaml:"StorageSize"`
	PrivateIp          string   `yaml:"PrivateIp"`
	UserData           []string `yaml:"UserData"`
}

func (nc Configuration) VPC() *string {
	client := amazon.EC2Client()
	dvo, err := client.DescribeVpcs(&ec2.DescribeVpcsInput{})
	util.MustExec(err)
	for _, vpc := range dvo.Vpcs {
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				if *tag.Value == nc.VPCNameTag {
					return vpc.VpcId
				}
			}
		}
	}
	panic("VPC not found.")
}

func (nc Configuration) StorageSizeToInt64() *int64 {
	ss := strings.Replace(nc.StorageSize, "gb", "", 1)
	num, err := strconv.Atoi(ss)
	util.MustExec(err)
	i64 := int64(num)
	return &i64
}

func (nc Configuration) DefaultDeviceName() *string {
	client := amazon.EC2Client()
	dio, err := client.DescribeImages(&ec2.DescribeImagesInput{
		ImageIds: []*string{&nc.AMI},
	})
	util.MustExec(err)

	return dio.Images[0].RootDeviceName
}

func (nc Configuration) SubnetId() *string {
	client := amazon.EC2Client()
	dso, err := client.DescribeSubnets(&ec2.DescribeSubnetsInput{})
	util.MustExec(err)

	if len(dso.Subnets) == 0 {
		panic("No subnets exist.")
	}

	for _, subnet := range dso.Subnets {
		for _, tag := range subnet.Tags {
			if *tag.Key == "Name" {
				if *tag.Value == nc.SubnetNameTag {
					return subnet.SubnetId
				}
			}
		}
	}

	panic("Subnet not found")
}

func (nc Configuration) SecurityGroupIds() []*string {
	client := amazon.EC2Client()
	dsgo, err := client.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{})
	util.MustExec(err)
	var ids []*string
	for _, sg := range dsgo.SecurityGroups {
		if contains(nc.SecurityGroupNames, *sg.GroupName) {
			ids = append(ids, sg.GroupId)
		}
	}
	return ids
}

func contains(names []string, sg string) bool {
	for i := range names {
		if names[i] == sg {
			return true
		}
	}
	return false
}

func (cv ConfigVars) Find(configName string) Configuration {
	for _, nc := range cv.Configurations {
		if nc.ConfigName == configName {
			return nc
		}
	}
	panic("Node configuration not found.")
}

func (nc Configuration) GetUserData() *string {
	var userData string
	for i := range nc.UserData {
		userData += nc.UserData[i] + "\n"
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(userData))
	return &b64
}

func (nc Configuration) GetPrivateIp() *string {
	if nc.PrivateIp == "" {
		return nil
	}
	return &nc.PrivateIp
}
