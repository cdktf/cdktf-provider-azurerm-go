// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesprovisionedcluster


type ArcKubernetesProvisionedClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/arc_kubernetes_provisioned_cluster#create ArcKubernetesProvisionedCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/arc_kubernetes_provisioned_cluster#delete ArcKubernetesProvisionedCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/arc_kubernetes_provisioned_cluster#read ArcKubernetesProvisionedCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/arc_kubernetes_provisioned_cluster#update ArcKubernetesProvisionedCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

