// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterApiServerAccessProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/kubernetes_cluster#authorized_ip_ranges KubernetesCluster#authorized_ip_ranges}.
	AuthorizedIpRanges *[]*string `field:"optional" json:"authorizedIpRanges" yaml:"authorizedIpRanges"`
}

