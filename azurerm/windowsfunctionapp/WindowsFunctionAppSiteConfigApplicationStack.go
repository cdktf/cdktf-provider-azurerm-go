package windowsfunctionapp


type WindowsFunctionAppSiteConfigApplicationStack struct {
	// The version of .Net. Possible values are `v3.0`, `v4.0`, `v6.0` and `v7.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#dotnet_version WindowsFunctionApp#dotnet_version}
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// The version of Java to use. Possible values are `1.8`, `11` and `17`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#java_version WindowsFunctionApp#java_version}
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// The version of Node to use. Possible values include `12`, `14`, `16` and `18`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#node_version WindowsFunctionApp#node_version}
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// The PowerShell Core version to use. Possible values are `7`, and `7.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#powershell_core_version WindowsFunctionApp#powershell_core_version}
	PowershellCoreVersion *string `field:"optional" json:"powershellCoreVersion" yaml:"powershellCoreVersion"`
	// Does the Function App use a custom Application Stack?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#use_custom_runtime WindowsFunctionApp#use_custom_runtime}
	UseCustomRuntime interface{} `field:"optional" json:"useCustomRuntime" yaml:"useCustomRuntime"`
	// Should the DotNet process use an isolated runtime. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_function_app#use_dotnet_isolated_runtime WindowsFunctionApp#use_dotnet_isolated_runtime}
	UseDotnetIsolatedRuntime interface{} `field:"optional" json:"useDotnetIsolatedRuntime" yaml:"useDotnetIsolatedRuntime"`
}

