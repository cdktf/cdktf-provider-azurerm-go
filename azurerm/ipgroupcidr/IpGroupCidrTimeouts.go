package ipgroupcidr


type IpGroupCidrTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/ip_group_cidr#create IpGroupCidr#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/ip_group_cidr#delete IpGroupCidr#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/ip_group_cidr#read IpGroupCidr#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
