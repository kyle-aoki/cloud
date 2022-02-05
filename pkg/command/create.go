package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Create() {
	launchTemp := args.Poll()
	instanceName := args.Poll()
	client := amazon.EC2Client()

	client.CreateLaunchTemplate(&ec2.CreateLaunchTemplateInput{
		LaunchTemplateData: &ec2.RequestLaunchTemplateData{
			
		},
	})

	rio, err := client.RunInstances(&ec2.RunInstancesInput{
		MinCount: util.IntToInt64Ptr(1),
		MaxCount: util.IntToInt64Ptr(1),
		LaunchTemplate: &ec2.LaunchTemplateSpecification{
			LaunchTemplateName: util.StrPtr(launchTemp),
		},
		TagSpecifications: []*ec2.TagSpecification{{
			ResourceType: util.StrPtr("instance"),
			Tags:         []*ec2.Tag{{Key: util.StrPtr("Name"), Value: util.StrPtr(instanceName)}},
		}},
	})
	util.MustExec(err)

	fmt.Println(*rio.Instances[0].InstanceId)
}
