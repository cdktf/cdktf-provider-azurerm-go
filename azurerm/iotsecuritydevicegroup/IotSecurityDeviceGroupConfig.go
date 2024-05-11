// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritydevicegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotSecurityDeviceGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#iothub_id IotSecurityDeviceGroup#iothub_id}.
	IothubId *string `field:"required" json:"iothubId" yaml:"iothubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#name IotSecurityDeviceGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// allow_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#allow_rule IotSecurityDeviceGroup#allow_rule}
	AllowRule *IotSecurityDeviceGroupAllowRule `field:"optional" json:"allowRule" yaml:"allowRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#id IotSecurityDeviceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// range_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#range_rule IotSecurityDeviceGroup#range_rule}
	RangeRule interface{} `field:"optional" json:"rangeRule" yaml:"rangeRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/iot_security_device_group#timeouts IotSecurityDeviceGroup#timeouts}
	Timeouts *IotSecurityDeviceGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

