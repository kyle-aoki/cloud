package create

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/config"
	"cloud/pkg/util"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateNode() {
	configurationName := args.Poll()
	instanceNames := args.Collect()
	
	cfgWithVars := config.FindInstanceConfiguration(configurationName)
	cfg := config.ReplaceVariables(cfgWithVars)
	client := amazon.EC2Client()
	defaultDeviceName := DefaultDeviceName(client, cfg)
	
	for _, instanceName := range instanceNames {
		rii := ec2.RunInstancesInput{
			BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
				Ebs:        &ec2.EbsBlockDevice{VolumeSize: ParseStorageSize(cfg)},
				DeviceName: defaultDeviceName,
			}},
			KeyName:          &cfg.KeyPair,
			ImageId:          &cfg.Ami,
			MinCount:         util.IntToInt64Ptr(1),
			MaxCount:         util.IntToInt64Ptr(1),
			SecurityGroupIds: util.StrSlicePtr(cfg.SecurityGroups),
			SubnetId:         &cfg.Subnet,
			InstanceType:     &cfg.InstanceType,
			TagSpecifications: []*ec2.TagSpecification{
				{
					ResourceType: util.StrPtr("instance"),
					Tags:         []*ec2.Tag{{Key: util.StrPtr("Name"), Value: &instanceName}},
				},
			},
		}

		_, err := client.RunInstances(&rii)
		util.Check(err)

		fmt.Println(instanceName)
	}
}

func DefaultDeviceName(client *ec2.EC2, cfg config.InstanceConfiguration) *string {
	defer handleDefaultDeviceNamePanic(cfg.Ami)
	dii := ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{Name: util.StrPtr("image-id"), Values: []*string{&cfg.Ami}},
		},
	}
	describeImagesOutput, err := client.DescribeImages(&dii)
	util.Check(err)
	return describeImagesOutput.Images[0].BlockDeviceMappings[0].DeviceName
}

func handleDefaultDeviceNamePanic(ami string) {
	if r := recover(); r != nil {
		util.PanicVerify("Failed to obtain default device name for AMI: %v", ami)
	}
}

func ParseStorageSize(cfg config.InstanceConfiguration) *int64 {
	storageSize := cfg.StorageSize
	l := len(storageSize)

	if l < 3 {
		util.PanicVerify("Invalid storage size. Found: '%v'.\nTry '8gb'.", storageSize)
	}

	endsInGB := storageSize[l-2] == 'g' && storageSize[l-1] == 'b'
	if !endsInGB {
		util.PanicVerify("Storage size must end in 'gb'. Example: '8gb'. Found: '%v'", storageSize)
	}

	storageSize = storageSize[:l-2]

	integer, err := strconv.Atoi(storageSize)
	if err != nil {
		util.PanicVerify("Failed to parse storage size. Found '%v'.\nTry '8gb'.", storageSize)
	}

	return util.IntToInt64Ptr(integer)
}
