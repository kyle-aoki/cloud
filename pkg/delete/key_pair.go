package delete

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func DeleteKeyPairs() {
	client := amazon.EC2Client()
	names := args.Collect()

	for _, name := range names {
		_, err := client.DeleteKeyPair(&ec2.DeleteKeyPairInput{
			KeyName: &name,
		})
		util.Check(err)
		fmt.Println(name)
	}
}
