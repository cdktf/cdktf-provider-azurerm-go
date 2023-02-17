package postgresqlconfiguration


type PostgresqlConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_configuration#create PostgresqlConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_configuration#delete PostgresqlConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_configuration#read PostgresqlConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_configuration#update PostgresqlConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

