package amazon

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ec2Session *ec2.EC2

func InitEC2Client() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Session = ec2.New(sess)
}

func EC2() *ec2.EC2 {
	return ec2Session
}
