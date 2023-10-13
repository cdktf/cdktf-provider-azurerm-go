// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclustertrustedaccessrolebinding


type KubernetesClusterTrustedAccessRoleBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/kubernetes_cluster_trusted_access_role_binding#create KubernetesClusterTrustedAccessRoleBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/kubernetes_cluster_trusted_access_role_binding#delete KubernetesClusterTrustedAccessRoleBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/kubernetes_cluster_trusted_access_role_binding#read KubernetesClusterTrustedAccessRoleBinding#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/kubernetes_cluster_trusted_access_role_binding#update KubernetesClusterTrustedAccessRoleBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

