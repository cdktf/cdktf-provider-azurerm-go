// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustocluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#location KustoCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#name KustoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#resource_group_name KustoCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#sku KustoCluster#sku}
	Sku *KustoClusterSku `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#allowed_fqdns KustoCluster#allowed_fqdns}.
	AllowedFqdns *[]*string `field:"optional" json:"allowedFqdns" yaml:"allowedFqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#allowed_ip_ranges KustoCluster#allowed_ip_ranges}.
	AllowedIpRanges *[]*string `field:"optional" json:"allowedIpRanges" yaml:"allowedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#auto_stop_enabled KustoCluster#auto_stop_enabled}.
	AutoStopEnabled interface{} `field:"optional" json:"autoStopEnabled" yaml:"autoStopEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#disk_encryption_enabled KustoCluster#disk_encryption_enabled}.
	DiskEncryptionEnabled interface{} `field:"optional" json:"diskEncryptionEnabled" yaml:"diskEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#double_encryption_enabled KustoCluster#double_encryption_enabled}.
	DoubleEncryptionEnabled interface{} `field:"optional" json:"doubleEncryptionEnabled" yaml:"doubleEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#engine KustoCluster#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#id KustoCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#identity KustoCluster#identity}
	Identity *KustoClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#language_extensions KustoCluster#language_extensions}.
	LanguageExtensions *[]*string `field:"optional" json:"languageExtensions" yaml:"languageExtensions"`
	// optimized_auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#optimized_auto_scale KustoCluster#optimized_auto_scale}
	OptimizedAutoScale *KustoClusterOptimizedAutoScale `field:"optional" json:"optimizedAutoScale" yaml:"optimizedAutoScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#outbound_network_access_restricted KustoCluster#outbound_network_access_restricted}.
	OutboundNetworkAccessRestricted interface{} `field:"optional" json:"outboundNetworkAccessRestricted" yaml:"outboundNetworkAccessRestricted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#public_ip_type KustoCluster#public_ip_type}.
	PublicIpType *string `field:"optional" json:"publicIpType" yaml:"publicIpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#public_network_access_enabled KustoCluster#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#purge_enabled KustoCluster#purge_enabled}.
	PurgeEnabled interface{} `field:"optional" json:"purgeEnabled" yaml:"purgeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#streaming_ingestion_enabled KustoCluster#streaming_ingestion_enabled}.
	StreamingIngestionEnabled interface{} `field:"optional" json:"streamingIngestionEnabled" yaml:"streamingIngestionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#tags KustoCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#timeouts KustoCluster#timeouts}
	Timeouts *KustoClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#trusted_external_tenants KustoCluster#trusted_external_tenants}.
	TrustedExternalTenants *[]*string `field:"optional" json:"trustedExternalTenants" yaml:"trustedExternalTenants"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#virtual_network_configuration KustoCluster#virtual_network_configuration}
	VirtualNetworkConfiguration *KustoClusterVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/kusto_cluster#zones KustoCluster#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

