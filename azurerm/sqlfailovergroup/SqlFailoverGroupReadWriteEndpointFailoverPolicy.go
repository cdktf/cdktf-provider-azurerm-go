// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlfailovergroup


type SqlFailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sql_failover_group#mode SqlFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sql_failover_group#grace_minutes SqlFailoverGroup#grace_minutes}.
	GraceMinutes *float64 `field:"optional" json:"graceMinutes" yaml:"graceMinutes"`
}

