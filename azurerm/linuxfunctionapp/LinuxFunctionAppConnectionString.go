// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppConnectionString struct {
	// The name which should be used for this Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/linux_function_app#name LinuxFunctionApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of database. Possible values include: `MySQL`, `SQLServer`, `SQLAzure`, `Custom`, `NotificationHub`, `ServiceBus`, `EventHub`, `APIHub`, `DocDb`, `RedisCache`, and `PostgreSQL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/linux_function_app#type LinuxFunctionApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The connection string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/linux_function_app#value LinuxFunctionApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

