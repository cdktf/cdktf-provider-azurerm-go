package maintenanceconfiguration


type MaintenanceConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#create MaintenanceConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#delete MaintenanceConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#read MaintenanceConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#update MaintenanceConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

