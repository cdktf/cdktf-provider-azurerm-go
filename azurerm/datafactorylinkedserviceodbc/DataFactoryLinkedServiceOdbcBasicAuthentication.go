// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceodbc


type DataFactoryLinkedServiceOdbcBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_odbc#password DataFactoryLinkedServiceOdbc#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_odbc#username DataFactoryLinkedServiceOdbc#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

