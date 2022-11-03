package nginxconfiguration


type NginxConfigurationConfigFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_configuration#content NginxConfiguration#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_configuration#virtual_path NginxConfiguration#virtual_path}.
	VirtualPath *string `field:"required" json:"virtualPath" yaml:"virtualPath"`
}

