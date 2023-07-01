package applicationgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayConfig struct {
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
	// backend_address_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#backend_address_pool ApplicationGateway#backend_address_pool}
	BackendAddressPool interface{} `field:"required" json:"backendAddressPool" yaml:"backendAddressPool"`
	// backend_http_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#backend_http_settings ApplicationGateway#backend_http_settings}
	BackendHttpSettings interface{} `field:"required" json:"backendHttpSettings" yaml:"backendHttpSettings"`
	// frontend_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#frontend_ip_configuration ApplicationGateway#frontend_ip_configuration}
	FrontendIpConfiguration interface{} `field:"required" json:"frontendIpConfiguration" yaml:"frontendIpConfiguration"`
	// frontend_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#frontend_port ApplicationGateway#frontend_port}
	FrontendPort interface{} `field:"required" json:"frontendPort" yaml:"frontendPort"`
	// gateway_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#gateway_ip_configuration ApplicationGateway#gateway_ip_configuration}
	GatewayIpConfiguration interface{} `field:"required" json:"gatewayIpConfiguration" yaml:"gatewayIpConfiguration"`
	// http_listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#http_listener ApplicationGateway#http_listener}
	HttpListener interface{} `field:"required" json:"httpListener" yaml:"httpListener"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#location ApplicationGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// request_routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#request_routing_rule ApplicationGateway#request_routing_rule}
	RequestRoutingRule interface{} `field:"required" json:"requestRoutingRule" yaml:"requestRoutingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#resource_group_name ApplicationGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#sku ApplicationGateway#sku}
	Sku *ApplicationGatewaySku `field:"required" json:"sku" yaml:"sku"`
	// authentication_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#authentication_certificate ApplicationGateway#authentication_certificate}
	AuthenticationCertificate interface{} `field:"optional" json:"authenticationCertificate" yaml:"authenticationCertificate"`
	// autoscale_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#autoscale_configuration ApplicationGateway#autoscale_configuration}
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfiguration `field:"optional" json:"autoscaleConfiguration" yaml:"autoscaleConfiguration"`
	// custom_error_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#custom_error_configuration ApplicationGateway#custom_error_configuration}
	CustomErrorConfiguration interface{} `field:"optional" json:"customErrorConfiguration" yaml:"customErrorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#enable_http2 ApplicationGateway#enable_http2}.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#fips_enabled ApplicationGateway#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#firewall_policy_id ApplicationGateway#firewall_policy_id}.
	FirewallPolicyId *string `field:"optional" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#force_firewall_policy_association ApplicationGateway#force_firewall_policy_association}.
	ForceFirewallPolicyAssociation interface{} `field:"optional" json:"forceFirewallPolicyAssociation" yaml:"forceFirewallPolicyAssociation"`
	// global block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#global ApplicationGateway#global}
	Global *ApplicationGatewayGlobal `field:"optional" json:"global" yaml:"global"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#id ApplicationGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#identity ApplicationGateway#identity}
	Identity *ApplicationGatewayIdentity `field:"optional" json:"identity" yaml:"identity"`
	// private_link_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#private_link_configuration ApplicationGateway#private_link_configuration}
	PrivateLinkConfiguration interface{} `field:"optional" json:"privateLinkConfiguration" yaml:"privateLinkConfiguration"`
	// probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#probe ApplicationGateway#probe}
	Probe interface{} `field:"optional" json:"probe" yaml:"probe"`
	// redirect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#redirect_configuration ApplicationGateway#redirect_configuration}
	RedirectConfiguration interface{} `field:"optional" json:"redirectConfiguration" yaml:"redirectConfiguration"`
	// rewrite_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#rewrite_rule_set ApplicationGateway#rewrite_rule_set}
	RewriteRuleSet interface{} `field:"optional" json:"rewriteRuleSet" yaml:"rewriteRuleSet"`
	// ssl_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#ssl_certificate ApplicationGateway#ssl_certificate}
	SslCertificate interface{} `field:"optional" json:"sslCertificate" yaml:"sslCertificate"`
	// ssl_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#ssl_policy ApplicationGateway#ssl_policy}
	SslPolicy *ApplicationGatewaySslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// ssl_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#ssl_profile ApplicationGateway#ssl_profile}
	SslProfile interface{} `field:"optional" json:"sslProfile" yaml:"sslProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#tags ApplicationGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#timeouts ApplicationGateway#timeouts}
	Timeouts *ApplicationGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trusted_client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#trusted_client_certificate ApplicationGateway#trusted_client_certificate}
	TrustedClientCertificate interface{} `field:"optional" json:"trustedClientCertificate" yaml:"trustedClientCertificate"`
	// trusted_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#trusted_root_certificate ApplicationGateway#trusted_root_certificate}
	TrustedRootCertificate interface{} `field:"optional" json:"trustedRootCertificate" yaml:"trustedRootCertificate"`
	// url_path_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#url_path_map ApplicationGateway#url_path_map}
	UrlPathMap interface{} `field:"optional" json:"urlPathMap" yaml:"urlPathMap"`
	// waf_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#waf_configuration ApplicationGateway#waf_configuration}
	WafConfiguration *ApplicationGatewayWafConfiguration `field:"optional" json:"wafConfiguration" yaml:"wafConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/application_gateway#zones ApplicationGateway#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

