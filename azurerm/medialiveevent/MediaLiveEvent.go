// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/medialiveevent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/media_live_event azurerm_media_live_event}.
type MediaLiveEvent interface {
	cdktf.TerraformResource
	AutoStartEnabled() interface{}
	SetAutoStartEnabled(val interface{})
	AutoStartEnabledInput() interface{}
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
	CrossSiteAccessPolicy() MediaLiveEventCrossSiteAccessPolicyOutputReference
	CrossSiteAccessPolicyInput() *MediaLiveEventCrossSiteAccessPolicy
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Encoding() MediaLiveEventEncodingOutputReference
	EncodingInput() *MediaLiveEventEncoding
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostnamePrefix() *string
	SetHostnamePrefix(val *string)
	HostnamePrefixInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Input() MediaLiveEventInputOutputReference
	InputInput() *MediaLiveEventInput
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MediaServicesAccountName() *string
	SetMediaServicesAccountName(val *string)
	MediaServicesAccountNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Preview() MediaLiveEventPreviewOutputReference
	PreviewInput() *MediaLiveEventPreview
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
	StreamOptions() *[]*string
	SetStreamOptions(val *[]*string)
	StreamOptionsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaLiveEventTimeoutsOutputReference
	TimeoutsInput() interface{}
	TranscriptionLanguages() *[]*string
	SetTranscriptionLanguages(val *[]*string)
	TranscriptionLanguagesInput() *[]*string
	UseStaticHostname() interface{}
	SetUseStaticHostname(val interface{})
	UseStaticHostnameInput() interface{}
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
	PutCrossSiteAccessPolicy(value *MediaLiveEventCrossSiteAccessPolicy)
	PutEncoding(value *MediaLiveEventEncoding)
	PutInput(value *MediaLiveEventInput)
	PutPreview(value *MediaLiveEventPreview)
	PutTimeouts(value *MediaLiveEventTimeouts)
	ResetAutoStartEnabled()
	ResetCrossSiteAccessPolicy()
	ResetDescription()
	ResetEncoding()
	ResetHostnamePrefix()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreview()
	ResetStreamOptions()
	ResetTags()
	ResetTimeouts()
	ResetTranscriptionLanguages()
	ResetUseStaticHostname()
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

// The jsii proxy struct for MediaLiveEvent
type jsiiProxy_MediaLiveEvent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaLiveEvent) AutoStartEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStartEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) AutoStartEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStartEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) CrossSiteAccessPolicy() MediaLiveEventCrossSiteAccessPolicyOutputReference {
	var returns MediaLiveEventCrossSiteAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"crossSiteAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) CrossSiteAccessPolicyInput() *MediaLiveEventCrossSiteAccessPolicy {
	var returns *MediaLiveEventCrossSiteAccessPolicy
	_jsii_.Get(
		j,
		"crossSiteAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Encoding() MediaLiveEventEncodingOutputReference {
	var returns MediaLiveEventEncodingOutputReference
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) EncodingInput() *MediaLiveEventEncoding {
	var returns *MediaLiveEventEncoding
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Input() MediaLiveEventInputOutputReference {
	var returns MediaLiveEventInputOutputReference
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) InputInput() *MediaLiveEventInput {
	var returns *MediaLiveEventInput
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) MediaServicesAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) MediaServicesAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Preview() MediaLiveEventPreviewOutputReference {
	var returns MediaLiveEventPreviewOutputReference
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) PreviewInput() *MediaLiveEventPreview {
	var returns *MediaLiveEventPreview
	_jsii_.Get(
		j,
		"previewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) StreamOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) StreamOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) Timeouts() MediaLiveEventTimeoutsOutputReference {
	var returns MediaLiveEventTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TranscriptionLanguages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transcriptionLanguages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) TranscriptionLanguagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transcriptionLanguagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) UseStaticHostname() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStaticHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEvent) UseStaticHostnameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStaticHostnameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/media_live_event azurerm_media_live_event} Resource.
func NewMediaLiveEvent(scope constructs.Construct, id *string, config *MediaLiveEventConfig) MediaLiveEvent {
	_init_.Initialize()

	if err := validateNewMediaLiveEventParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEvent{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/media_live_event azurerm_media_live_event} Resource.
func NewMediaLiveEvent_Override(m MediaLiveEvent, scope constructs.Construct, id *string, config *MediaLiveEventConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetAutoStartEnabled(val interface{}) {
	if err := j.validateSetAutoStartEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStartEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetMediaServicesAccountName(val *string) {
	if err := j.validateSetMediaServicesAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaServicesAccountName",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetStreamOptions(val *[]*string) {
	if err := j.validateSetStreamOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamOptions",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetTranscriptionLanguages(val *[]*string) {
	if err := j.validateSetTranscriptionLanguagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transcriptionLanguages",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEvent)SetUseStaticHostname(val interface{}) {
	if err := j.validateSetUseStaticHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStaticHostname",
		val,
	)
}

// Generates CDKTF code for importing a MediaLiveEvent resource upon running "cdktf plan <stack-name>".
func MediaLiveEvent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediaLiveEvent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
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
func MediaLiveEvent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEvent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaLiveEvent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEvent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaLiveEvent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEvent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaLiveEvent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEvent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaLiveEvent) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediaLiveEvent) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaLiveEvent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEvent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaLiveEvent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaLiveEvent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaLiveEvent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaLiveEvent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaLiveEvent) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaLiveEvent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaLiveEvent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediaLiveEvent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEvent) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaLiveEvent) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediaLiveEvent) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediaLiveEvent) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaLiveEvent) PutCrossSiteAccessPolicy(value *MediaLiveEventCrossSiteAccessPolicy) {
	if err := m.validatePutCrossSiteAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCrossSiteAccessPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) PutEncoding(value *MediaLiveEventEncoding) {
	if err := m.validatePutEncodingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncoding",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) PutInput(value *MediaLiveEventInput) {
	if err := m.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInput",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) PutPreview(value *MediaLiveEventPreview) {
	if err := m.validatePutPreviewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPreview",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) PutTimeouts(value *MediaLiveEventTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetAutoStartEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoStartEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetCrossSiteAccessPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCrossSiteAccessPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetEncoding() {
	_jsii_.InvokeVoid(
		m,
		"resetEncoding",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetHostnamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetHostnamePrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetPreview() {
	_jsii_.InvokeVoid(
		m,
		"resetPreview",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetStreamOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetTranscriptionLanguages() {
	_jsii_.InvokeVoid(
		m,
		"resetTranscriptionLanguages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) ResetUseStaticHostname() {
	_jsii_.InvokeVoid(
		m,
		"resetUseStaticHostname",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEvent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEvent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

