// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearninginferencecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningInferenceClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#kubernetes_cluster_id MachineLearningInferenceCluster#kubernetes_cluster_id}.
	KubernetesClusterId *string `field:"required" json:"kubernetesClusterId" yaml:"kubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#location MachineLearningInferenceCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#machine_learning_workspace_id MachineLearningInferenceCluster#machine_learning_workspace_id}.
	MachineLearningWorkspaceId *string `field:"required" json:"machineLearningWorkspaceId" yaml:"machineLearningWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#name MachineLearningInferenceCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#cluster_purpose MachineLearningInferenceCluster#cluster_purpose}.
	ClusterPurpose *string `field:"optional" json:"clusterPurpose" yaml:"clusterPurpose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#description MachineLearningInferenceCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#id MachineLearningInferenceCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#identity MachineLearningInferenceCluster#identity}
	Identity *MachineLearningInferenceClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#ssl MachineLearningInferenceCluster#ssl}
	Ssl *MachineLearningInferenceClusterSsl `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#tags MachineLearningInferenceCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/machine_learning_inference_cluster#timeouts MachineLearningInferenceCluster#timeouts}
	Timeouts *MachineLearningInferenceClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

