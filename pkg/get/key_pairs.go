package get

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func PrintKeyPairs() {
	client := amazon.EC2Client()

	dkp, err := client.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	util.Check(err)

	for _, kp := range dkp.KeyPairs {
		fmt.Println(*kp.KeyName)
	}
}