package apimanagementidentityprovideraadb2c

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementIdentityProviderAadb2CConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#allowed_tenant ApiManagementIdentityProviderAadb2C#allowed_tenant}.
	AllowedTenant *string `field:"required" json:"allowedTenant" yaml:"allowedTenant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#api_management_name ApiManagementIdentityProviderAadb2C#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#authority ApiManagementIdentityProviderAadb2C#authority}.
	Authority *string `field:"required" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#client_id ApiManagementIdentityProviderAadb2C#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#client_secret ApiManagementIdentityProviderAadb2C#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#resource_group_name ApiManagementIdentityProviderAadb2C#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#signin_policy ApiManagementIdentityProviderAadb2C#signin_policy}.
	SigninPolicy *string `field:"required" json:"signinPolicy" yaml:"signinPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#signin_tenant ApiManagementIdentityProviderAadb2C#signin_tenant}.
	SigninTenant *string `field:"required" json:"signinTenant" yaml:"signinTenant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#signup_policy ApiManagementIdentityProviderAadb2C#signup_policy}.
	SignupPolicy *string `field:"required" json:"signupPolicy" yaml:"signupPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#id ApiManagementIdentityProviderAadb2C#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#password_reset_policy ApiManagementIdentityProviderAadb2C#password_reset_policy}.
	PasswordResetPolicy *string `field:"optional" json:"passwordResetPolicy" yaml:"passwordResetPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#profile_editing_policy ApiManagementIdentityProviderAadb2C#profile_editing_policy}.
	ProfileEditingPolicy *string `field:"optional" json:"profileEditingPolicy" yaml:"profileEditingPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/api_management_identity_provider_aadb2c#timeouts ApiManagementIdentityProviderAadb2C#timeouts}
	Timeouts *ApiManagementIdentityProviderAadb2CTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

