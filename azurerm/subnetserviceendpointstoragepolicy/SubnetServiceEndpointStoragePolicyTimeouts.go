package subnetserviceendpointstoragepolicy


type SubnetServiceEndpointStoragePolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet_service_endpoint_storage_policy#create SubnetServiceEndpointStoragePolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet_service_endpoint_storage_policy#delete SubnetServiceEndpointStoragePolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet_service_endpoint_storage_policy#read SubnetServiceEndpointStoragePolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet_service_endpoint_storage_policy#update SubnetServiceEndpointStoragePolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
