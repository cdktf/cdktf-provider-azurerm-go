// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package portaldashboard


type PortalDashboardTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/portal_dashboard#create PortalDashboard#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/portal_dashboard#delete PortalDashboard#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/portal_dashboard#read PortalDashboard#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/portal_dashboard#update PortalDashboard#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

