package appservice


type AppServiceLogsApplicationLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#azure_blob_storage AppService#azure_blob_storage}
	AzureBlobStorage *AppServiceLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#file_system_level AppService#file_system_level}.
	FileSystemLevel *string `field:"optional" json:"fileSystemLevel" yaml:"fileSystemLevel"`
}
