package privatednssrvrecord


type PrivateDnsSrvRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/private_dns_srv_record#create PrivateDnsSrvRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/private_dns_srv_record#delete PrivateDnsSrvRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/private_dns_srv_record#read PrivateDnsSrvRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/private_dns_srv_record#update PrivateDnsSrvRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

