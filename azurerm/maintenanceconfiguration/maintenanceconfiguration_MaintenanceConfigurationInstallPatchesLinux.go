package maintenanceconfiguration


type MaintenanceConfigurationInstallPatchesLinux struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#classifications_to_include MaintenanceConfiguration#classifications_to_include}.
	ClassificationsToInclude *[]*string `field:"optional" json:"classificationsToInclude" yaml:"classificationsToInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#package_names_mask_to_exclude MaintenanceConfiguration#package_names_mask_to_exclude}.
	PackageNamesMaskToExclude *[]*string `field:"optional" json:"packageNamesMaskToExclude" yaml:"packageNamesMaskToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/maintenance_configuration#package_names_mask_to_include MaintenanceConfiguration#package_names_mask_to_include}.
	PackageNamesMaskToInclude *[]*string `field:"optional" json:"packageNamesMaskToInclude" yaml:"packageNamesMaskToInclude"`
}

