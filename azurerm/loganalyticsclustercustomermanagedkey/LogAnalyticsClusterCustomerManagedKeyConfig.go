package loganalyticsclustercustomermanagedkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogAnalyticsClusterCustomerManagedKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/log_analytics_cluster_customer_managed_key#key_vault_key_id LogAnalyticsClusterCustomerManagedKey#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/log_analytics_cluster_customer_managed_key#log_analytics_cluster_id LogAnalyticsClusterCustomerManagedKey#log_analytics_cluster_id}.
	LogAnalyticsClusterId *string `field:"required" json:"logAnalyticsClusterId" yaml:"logAnalyticsClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/log_analytics_cluster_customer_managed_key#id LogAnalyticsClusterCustomerManagedKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/log_analytics_cluster_customer_managed_key#timeouts LogAnalyticsClusterCustomerManagedKey#timeouts}
	Timeouts *LogAnalyticsClusterCustomerManagedKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

