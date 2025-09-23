// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SearchServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#location SearchService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#name SearchService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#resource_group_name SearchService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#sku SearchService#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#allowed_ips SearchService#allowed_ips}.
	AllowedIps *[]*string `field:"optional" json:"allowedIps" yaml:"allowedIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#authentication_failure_mode SearchService#authentication_failure_mode}.
	AuthenticationFailureMode *string `field:"optional" json:"authenticationFailureMode" yaml:"authenticationFailureMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#customer_managed_key_enforcement_enabled SearchService#customer_managed_key_enforcement_enabled}.
	CustomerManagedKeyEnforcementEnabled interface{} `field:"optional" json:"customerManagedKeyEnforcementEnabled" yaml:"customerManagedKeyEnforcementEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#hosting_mode SearchService#hosting_mode}.
	HostingMode *string `field:"optional" json:"hostingMode" yaml:"hostingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#id SearchService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#identity SearchService#identity}
	Identity *SearchServiceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#local_authentication_enabled SearchService#local_authentication_enabled}.
	LocalAuthenticationEnabled interface{} `field:"optional" json:"localAuthenticationEnabled" yaml:"localAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#network_rule_bypass_option SearchService#network_rule_bypass_option}.
	NetworkRuleBypassOption *string `field:"optional" json:"networkRuleBypassOption" yaml:"networkRuleBypassOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#partition_count SearchService#partition_count}.
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#public_network_access_enabled SearchService#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#replica_count SearchService#replica_count}.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#semantic_search_sku SearchService#semantic_search_sku}.
	SemanticSearchSku *string `field:"optional" json:"semanticSearchSku" yaml:"semanticSearchSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#tags SearchService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/search_service#timeouts SearchService#timeouts}
	Timeouts *SearchServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

