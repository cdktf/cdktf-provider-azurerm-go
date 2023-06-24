package datafactorylinkedservicesftp


type DataFactoryLinkedServiceSftpTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_sftp#create DataFactoryLinkedServiceSftp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_sftp#delete DataFactoryLinkedServiceSftp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_sftp#read DataFactoryLinkedServiceSftp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_sftp#update DataFactoryLinkedServiceSftp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

