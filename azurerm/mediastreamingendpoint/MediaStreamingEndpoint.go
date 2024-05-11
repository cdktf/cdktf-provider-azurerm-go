// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mediastreamingendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/media_streaming_endpoint azurerm_media_streaming_endpoint}.
type MediaStreamingEndpoint interface {
	cdktf.TerraformResource
	AccessControl() MediaStreamingEndpointAccessControlOutputReference
	AccessControlInput() *MediaStreamingEndpointAccessControl
	AutoStartEnabled() interface{}
	SetAutoStartEnabled(val interface{})
	AutoStartEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnEnabled() interface{}
	SetCdnEnabled(val interface{})
	CdnEnabledInput() interface{}
	CdnProfile() *string
	SetCdnProfile(val *string)
	CdnProfileInput() *string
	CdnProvider() *string
	SetCdnProvider(val *string)
	CdnProviderInput() *string
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
	CrossSiteAccessPolicy() MediaStreamingEndpointCrossSiteAccessPolicyOutputReference
	CrossSiteAccessPolicyInput() *MediaStreamingEndpointCrossSiteAccessPolicy
	CustomHostNames() *[]*string
	SetCustomHostNames(val *[]*string)
	CustomHostNamesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxCacheAgeSeconds() *float64
	SetMaxCacheAgeSeconds(val *float64)
	MaxCacheAgeSecondsInput() *float64
	MediaServicesAccountName() *string
	SetMediaServicesAccountName(val *string)
	MediaServicesAccountNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScaleUnits() *float64
	SetScaleUnits(val *float64)
	ScaleUnitsInput() *float64
	Sku() MediaStreamingEndpointSkuList
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaStreamingEndpointTimeoutsOutputReference
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
	PutAccessControl(value *MediaStreamingEndpointAccessControl)
	PutCrossSiteAccessPolicy(value *MediaStreamingEndpointCrossSiteAccessPolicy)
	PutTimeouts(value *MediaStreamingEndpointTimeouts)
	ResetAccessControl()
	ResetAutoStartEnabled()
	ResetCdnEnabled()
	ResetCdnProfile()
	ResetCdnProvider()
	ResetCrossSiteAccessPolicy()
	ResetCustomHostNames()
	ResetDescription()
	ResetId()
	ResetMaxCacheAgeSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for MediaStreamingEndpoint
type jsiiProxy_MediaStreamingEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaStreamingEndpoint) AccessControl() MediaStreamingEndpointAccessControlOutputReference {
	var returns MediaStreamingEndpointAccessControlOutputReference
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) AccessControlInput() *MediaStreamingEndpointAccessControl {
	var returns *MediaStreamingEndpointAccessControl
	_jsii_.Get(
		j,
		"accessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) AutoStartEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStartEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) AutoStartEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStartEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdnEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdnEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CdnProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CrossSiteAccessPolicy() MediaStreamingEndpointCrossSiteAccessPolicyOutputReference {
	var returns MediaStreamingEndpointCrossSiteAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"crossSiteAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CrossSiteAccessPolicyInput() *MediaStreamingEndpointCrossSiteAccessPolicy {
	var returns *MediaStreamingEndpointCrossSiteAccessPolicy
	_jsii_.Get(
		j,
		"crossSiteAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CustomHostNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customHostNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) CustomHostNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customHostNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) MaxCacheAgeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCacheAgeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) MaxCacheAgeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCacheAgeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) MediaServicesAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) MediaServicesAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ScaleUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) ScaleUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Sku() MediaStreamingEndpointSkuList {
	var returns MediaStreamingEndpointSkuList
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) Timeouts() MediaStreamingEndpointTimeoutsOutputReference {
	var returns MediaStreamingEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/media_streaming_endpoint azurerm_media_streaming_endpoint} Resource.
func NewMediaStreamingEndpoint(scope constructs.Construct, id *string, config *MediaStreamingEndpointConfig) MediaStreamingEndpoint {
	_init_.Initialize()

	if err := validateNewMediaStreamingEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaStreamingEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/media_streaming_endpoint azurerm_media_streaming_endpoint} Resource.
func NewMediaStreamingEndpoint_Override(m MediaStreamingEndpoint, scope constructs.Construct, id *string, config *MediaStreamingEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetAutoStartEnabled(val interface{}) {
	if err := j.validateSetAutoStartEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStartEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetCdnEnabled(val interface{}) {
	if err := j.validateSetCdnEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetCdnProfile(val *string) {
	if err := j.validateSetCdnProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnProfile",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetCdnProvider(val *string) {
	if err := j.validateSetCdnProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnProvider",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetCustomHostNames(val *[]*string) {
	if err := j.validateSetCustomHostNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHostNames",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetMaxCacheAgeSeconds(val *float64) {
	if err := j.validateSetMaxCacheAgeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCacheAgeSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetMediaServicesAccountName(val *string) {
	if err := j.validateSetMediaServicesAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaServicesAccountName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetScaleUnits(val *float64) {
	if err := j.validateSetScaleUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleUnits",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a MediaStreamingEndpoint resource upon running "cdktf plan <stack-name>".
func MediaStreamingEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediaStreamingEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
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
func MediaStreamingEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaStreamingEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaStreamingEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaStreamingEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaStreamingEndpoint.MediaStreamingEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) PutAccessControl(value *MediaStreamingEndpointAccessControl) {
	if err := m.validatePutAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAccessControl",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) PutCrossSiteAccessPolicy(value *MediaStreamingEndpointCrossSiteAccessPolicy) {
	if err := m.validatePutCrossSiteAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCrossSiteAccessPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) PutTimeouts(value *MediaStreamingEndpointTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetAccessControl() {
	_jsii_.InvokeVoid(
		m,
		"resetAccessControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetAutoStartEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoStartEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetCdnEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetCdnEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetCdnProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetCdnProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetCdnProvider() {
	_jsii_.InvokeVoid(
		m,
		"resetCdnProvider",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetCrossSiteAccessPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCrossSiteAccessPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetCustomHostNames() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomHostNames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetMaxCacheAgeSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxCacheAgeSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

