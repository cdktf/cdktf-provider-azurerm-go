// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggertumblingwindow


type DataFactoryTriggerTumblingWindowRetry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/data_factory_trigger_tumbling_window#count DataFactoryTriggerTumblingWindow#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/data_factory_trigger_tumbling_window#interval DataFactoryTriggerTumblingWindow#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

