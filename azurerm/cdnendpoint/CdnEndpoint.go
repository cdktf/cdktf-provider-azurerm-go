// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/cdnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/cdn_endpoint azurerm_cdn_endpoint}.
type CdnEndpoint interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentTypesToCompress() *[]*string
	SetContentTypesToCompress(val *[]*string)
	ContentTypesToCompressInput() *[]*string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeliveryRule() CdnEndpointDeliveryRuleList
	DeliveryRuleInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoFilter() CdnEndpointGeoFilterList
	GeoFilterInput() interface{}
	GlobalDeliveryRule() CdnEndpointGlobalDeliveryRuleOutputReference
	GlobalDeliveryRuleInput() *CdnEndpointGlobalDeliveryRule
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsCompressionEnabled() interface{}
	SetIsCompressionEnabled(val interface{})
	IsCompressionEnabledInput() interface{}
	IsHttpAllowed() interface{}
	SetIsHttpAllowed(val interface{})
	IsHttpAllowedInput() interface{}
	IsHttpsAllowed() interface{}
	SetIsHttpsAllowed(val interface{})
	IsHttpsAllowedInput() interface{}
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
	OptimizationType() *string
	SetOptimizationType(val *string)
	OptimizationTypeInput() *string
	Origin() CdnEndpointOriginList
	OriginHostHeader() *string
	SetOriginHostHeader(val *string)
	OriginHostHeaderInput() *string
	OriginInput() interface{}
	OriginPath() *string
	SetOriginPath(val *string)
	OriginPathInput() *string
	ProbePath() *string
	SetProbePath(val *string)
	ProbePathInput() *string
	ProfileName() *string
	SetProfileName(val *string)
	ProfileNameInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuerystringCachingBehaviour() *string
	SetQuerystringCachingBehaviour(val *string)
	QuerystringCachingBehaviourInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnEndpointTimeoutsOutputReference
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
	PutDeliveryRule(value interface{})
	PutGeoFilter(value interface{})
	PutGlobalDeliveryRule(value *CdnEndpointGlobalDeliveryRule)
	PutOrigin(value interface{})
	PutTimeouts(value *CdnEndpointTimeouts)
	ResetContentTypesToCompress()
	ResetDeliveryRule()
	ResetGeoFilter()
	ResetGlobalDeliveryRule()
	ResetId()
	ResetIsCompressionEnabled()
	ResetIsHttpAllowed()
	ResetIsHttpsAllowed()
	ResetOptimizationType()
	ResetOriginHostHeader()
	ResetOriginPath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProbePath()
	ResetQuerystringCachingBehaviour()
	ResetTags()
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

// The jsii proxy struct for CdnEndpoint
type jsiiProxy_CdnEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ContentTypesToCompress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ContentTypesToCompressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DeliveryRule() CdnEndpointDeliveryRuleList {
	var returns CdnEndpointDeliveryRuleList
	_jsii_.Get(
		j,
		"deliveryRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DeliveryRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GeoFilter() CdnEndpointGeoFilterList {
	var returns CdnEndpointGeoFilterList
	_jsii_.Get(
		j,
		"geoFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GeoFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GlobalDeliveryRule() CdnEndpointGlobalDeliveryRuleOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleOutputReference
	_jsii_.Get(
		j,
		"globalDeliveryRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) GlobalDeliveryRuleInput() *CdnEndpointGlobalDeliveryRule {
	var returns *CdnEndpointGlobalDeliveryRule
	_jsii_.Get(
		j,
		"globalDeliveryRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsCompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCompressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsCompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCompressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpsAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) IsHttpsAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHttpsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OptimizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OptimizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Origin() CdnEndpointOriginList {
	var returns CdnEndpointOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProbePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProbePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) QuerystringCachingBehaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querystringCachingBehaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) QuerystringCachingBehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querystringCachingBehaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) Timeouts() CdnEndpointTimeoutsOutputReference {
	var returns CdnEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/cdn_endpoint azurerm_cdn_endpoint} Resource.
func NewCdnEndpoint(scope constructs.Construct, id *string, config *CdnEndpointConfig) CdnEndpoint {
	_init_.Initialize()

	if err := validateNewCdnEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/cdn_endpoint azurerm_cdn_endpoint} Resource.
func NewCdnEndpoint_Override(c CdnEndpoint, scope constructs.Construct, id *string, config *CdnEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetContentTypesToCompress(val *[]*string) {
	if err := j.validateSetContentTypesToCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentTypesToCompress",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetIsCompressionEnabled(val interface{}) {
	if err := j.validateSetIsCompressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCompressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetIsHttpAllowed(val interface{}) {
	if err := j.validateSetIsHttpAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHttpAllowed",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetIsHttpsAllowed(val interface{}) {
	if err := j.validateSetIsHttpsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHttpsAllowed",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetOptimizationType(val *string) {
	if err := j.validateSetOptimizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationType",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetOriginHostHeader(val *string) {
	if err := j.validateSetOriginHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originHostHeader",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetProbePath(val *string) {
	if err := j.validateSetProbePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probePath",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetProfileName(val *string) {
	if err := j.validateSetProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileName",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetQuerystringCachingBehaviour(val *string) {
	if err := j.validateSetQuerystringCachingBehaviourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"querystringCachingBehaviour",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CdnEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a CdnEndpoint resource upon running "cdktf plan <stack-name>".
func CdnEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCdnEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
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
func CdnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnEndpoint) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CdnEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CdnEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpoint) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CdnEndpoint) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutDeliveryRule(value interface{}) {
	if err := c.validatePutDeliveryRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeliveryRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutGeoFilter(value interface{}) {
	if err := c.validatePutGeoFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGeoFilter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutGlobalDeliveryRule(value *CdnEndpointGlobalDeliveryRule) {
	if err := c.validatePutGlobalDeliveryRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGlobalDeliveryRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutOrigin(value interface{}) {
	if err := c.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) PutTimeouts(value *CdnEndpointTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetContentTypesToCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypesToCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetDeliveryRule() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetGeoFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetGeoFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetGlobalDeliveryRule() {
	_jsii_.InvokeVoid(
		c,
		"resetGlobalDeliveryRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsHttpAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetIsHttpAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetIsHttpsAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetIsHttpsAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOptimizationType() {
	_jsii_.InvokeVoid(
		c,
		"resetOptimizationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOriginHostHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginHostHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetProbePath() {
	_jsii_.InvokeVoid(
		c,
		"resetProbePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetQuerystringCachingBehaviour() {
	_jsii_.InvokeVoid(
		c,
		"resetQuerystringCachingBehaviour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

