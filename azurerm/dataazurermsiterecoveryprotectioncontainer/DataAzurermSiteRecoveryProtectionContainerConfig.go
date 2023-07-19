package dataazurermsiterecoveryprotectioncontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermSiteRecoveryProtectionContainerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#name DataAzurermSiteRecoveryProtectionContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#recovery_fabric_name DataAzurermSiteRecoveryProtectionContainer#recovery_fabric_name}.
	RecoveryFabricName *string `field:"required" json:"recoveryFabricName" yaml:"recoveryFabricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#recovery_vault_name DataAzurermSiteRecoveryProtectionContainer#recovery_vault_name}.
	RecoveryVaultName *string `field:"required" json:"recoveryVaultName" yaml:"recoveryVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#resource_group_name DataAzurermSiteRecoveryProtectionContainer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#id DataAzurermSiteRecoveryProtectionContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/data-sources/site_recovery_protection_container#timeouts DataAzurermSiteRecoveryProtectionContainer#timeouts}
	Timeouts *DataAzurermSiteRecoveryProtectionContainerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

