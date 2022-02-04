package command

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

	cfg := config.Find(configurationName)

	client := amazon.EC2Client()
	defaultDeviceName := DefaultDeviceName(client, cfg)

	for _, instanceName := range instanceNames {
		rio, err := client.RunInstances(&ec2.RunInstancesInput{
			BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
				Ebs:        &ec2.EbsBlockDevice{VolumeSize: ParseStorageSize(cfg)},
				DeviceName: defaultDeviceName,
			}},
			KeyName:        &cfg.KeyPair,
			ImageId:        &cfg.AMI,
			MinCount:       util.IntToInt64Ptr(1),
			MaxCount:       util.IntToInt64Ptr(1),
			SecurityGroups: util.StrSlicePtr(cfg.SecurityGroupNames),
			SubnetId:       &cfg.SubnetNameTag,
			InstanceType:   &cfg.InstanceType,
			TagSpecifications: []*ec2.TagSpecification{
				{
					ResourceType: util.StrPtr("instance"),
					Tags:         []*ec2.Tag{{Key: util.StrPtr("Name"), Value: &instanceName}},
				},
			},
		})
		util.MustExec(err)

		fmt.Println(*rio.Instances[0].InstanceId)
	}
}

func DefaultDeviceName(client *ec2.EC2, cfg config.NodeConfiguration) *string {
	defer func() {
		if r := recover(); r != nil {
			panic(f("Failed to obtain default device name for AMI: %v", cfg.AMI))
		}
	}()

	describeImagesOutput, err := client.DescribeImages(&ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{Name: util.StrPtr("image-id"), Values: []*string{&cfg.AMI}},
		},
	})
	util.MustExec(err)

	return describeImagesOutput.Images[0].BlockDeviceMappings[0].DeviceName
}

func ParseStorageSize(cfg config.NodeConfiguration) *int64 {
	storageSize := cfg.StorageSize
	l := len(storageSize)

	if l < 3 {
		panic(f("Invalid storage size. Found: '%v'.\nTry '8gb'.", storageSize))
	}

	endsInGB := storageSize[l-2] == 'g' && storageSize[l-1] == 'b'
	if !endsInGB {
		panic(f("Storage size must end in 'gb'. Example: '8gb'. Found: '%v'", storageSize))
	}

	storageSize = storageSize[:l-2]

	integer, err := strconv.Atoi(storageSize)
	if err != nil {
		panic(f("Failed to parse storage size. Found '%v'.\nTry '8gb'.", storageSize))
	}

	return util.IntToInt64Ptr(integer)
}
