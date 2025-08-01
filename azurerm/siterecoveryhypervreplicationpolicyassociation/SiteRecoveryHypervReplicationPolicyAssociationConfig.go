// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryhypervreplicationpolicyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryHypervReplicationPolicyAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_hyperv_replication_policy_association#hyperv_site_id SiteRecoveryHypervReplicationPolicyAssociation#hyperv_site_id}.
	HypervSiteId *string `field:"required" json:"hypervSiteId" yaml:"hypervSiteId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_hyperv_replication_policy_association#name SiteRecoveryHypervReplicationPolicyAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_hyperv_replication_policy_association#policy_id SiteRecoveryHypervReplicationPolicyAssociation#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_hyperv_replication_policy_association#id SiteRecoveryHypervReplicationPolicyAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_hyperv_replication_policy_association#timeouts SiteRecoveryHypervReplicationPolicyAssociation#timeouts}
	Timeouts *SiteRecoveryHypervReplicationPolicyAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

