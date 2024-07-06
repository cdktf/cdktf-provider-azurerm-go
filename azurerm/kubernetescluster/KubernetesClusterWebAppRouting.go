// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterWebAppRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/kubernetes_cluster#dns_zone_id KubernetesCluster#dns_zone_id}.
	DnsZoneId *string `field:"optional" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/kubernetes_cluster#dns_zone_ids KubernetesCluster#dns_zone_ids}.
	DnsZoneIds *[]*string `field:"optional" json:"dnsZoneIds" yaml:"dnsZoneIds"`
}

