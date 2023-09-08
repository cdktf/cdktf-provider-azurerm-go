// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningsynapsespark

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningSynapseSparkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#location MachineLearningSynapseSpark#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#machine_learning_workspace_id MachineLearningSynapseSpark#machine_learning_workspace_id}.
	MachineLearningWorkspaceId *string `field:"required" json:"machineLearningWorkspaceId" yaml:"machineLearningWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#name MachineLearningSynapseSpark#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#synapse_spark_pool_id MachineLearningSynapseSpark#synapse_spark_pool_id}.
	SynapseSparkPoolId *string `field:"required" json:"synapseSparkPoolId" yaml:"synapseSparkPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#description MachineLearningSynapseSpark#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#id MachineLearningSynapseSpark#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#identity MachineLearningSynapseSpark#identity}
	Identity *MachineLearningSynapseSparkIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#local_auth_enabled MachineLearningSynapseSpark#local_auth_enabled}.
	LocalAuthEnabled interface{} `field:"optional" json:"localAuthEnabled" yaml:"localAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#tags MachineLearningSynapseSpark#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/machine_learning_synapse_spark#timeouts MachineLearningSynapseSpark#timeouts}
	Timeouts *MachineLearningSynapseSparkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

