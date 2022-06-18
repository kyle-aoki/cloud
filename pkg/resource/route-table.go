package resource

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func internetGatewayRouteExistsOnRouteTable(rt *ec2.RouteTable, ig *ec2.InternetGateway) bool {
	for _, route := range rt.Routes {
		if *route.GatewayId == *ig.InternetGatewayId {
			return true
		}
	}
	return false
}

func addInternetGatewayRoute(rt *ec2.RouteTable, ig *ec2.InternetGateway, cidr string) {
	_, err := amazon.EC2().CreateRoute(&ec2.CreateRouteInput{
		RouteTableId:         rt.RouteTableId,
		GatewayId:            ig.InternetGatewayId,
		DestinationCidrBlock: util.StrPtr(cidr),
	})
	util.MustExec(err)
}

func subnetAssociationExistsOnRouteTable(rt *ec2.RouteTable, subnet *ec2.Subnet) bool {
	for _, assoc := range rt.Associations {
		if assoc.SubnetId != nil && *assoc.SubnetId == *subnet.SubnetId {
			return true
		}
	}
	return false
}

func associateSubnetWithRouteTable(rt *ec2.RouteTable, subnet *ec2.Subnet) {
	_, err := amazon.EC2().AssociateRouteTable(&ec2.AssociateRouteTableInput{
		RouteTableId: rt.RouteTableId,
		SubnetId:     subnet.SubnetId,
	})
	util.MustExec(err)
}

func disassociateSubnetsFromRouteTable(rt *ec2.RouteTable) {
	for _, assoc := range rt.Associations {
		if assoc.SubnetId != nil {
			disassociateRouteTable(assoc.RouteTableAssociationId)
		}
	}
}

func disassociateRouteTable(assocId *string) {
	_, err := amazon.EC2().DisassociateRouteTable(&ec2.DisassociateRouteTableInput{
		AssociationId: assocId,
	})
	util.MustExec(err)
}
