package storagemanagementpolicy


type StorageManagementPolicyRuleFiltersMatchBlobIndexTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/storage_management_policy#name StorageManagementPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/storage_management_policy#value StorageManagementPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/storage_management_policy#operation StorageManagementPolicy#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

