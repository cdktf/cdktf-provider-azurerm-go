package privatednscnamerecord


type PrivateDnsCnameRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_cname_record#create PrivateDnsCnameRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_cname_record#delete PrivateDnsCnameRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_cname_record#read PrivateDnsCnameRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_cname_record#update PrivateDnsCnameRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
