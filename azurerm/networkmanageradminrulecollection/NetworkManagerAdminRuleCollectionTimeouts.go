// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanageradminrulecollection


type NetworkManagerAdminRuleCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/network_manager_admin_rule_collection#create NetworkManagerAdminRuleCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/network_manager_admin_rule_collection#delete NetworkManagerAdminRuleCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/network_manager_admin_rule_collection#read NetworkManagerAdminRuleCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/network_manager_admin_rule_collection#update NetworkManagerAdminRuleCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

