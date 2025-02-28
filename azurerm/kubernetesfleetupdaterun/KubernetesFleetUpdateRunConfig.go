// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdaterun

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFleetUpdateRunConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#kubernetes_fleet_manager_id KubernetesFleetUpdateRun#kubernetes_fleet_manager_id}.
	KubernetesFleetManagerId *string `field:"required" json:"kubernetesFleetManagerId" yaml:"kubernetesFleetManagerId"`
	// managed_cluster_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#managed_cluster_update KubernetesFleetUpdateRun#managed_cluster_update}
	ManagedClusterUpdate *KubernetesFleetUpdateRunManagedClusterUpdate `field:"required" json:"managedClusterUpdate" yaml:"managedClusterUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#name KubernetesFleetUpdateRun#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#fleet_update_strategy_id KubernetesFleetUpdateRun#fleet_update_strategy_id}.
	FleetUpdateStrategyId *string `field:"optional" json:"fleetUpdateStrategyId" yaml:"fleetUpdateStrategyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#id KubernetesFleetUpdateRun#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#stage KubernetesFleetUpdateRun#stage}
	Stage interface{} `field:"optional" json:"stage" yaml:"stage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kubernetes_fleet_update_run#timeouts KubernetesFleetUpdateRun#timeouts}
	Timeouts *KubernetesFleetUpdateRunTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

