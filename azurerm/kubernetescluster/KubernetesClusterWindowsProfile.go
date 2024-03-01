// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterWindowsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#admin_username KubernetesCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#admin_password KubernetesCluster#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// gmsa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#gmsa KubernetesCluster#gmsa}
	Gmsa *KubernetesClusterWindowsProfileGmsa `field:"optional" json:"gmsa" yaml:"gmsa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#license KubernetesCluster#license}.
	License *string `field:"optional" json:"license" yaml:"license"`
}

