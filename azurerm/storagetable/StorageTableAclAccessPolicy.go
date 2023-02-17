package storagetable


type StorageTableAclAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_table#expiry StorageTable#expiry}.
	Expiry *string `field:"required" json:"expiry" yaml:"expiry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_table#permissions StorageTable#permissions}.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_table#start StorageTable#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

