package datafactorylinkedserviceodata


type DataFactoryLinkedServiceOdataBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_odata#password DataFactoryLinkedServiceOdata#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_odata#username DataFactoryLinkedServiceOdata#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

