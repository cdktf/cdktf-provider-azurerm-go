package batchpool


type BatchPoolNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#subnet_id BatchPool#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#dynamic_vnet_assignment_scope BatchPool#dynamic_vnet_assignment_scope}.
	DynamicVnetAssignmentScope *string `field:"optional" json:"dynamicVnetAssignmentScope" yaml:"dynamicVnetAssignmentScope"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#endpoint_configuration BatchPool#endpoint_configuration}
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#public_address_provisioning_type BatchPool#public_address_provisioning_type}.
	PublicAddressProvisioningType *string `field:"optional" json:"publicAddressProvisioningType" yaml:"publicAddressProvisioningType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/batch_pool#public_ips BatchPool#public_ips}.
	PublicIps *[]*string `field:"optional" json:"publicIps" yaml:"publicIps"`
}

