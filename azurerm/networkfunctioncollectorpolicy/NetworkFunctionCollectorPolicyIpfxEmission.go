// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfunctioncollectorpolicy


type NetworkFunctionCollectorPolicyIpfxEmission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/network_function_collector_policy#destination_types NetworkFunctionCollectorPolicy#destination_types}.
	DestinationTypes *[]*string `field:"required" json:"destinationTypes" yaml:"destinationTypes"`
}

