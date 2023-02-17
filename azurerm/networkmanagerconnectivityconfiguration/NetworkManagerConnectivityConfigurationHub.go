package networkmanagerconnectivityconfiguration


type NetworkManagerConnectivityConfigurationHub struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_connectivity_configuration#resource_id NetworkManagerConnectivityConfiguration#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_connectivity_configuration#resource_type NetworkManagerConnectivityConfiguration#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

