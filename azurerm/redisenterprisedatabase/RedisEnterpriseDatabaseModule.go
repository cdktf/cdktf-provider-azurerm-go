// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redisenterprisedatabase


type RedisEnterpriseDatabaseModule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/redis_enterprise_database#name RedisEnterpriseDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/redis_enterprise_database#args RedisEnterpriseDatabase#args}.
	Args *string `field:"optional" json:"args" yaml:"args"`
}

