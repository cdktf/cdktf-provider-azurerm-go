package windowswebappslot


type WindowsWebAppSlotSiteConfigApplicationStack struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#current_stack WindowsWebAppSlot#current_stack}.
	CurrentStack *string `field:"optional" json:"currentStack" yaml:"currentStack"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#docker_container_name WindowsWebAppSlot#docker_container_name}.
	DockerContainerName *string `field:"optional" json:"dockerContainerName" yaml:"dockerContainerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#docker_container_registry WindowsWebAppSlot#docker_container_registry}.
	DockerContainerRegistry *string `field:"optional" json:"dockerContainerRegistry" yaml:"dockerContainerRegistry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#docker_container_tag WindowsWebAppSlot#docker_container_tag}.
	DockerContainerTag *string `field:"optional" json:"dockerContainerTag" yaml:"dockerContainerTag"`
	// The version of DotNetCore to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#dotnet_core_version WindowsWebAppSlot#dotnet_core_version}
	DotnetCoreVersion *string `field:"optional" json:"dotnetCoreVersion" yaml:"dotnetCoreVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#dotnet_version WindowsWebAppSlot#dotnet_version}.
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#java_container WindowsWebAppSlot#java_container}.
	JavaContainer *string `field:"optional" json:"javaContainer" yaml:"javaContainer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#java_container_version WindowsWebAppSlot#java_container_version}.
	JavaContainerVersion *string `field:"optional" json:"javaContainerVersion" yaml:"javaContainerVersion"`
	// Should the application use the embedded web server for the version of Java in use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#java_embedded_server_enabled WindowsWebAppSlot#java_embedded_server_enabled}
	JavaEmbeddedServerEnabled interface{} `field:"optional" json:"javaEmbeddedServerEnabled" yaml:"javaEmbeddedServerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#java_version WindowsWebAppSlot#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#node_version WindowsWebAppSlot#node_version}.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#php_version WindowsWebAppSlot#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#python WindowsWebAppSlot#python}.
	Python interface{} `field:"optional" json:"python" yaml:"python"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#python_version WindowsWebAppSlot#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#tomcat_version WindowsWebAppSlot#tomcat_version}.
	TomcatVersion *string `field:"optional" json:"tomcatVersion" yaml:"tomcatVersion"`
}
