package storageblobinventorypolicy


type StorageBlobInventoryPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_blob_inventory_policy#create StorageBlobInventoryPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_blob_inventory_policy#delete StorageBlobInventoryPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_blob_inventory_policy#read StorageBlobInventoryPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_blob_inventory_policy#update StorageBlobInventoryPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
