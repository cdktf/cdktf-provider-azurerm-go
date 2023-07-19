package storagemover


type StorageMoverTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/storage_mover#create StorageMover#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/storage_mover#delete StorageMover#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/storage_mover#read StorageMover#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/storage_mover#update StorageMover#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

