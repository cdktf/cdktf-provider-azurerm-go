// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdaterun


type KubernetesFleetUpdateRunManagedClusterUpdate struct {
	// upgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/kubernetes_fleet_update_run#upgrade KubernetesFleetUpdateRun#upgrade}
	Upgrade *KubernetesFleetUpdateRunManagedClusterUpdateUpgrade `field:"required" json:"upgrade" yaml:"upgrade"`
	// node_image_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/kubernetes_fleet_update_run#node_image_selection KubernetesFleetUpdateRun#node_image_selection}
	NodeImageSelection *KubernetesFleetUpdateRunManagedClusterUpdateNodeImageSelection `field:"optional" json:"nodeImageSelection" yaml:"nodeImageSelection"`
}

