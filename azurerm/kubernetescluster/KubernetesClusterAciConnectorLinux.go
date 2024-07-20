// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterAciConnectorLinux struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/kubernetes_cluster#subnet_name KubernetesCluster#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
}

