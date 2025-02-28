// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoclustercustomermanagedkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoClusterCustomerManagedKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#cluster_id KustoClusterCustomerManagedKey#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#key_name KustoClusterCustomerManagedKey#key_name}.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#key_vault_id KustoClusterCustomerManagedKey#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#id KustoClusterCustomerManagedKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#key_version KustoClusterCustomerManagedKey#key_version}.
	KeyVersion *string `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#timeouts KustoClusterCustomerManagedKey#timeouts}
	Timeouts *KustoClusterCustomerManagedKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/kusto_cluster_customer_managed_key#user_identity KustoClusterCustomerManagedKey#user_identity}.
	UserIdentity *string `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

