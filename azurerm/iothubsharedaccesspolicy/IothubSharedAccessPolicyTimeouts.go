package iothubsharedaccesspolicy


type IothubSharedAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/iothub_shared_access_policy#create IothubSharedAccessPolicyA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/iothub_shared_access_policy#delete IothubSharedAccessPolicyA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/iothub_shared_access_policy#read IothubSharedAccessPolicyA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/iothub_shared_access_policy#update IothubSharedAccessPolicyA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

