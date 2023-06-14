package batchpool


type BatchPoolMountNfsMount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/batch_pool#relative_mount_path BatchPool#relative_mount_path}.
	RelativeMountPath *string `field:"required" json:"relativeMountPath" yaml:"relativeMountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/batch_pool#source BatchPool#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/batch_pool#mount_options BatchPool#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

