package batchpool


type BatchPoolUserAccounts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#elevation_level BatchPool#elevation_level}.
	ElevationLevel *string `field:"required" json:"elevationLevel" yaml:"elevationLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#password BatchPool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// linux_user_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#linux_user_configuration BatchPool#linux_user_configuration}
	LinuxUserConfiguration interface{} `field:"optional" json:"linuxUserConfiguration" yaml:"linuxUserConfiguration"`
	// windows_user_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#windows_user_configuration BatchPool#windows_user_configuration}
	WindowsUserConfiguration interface{} `field:"optional" json:"windowsUserConfiguration" yaml:"windowsUserConfiguration"`
}
