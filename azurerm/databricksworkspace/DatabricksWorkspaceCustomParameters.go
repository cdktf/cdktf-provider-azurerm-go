// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databricksworkspace


type DatabricksWorkspaceCustomParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#machine_learning_workspace_id DatabricksWorkspace#machine_learning_workspace_id}.
	MachineLearningWorkspaceId *string `field:"optional" json:"machineLearningWorkspaceId" yaml:"machineLearningWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#nat_gateway_name DatabricksWorkspace#nat_gateway_name}.
	NatGatewayName *string `field:"optional" json:"natGatewayName" yaml:"natGatewayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#no_public_ip DatabricksWorkspace#no_public_ip}.
	NoPublicIp interface{} `field:"optional" json:"noPublicIp" yaml:"noPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#private_subnet_name DatabricksWorkspace#private_subnet_name}.
	PrivateSubnetName *string `field:"optional" json:"privateSubnetName" yaml:"privateSubnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#private_subnet_network_security_group_association_id DatabricksWorkspace#private_subnet_network_security_group_association_id}.
	PrivateSubnetNetworkSecurityGroupAssociationId *string `field:"optional" json:"privateSubnetNetworkSecurityGroupAssociationId" yaml:"privateSubnetNetworkSecurityGroupAssociationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#public_ip_name DatabricksWorkspace#public_ip_name}.
	PublicIpName *string `field:"optional" json:"publicIpName" yaml:"publicIpName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#public_subnet_name DatabricksWorkspace#public_subnet_name}.
	PublicSubnetName *string `field:"optional" json:"publicSubnetName" yaml:"publicSubnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#public_subnet_network_security_group_association_id DatabricksWorkspace#public_subnet_network_security_group_association_id}.
	PublicSubnetNetworkSecurityGroupAssociationId *string `field:"optional" json:"publicSubnetNetworkSecurityGroupAssociationId" yaml:"publicSubnetNetworkSecurityGroupAssociationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#storage_account_name DatabricksWorkspace#storage_account_name}.
	StorageAccountName *string `field:"optional" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#storage_account_sku_name DatabricksWorkspace#storage_account_sku_name}.
	StorageAccountSkuName *string `field:"optional" json:"storageAccountSkuName" yaml:"storageAccountSkuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#virtual_network_id DatabricksWorkspace#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/databricks_workspace#vnet_address_prefix DatabricksWorkspace#vnet_address_prefix}.
	VnetAddressPrefix *string `field:"optional" json:"vnetAddressPrefix" yaml:"vnetAddressPrefix"`
}

