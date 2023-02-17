package image


type ImageOsDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#blob_uri Image#blob_uri}.
	BlobUri *string `field:"optional" json:"blobUri" yaml:"blobUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#caching Image#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#managed_disk_id Image#managed_disk_id}.
	ManagedDiskId *string `field:"optional" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#os_state Image#os_state}.
	OsState *string `field:"optional" json:"osState" yaml:"osState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#os_type Image#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/image#size_gb Image#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

