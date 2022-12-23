package resource

import (
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

const cloudLabResource = "cloudlab-resource"

func CreateTagSpecs(resourceType string, kvs map[string]string) []*ec2.TagSpecification {
	kvs[cloudLabResource] = "true"
	var tagSlice []*ec2.Tag
	for key := range kvs {
		tagSlice = append(tagSlice, &ec2.Tag{
			Key: util.Ptr(key), Value: util.Ptr(kvs[key]),
		})
	}
	return []*ec2.TagSpecification{{
		ResourceType: &resourceType,
		Tags:         tagSlice,
	}}
}

func FindNameTagValue(tags []*ec2.Tag) *string {
	for _, tag := range tags {
		if tag.Key != nil && *tag.Key == "Name" {
			return tag.Value
		}
	}
	return nil
}

func TagEquals(tags []*ec2.Tag, key string, value string) bool {
	for _, tag := range tags {
		if tag.Key != nil && *tag.Key == key {
			if tag.Value != nil && *tag.Value == value {
				return true
			}
		}
	}
	return false
}

func NameTagEquals(tags []*ec2.Tag, name string) bool {
	nameTagValue := FindNameTagValue(tags)
	if nameTagValue != nil && *nameTagValue == name {
		return true
	}
	return false
}

func CreateNameTagSpec(resourceType string, name string) []*ec2.TagSpecification {
	return []*ec2.TagSpecification{{
		ResourceType: util.StrPtr(resourceType),
		Tags:         CreateNameTagArray(name),
	}}
}

func CreateNameTagArray(name string) []*ec2.Tag {
	return []*ec2.Tag{{
		Key:   util.StrPtr("Name"),
		Value: util.StrPtr(name),
	}}
}
