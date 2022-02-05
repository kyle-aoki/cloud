package command

import "github.com/aws/aws-sdk-go/service/ec2"

func GetName(instance *ec2.Instance) string {
	for _, tag := range instance.Tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}
	return ""
}
