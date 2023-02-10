package containerapp


type ContainerAppDapr struct {
	// The Dapr Application Identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#app_id ContainerApp#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The port which the application is listening on. This is the same as the `ingress` port.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#app_port ContainerApp#app_port}
	AppPort *float64 `field:"required" json:"appPort" yaml:"appPort"`
	// The protocol for the app. Possible values include `http` and `grpc`. Defaults to `http`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#app_protocol ContainerApp#app_protocol}
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
}

