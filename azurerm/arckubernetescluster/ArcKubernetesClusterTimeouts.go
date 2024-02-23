// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetescluster


type ArcKubernetesClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/arc_kubernetes_cluster#create ArcKubernetesCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/arc_kubernetes_cluster#delete ArcKubernetesCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/arc_kubernetes_cluster#read ArcKubernetesCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/arc_kubernetes_cluster#update ArcKubernetesCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

