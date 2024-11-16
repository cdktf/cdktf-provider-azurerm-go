// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#docker_image_name LinuxWebApp#docker_image_name}.
	DockerImageName *string `field:"optional" json:"dockerImageName" yaml:"dockerImageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#docker_registry_password LinuxWebApp#docker_registry_password}.
	DockerRegistryPassword *string `field:"optional" json:"dockerRegistryPassword" yaml:"dockerRegistryPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#docker_registry_url LinuxWebApp#docker_registry_url}.
	DockerRegistryUrl *string `field:"optional" json:"dockerRegistryUrl" yaml:"dockerRegistryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#docker_registry_username LinuxWebApp#docker_registry_username}.
	DockerRegistryUsername *string `field:"optional" json:"dockerRegistryUsername" yaml:"dockerRegistryUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#dotnet_version LinuxWebApp#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#go_version LinuxWebApp#go_version}.
	GoVersion *string `field:"optional" json:"goVersion" yaml:"goVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#java_server LinuxWebApp#java_server}.
	JavaServer *string `field:"optional" json:"javaServer" yaml:"javaServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#java_server_version LinuxWebApp#java_server_version}.
	JavaServerVersion *string `field:"optional" json:"javaServerVersion" yaml:"javaServerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#java_version LinuxWebApp#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#node_version LinuxWebApp#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#php_version LinuxWebApp#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#python_version LinuxWebApp#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/linux_web_app#ruby_version LinuxWebApp#ruby_version}.
	RubyVersion *string `field:"optional" json:"rubyVersion" yaml:"rubyVersion"`
}

