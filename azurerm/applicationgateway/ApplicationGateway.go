package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/application_gateway azurerm_application_gateway}.
type ApplicationGateway interface {
	cdktf.TerraformResource
	AuthenticationCertificate() ApplicationGatewayAuthenticationCertificateList
	AuthenticationCertificateInput() interface{}
	AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationOutputReference
	AutoscaleConfigurationInput() *ApplicationGatewayAutoscaleConfiguration
	BackendAddressPool() ApplicationGatewayBackendAddressPoolList
	BackendAddressPoolInput() interface{}
	BackendHttpSettings() ApplicationGatewayBackendHttpSettingsList
	BackendHttpSettingsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomErrorConfiguration() ApplicationGatewayCustomErrorConfigurationList
	CustomErrorConfigurationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableHttp2() interface{}
	SetEnableHttp2(val interface{})
	EnableHttp2Input() interface{}
	FipsEnabled() interface{}
	SetFipsEnabled(val interface{})
	FipsEnabledInput() interface{}
	FirewallPolicyId() *string
	SetFirewallPolicyId(val *string)
	FirewallPolicyIdInput() *string
	ForceFirewallPolicyAssociation() interface{}
	SetForceFirewallPolicyAssociation(val interface{})
	ForceFirewallPolicyAssociationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontendIpConfiguration() ApplicationGatewayFrontendIpConfigurationList
	FrontendIpConfigurationInput() interface{}
	FrontendPort() ApplicationGatewayFrontendPortList
	FrontendPortInput() interface{}
	GatewayIpConfiguration() ApplicationGatewayGatewayIpConfigurationList
	GatewayIpConfigurationInput() interface{}
	Global() ApplicationGatewayGlobalOutputReference
	GlobalInput() *ApplicationGatewayGlobal
	HttpListener() ApplicationGatewayHttpListenerList
	HttpListenerInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() ApplicationGatewayIdentityOutputReference
	IdentityInput() *ApplicationGatewayIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateEndpointConnection() ApplicationGatewayPrivateEndpointConnectionList
	PrivateLinkConfiguration() ApplicationGatewayPrivateLinkConfigurationList
	PrivateLinkConfigurationInput() interface{}
	Probe() ApplicationGatewayProbeList
	ProbeInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RedirectConfiguration() ApplicationGatewayRedirectConfigurationList
	RedirectConfigurationInput() interface{}
	RequestRoutingRule() ApplicationGatewayRequestRoutingRuleList
	RequestRoutingRuleInput() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RewriteRuleSet() ApplicationGatewayRewriteRuleSetList
	RewriteRuleSetInput() interface{}
	Sku() ApplicationGatewaySkuOutputReference
	SkuInput() *ApplicationGatewaySku
	SslCertificate() ApplicationGatewaySslCertificateList
	SslCertificateInput() interface{}
	SslPolicy() ApplicationGatewaySslPolicyOutputReference
	SslPolicyInput() *ApplicationGatewaySslPolicy
	SslProfile() ApplicationGatewaySslProfileList
	SslProfileInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApplicationGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustedClientCertificate() ApplicationGatewayTrustedClientCertificateList
	TrustedClientCertificateInput() interface{}
	TrustedRootCertificate() ApplicationGatewayTrustedRootCertificateList
	TrustedRootCertificateInput() interface{}
	UrlPathMap() ApplicationGatewayUrlPathMapList
	UrlPathMapInput() interface{}
	WafConfiguration() ApplicationGatewayWafConfigurationOutputReference
	WafConfigurationInput() *ApplicationGatewayWafConfiguration
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAuthenticationCertificate(value interface{})
	PutAutoscaleConfiguration(value *ApplicationGatewayAutoscaleConfiguration)
	PutBackendAddressPool(value interface{})
	PutBackendHttpSettings(value interface{})
	PutCustomErrorConfiguration(value interface{})
	PutFrontendIpConfiguration(value interface{})
	PutFrontendPort(value interface{})
	PutGatewayIpConfiguration(value interface{})
	PutGlobal(value *ApplicationGatewayGlobal)
	PutHttpListener(value interface{})
	PutIdentity(value *ApplicationGatewayIdentity)
	PutPrivateLinkConfiguration(value interface{})
	PutProbe(value interface{})
	PutRedirectConfiguration(value interface{})
	PutRequestRoutingRule(value interface{})
	PutRewriteRuleSet(value interface{})
	PutSku(value *ApplicationGatewaySku)
	PutSslCertificate(value interface{})
	PutSslPolicy(value *ApplicationGatewaySslPolicy)
	PutSslProfile(value interface{})
	PutTimeouts(value *ApplicationGatewayTimeouts)
	PutTrustedClientCertificate(value interface{})
	PutTrustedRootCertificate(value interface{})
	PutUrlPathMap(value interface{})
	PutWafConfiguration(value *ApplicationGatewayWafConfiguration)
	ResetAuthenticationCertificate()
	ResetAutoscaleConfiguration()
	ResetCustomErrorConfiguration()
	ResetEnableHttp2()
	ResetFipsEnabled()
	ResetFirewallPolicyId()
	ResetForceFirewallPolicyAssociation()
	ResetGlobal()
	ResetId()
	ResetIdentity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateLinkConfiguration()
	ResetProbe()
	ResetRedirectConfiguration()
	ResetRewriteRuleSet()
	ResetSslCertificate()
	ResetSslPolicy()
	ResetSslProfile()
	ResetTags()
	ResetTimeouts()
	ResetTrustedClientCertificate()
	ResetTrustedRootCertificate()
	ResetUrlPathMap()
	ResetWafConfiguration()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApplicationGateway
type jsiiProxy_ApplicationGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationGateway) AuthenticationCertificate() ApplicationGatewayAuthenticationCertificateList {
	var returns ApplicationGatewayAuthenticationCertificateList
	_jsii_.Get(
		j,
		"authenticationCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) AuthenticationCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationOutputReference {
	var returns ApplicationGatewayAutoscaleConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoscaleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) AutoscaleConfigurationInput() *ApplicationGatewayAutoscaleConfiguration {
	var returns *ApplicationGatewayAutoscaleConfiguration
	_jsii_.Get(
		j,
		"autoscaleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) BackendAddressPool() ApplicationGatewayBackendAddressPoolList {
	var returns ApplicationGatewayBackendAddressPoolList
	_jsii_.Get(
		j,
		"backendAddressPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) BackendAddressPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendAddressPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) BackendHttpSettings() ApplicationGatewayBackendHttpSettingsList {
	var returns ApplicationGatewayBackendHttpSettingsList
	_jsii_.Get(
		j,
		"backendHttpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) BackendHttpSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendHttpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) CustomErrorConfiguration() ApplicationGatewayCustomErrorConfigurationList {
	var returns ApplicationGatewayCustomErrorConfigurationList
	_jsii_.Get(
		j,
		"customErrorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) CustomErrorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FipsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FipsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FirewallPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FirewallPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ForceFirewallPolicyAssociation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceFirewallPolicyAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ForceFirewallPolicyAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceFirewallPolicyAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FrontendIpConfiguration() ApplicationGatewayFrontendIpConfigurationList {
	var returns ApplicationGatewayFrontendIpConfigurationList
	_jsii_.Get(
		j,
		"frontendIpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FrontendIpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendIpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FrontendPort() ApplicationGatewayFrontendPortList {
	var returns ApplicationGatewayFrontendPortList
	_jsii_.Get(
		j,
		"frontendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) FrontendPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frontendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) GatewayIpConfiguration() ApplicationGatewayGatewayIpConfigurationList {
	var returns ApplicationGatewayGatewayIpConfigurationList
	_jsii_.Get(
		j,
		"gatewayIpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) GatewayIpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayIpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Global() ApplicationGatewayGlobalOutputReference {
	var returns ApplicationGatewayGlobalOutputReference
	_jsii_.Get(
		j,
		"global",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) GlobalInput() *ApplicationGatewayGlobal {
	var returns *ApplicationGatewayGlobal
	_jsii_.Get(
		j,
		"globalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) HttpListener() ApplicationGatewayHttpListenerList {
	var returns ApplicationGatewayHttpListenerList
	_jsii_.Get(
		j,
		"httpListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) HttpListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Identity() ApplicationGatewayIdentityOutputReference {
	var returns ApplicationGatewayIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) IdentityInput() *ApplicationGatewayIdentity {
	var returns *ApplicationGatewayIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) PrivateEndpointConnection() ApplicationGatewayPrivateEndpointConnectionList {
	var returns ApplicationGatewayPrivateEndpointConnectionList
	_jsii_.Get(
		j,
		"privateEndpointConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) PrivateLinkConfiguration() ApplicationGatewayPrivateLinkConfigurationList {
	var returns ApplicationGatewayPrivateLinkConfigurationList
	_jsii_.Get(
		j,
		"privateLinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) PrivateLinkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Probe() ApplicationGatewayProbeList {
	var returns ApplicationGatewayProbeList
	_jsii_.Get(
		j,
		"probe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ProbeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"probeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RedirectConfiguration() ApplicationGatewayRedirectConfigurationList {
	var returns ApplicationGatewayRedirectConfigurationList
	_jsii_.Get(
		j,
		"redirectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RedirectConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RequestRoutingRule() ApplicationGatewayRequestRoutingRuleList {
	var returns ApplicationGatewayRequestRoutingRuleList
	_jsii_.Get(
		j,
		"requestRoutingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RequestRoutingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestRoutingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RewriteRuleSet() ApplicationGatewayRewriteRuleSetList {
	var returns ApplicationGatewayRewriteRuleSetList
	_jsii_.Get(
		j,
		"rewriteRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) RewriteRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rewriteRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Sku() ApplicationGatewaySkuOutputReference {
	var returns ApplicationGatewaySkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SkuInput() *ApplicationGatewaySku {
	var returns *ApplicationGatewaySku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslCertificate() ApplicationGatewaySslCertificateList {
	var returns ApplicationGatewaySslCertificateList
	_jsii_.Get(
		j,
		"sslCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslPolicy() ApplicationGatewaySslPolicyOutputReference {
	var returns ApplicationGatewaySslPolicyOutputReference
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslPolicyInput() *ApplicationGatewaySslPolicy {
	var returns *ApplicationGatewaySslPolicy
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslProfile() ApplicationGatewaySslProfileList {
	var returns ApplicationGatewaySslProfileList
	_jsii_.Get(
		j,
		"sslProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) SslProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Timeouts() ApplicationGatewayTimeoutsOutputReference {
	var returns ApplicationGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TrustedClientCertificate() ApplicationGatewayTrustedClientCertificateList {
	var returns ApplicationGatewayTrustedClientCertificateList
	_jsii_.Get(
		j,
		"trustedClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TrustedClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TrustedRootCertificate() ApplicationGatewayTrustedRootCertificateList {
	var returns ApplicationGatewayTrustedRootCertificateList
	_jsii_.Get(
		j,
		"trustedRootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) TrustedRootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedRootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) UrlPathMap() ApplicationGatewayUrlPathMapList {
	var returns ApplicationGatewayUrlPathMapList
	_jsii_.Get(
		j,
		"urlPathMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) UrlPathMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlPathMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) WafConfiguration() ApplicationGatewayWafConfigurationOutputReference {
	var returns ApplicationGatewayWafConfigurationOutputReference
	_jsii_.Get(
		j,
		"wafConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) WafConfigurationInput() *ApplicationGatewayWafConfiguration {
	var returns *ApplicationGatewayWafConfiguration
	_jsii_.Get(
		j,
		"wafConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGateway) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/application_gateway azurerm_application_gateway} Resource.
func NewApplicationGateway(scope constructs.Construct, id *string, config *ApplicationGatewayConfig) ApplicationGateway {
	_init_.Initialize()

	if err := validateNewApplicationGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGateway{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/application_gateway azurerm_application_gateway} Resource.
func NewApplicationGateway_Override(a ApplicationGateway, scope constructs.Construct, id *string, config *ApplicationGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetEnableHttp2(val interface{}) {
	if err := j.validateSetEnableHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetFipsEnabled(val interface{}) {
	if err := j.validateSetFipsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fipsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetFirewallPolicyId(val *string) {
	if err := j.validateSetFirewallPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyId",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetForceFirewallPolicyAssociation(val interface{}) {
	if err := j.validateSetForceFirewallPolicyAssociationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceFirewallPolicyAssociation",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApplicationGateway)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApplicationGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationGateway) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutAuthenticationCertificate(value interface{}) {
	if err := a.validatePutAuthenticationCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutAutoscaleConfiguration(value *ApplicationGatewayAutoscaleConfiguration) {
	if err := a.validatePutAutoscaleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoscaleConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutBackendAddressPool(value interface{}) {
	if err := a.validatePutBackendAddressPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendAddressPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutBackendHttpSettings(value interface{}) {
	if err := a.validatePutBackendHttpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendHttpSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutCustomErrorConfiguration(value interface{}) {
	if err := a.validatePutCustomErrorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutFrontendIpConfiguration(value interface{}) {
	if err := a.validatePutFrontendIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendIpConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutFrontendPort(value interface{}) {
	if err := a.validatePutFrontendPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendPort",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutGatewayIpConfiguration(value interface{}) {
	if err := a.validatePutGatewayIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGatewayIpConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutGlobal(value *ApplicationGatewayGlobal) {
	if err := a.validatePutGlobalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutHttpListener(value interface{}) {
	if err := a.validatePutHttpListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpListener",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutIdentity(value *ApplicationGatewayIdentity) {
	if err := a.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutPrivateLinkConfiguration(value interface{}) {
	if err := a.validatePutPrivateLinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrivateLinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutProbe(value interface{}) {
	if err := a.validatePutProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProbe",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutRedirectConfiguration(value interface{}) {
	if err := a.validatePutRedirectConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedirectConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutRequestRoutingRule(value interface{}) {
	if err := a.validatePutRequestRoutingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestRoutingRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutRewriteRuleSet(value interface{}) {
	if err := a.validatePutRewriteRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRewriteRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutSku(value *ApplicationGatewaySku) {
	if err := a.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSku",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutSslCertificate(value interface{}) {
	if err := a.validatePutSslCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSslCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutSslPolicy(value *ApplicationGatewaySslPolicy) {
	if err := a.validatePutSslPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSslPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutSslProfile(value interface{}) {
	if err := a.validatePutSslProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSslProfile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutTimeouts(value *ApplicationGatewayTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutTrustedClientCertificate(value interface{}) {
	if err := a.validatePutTrustedClientCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrustedClientCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutTrustedRootCertificate(value interface{}) {
	if err := a.validatePutTrustedRootCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrustedRootCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutUrlPathMap(value interface{}) {
	if err := a.validatePutUrlPathMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUrlPathMap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) PutWafConfiguration(value *ApplicationGatewayWafConfiguration) {
	if err := a.validatePutWafConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWafConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetAuthenticationCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetAutoscaleConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoscaleConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetCustomErrorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetFipsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFipsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetFirewallPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetForceFirewallPolicyAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetForceFirewallPolicyAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetGlobal() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetPrivateLinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateLinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetProbe() {
	_jsii_.InvokeVoid(
		a,
		"resetProbe",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetRedirectConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetRewriteRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetRewriteRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetSslCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetSslProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetSslProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetTrustedClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetTrustedRootCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedRootCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetUrlPathMap() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlPathMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetWafConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetWafConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) ResetZones() {
	_jsii_.InvokeVoid(
		a,
		"resetZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

