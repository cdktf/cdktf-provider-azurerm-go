package manageddisksastoken


type ManagedDiskSasTokenTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk_sas_token#create ManagedDiskSasToken#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk_sas_token#delete ManagedDiskSasToken#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/managed_disk_sas_token#read ManagedDiskSasToken#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

