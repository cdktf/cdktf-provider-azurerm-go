// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetmanager


type KubernetesFleetManagerHubProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/kubernetes_fleet_manager#dns_prefix KubernetesFleetManager#dns_prefix}.
	DnsPrefix *string `field:"required" json:"dnsPrefix" yaml:"dnsPrefix"`
}

