// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitHostNetworkIntent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#adapter StackHciDeploymentSetting#adapter}.
	Adapter *[]*string `field:"required" json:"adapter" yaml:"adapter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#name StackHciDeploymentSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#traffic_type StackHciDeploymentSetting#traffic_type}.
	TrafficType *[]*string `field:"required" json:"trafficType" yaml:"trafficType"`
	// adapter_property_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#adapter_property_override StackHciDeploymentSetting#adapter_property_override}
	AdapterPropertyOverride *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride `field:"optional" json:"adapterPropertyOverride" yaml:"adapterPropertyOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#adapter_property_override_enabled StackHciDeploymentSetting#adapter_property_override_enabled}.
	AdapterPropertyOverrideEnabled interface{} `field:"optional" json:"adapterPropertyOverrideEnabled" yaml:"adapterPropertyOverrideEnabled"`
	// qos_policy_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#qos_policy_override StackHciDeploymentSetting#qos_policy_override}
	QosPolicyOverride *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride `field:"optional" json:"qosPolicyOverride" yaml:"qosPolicyOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#qos_policy_override_enabled StackHciDeploymentSetting#qos_policy_override_enabled}.
	QosPolicyOverrideEnabled interface{} `field:"optional" json:"qosPolicyOverrideEnabled" yaml:"qosPolicyOverrideEnabled"`
	// virtual_switch_configuration_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#virtual_switch_configuration_override StackHciDeploymentSetting#virtual_switch_configuration_override}
	VirtualSwitchConfigurationOverride *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride `field:"optional" json:"virtualSwitchConfigurationOverride" yaml:"virtualSwitchConfigurationOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/stack_hci_deployment_setting#virtual_switch_configuration_override_enabled StackHciDeploymentSetting#virtual_switch_configuration_override_enabled}.
	VirtualSwitchConfigurationOverrideEnabled interface{} `field:"optional" json:"virtualSwitchConfigurationOverrideEnabled" yaml:"virtualSwitchConfigurationOverrideEnabled"`
}

