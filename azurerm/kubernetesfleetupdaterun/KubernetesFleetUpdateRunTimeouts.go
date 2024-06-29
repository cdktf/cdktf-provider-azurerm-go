// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdaterun


type KubernetesFleetUpdateRunTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/kubernetes_fleet_update_run#create KubernetesFleetUpdateRun#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/kubernetes_fleet_update_run#delete KubernetesFleetUpdateRun#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/kubernetes_fleet_update_run#read KubernetesFleetUpdateRun#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/kubernetes_fleet_update_run#update KubernetesFleetUpdateRun#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

