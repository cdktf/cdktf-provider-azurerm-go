package apimanagement


type ApiManagementProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_http2 ApiManagement#enable_http2}.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
}
