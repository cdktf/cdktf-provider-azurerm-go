package apimanagementgatewaycertificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementGatewayCertificateAuthorityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#api_management_id ApiManagementGatewayCertificateAuthority#api_management_id}.
	ApiManagementId *string `field:"required" json:"apiManagementId" yaml:"apiManagementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#certificate_name ApiManagementGatewayCertificateAuthority#certificate_name}.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#gateway_name ApiManagementGatewayCertificateAuthority#gateway_name}.
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#id ApiManagementGatewayCertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#is_trusted ApiManagementGatewayCertificateAuthority#is_trusted}.
	IsTrusted interface{} `field:"optional" json:"isTrusted" yaml:"isTrusted"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/api_management_gateway_certificate_authority#timeouts ApiManagementGatewayCertificateAuthority#timeouts}
	Timeouts *ApiManagementGatewayCertificateAuthorityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

