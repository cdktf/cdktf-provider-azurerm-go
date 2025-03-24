// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspacenetworkoutboundruleprivateendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#name MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#service_resource_id MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#service_resource_id}.
	ServiceResourceId *string `field:"required" json:"serviceResourceId" yaml:"serviceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#sub_resource_target MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#sub_resource_target}.
	SubResourceTarget *string `field:"required" json:"subResourceTarget" yaml:"subResourceTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#workspace_id MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#id MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#spark_enabled MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#spark_enabled}.
	SparkEnabled interface{} `field:"optional" json:"sparkEnabled" yaml:"sparkEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/machine_learning_workspace_network_outbound_rule_private_endpoint#timeouts MachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint#timeouts}
	Timeouts *MachineLearningWorkspaceNetworkOutboundRulePrivateEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

