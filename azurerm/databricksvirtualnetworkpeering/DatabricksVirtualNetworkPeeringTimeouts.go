package databricksvirtualnetworkpeering


type DatabricksVirtualNetworkPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/databricks_virtual_network_peering#create DatabricksVirtualNetworkPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/databricks_virtual_network_peering#delete DatabricksVirtualNetworkPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/databricks_virtual_network_peering#read DatabricksVirtualNetworkPeering#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/databricks_virtual_network_peering#update DatabricksVirtualNetworkPeering#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

