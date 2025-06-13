// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arcresourcebridgeappliance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcResourceBridgeApplianceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#distro ArcResourceBridgeAppliance#distro}.
	Distro *string `field:"required" json:"distro" yaml:"distro"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#identity ArcResourceBridgeAppliance#identity}
	Identity *ArcResourceBridgeApplianceIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#infrastructure_provider ArcResourceBridgeAppliance#infrastructure_provider}.
	InfrastructureProvider *string `field:"required" json:"infrastructureProvider" yaml:"infrastructureProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#location ArcResourceBridgeAppliance#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#name ArcResourceBridgeAppliance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#resource_group_name ArcResourceBridgeAppliance#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#id ArcResourceBridgeAppliance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#public_key_base64 ArcResourceBridgeAppliance#public_key_base64}.
	PublicKeyBase64 *string `field:"optional" json:"publicKeyBase64" yaml:"publicKeyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#tags ArcResourceBridgeAppliance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/arc_resource_bridge_appliance#timeouts ArcResourceBridgeAppliance#timeouts}
	Timeouts *ArcResourceBridgeApplianceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

