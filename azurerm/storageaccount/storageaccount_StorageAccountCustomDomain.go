package storageaccount


type StorageAccountCustomDomain struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#use_subdomain StorageAccount#use_subdomain}.
	UseSubdomain interface{} `field:"optional" json:"useSubdomain" yaml:"useSubdomain"`
}

