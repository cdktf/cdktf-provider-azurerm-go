package iothubdpssharedaccesspolicy


type IothubDpsSharedAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_dps_shared_access_policy#create IothubDpsSharedAccessPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_dps_shared_access_policy#delete IothubDpsSharedAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_dps_shared_access_policy#read IothubDpsSharedAccessPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub_dps_shared_access_policy#update IothubDpsSharedAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

