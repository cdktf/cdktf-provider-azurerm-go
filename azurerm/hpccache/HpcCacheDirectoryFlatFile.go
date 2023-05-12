package hpccache


type HpcCacheDirectoryFlatFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hpc_cache#group_file_uri HpcCache#group_file_uri}.
	GroupFileUri *string `field:"required" json:"groupFileUri" yaml:"groupFileUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hpc_cache#password_file_uri HpcCache#password_file_uri}.
	PasswordFileUri *string `field:"required" json:"passwordFileUri" yaml:"passwordFileUri"`
}

