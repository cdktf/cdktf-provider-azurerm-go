package batchpool


type BatchPoolUserAccountsWindowsUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#login_mode BatchPool#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}

