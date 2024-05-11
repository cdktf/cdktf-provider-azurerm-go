// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterNodeTypeApplicationPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/service_fabric_cluster#end_port ServiceFabricCluster#end_port}.
	EndPort *float64 `field:"required" json:"endPort" yaml:"endPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/service_fabric_cluster#start_port ServiceFabricCluster#start_port}.
	StartPort *float64 `field:"required" json:"startPort" yaml:"startPort"`
}

