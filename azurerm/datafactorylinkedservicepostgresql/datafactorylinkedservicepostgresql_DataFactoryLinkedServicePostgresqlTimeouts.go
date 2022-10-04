package datafactorylinkedservicepostgresql


type DataFactoryLinkedServicePostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_postgresql#create DataFactoryLinkedServicePostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_postgresql#delete DataFactoryLinkedServicePostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_postgresql#read DataFactoryLinkedServicePostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_postgresql#update DataFactoryLinkedServicePostgresql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

