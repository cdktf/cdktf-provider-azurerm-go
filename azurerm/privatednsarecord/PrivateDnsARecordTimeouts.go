package privatednsarecord


type PrivateDnsARecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_a_record#create PrivateDnsARecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_a_record#delete PrivateDnsARecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_a_record#read PrivateDnsARecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_a_record#update PrivateDnsARecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
