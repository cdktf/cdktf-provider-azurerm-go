// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleexadatainfrastructure


type OracleExadataInfrastructureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_exadata_infrastructure#create OracleExadataInfrastructure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_exadata_infrastructure#delete OracleExadataInfrastructure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_exadata_infrastructure#read OracleExadataInfrastructure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_exadata_infrastructure#update OracleExadataInfrastructure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

