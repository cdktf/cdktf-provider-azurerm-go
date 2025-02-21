// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridnamespace


type EventgridNamespaceTopicSpacesConfigurationStaticRoutingEnrichment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/eventgrid_namespace#key EventgridNamespace#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/eventgrid_namespace#value EventgridNamespace#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

