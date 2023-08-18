package storagemanagementpolicy


type StorageManagementPolicyRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/storage_management_policy#actions StorageManagementPolicy#actions}
	Actions *StorageManagementPolicyRuleActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/storage_management_policy#enabled StorageManagementPolicy#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/storage_management_policy#filters StorageManagementPolicy#filters}
	Filters *StorageManagementPolicyRuleFilters `field:"required" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/storage_management_policy#name StorageManagementPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

