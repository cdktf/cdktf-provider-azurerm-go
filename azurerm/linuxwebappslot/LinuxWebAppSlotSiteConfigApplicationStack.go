package linuxwebappslot


type LinuxWebAppSlotSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_image LinuxWebAppSlot#docker_image}.
	DockerImage *string `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_image_name LinuxWebAppSlot#docker_image_name}.
	DockerImageName *string `field:"optional" json:"dockerImageName" yaml:"dockerImageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_image_tag LinuxWebAppSlot#docker_image_tag}.
	DockerImageTag *string `field:"optional" json:"dockerImageTag" yaml:"dockerImageTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_registry_password LinuxWebAppSlot#docker_registry_password}.
	DockerRegistryPassword *string `field:"optional" json:"dockerRegistryPassword" yaml:"dockerRegistryPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_registry_url LinuxWebAppSlot#docker_registry_url}.
	DockerRegistryUrl *string `field:"optional" json:"dockerRegistryUrl" yaml:"dockerRegistryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#docker_registry_username LinuxWebAppSlot#docker_registry_username}.
	DockerRegistryUsername *string `field:"optional" json:"dockerRegistryUsername" yaml:"dockerRegistryUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#dotnet_version LinuxWebAppSlot#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#go_version LinuxWebAppSlot#go_version}.
	GoVersion *string `field:"optional" json:"goVersion" yaml:"goVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#java_server LinuxWebAppSlot#java_server}.
	JavaServer *string `field:"optional" json:"javaServer" yaml:"javaServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#java_server_version LinuxWebAppSlot#java_server_version}.
	JavaServerVersion *string `field:"optional" json:"javaServerVersion" yaml:"javaServerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#java_version LinuxWebAppSlot#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#node_version LinuxWebAppSlot#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#php_version LinuxWebAppSlot#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#python_version LinuxWebAppSlot#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_web_app_slot#ruby_version LinuxWebAppSlot#ruby_version}.
	RubyVersion *string `field:"optional" json:"rubyVersion" yaml:"rubyVersion"`
}

