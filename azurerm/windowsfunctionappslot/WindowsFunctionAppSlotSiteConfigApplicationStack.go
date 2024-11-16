// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotSiteConfigApplicationStack struct {
	// The version of .Net. Possible values are `v3.0`, `v4.0`, `v6.0`, `v7.0`, `v8.0` and `v9.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#dotnet_version WindowsFunctionAppSlot#dotnet_version}
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// The version of Java to use. Possible values are `1.8`, `11` and `17`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#java_version WindowsFunctionAppSlot#java_version}
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// The version of Node to use. Possible values include `12`, `14`, `16` and `18`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#node_version WindowsFunctionAppSlot#node_version}
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// The PowerShell Core version to use. Possible values are `7`, `7.2`, and `7.4`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#powershell_core_version WindowsFunctionAppSlot#powershell_core_version}
	PowershellCoreVersion *string `field:"optional" json:"powershellCoreVersion" yaml:"powershellCoreVersion"`
	// Does the Function App use a custom Application Stack?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#use_custom_runtime WindowsFunctionAppSlot#use_custom_runtime}
	UseCustomRuntime interface{} `field:"optional" json:"useCustomRuntime" yaml:"useCustomRuntime"`
	// Should the DotNet process use an isolated runtime. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#use_dotnet_isolated_runtime WindowsFunctionAppSlot#use_dotnet_isolated_runtime}
	UseDotnetIsolatedRuntime interface{} `field:"optional" json:"useDotnetIsolatedRuntime" yaml:"useDotnetIsolatedRuntime"`
}

