// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgriddomaintopic


type EventgridDomainTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/eventgrid_domain_topic#create EventgridDomainTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/eventgrid_domain_topic#delete EventgridDomainTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/eventgrid_domain_topic#read EventgridDomainTopic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

