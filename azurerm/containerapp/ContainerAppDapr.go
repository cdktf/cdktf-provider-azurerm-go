// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppDapr struct {
	// The Dapr Application Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app#app_id ContainerApp#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The port which the application is listening on. This is the same as the `ingress` port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app#app_port ContainerApp#app_port}
	AppPort *float64 `field:"optional" json:"appPort" yaml:"appPort"`
	// The protocol for the app. Possible values include `http` and `grpc`. Defaults to `http`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app#app_protocol ContainerApp#app_protocol}
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
}

