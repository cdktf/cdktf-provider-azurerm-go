package storageshare


type StorageShareAclAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/storage_share#permissions StorageShare#permissions}.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/storage_share#expiry StorageShare#expiry}.
	Expiry *string `field:"optional" json:"expiry" yaml:"expiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/storage_share#start StorageShare#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

