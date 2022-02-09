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
	ShowTerminatedNodes bool
	NodeConfigurations  []NodeConfiguration
}

type NodeConfiguration struct {
	ConfigName         string   `json:"ConfigName"`
	VPCNameTag         string   `json:"VPCNameTag"`
	SubnetNameTag      string   `json:"SubnetNameTag"`
	SecurityGroupNames []string `json:"SecurityGroupNames"`
	AMI                string   `json:"AMI"`
	KeyPair            string   `json:"KeyPair"`
	InstanceType       string   `json:"InstanceType"`
	StorageSize        string   `json:"StorageSize"`
	PrivateIp          string   `json:"PrivateIp"`
	UserData           []string `json:"UserData"`
}

func (nc NodeConfiguration) VPC() *string {
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

func (nc NodeConfiguration) StorageSizeToInt64() *int64 {
	ss := strings.Replace(nc.StorageSize, "gb", "", 1)
	num, err := strconv.Atoi(ss)
	util.MustExec(err)
	i64 := int64(num)
	return &i64
}

func (nc NodeConfiguration) DefaultDeviceName() *string {
	client := amazon.EC2Client()
	dio, err := client.DescribeImages(&ec2.DescribeImagesInput{
		ImageIds: []*string{&nc.AMI},
	})
	util.MustExec(err)

	return dio.Images[0].RootDeviceName
}

func (nc NodeConfiguration) SubnetId() *string {
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

func (nc NodeConfiguration) SecurityGroupIds() []*string {
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

func (cv ConfigVars) Find(configName string) NodeConfiguration {
	for _, nc := range cv.NodeConfigurations {
		if nc.ConfigName == configName {
			return nc
		}
	}
	panic("Node configuration not found.")
}

func (nc NodeConfiguration) GetUserData() *string {
	var userData string
	for i := range nc.UserData {
		userData += nc.UserData[i] + "\n"
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(userData))
	return &b64
}

func (nc NodeConfiguration) GetPrivateIp() *string {
	if nc.PrivateIp == "" {
		return nil
	}
	return &nc.PrivateIp
}
