package eventhubnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventhubNamespaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#location EventhubNamespace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#name EventhubNamespace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#resource_group_name EventhubNamespace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#sku EventhubNamespace#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#auto_inflate_enabled EventhubNamespace#auto_inflate_enabled}.
	AutoInflateEnabled interface{} `field:"optional" json:"autoInflateEnabled" yaml:"autoInflateEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#capacity EventhubNamespace#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#dedicated_cluster_id EventhubNamespace#dedicated_cluster_id}.
	DedicatedClusterId *string `field:"optional" json:"dedicatedClusterId" yaml:"dedicatedClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#id EventhubNamespace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#identity EventhubNamespace#identity}
	Identity *EventhubNamespaceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#local_authentication_enabled EventhubNamespace#local_authentication_enabled}.
	LocalAuthenticationEnabled interface{} `field:"optional" json:"localAuthenticationEnabled" yaml:"localAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#maximum_throughput_units EventhubNamespace#maximum_throughput_units}.
	MaximumThroughputUnits *float64 `field:"optional" json:"maximumThroughputUnits" yaml:"maximumThroughputUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#minimum_tls_version EventhubNamespace#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#network_rulesets EventhubNamespace#network_rulesets}.
	NetworkRulesets interface{} `field:"optional" json:"networkRulesets" yaml:"networkRulesets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#public_network_access_enabled EventhubNamespace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#tags EventhubNamespace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#timeouts EventhubNamespace#timeouts}
	Timeouts *EventhubNamespaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventhub_namespace#zone_redundant EventhubNamespace#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

