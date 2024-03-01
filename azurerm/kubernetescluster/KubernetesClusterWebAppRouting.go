// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterWebAppRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#dns_zone_id KubernetesCluster#dns_zone_id}.
	DnsZoneId *string `field:"required" json:"dnsZoneId" yaml:"dnsZoneId"`
}

