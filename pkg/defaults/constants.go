package defaults

const DefaultVpcName = "cloudlab-vpc"
const CloudLabRouteTable = "cloudlab-route-table"
const CloudLabPublicSubnetName = "cloudlab-public-subnet"
const CloudLabPrivateSubnetName = "cloudlab-private-subnet"
const CloudLabInternetGateway = "cloudlab-internet-gateway"
const CloudLabSecutiyGroup = "cloudlab-security-group"
const CloudLabKeyPair = "cloud-lab-key-pair"
const CloudLabKeyPairNameTemplate = "cloudlab-key-pair-"
const CloudLabInstance = "cloud-lab-instance"

const MissingVpcErrorMessage = `
You don't have a VPC named 'cloudlab'.
Run 'cloudlab init' to create it.
`

const DefaultVpcCidrBlock = "10.0.0.0/16"
const AllIpsCidr = "0.0.0.0/0"
const RouteTablePublicSubnetCidr = "0.0.0.0/0"

const PublicSubnetCidrBlock = "10.0.0.0/24"
const PublicSubnetNameTagValue = "public"

const PrivateSubnetCidrBlock = "10.0.1.0/24"
const PrivateSubnetNameTagValue = "private"

type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)
