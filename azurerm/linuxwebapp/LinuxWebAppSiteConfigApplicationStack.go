package linuxwebapp


type LinuxWebAppSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#docker_image LinuxWebApp#docker_image}.
	DockerImage *string `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#docker_image_tag LinuxWebApp#docker_image_tag}.
	DockerImageTag *string `field:"optional" json:"dockerImageTag" yaml:"dockerImageTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#dotnet_version LinuxWebApp#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#go_version LinuxWebApp#go_version}.
	GoVersion *string `field:"optional" json:"goVersion" yaml:"goVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#java_server LinuxWebApp#java_server}.
	JavaServer *string `field:"optional" json:"javaServer" yaml:"javaServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#java_server_version LinuxWebApp#java_server_version}.
	JavaServerVersion *string `field:"optional" json:"javaServerVersion" yaml:"javaServerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#java_version LinuxWebApp#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#node_version LinuxWebApp#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#php_version LinuxWebApp#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#python_version LinuxWebApp#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/linux_web_app#ruby_version LinuxWebApp#ruby_version}.
	RubyVersion *string `field:"optional" json:"rubyVersion" yaml:"rubyVersion"`
}

