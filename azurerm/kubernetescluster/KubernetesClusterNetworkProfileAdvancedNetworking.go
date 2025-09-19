// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterNetworkProfileAdvancedNetworking struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/kubernetes_cluster#observability_enabled KubernetesCluster#observability_enabled}.
	ObservabilityEnabled interface{} `field:"optional" json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/kubernetes_cluster#security_enabled KubernetesCluster#security_enabled}.
	SecurityEnabled interface{} `field:"optional" json:"securityEnabled" yaml:"securityEnabled"`
}

