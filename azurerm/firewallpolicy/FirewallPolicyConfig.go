package firewallpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#location FirewallPolicy#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#name FirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#resource_group_name FirewallPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#auto_learn_private_ranges_enabled FirewallPolicy#auto_learn_private_ranges_enabled}.
	AutoLearnPrivateRangesEnabled interface{} `field:"optional" json:"autoLearnPrivateRangesEnabled" yaml:"autoLearnPrivateRangesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#base_policy_id FirewallPolicy#base_policy_id}.
	BasePolicyId *string `field:"optional" json:"basePolicyId" yaml:"basePolicyId"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#dns FirewallPolicy#dns}
	Dns *FirewallPolicyDns `field:"optional" json:"dns" yaml:"dns"`
	// explicit_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#explicit_proxy FirewallPolicy#explicit_proxy}
	ExplicitProxy *FirewallPolicyExplicitProxy `field:"optional" json:"explicitProxy" yaml:"explicitProxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#id FirewallPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#identity FirewallPolicy#identity}
	Identity *FirewallPolicyIdentity `field:"optional" json:"identity" yaml:"identity"`
	// insights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#insights FirewallPolicy#insights}
	Insights *FirewallPolicyInsights `field:"optional" json:"insights" yaml:"insights"`
	// intrusion_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#intrusion_detection FirewallPolicy#intrusion_detection}
	IntrusionDetection *FirewallPolicyIntrusionDetection `field:"optional" json:"intrusionDetection" yaml:"intrusionDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#private_ip_ranges FirewallPolicy#private_ip_ranges}.
	PrivateIpRanges *[]*string `field:"optional" json:"privateIpRanges" yaml:"privateIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#sku FirewallPolicy#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#sql_redirect_allowed FirewallPolicy#sql_redirect_allowed}.
	SqlRedirectAllowed interface{} `field:"optional" json:"sqlRedirectAllowed" yaml:"sqlRedirectAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#tags FirewallPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_intelligence_allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#threat_intelligence_allowlist FirewallPolicy#threat_intelligence_allowlist}
	ThreatIntelligenceAllowlist *FirewallPolicyThreatIntelligenceAllowlistStruct `field:"optional" json:"threatIntelligenceAllowlist" yaml:"threatIntelligenceAllowlist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#threat_intelligence_mode FirewallPolicy#threat_intelligence_mode}.
	ThreatIntelligenceMode *string `field:"optional" json:"threatIntelligenceMode" yaml:"threatIntelligenceMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#timeouts FirewallPolicy#timeouts}
	Timeouts *FirewallPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tls_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/firewall_policy#tls_certificate FirewallPolicy#tls_certificate}
	TlsCertificate *FirewallPolicyTlsCertificate `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
}

