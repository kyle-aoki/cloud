package defaults

import (
	"cloud/pkg/amazon"
	"cloud/pkg/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// func (cldo *CloudLabDefaultsOperator) createRouteTable() {
// 	if cldo.foundRouteTable() {
// 		return
// 	}
// 	crto, err := amazon.EC2().CreateRouteTable(&ec2.CreateRouteTableInput{
// 		VpcId:             cldo.Vpc.VpcId,
// 		TagSpecifications: createNameTag("route-table", CloudLabRouteTable),
// 	})
// 	util.MustExec(err)
// 	util.VMessage("created", CloudLabRouteTable, *crto.RouteTable.RouteTableId)
// }

// func (cldo *CloudLabDefaultsOperator) deleteRouteTable() {
// 	if !cldo.foundRouteTable() {
// 		return
// 	}
// 	_, err := amazon.EC2().DeleteRouteTable(&ec2.DeleteRouteTableInput{
// 		RouteTableId: cldo.RouteTable.RouteTableId,
// 	})
// 	util.MustExec(err)
// 	util.VMessage("deleted", CloudLabRouteTable, *cldo.RouteTable.RouteTableId)
// }

// func (cldo *CloudLabDefaultsOperator) foundRouteTable() bool {
// 	return cldo.RouteTable != nil
// }

func (cldo *CloudLabDefaultsOperator) findCloudLabRouteTable() {
	if cldo.Vpc == nil {
		return
	}
	err := amazon.EC2().DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(drto *ec2.DescribeRouteTablesOutput, b bool) bool {
			for _, rt := range drto.RouteTables {
				if rt.VpcId != nil && *rt.VpcId == *cldo.Vpc.VpcId {
					cldo.RouteTable = rt
					return false
				}
			}
			return true
		},
	)
	util.MustExec(err)
	cldo.nameRouteTable()
}

func (cldo *CloudLabDefaultsOperator) nameRouteTable() {
	if cldo.RouteTable == nil {
		return
	}
	routeTableNameTagValue := findNameTagValue(cldo.RouteTable.Tags)
	if routeTableNameTagValue != nil && *routeTableNameTagValue == CloudLabRouteTable {
		return
	}
	_, err := amazon.EC2().CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{cldo.RouteTable.RouteTableId},
		Tags: []*ec2.Tag{{
			Key:   util.StrPtr("Name"),
			Value: util.StrPtr(CloudLabRouteTable),
		}},
	})
	util.MustExec(err)
}

func (cldo *CloudLabDefaultsOperator) addInternetGatewayRoute() {
	for _, route := range cldo.RouteTable.Routes {
		if *route.GatewayId == *cldo.InternetGateway.InternetGatewayId {
			return
		}
	}
	_, err := amazon.EC2().CreateRoute(&ec2.CreateRouteInput{
		RouteTableId:         cldo.RouteTable.RouteTableId,
		GatewayId:            cldo.InternetGateway.InternetGatewayId,
		DestinationCidrBlock: util.StrPtr(RouteTablePublicSubnetCidr),
	})
	util.MustExec(err)
	fmt.Println(
		fmt.Sprintf("added %s route to %s", CloudLabInternetGateway, CloudLabRouteTable),
	)
}

func (cldo *CloudLabDefaultsOperator) associatePublicSubnetWithRouteTable() {
	for _, assoc := range cldo.RouteTable.Associations {
		if assoc.SubnetId != nil && *assoc.SubnetId == *cldo.PublicSubnet.SubnetId {
			return
		}
	}
	_, err := amazon.EC2().AssociateRouteTable(&ec2.AssociateRouteTableInput{
		RouteTableId: cldo.RouteTable.RouteTableId,
		SubnetId:     cldo.PublicSubnet.SubnetId,
	})
	util.MustExec(err)
	fmt.Println(
		fmt.Sprintf("added %s association to %s", CloudLabPublicSubnetName, CloudLabRouteTable),
	)
}
