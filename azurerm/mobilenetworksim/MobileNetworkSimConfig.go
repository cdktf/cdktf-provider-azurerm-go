// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksim

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkSimConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#authentication_key MobileNetworkSim#authentication_key}.
	AuthenticationKey *string `field:"required" json:"authenticationKey" yaml:"authenticationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#integrated_circuit_card_identifier MobileNetworkSim#integrated_circuit_card_identifier}.
	IntegratedCircuitCardIdentifier *string `field:"required" json:"integratedCircuitCardIdentifier" yaml:"integratedCircuitCardIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#international_mobile_subscriber_identity MobileNetworkSim#international_mobile_subscriber_identity}.
	InternationalMobileSubscriberIdentity *string `field:"required" json:"internationalMobileSubscriberIdentity" yaml:"internationalMobileSubscriberIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#mobile_network_sim_group_id MobileNetworkSim#mobile_network_sim_group_id}.
	MobileNetworkSimGroupId *string `field:"required" json:"mobileNetworkSimGroupId" yaml:"mobileNetworkSimGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#name MobileNetworkSim#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#operator_key_code MobileNetworkSim#operator_key_code}.
	OperatorKeyCode *string `field:"required" json:"operatorKeyCode" yaml:"operatorKeyCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#device_type MobileNetworkSim#device_type}.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#id MobileNetworkSim#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#sim_policy_id MobileNetworkSim#sim_policy_id}.
	SimPolicyId *string `field:"optional" json:"simPolicyId" yaml:"simPolicyId"`
	// static_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#static_ip_configuration MobileNetworkSim#static_ip_configuration}
	StaticIpConfiguration interface{} `field:"optional" json:"staticIpConfiguration" yaml:"staticIpConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mobile_network_sim#timeouts MobileNetworkSim#timeouts}
	Timeouts *MobileNetworkSimTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

