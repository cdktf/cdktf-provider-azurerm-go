package databasemigrationproject


type DatabaseMigrationProjectTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_project#create DatabaseMigrationProject#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_project#delete DatabaseMigrationProject#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_project#read DatabaseMigrationProject#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_project#update DatabaseMigrationProject#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
