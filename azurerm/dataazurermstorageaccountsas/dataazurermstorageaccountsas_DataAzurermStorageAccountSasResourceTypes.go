package dataazurermstorageaccountsas


type DataAzurermStorageAccountSasResourceTypes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_sas#container DataAzurermStorageAccountSas#container}.
	Container interface{} `field:"required" json:"container" yaml:"container"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_sas#object DataAzurermStorageAccountSas#object}.
	Object interface{} `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_sas#service DataAzurermStorageAccountSas#service}.
	Service interface{} `field:"required" json:"service" yaml:"service"`
}

