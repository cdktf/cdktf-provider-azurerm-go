// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanagementprivatelinkassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceManagementPrivateLinkAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#management_group_id ResourceManagementPrivateLinkAssociation#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#public_network_access_enabled ResourceManagementPrivateLinkAssociation#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"required" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#resource_management_private_link_id ResourceManagementPrivateLinkAssociation#resource_management_private_link_id}.
	ResourceManagementPrivateLinkId *string `field:"required" json:"resourceManagementPrivateLinkId" yaml:"resourceManagementPrivateLinkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#id ResourceManagementPrivateLinkAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#name ResourceManagementPrivateLinkAssociation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/resource_management_private_link_association#timeouts ResourceManagementPrivateLinkAssociation#timeouts}
	Timeouts *ResourceManagementPrivateLinkAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

