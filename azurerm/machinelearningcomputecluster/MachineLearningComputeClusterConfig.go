// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningcomputecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningComputeClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#location MachineLearningComputeCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#machine_learning_workspace_id MachineLearningComputeCluster#machine_learning_workspace_id}.
	MachineLearningWorkspaceId *string `field:"required" json:"machineLearningWorkspaceId" yaml:"machineLearningWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#name MachineLearningComputeCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// scale_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#scale_settings MachineLearningComputeCluster#scale_settings}
	ScaleSettings *MachineLearningComputeClusterScaleSettings `field:"required" json:"scaleSettings" yaml:"scaleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#vm_priority MachineLearningComputeCluster#vm_priority}.
	VmPriority *string `field:"required" json:"vmPriority" yaml:"vmPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#vm_size MachineLearningComputeCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#description MachineLearningComputeCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#id MachineLearningComputeCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#identity MachineLearningComputeCluster#identity}
	Identity *MachineLearningComputeClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#local_auth_enabled MachineLearningComputeCluster#local_auth_enabled}.
	LocalAuthEnabled interface{} `field:"optional" json:"localAuthEnabled" yaml:"localAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#node_public_ip_enabled MachineLearningComputeCluster#node_public_ip_enabled}.
	NodePublicIpEnabled interface{} `field:"optional" json:"nodePublicIpEnabled" yaml:"nodePublicIpEnabled"`
	// ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#ssh MachineLearningComputeCluster#ssh}
	Ssh *MachineLearningComputeClusterSsh `field:"optional" json:"ssh" yaml:"ssh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#ssh_public_access_enabled MachineLearningComputeCluster#ssh_public_access_enabled}.
	SshPublicAccessEnabled interface{} `field:"optional" json:"sshPublicAccessEnabled" yaml:"sshPublicAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#subnet_resource_id MachineLearningComputeCluster#subnet_resource_id}.
	SubnetResourceId *string `field:"optional" json:"subnetResourceId" yaml:"subnetResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#tags MachineLearningComputeCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/machine_learning_compute_cluster#timeouts MachineLearningComputeCluster#timeouts}
	Timeouts *MachineLearningComputeClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

