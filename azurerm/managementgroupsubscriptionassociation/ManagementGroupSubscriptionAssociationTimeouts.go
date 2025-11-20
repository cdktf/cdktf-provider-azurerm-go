// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgroupsubscriptionassociation


type ManagementGroupSubscriptionAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_subscription_association#create ManagementGroupSubscriptionAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_subscription_association#delete ManagementGroupSubscriptionAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/management_group_subscription_association#read ManagementGroupSubscriptionAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

