// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/cdnfrontdoorroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/cdn_frontdoor_route azurerm_cdn_frontdoor_route}.
type CdnFrontdoorRoute interface {
	cdktf.TerraformResource
	Cache() CdnFrontdoorRouteCacheOutputReference
	CacheInput() *CdnFrontdoorRouteCache
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnFrontdoorCustomDomainIds() *[]*string
	SetCdnFrontdoorCustomDomainIds(val *[]*string)
	CdnFrontdoorCustomDomainIdsInput() *[]*string
	CdnFrontdoorEndpointId() *string
	SetCdnFrontdoorEndpointId(val *string)
	CdnFrontdoorEndpointIdInput() *string
	CdnFrontdoorOriginGroupId() *string
	SetCdnFrontdoorOriginGroupId(val *string)
	CdnFrontdoorOriginGroupIdInput() *string
	CdnFrontdoorOriginIds() *[]*string
	SetCdnFrontdoorOriginIds(val *[]*string)
	CdnFrontdoorOriginIdsInput() *[]*string
	CdnFrontdoorOriginPath() *string
	SetCdnFrontdoorOriginPath(val *string)
	CdnFrontdoorOriginPathInput() *string
	CdnFrontdoorRuleSetIds() *[]*string
	SetCdnFrontdoorRuleSetIds(val *[]*string)
	CdnFrontdoorRuleSetIdsInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardingProtocol() *string
	SetForwardingProtocol(val *string)
	ForwardingProtocolInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsRedirectEnabled() interface{}
	SetHttpsRedirectEnabled(val interface{})
	HttpsRedirectEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkToDefaultDomain() interface{}
	SetLinkToDefaultDomain(val interface{})
	LinkToDefaultDomainInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PatternsToMatch() *[]*string
	SetPatternsToMatch(val *[]*string)
	PatternsToMatchInput() *[]*string
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
	SupportedProtocols() *[]*string
	SetSupportedProtocols(val *[]*string)
	SupportedProtocolsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnFrontdoorRouteTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCache(value *CdnFrontdoorRouteCache)
	PutTimeouts(value *CdnFrontdoorRouteTimeouts)
	ResetCache()
	ResetCdnFrontdoorCustomDomainIds()
	ResetCdnFrontdoorOriginPath()
	ResetCdnFrontdoorRuleSetIds()
	ResetEnabled()
	ResetForwardingProtocol()
	ResetHttpsRedirectEnabled()
	ResetId()
	ResetLinkToDefaultDomain()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for CdnFrontdoorRoute
type jsiiProxy_CdnFrontdoorRoute struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnFrontdoorRoute) Cache() CdnFrontdoorRouteCacheOutputReference {
	var returns CdnFrontdoorRouteCacheOutputReference
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CacheInput() *CdnFrontdoorRouteCache {
	var returns *CdnFrontdoorRouteCache
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorCustomDomainIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorCustomDomainIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorCustomDomainIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorCustomDomainIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorOriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorRuleSetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorRuleSetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) CdnFrontdoorRuleSetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cdnFrontdoorRuleSetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForwardingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) ForwardingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) HttpsRedirectEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) HttpsRedirectEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) LinkToDefaultDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkToDefaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) LinkToDefaultDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkToDefaultDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) PatternsToMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) PatternsToMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) SupportedProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) SupportedProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) Timeouts() CdnFrontdoorRouteTimeoutsOutputReference {
	var returns CdnFrontdoorRouteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/cdn_frontdoor_route azurerm_cdn_frontdoor_route} Resource.
func NewCdnFrontdoorRoute(scope constructs.Construct, id *string, config *CdnFrontdoorRouteConfig) CdnFrontdoorRoute {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRouteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRoute{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.6.0/docs/resources/cdn_frontdoor_route azurerm_cdn_frontdoor_route} Resource.
func NewCdnFrontdoorRoute_Override(c CdnFrontdoorRoute, scope constructs.Construct, id *string, config *CdnFrontdoorRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorCustomDomainIds(val *[]*string) {
	if err := j.validateSetCdnFrontdoorCustomDomainIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorCustomDomainIds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorEndpointId(val *string) {
	if err := j.validateSetCdnFrontdoorEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorEndpointId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorOriginGroupId(val *string) {
	if err := j.validateSetCdnFrontdoorOriginGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginGroupId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorOriginIds(val *[]*string) {
	if err := j.validateSetCdnFrontdoorOriginIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginIds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorOriginPath(val *string) {
	if err := j.validateSetCdnFrontdoorOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginPath",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCdnFrontdoorRuleSetIds(val *[]*string) {
	if err := j.validateSetCdnFrontdoorRuleSetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorRuleSetIds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetForwardingProtocol(val *string) {
	if err := j.validateSetForwardingProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardingProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetHttpsRedirectEnabled(val interface{}) {
	if err := j.validateSetHttpsRedirectEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsRedirectEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetLinkToDefaultDomain(val interface{}) {
	if err := j.validateSetLinkToDefaultDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkToDefaultDomain",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetPatternsToMatch(val *[]*string) {
	if err := j.validateSetPatternsToMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternsToMatch",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRoute)SetSupportedProtocols(val *[]*string) {
	if err := j.validateSetSupportedProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedProtocols",
		val,
	)
}

// Generates CDKTF code for importing a CdnFrontdoorRoute resource upon running "cdktf plan <stack-name>".
func CdnFrontdoorRoute_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCdnFrontdoorRoute_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
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
func CdnFrontdoorRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnFrontdoorRoute_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorRoute_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnFrontdoorRoute_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorRoute_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnFrontdoorRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) PutCache(value *CdnFrontdoorRouteCache) {
	if err := c.validatePutCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) PutTimeouts(value *CdnFrontdoorRouteTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCache() {
	_jsii_.InvokeVoid(
		c,
		"resetCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCdnFrontdoorCustomDomainIds() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnFrontdoorCustomDomainIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCdnFrontdoorOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnFrontdoorOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetCdnFrontdoorRuleSetIds() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnFrontdoorRuleSetIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetForwardingProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardingProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetHttpsRedirectEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsRedirectEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetLinkToDefaultDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetLinkToDefaultDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

