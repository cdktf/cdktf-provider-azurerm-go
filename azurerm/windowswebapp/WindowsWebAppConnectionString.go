// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppConnectionString struct {
	// The name which should be used for this Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app#name WindowsWebApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of database. Possible values include: `MySQL`, `SQLServer`, `SQLAzure`, `Custom`, `NotificationHub`, `ServiceBus`, `EventHub`, `APIHub`, `DocDb`, `RedisCache`, and `PostgreSQL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app#type WindowsWebApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The connection string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/windows_web_app#value WindowsWebApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

