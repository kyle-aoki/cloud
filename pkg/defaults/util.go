package defaults

import (
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func findNameTagValue(tags []*ec2.Tag) *string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return tag.Value
		}
	}
	return nil
}

func nameTagEquals(tags []*ec2.Tag, name string) bool {
	nameTagValue := findNameTagValue(tags)
	if nameTagValue != nil && *nameTagValue == name {
		return true
	}
	return false
}

func createNameTag(resourceType string, name string) []*ec2.TagSpecification {
	return []*ec2.TagSpecification{{
		ResourceType: util.StrPtr(resourceType),
		Tags: []*ec2.Tag{{
			Key:   util.StrPtr("Name"),
			Value: util.StrPtr(name),
		}}},
	}
}
