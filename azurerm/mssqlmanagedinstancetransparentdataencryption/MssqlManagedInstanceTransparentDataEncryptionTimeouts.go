package mssqlmanagedinstancetransparentdataencryption


type MssqlManagedInstanceTransparentDataEncryptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_transparent_data_encryption#create MssqlManagedInstanceTransparentDataEncryption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_transparent_data_encryption#delete MssqlManagedInstanceTransparentDataEncryption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_transparent_data_encryption#read MssqlManagedInstanceTransparentDataEncryption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_transparent_data_encryption#update MssqlManagedInstanceTransparentDataEncryption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

