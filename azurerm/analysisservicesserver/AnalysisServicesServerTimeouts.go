// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package analysisservicesserver


type AnalysisServicesServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/analysis_services_server#create AnalysisServicesServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/analysis_services_server#delete AnalysisServicesServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/analysis_services_server#read AnalysisServicesServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/analysis_services_server#update AnalysisServicesServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

