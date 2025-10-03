// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermapplicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermapplicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/data-sources/application_gateway azurerm_application_gateway}.
type DataAzurermApplicationGateway interface {
	cdktf.TerraformDataSource
	AuthenticationCertificate() DataAzurermApplicationGatewayAuthenticationCertificateList
	AutoscaleConfiguration() DataAzurermApplicationGatewayAutoscaleConfigurationList
	BackendAddressPool() DataAzurermApplicationGatewayBackendAddressPoolList
	BackendHttpSettings() DataAzurermApplicationGatewayBackendHttpSettingsList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomErrorConfiguration() DataAzurermApplicationGatewayCustomErrorConfigurationList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FipsEnabled() cdktf.IResolvable
	FirewallPolicyId() *string
	ForceFirewallPolicyAssociation() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontendIpConfiguration() DataAzurermApplicationGatewayFrontendIpConfigurationList
	FrontendPort() DataAzurermApplicationGatewayFrontendPortList
	GatewayIpConfiguration() DataAzurermApplicationGatewayGatewayIpConfigurationList
	Global() DataAzurermApplicationGatewayGlobalList
	Http2Enabled() cdktf.IResolvable
	HttpListener() DataAzurermApplicationGatewayHttpListenerList
	Id() *string
	SetId(val *string)
	Identity() DataAzurermApplicationGatewayIdentityList
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateEndpointConnection() DataAzurermApplicationGatewayPrivateEndpointConnectionList
	PrivateLinkConfiguration() DataAzurermApplicationGatewayPrivateLinkConfigurationList
	Probe() DataAzurermApplicationGatewayProbeList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RedirectConfiguration() DataAzurermApplicationGatewayRedirectConfigurationList
	RequestRoutingRule() DataAzurermApplicationGatewayRequestRoutingRuleList
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RewriteRuleSet() DataAzurermApplicationGatewayRewriteRuleSetList
	Sku() DataAzurermApplicationGatewaySkuList
	SslCertificate() DataAzurermApplicationGatewaySslCertificateList
	SslPolicy() DataAzurermApplicationGatewaySslPolicyList
	SslProfile() DataAzurermApplicationGatewaySslProfileList
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermApplicationGatewayTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustedClientCertificate() DataAzurermApplicationGatewayTrustedClientCertificateList
	TrustedRootCertificate() DataAzurermApplicationGatewayTrustedRootCertificateList
	UrlPathMap() DataAzurermApplicationGatewayUrlPathMapList
	WafConfiguration() DataAzurermApplicationGatewayWafConfigurationList
	Zones() *[]*string
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
	PutTimeouts(value *DataAzurermApplicationGatewayTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzurermApplicationGateway
type jsiiProxy_DataAzurermApplicationGateway struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermApplicationGateway) AuthenticationCertificate() DataAzurermApplicationGatewayAuthenticationCertificateList {
	var returns DataAzurermApplicationGatewayAuthenticationCertificateList
	_jsii_.Get(
		j,
		"authenticationCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) AutoscaleConfiguration() DataAzurermApplicationGatewayAutoscaleConfigurationList {
	var returns DataAzurermApplicationGatewayAutoscaleConfigurationList
	_jsii_.Get(
		j,
		"autoscaleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) BackendAddressPool() DataAzurermApplicationGatewayBackendAddressPoolList {
	var returns DataAzurermApplicationGatewayBackendAddressPoolList
	_jsii_.Get(
		j,
		"backendAddressPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) BackendHttpSettings() DataAzurermApplicationGatewayBackendHttpSettingsList {
	var returns DataAzurermApplicationGatewayBackendHttpSettingsList
	_jsii_.Get(
		j,
		"backendHttpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) CustomErrorConfiguration() DataAzurermApplicationGatewayCustomErrorConfigurationList {
	var returns DataAzurermApplicationGatewayCustomErrorConfigurationList
	_jsii_.Get(
		j,
		"customErrorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) FipsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) FirewallPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) ForceFirewallPolicyAssociation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"forceFirewallPolicyAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) FrontendIpConfiguration() DataAzurermApplicationGatewayFrontendIpConfigurationList {
	var returns DataAzurermApplicationGatewayFrontendIpConfigurationList
	_jsii_.Get(
		j,
		"frontendIpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) FrontendPort() DataAzurermApplicationGatewayFrontendPortList {
	var returns DataAzurermApplicationGatewayFrontendPortList
	_jsii_.Get(
		j,
		"frontendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) GatewayIpConfiguration() DataAzurermApplicationGatewayGatewayIpConfigurationList {
	var returns DataAzurermApplicationGatewayGatewayIpConfigurationList
	_jsii_.Get(
		j,
		"gatewayIpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Global() DataAzurermApplicationGatewayGlobalList {
	var returns DataAzurermApplicationGatewayGlobalList
	_jsii_.Get(
		j,
		"global",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Http2Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) HttpListener() DataAzurermApplicationGatewayHttpListenerList {
	var returns DataAzurermApplicationGatewayHttpListenerList
	_jsii_.Get(
		j,
		"httpListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Identity() DataAzurermApplicationGatewayIdentityList {
	var returns DataAzurermApplicationGatewayIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) PrivateEndpointConnection() DataAzurermApplicationGatewayPrivateEndpointConnectionList {
	var returns DataAzurermApplicationGatewayPrivateEndpointConnectionList
	_jsii_.Get(
		j,
		"privateEndpointConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) PrivateLinkConfiguration() DataAzurermApplicationGatewayPrivateLinkConfigurationList {
	var returns DataAzurermApplicationGatewayPrivateLinkConfigurationList
	_jsii_.Get(
		j,
		"privateLinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Probe() DataAzurermApplicationGatewayProbeList {
	var returns DataAzurermApplicationGatewayProbeList
	_jsii_.Get(
		j,
		"probe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) RedirectConfiguration() DataAzurermApplicationGatewayRedirectConfigurationList {
	var returns DataAzurermApplicationGatewayRedirectConfigurationList
	_jsii_.Get(
		j,
		"redirectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) RequestRoutingRule() DataAzurermApplicationGatewayRequestRoutingRuleList {
	var returns DataAzurermApplicationGatewayRequestRoutingRuleList
	_jsii_.Get(
		j,
		"requestRoutingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) RewriteRuleSet() DataAzurermApplicationGatewayRewriteRuleSetList {
	var returns DataAzurermApplicationGatewayRewriteRuleSetList
	_jsii_.Get(
		j,
		"rewriteRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Sku() DataAzurermApplicationGatewaySkuList {
	var returns DataAzurermApplicationGatewaySkuList
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) SslCertificate() DataAzurermApplicationGatewaySslCertificateList {
	var returns DataAzurermApplicationGatewaySslCertificateList
	_jsii_.Get(
		j,
		"sslCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) SslPolicy() DataAzurermApplicationGatewaySslPolicyList {
	var returns DataAzurermApplicationGatewaySslPolicyList
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) SslProfile() DataAzurermApplicationGatewaySslProfileList {
	var returns DataAzurermApplicationGatewaySslProfileList
	_jsii_.Get(
		j,
		"sslProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Timeouts() DataAzurermApplicationGatewayTimeoutsOutputReference {
	var returns DataAzurermApplicationGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TrustedClientCertificate() DataAzurermApplicationGatewayTrustedClientCertificateList {
	var returns DataAzurermApplicationGatewayTrustedClientCertificateList
	_jsii_.Get(
		j,
		"trustedClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) TrustedRootCertificate() DataAzurermApplicationGatewayTrustedRootCertificateList {
	var returns DataAzurermApplicationGatewayTrustedRootCertificateList
	_jsii_.Get(
		j,
		"trustedRootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) UrlPathMap() DataAzurermApplicationGatewayUrlPathMapList {
	var returns DataAzurermApplicationGatewayUrlPathMapList
	_jsii_.Get(
		j,
		"urlPathMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) WafConfiguration() DataAzurermApplicationGatewayWafConfigurationList {
	var returns DataAzurermApplicationGatewayWafConfigurationList
	_jsii_.Get(
		j,
		"wafConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermApplicationGateway) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/data-sources/application_gateway azurerm_application_gateway} Data Source.
func NewDataAzurermApplicationGateway(scope constructs.Construct, id *string, config *DataAzurermApplicationGatewayConfig) DataAzurermApplicationGateway {
	_init_.Initialize()

	if err := validateNewDataAzurermApplicationGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermApplicationGateway{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/data-sources/application_gateway azurerm_application_gateway} Data Source.
func NewDataAzurermApplicationGateway_Override(d DataAzurermApplicationGateway, scope constructs.Construct, id *string, config *DataAzurermApplicationGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermApplicationGateway)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermApplicationGateway resource upon running "cdktf plan <stack-name>".
func DataAzurermApplicationGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermApplicationGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataAzurermApplicationGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermApplicationGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermApplicationGateway_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermApplicationGateway_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermApplicationGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermApplicationGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermApplicationGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermApplicationGateway.DataAzurermApplicationGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) PutTimeouts(value *DataAzurermApplicationGatewayTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermApplicationGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermApplicationGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

