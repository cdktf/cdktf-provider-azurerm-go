package dataprotectionresourceguard


type DataProtectionResourceGuardTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_resource_guard#create DataProtectionResourceGuard#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_resource_guard#delete DataProtectionResourceGuard#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_resource_guard#read DataProtectionResourceGuard#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_resource_guard#update DataProtectionResourceGuard#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
