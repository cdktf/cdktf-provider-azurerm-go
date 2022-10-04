package batchpool


type BatchPoolUserAccountsLinuxUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#gid BatchPool#gid}.
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#ssh_private_key BatchPool#ssh_private_key}.
	SshPrivateKey *string `field:"optional" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#uid BatchPool#uid}.
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

