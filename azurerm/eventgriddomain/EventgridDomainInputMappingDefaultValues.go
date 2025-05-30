// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgriddomain


type EventgridDomainInputMappingDefaultValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/eventgrid_domain#data_version EventgridDomain#data_version}.
	DataVersion *string `field:"optional" json:"dataVersion" yaml:"dataVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/eventgrid_domain#event_type EventgridDomain#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/eventgrid_domain#subject EventgridDomain#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

