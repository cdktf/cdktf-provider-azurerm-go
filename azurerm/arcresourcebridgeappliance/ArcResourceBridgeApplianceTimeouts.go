// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arcresourcebridgeappliance


type ArcResourceBridgeApplianceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/arc_resource_bridge_appliance#create ArcResourceBridgeAppliance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/arc_resource_bridge_appliance#delete ArcResourceBridgeAppliance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/arc_resource_bridge_appliance#read ArcResourceBridgeAppliance#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/arc_resource_bridge_appliance#update ArcResourceBridgeAppliance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

