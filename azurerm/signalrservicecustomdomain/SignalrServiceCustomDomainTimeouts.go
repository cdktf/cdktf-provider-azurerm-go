// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservicecustomdomain


type SignalrServiceCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/signalr_service_custom_domain#create SignalrServiceCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/signalr_service_custom_domain#delete SignalrServiceCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/signalr_service_custom_domain#read SignalrServiceCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

