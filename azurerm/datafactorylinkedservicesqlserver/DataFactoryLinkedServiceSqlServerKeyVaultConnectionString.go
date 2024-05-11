// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesqlserver


type DataFactoryLinkedServiceSqlServerKeyVaultConnectionString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_linked_service_sql_server#linked_service_name DataFactoryLinkedServiceSqlServer#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_linked_service_sql_server#secret_name DataFactoryLinkedServiceSqlServer#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

