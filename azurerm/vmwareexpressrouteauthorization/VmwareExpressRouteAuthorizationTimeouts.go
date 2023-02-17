package vmwareexpressrouteauthorization


type VmwareExpressRouteAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_express_route_authorization#create VmwareExpressRouteAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_express_route_authorization#delete VmwareExpressRouteAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vmware_express_route_authorization#read VmwareExpressRouteAuthorization#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

