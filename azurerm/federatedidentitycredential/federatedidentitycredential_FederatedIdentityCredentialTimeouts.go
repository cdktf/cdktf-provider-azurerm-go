package federatedidentitycredential


type FederatedIdentityCredentialTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/federated_identity_credential#create FederatedIdentityCredential#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/federated_identity_credential#delete FederatedIdentityCredential#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/federated_identity_credential#read FederatedIdentityCredential#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
