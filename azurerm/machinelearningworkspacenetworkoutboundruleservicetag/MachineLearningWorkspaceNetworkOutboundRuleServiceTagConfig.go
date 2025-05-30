// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspacenetworkoutboundruleservicetag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningWorkspaceNetworkOutboundRuleServiceTagConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#name MachineLearningWorkspaceNetworkOutboundRuleServiceTag#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#port_ranges MachineLearningWorkspaceNetworkOutboundRuleServiceTag#port_ranges}.
	PortRanges *string `field:"required" json:"portRanges" yaml:"portRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#protocol MachineLearningWorkspaceNetworkOutboundRuleServiceTag#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#service_tag MachineLearningWorkspaceNetworkOutboundRuleServiceTag#service_tag}.
	ServiceTag *string `field:"required" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#workspace_id MachineLearningWorkspaceNetworkOutboundRuleServiceTag#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#id MachineLearningWorkspaceNetworkOutboundRuleServiceTag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/machine_learning_workspace_network_outbound_rule_service_tag#timeouts MachineLearningWorkspaceNetworkOutboundRuleServiceTag#timeouts}
	Timeouts *MachineLearningWorkspaceNetworkOutboundRuleServiceTagTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

