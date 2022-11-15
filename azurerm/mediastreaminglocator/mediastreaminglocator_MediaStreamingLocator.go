package mediastreaminglocator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/mediastreaminglocator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator azurerm_media_streaming_locator}.
type MediaStreamingLocator interface {
	cdktf.TerraformResource
	AlternativeMediaId() *string
	SetAlternativeMediaId(val *string)
	AlternativeMediaIdInput() *string
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
	ContentKey() MediaStreamingLocatorContentKeyList
	ContentKeyInput() interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultContentKeyPolicyName() *string
	SetDefaultContentKeyPolicyName(val *string)
	DefaultContentKeyPolicyNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	StreamingLocatorId() *string
	SetStreamingLocatorId(val *string)
	StreamingLocatorIdInput() *string
	StreamingPolicyName() *string
	SetStreamingPolicyName(val *string)
	StreamingPolicyNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaStreamingLocatorTimeoutsOutputReference
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
	PutContentKey(value interface{})
	PutTimeouts(value *MediaStreamingLocatorTimeouts)
	ResetAlternativeMediaId()
	ResetContentKey()
	ResetDefaultContentKeyPolicyName()
	ResetEndTime()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartTime()
	ResetStreamingLocatorId()
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

// The jsii proxy struct for MediaStreamingLocator
type jsiiProxy_MediaStreamingLocator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaStreamingLocator) AlternativeMediaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeMediaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) AlternativeMediaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeMediaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) AssetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) AssetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ContentKey() MediaStreamingLocatorContentKeyList {
	var returns MediaStreamingLocatorContentKeyList
	_jsii_.Get(
		j,
		"contentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ContentKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) DefaultContentKeyPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContentKeyPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) DefaultContentKeyPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContentKeyPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) MediaServicesAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) MediaServicesAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StreamingLocatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingLocatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StreamingLocatorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingLocatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StreamingPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) StreamingPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) Timeouts() MediaStreamingLocatorTimeoutsOutputReference {
	var returns MediaStreamingLocatorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingLocator) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator azurerm_media_streaming_locator} Resource.
func NewMediaStreamingLocator(scope constructs.Construct, id *string, config *MediaStreamingLocatorConfig) MediaStreamingLocator {
	_init_.Initialize()

	if err := validateNewMediaStreamingLocatorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaStreamingLocator{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator azurerm_media_streaming_locator} Resource.
func NewMediaStreamingLocator_Override(m MediaStreamingLocator, scope constructs.Construct, id *string, config *MediaStreamingLocatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetAlternativeMediaId(val *string) {
	if err := j.validateSetAlternativeMediaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeMediaId",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetAssetName(val *string) {
	if err := j.validateSetAssetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetDefaultContentKeyPolicyName(val *string) {
	if err := j.validateSetDefaultContentKeyPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultContentKeyPolicyName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetMediaServicesAccountName(val *string) {
	if err := j.validateSetMediaServicesAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaServicesAccountName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetStreamingLocatorId(val *string) {
	if err := j.validateSetStreamingLocatorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingLocatorId",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingLocator)SetStreamingPolicyName(val *string) {
	if err := j.validateSetStreamingPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingPolicyName",
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
func MediaStreamingLocator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingLocator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaStreamingLocator_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingLocator_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediaStreamingLocator_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaStreamingLocator_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaStreamingLocator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaStreamingLocator.MediaStreamingLocator",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaStreamingLocator) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaStreamingLocator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaStreamingLocator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingLocator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaStreamingLocator) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaStreamingLocator) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaStreamingLocator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaStreamingLocator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaStreamingLocator) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaStreamingLocator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaStreamingLocator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingLocator) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaStreamingLocator) PutContentKey(value interface{}) {
	if err := m.validatePutContentKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putContentKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingLocator) PutTimeouts(value *MediaStreamingLocatorTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetAlternativeMediaId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlternativeMediaId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetContentKey() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetDefaultContentKeyPolicyName() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultContentKeyPolicyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetEndTime() {
	_jsii_.InvokeVoid(
		m,
		"resetEndTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetStreamingLocatorId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamingLocatorId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingLocator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingLocator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingLocator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingLocator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

