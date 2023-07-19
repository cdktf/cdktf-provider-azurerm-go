package publicip

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PublicIpConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#allocation_method PublicIp#allocation_method}.
	AllocationMethod *string `field:"required" json:"allocationMethod" yaml:"allocationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#location PublicIp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#name PublicIp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#resource_group_name PublicIp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#ddos_protection_mode PublicIp#ddos_protection_mode}.
	DdosProtectionMode *string `field:"optional" json:"ddosProtectionMode" yaml:"ddosProtectionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#ddos_protection_plan_id PublicIp#ddos_protection_plan_id}.
	DdosProtectionPlanId *string `field:"optional" json:"ddosProtectionPlanId" yaml:"ddosProtectionPlanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#domain_name_label PublicIp#domain_name_label}.
	DomainNameLabel *string `field:"optional" json:"domainNameLabel" yaml:"domainNameLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#edge_zone PublicIp#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#id PublicIp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#idle_timeout_in_minutes PublicIp#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#ip_tags PublicIp#ip_tags}.
	IpTags *map[string]*string `field:"optional" json:"ipTags" yaml:"ipTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#ip_version PublicIp#ip_version}.
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#public_ip_prefix_id PublicIp#public_ip_prefix_id}.
	PublicIpPrefixId *string `field:"optional" json:"publicIpPrefixId" yaml:"publicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#reverse_fqdn PublicIp#reverse_fqdn}.
	ReverseFqdn *string `field:"optional" json:"reverseFqdn" yaml:"reverseFqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#sku PublicIp#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#sku_tier PublicIp#sku_tier}.
	SkuTier *string `field:"optional" json:"skuTier" yaml:"skuTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#tags PublicIp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#timeouts PublicIp#timeouts}
	Timeouts *PublicIpTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/public_ip#zones PublicIp#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

