package springcloudapp


type SpringCloudAppIngressSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/spring_cloud_app#backend_protocol SpringCloudApp#backend_protocol}.
	BackendProtocol *string `field:"optional" json:"backendProtocol" yaml:"backendProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/spring_cloud_app#read_timeout_in_seconds SpringCloudApp#read_timeout_in_seconds}.
	ReadTimeoutInSeconds *float64 `field:"optional" json:"readTimeoutInSeconds" yaml:"readTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/spring_cloud_app#send_timeout_in_seconds SpringCloudApp#send_timeout_in_seconds}.
	SendTimeoutInSeconds *float64 `field:"optional" json:"sendTimeoutInSeconds" yaml:"sendTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/spring_cloud_app#session_affinity SpringCloudApp#session_affinity}.
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/spring_cloud_app#session_cookie_max_age SpringCloudApp#session_cookie_max_age}.
	SessionCookieMaxAge *float64 `field:"optional" json:"sessionCookieMaxAge" yaml:"sessionCookieMaxAge"`
}

