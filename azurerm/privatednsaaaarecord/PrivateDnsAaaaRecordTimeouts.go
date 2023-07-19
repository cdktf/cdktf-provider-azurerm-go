package privatednsaaaarecord


type PrivateDnsAaaaRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/private_dns_aaaa_record#create PrivateDnsAaaaRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/private_dns_aaaa_record#delete PrivateDnsAaaaRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/private_dns_aaaa_record#read PrivateDnsAaaaRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/private_dns_aaaa_record#update PrivateDnsAaaaRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

