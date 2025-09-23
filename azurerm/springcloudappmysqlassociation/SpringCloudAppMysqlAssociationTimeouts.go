// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudappmysqlassociation


type SpringCloudAppMysqlAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/spring_cloud_app_mysql_association#create SpringCloudAppMysqlAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/spring_cloud_app_mysql_association#delete SpringCloudAppMysqlAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/spring_cloud_app_mysql_association#read SpringCloudAppMysqlAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/spring_cloud_app_mysql_association#update SpringCloudAppMysqlAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

