package natgatewaypublicipassociation


type NatGatewayPublicIpAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nat_gateway_public_ip_association#create NatGatewayPublicIpAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nat_gateway_public_ip_association#delete NatGatewayPublicIpAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nat_gateway_public_ip_association#read NatGatewayPublicIpAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
