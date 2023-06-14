package medialiveeventoutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/medialiveeventoutput/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_live_event_output azurerm_media_live_event_output}.
type MediaLiveEventOutput interface {
	cdktf.TerraformResource
	ArchiveWindowDuration() *string
	SetArchiveWindowDuration(val *string)
	ArchiveWindowDurationInput() *string
	AssetName() *string
	SetAssetName(val *string)
	AssetNameInput() *string
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
	HlsFragmentsPerTsSegment() *float64
	SetHlsFragmentsPerTsSegment(val *float64)
	HlsFragmentsPerTsSegmentInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LiveEventId() *string
	SetLiveEventId(val *string)
	LiveEventIdInput() *string
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputSnapTimeInSeconds() *float64
	SetOutputSnapTimeInSeconds(val *float64)
	OutputSnapTimeInSecondsInput() *float64
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
	RewindWindowDuration() *string
	SetRewindWindowDuration(val *string)
	RewindWindowDurationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaLiveEventOutputTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *MediaLiveEventOutputTimeouts)
	ResetDescription()
	ResetHlsFragmentsPerTsSegment()
	ResetId()
	ResetManifestName()
	ResetOutputSnapTimeInSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRewindWindowDuration()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MediaLiveEventOutput
type jsiiProxy_MediaLiveEventOutput struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaLiveEventOutput) ArchiveWindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveWindowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) ArchiveWindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveWindowDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) AssetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) AssetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) HlsFragmentsPerTsSegment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hlsFragmentsPerTsSegment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) HlsFragmentsPerTsSegmentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hlsFragmentsPerTsSegmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) LiveEventId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveEventId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) LiveEventIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveEventIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) OutputSnapTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputSnapTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) OutputSnapTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputSnapTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) RewindWindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewindWindowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) RewindWindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewindWindowDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) Timeouts() MediaLiveEventOutputTimeoutsOutputReference {
	var returns MediaLiveEventOutputTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventOutput) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_live_event_output azurerm_media_live_event_output} Resource.
func NewMediaLiveEventOutput(scope constructs.Construct, id *string, config *MediaLiveEventOutputConfig) MediaLiveEventOutput {
	_init_.Initialize()

	if err := validateNewMediaLiveEventOutputParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEventOutput{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_live_event_output azurerm_media_live_event_output} Resource.
func NewMediaLiveEventOutput_Override(m MediaLiveEventOutput, scope constructs.Construct, id *string, config *MediaLiveEventOutputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetArchiveWindowDuration(val *string) {
	if err := j.validateSetArchiveWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveWindowDuration",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetAssetName(val *string) {
	if err := j.validateSetAssetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetName",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetHlsFragmentsPerTsSegment(val *float64) {
	if err := j.validateSetHlsFragmentsPerTsSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsFragmentsPerTsSegment",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetLiveEventId(val *string) {
	if err := j.validateSetLiveEventIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveEventId",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetOutputSnapTimeInSeconds(val *float64) {
	if err := j.validateSetOutputSnapTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSnapTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventOutput)SetRewindWindowDuration(val *string) {
	if err := j.validateSetRewindWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rewindWindowDuration",
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
func MediaLiveEventOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEventOutput_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaLiveEventOutput_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEventOutput_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaLiveEventOutput_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaLiveEventOutput_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaLiveEventOutput_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaLiveEventOutput.MediaLiveEventOutput",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaLiveEventOutput) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaLiveEventOutput) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaLiveEventOutput) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventOutput) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) PutTimeouts(value *MediaLiveEventOutputTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetHlsFragmentsPerTsSegment() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsFragmentsPerTsSegment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetOutputSnapTimeInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputSnapTimeInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetRewindWindowDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetRewindWindowDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventOutput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventOutput) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventOutput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

