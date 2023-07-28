package confidentialledger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfidentialLedgerConfig struct {
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
	// azuread_based_service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#azuread_based_service_principal ConfidentialLedger#azuread_based_service_principal}
	AzureadBasedServicePrincipal interface{} `field:"required" json:"azureadBasedServicePrincipal" yaml:"azureadBasedServicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#ledger_type ConfidentialLedger#ledger_type}.
	LedgerType *string `field:"required" json:"ledgerType" yaml:"ledgerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#location ConfidentialLedger#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#name ConfidentialLedger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#resource_group_name ConfidentialLedger#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// certificate_based_security_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#certificate_based_security_principal ConfidentialLedger#certificate_based_security_principal}
	CertificateBasedSecurityPrincipal interface{} `field:"optional" json:"certificateBasedSecurityPrincipal" yaml:"certificateBasedSecurityPrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#id ConfidentialLedger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#tags ConfidentialLedger#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/confidential_ledger#timeouts ConfidentialLedger#timeouts}
	Timeouts *ConfidentialLedgerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

