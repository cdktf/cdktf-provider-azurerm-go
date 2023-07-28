package recoveryservicesvaultresourceguardassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecoveryServicesVaultResourceGuardAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/recovery_services_vault_resource_guard_association#name RecoveryServicesVaultResourceGuardAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/recovery_services_vault_resource_guard_association#resource_guard_id RecoveryServicesVaultResourceGuardAssociation#resource_guard_id}.
	ResourceGuardId *string `field:"required" json:"resourceGuardId" yaml:"resourceGuardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/recovery_services_vault_resource_guard_association#vault_id RecoveryServicesVaultResourceGuardAssociation#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/recovery_services_vault_resource_guard_association#id RecoveryServicesVaultResourceGuardAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/recovery_services_vault_resource_guard_association#timeouts RecoveryServicesVaultResourceGuardAssociation#timeouts}
	Timeouts *RecoveryServicesVaultResourceGuardAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

