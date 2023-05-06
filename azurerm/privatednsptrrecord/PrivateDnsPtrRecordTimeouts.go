package privatednsptrrecord


type PrivateDnsPtrRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/private_dns_ptr_record#create PrivateDnsPtrRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/private_dns_ptr_record#delete PrivateDnsPtrRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/private_dns_ptr_record#read PrivateDnsPtrRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/private_dns_ptr_record#update PrivateDnsPtrRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

