// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolFixedScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/batch_pool#node_deallocation_method BatchPool#node_deallocation_method}.
	NodeDeallocationMethod *string `field:"optional" json:"nodeDeallocationMethod" yaml:"nodeDeallocationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/batch_pool#resize_timeout BatchPool#resize_timeout}.
	ResizeTimeout *string `field:"optional" json:"resizeTimeout" yaml:"resizeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/batch_pool#target_dedicated_nodes BatchPool#target_dedicated_nodes}.
	TargetDedicatedNodes *float64 `field:"optional" json:"targetDedicatedNodes" yaml:"targetDedicatedNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/batch_pool#target_low_priority_nodes BatchPool#target_low_priority_nodes}.
	TargetLowPriorityNodes *float64 `field:"optional" json:"targetLowPriorityNodes" yaml:"targetLowPriorityNodes"`
}

