package databasemigrationservice


type DatabaseMigrationServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_service#create DatabaseMigrationService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_service#delete DatabaseMigrationService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_service#read DatabaseMigrationService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/database_migration_service#update DatabaseMigrationService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

