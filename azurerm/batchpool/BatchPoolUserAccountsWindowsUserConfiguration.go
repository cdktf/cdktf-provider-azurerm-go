package batchpool


type BatchPoolUserAccountsWindowsUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/batch_pool#login_mode BatchPool#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}

