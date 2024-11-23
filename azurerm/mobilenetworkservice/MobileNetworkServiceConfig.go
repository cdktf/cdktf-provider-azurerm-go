// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#location MobileNetworkService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#mobile_network_id MobileNetworkService#mobile_network_id}.
	MobileNetworkId *string `field:"required" json:"mobileNetworkId" yaml:"mobileNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#name MobileNetworkService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// pcc_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#pcc_rule MobileNetworkService#pcc_rule}
	PccRule interface{} `field:"required" json:"pccRule" yaml:"pccRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#service_precedence MobileNetworkService#service_precedence}.
	ServicePrecedence *float64 `field:"required" json:"servicePrecedence" yaml:"servicePrecedence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#id MobileNetworkService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// service_qos_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#service_qos_policy MobileNetworkService#service_qos_policy}
	ServiceQosPolicy *MobileNetworkServiceServiceQosPolicy `field:"optional" json:"serviceQosPolicy" yaml:"serviceQosPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#tags MobileNetworkService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/mobile_network_service#timeouts MobileNetworkService#timeouts}
	Timeouts *MobileNetworkServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

