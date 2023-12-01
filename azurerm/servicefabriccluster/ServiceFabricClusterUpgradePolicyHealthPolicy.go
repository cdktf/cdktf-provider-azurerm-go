// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterUpgradePolicyHealthPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/service_fabric_cluster#max_unhealthy_applications_percent ServiceFabricCluster#max_unhealthy_applications_percent}.
	MaxUnhealthyApplicationsPercent *float64 `field:"optional" json:"maxUnhealthyApplicationsPercent" yaml:"maxUnhealthyApplicationsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/service_fabric_cluster#max_unhealthy_nodes_percent ServiceFabricCluster#max_unhealthy_nodes_percent}.
	MaxUnhealthyNodesPercent *float64 `field:"optional" json:"maxUnhealthyNodesPercent" yaml:"maxUnhealthyNodesPercent"`
}

