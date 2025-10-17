// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfunctioncollectorpolicy


type NetworkFunctionCollectorPolicyIpfxIngestion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/network_function_collector_policy#source_resource_ids NetworkFunctionCollectorPolicy#source_resource_ids}.
	SourceResourceIds *[]*string `field:"required" json:"sourceResourceIds" yaml:"sourceResourceIds"`
}

