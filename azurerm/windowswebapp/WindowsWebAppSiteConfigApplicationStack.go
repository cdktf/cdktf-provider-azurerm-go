// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#current_stack WindowsWebApp#current_stack}.
	CurrentStack *string `field:"optional" json:"currentStack" yaml:"currentStack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#docker_image_name WindowsWebApp#docker_image_name}.
	DockerImageName *string `field:"optional" json:"dockerImageName" yaml:"dockerImageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#docker_registry_password WindowsWebApp#docker_registry_password}.
	DockerRegistryPassword *string `field:"optional" json:"dockerRegistryPassword" yaml:"dockerRegistryPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#docker_registry_url WindowsWebApp#docker_registry_url}.
	DockerRegistryUrl *string `field:"optional" json:"dockerRegistryUrl" yaml:"dockerRegistryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#docker_registry_username WindowsWebApp#docker_registry_username}.
	DockerRegistryUsername *string `field:"optional" json:"dockerRegistryUsername" yaml:"dockerRegistryUsername"`
	// The version of DotNetCore to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#dotnet_core_version WindowsWebApp#dotnet_core_version}
	DotnetCoreVersion *string `field:"optional" json:"dotnetCoreVersion" yaml:"dotnetCoreVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#dotnet_version WindowsWebApp#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#java_container WindowsWebApp#java_container}.
	JavaContainer *string `field:"optional" json:"javaContainer" yaml:"javaContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#java_container_version WindowsWebApp#java_container_version}.
	JavaContainerVersion *string `field:"optional" json:"javaContainerVersion" yaml:"javaContainerVersion"`
	// Should the application use the embedded web server for the version of Java in use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#java_embedded_server_enabled WindowsWebApp#java_embedded_server_enabled}
	JavaEmbeddedServerEnabled interface{} `field:"optional" json:"javaEmbeddedServerEnabled" yaml:"javaEmbeddedServerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#java_version WindowsWebApp#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#node_version WindowsWebApp#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#php_version WindowsWebApp#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#python WindowsWebApp#python}.
	Python interface{} `field:"optional" json:"python" yaml:"python"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app#tomcat_version WindowsWebApp#tomcat_version}.
	TomcatVersion *string `field:"optional" json:"tomcatVersion" yaml:"tomcatVersion"`
}

