// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhubnamespacedisasterrecoveryconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventhubNamespaceDisasterRecoveryConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#name EventhubNamespaceDisasterRecoveryConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#namespace_name EventhubNamespaceDisasterRecoveryConfig#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#partner_namespace_id EventhubNamespaceDisasterRecoveryConfig#partner_namespace_id}.
	PartnerNamespaceId *string `field:"required" json:"partnerNamespaceId" yaml:"partnerNamespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#resource_group_name EventhubNamespaceDisasterRecoveryConfig#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#id EventhubNamespaceDisasterRecoveryConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/eventhub_namespace_disaster_recovery_config#timeouts EventhubNamespaceDisasterRecoveryConfig#timeouts}
	Timeouts *EventhubNamespaceDisasterRecoveryConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

