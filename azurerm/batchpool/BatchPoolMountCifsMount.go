package batchpool


type BatchPoolMountCifsMount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#password BatchPool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#source BatchPool#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

