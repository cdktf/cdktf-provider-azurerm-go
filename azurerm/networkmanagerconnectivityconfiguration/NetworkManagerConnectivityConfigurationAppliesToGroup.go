package networkmanagerconnectivityconfiguration


type NetworkManagerConnectivityConfigurationAppliesToGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/network_manager_connectivity_configuration#group_connectivity NetworkManagerConnectivityConfiguration#group_connectivity}.
	GroupConnectivity *string `field:"required" json:"groupConnectivity" yaml:"groupConnectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/network_manager_connectivity_configuration#network_group_id NetworkManagerConnectivityConfiguration#network_group_id}.
	NetworkGroupId *string `field:"required" json:"networkGroupId" yaml:"networkGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/network_manager_connectivity_configuration#global_mesh_enabled NetworkManagerConnectivityConfiguration#global_mesh_enabled}.
	GlobalMeshEnabled interface{} `field:"optional" json:"globalMeshEnabled" yaml:"globalMeshEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/network_manager_connectivity_configuration#use_hub_gateway NetworkManagerConnectivityConfiguration#use_hub_gateway}.
	UseHubGateway interface{} `field:"optional" json:"useHubGateway" yaml:"useHubGateway"`
}

