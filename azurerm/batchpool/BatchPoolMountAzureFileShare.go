package batchpool


type BatchPoolMountAzureFileShare struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/batch_pool#account_key BatchPool#account_key}.
	AccountKey *string `field:"required" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/batch_pool#azure_file_url BatchPool#azure_file_url}.
	AzureFileUrl *string `field:"required" json:"azureFileUrl" yaml:"azureFileUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

