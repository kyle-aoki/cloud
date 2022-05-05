package initialize

const DefaultVpcName = "cloudlab-vpc"
const CloudLabRouteTable = "cloudlab-route-table"

const MissingVpcErrorMessage = `
You don't have a VPC named 'cloudlab'.
Run 'cloudlab init' to create it.
`

const DefaultVpcCidrBlock = "10.0.0.0/16"

const PublicSubnetCidrBlock = "10.0.0.0/24"
const PublicSubnetNameTagValue = "public"

const PrivateSubnetCidrBlock = "10.0.1.0/24"
const PrivateSubnetNameTagValue = "private"
