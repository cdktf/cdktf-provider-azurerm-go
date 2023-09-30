// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesparkpool


type SynapseSparkPoolAutoPause struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/synapse_spark_pool#delay_in_minutes SynapseSparkPool#delay_in_minutes}.
	DelayInMinutes *float64 `field:"required" json:"delayInMinutes" yaml:"delayInMinutes"`
}

