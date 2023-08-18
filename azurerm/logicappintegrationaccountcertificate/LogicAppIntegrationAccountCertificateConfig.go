package logicappintegrationaccountcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppIntegrationAccountCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#integration_account_name LogicAppIntegrationAccountCertificate#integration_account_name}.
	IntegrationAccountName *string `field:"required" json:"integrationAccountName" yaml:"integrationAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#name LogicAppIntegrationAccountCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#resource_group_name LogicAppIntegrationAccountCertificate#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#id LogicAppIntegrationAccountCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key_vault_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#key_vault_key LogicAppIntegrationAccountCertificate#key_vault_key}
	KeyVaultKey *LogicAppIntegrationAccountCertificateKeyVaultKey `field:"optional" json:"keyVaultKey" yaml:"keyVaultKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#metadata LogicAppIntegrationAccountCertificate#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#public_certificate LogicAppIntegrationAccountCertificate#public_certificate}.
	PublicCertificate *string `field:"optional" json:"publicCertificate" yaml:"publicCertificate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/logic_app_integration_account_certificate#timeouts LogicAppIntegrationAccountCertificate#timeouts}
	Timeouts *LogicAppIntegrationAccountCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

