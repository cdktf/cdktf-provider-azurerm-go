package maintenanceconfiguration


type MaintenanceConfigurationWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#start_date_time MaintenanceConfiguration#start_date_time}.
	StartDateTime *string `field:"required" json:"startDateTime" yaml:"startDateTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#time_zone MaintenanceConfiguration#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#duration MaintenanceConfiguration#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#expiration_date_time MaintenanceConfiguration#expiration_date_time}.
	ExpirationDateTime *string `field:"optional" json:"expirationDateTime" yaml:"expirationDateTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#recur_every MaintenanceConfiguration#recur_every}.
	RecurEvery *string `field:"optional" json:"recurEvery" yaml:"recurEvery"`
}
