package storagesynccloudendpoint


type StorageSyncCloudEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_sync_cloud_endpoint#create StorageSyncCloudEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_sync_cloud_endpoint#delete StorageSyncCloudEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_sync_cloud_endpoint#read StorageSyncCloudEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
