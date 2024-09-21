// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceodata


type DataFactoryLinkedServiceOdataBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_odata#password DataFactoryLinkedServiceOdata#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_odata#username DataFactoryLinkedServiceOdata#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

