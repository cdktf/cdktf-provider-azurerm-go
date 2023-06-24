package provider


type AzurermProviderFeaturesVirtualMachine struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs#delete_os_disk_on_deletion AzurermProvider#delete_os_disk_on_deletion}.
	DeleteOsDiskOnDeletion interface{} `field:"optional" json:"deleteOsDiskOnDeletion" yaml:"deleteOsDiskOnDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs#graceful_shutdown AzurermProvider#graceful_shutdown}.
	GracefulShutdown interface{} `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs#skip_shutdown_and_force_delete AzurermProvider#skip_shutdown_and_force_delete}.
	SkipShutdownAndForceDelete interface{} `field:"optional" json:"skipShutdownAndForceDelete" yaml:"skipShutdownAndForceDelete"`
}

