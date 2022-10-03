package mediastreamingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/mediastreamingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy azurerm_media_streaming_policy}.
type MediaStreamingPolicy interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommonEncryptionCbcs() MediaStreamingPolicyCommonEncryptionCbcsOutputReference
	CommonEncryptionCbcsInput() *MediaStreamingPolicyCommonEncryptionCbcs
	CommonEncryptionCenc() MediaStreamingPolicyCommonEncryptionCencOutputReference
	CommonEncryptionCencInput() *MediaStreamingPolicyCommonEncryptionCenc
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
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
	NoEncryptionEnabledProtocols() MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference
	NoEncryptionEnabledProtocolsInput() *MediaStreamingPolicyNoEncryptionEnabledProtocols
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaStreamingPolicyTimeoutsOutputReference
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
	PutCommonEncryptionCbcs(value *MediaStreamingPolicyCommonEncryptionCbcs)
	PutCommonEncryptionCenc(value *MediaStreamingPolicyCommonEncryptionCenc)
	PutNoEncryptionEnabledProtocols(value *MediaStreamingPolicyNoEncryptionEnabledProtocols)
	PutTimeouts(value *MediaStreamingPolicyTimeouts)
	ResetCommonEncryptionCbcs()
	ResetCommonEncryptionCenc()
	ResetDefaultContentKeyPolicyName()
	ResetId()
	ResetNoEncryptionEnabledProtocols()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for MediaStreamingPolicy
type jsiiProxy_MediaStreamingPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaStreamingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) CommonEncryptionCbcs() MediaStreamingPolicyCommonEncryptionCbcsOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCbcsOutputReference
	_jsii_.Get(
		j,
		"commonEncryptionCbcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) CommonEncryptionCbcsInput() *MediaStreamingPolicyCommonEncryptionCbcs {
	var returns *MediaStreamingPolicyCommonEncryptionCbcs
	_jsii_.Get(
		j,
		"commonEncryptionCbcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) CommonEncryptionCenc() MediaStreamingPolicyCommonEncryptionCencOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCencOutputReference
	_jsii_.Get(
		j,
		"commonEncryptionCenc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) CommonEncryptionCencInput() *MediaStreamingPolicyCommonEncryptionCenc {
	var returns *MediaStreamingPolicyCommonEncryptionCenc
	_jsii_.Get(
		j,
		"commonEncryptionCencInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) DefaultContentKeyPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContentKeyPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) DefaultContentKeyPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContentKeyPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) MediaServicesAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) MediaServicesAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) NoEncryptionEnabledProtocols() MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference {
	var returns MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference
	_jsii_.Get(
		j,
		"noEncryptionEnabledProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) NoEncryptionEnabledProtocolsInput() *MediaStreamingPolicyNoEncryptionEnabledProtocols {
	var returns *MediaStreamingPolicyNoEncryptionEnabledProtocols
	_jsii_.Get(
		j,
		"noEncryptionEnabledProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) Timeouts() MediaStreamingPolicyTimeoutsOutputReference {
	var returns MediaStreamingPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy azurerm_media_streaming_policy} Resource.
func NewMediaStreamingPolicy(scope constructs.Construct, id *string, config *MediaStreamingPolicyConfig) MediaStreamingPolicy {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy azurerm_media_streaming_policy} Resource.
func NewMediaStreamingPolicy_Override(m MediaStreamingPolicy, scope constructs.Construct, id *string, config *MediaStreamingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicy",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetDefaultContentKeyPolicyName(val *string) {
	_jsii_.Set(
		j,
		"defaultContentKeyPolicyName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetMediaServicesAccountName(val *string) {
	_jsii_.Set(
		j,
		"mediaServicesAccountName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicy) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func MediaStreamingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaStreamingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) PutCommonEncryptionCbcs(value *MediaStreamingPolicyCommonEncryptionCbcs) {
	_jsii_.InvokeVoid(
		m,
		"putCommonEncryptionCbcs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) PutCommonEncryptionCenc(value *MediaStreamingPolicyCommonEncryptionCenc) {
	_jsii_.InvokeVoid(
		m,
		"putCommonEncryptionCenc",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) PutNoEncryptionEnabledProtocols(value *MediaStreamingPolicyNoEncryptionEnabledProtocols) {
	_jsii_.InvokeVoid(
		m,
		"putNoEncryptionEnabledProtocols",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) PutTimeouts(value *MediaStreamingPolicyTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetCommonEncryptionCbcs() {
	_jsii_.InvokeVoid(
		m,
		"resetCommonEncryptionCbcs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetCommonEncryptionCenc() {
	_jsii_.InvokeVoid(
		m,
		"resetCommonEncryptionCenc",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetDefaultContentKeyPolicyName() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultContentKeyPolicyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetNoEncryptionEnabledProtocols() {
	_jsii_.InvokeVoid(
		m,
		"resetNoEncryptionEnabledProtocols",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCbcs struct {
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// drm_fairplay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#drm_fairplay MediaStreamingPolicy#drm_fairplay}
	DrmFairplay *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay `field:"optional" json:"drmFairplay" yaml:"drmFairplay"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

type MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#label MediaStreamingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#policy_name MediaStreamingPolicy#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

type MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetLabel()
	ResetPolicyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ResetPolicyName() {
	_jsii_.InvokeVoid(
		m,
		"resetPolicyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#allow_persistent_license MediaStreamingPolicy#allow_persistent_license}.
	AllowPersistentLicense interface{} `field:"optional" json:"allowPersistentLicense" yaml:"allowPersistentLicense"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#custom_license_acquisition_url_template MediaStreamingPolicy#custom_license_acquisition_url_template}.
	CustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"customLicenseAcquisitionUrlTemplate" yaml:"customLicenseAcquisitionUrlTemplate"`
}

type MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference interface {
	cdktf.ComplexObject
	AllowPersistentLicense() interface{}
	SetAllowPersistentLicense(val interface{})
	AllowPersistentLicenseInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomLicenseAcquisitionUrlTemplate() *string
	SetCustomLicenseAcquisitionUrlTemplate(val *string)
	CustomLicenseAcquisitionUrlTemplateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowPersistentLicense()
	ResetCustomLicenseAcquisitionUrlTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) AllowPersistentLicense() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPersistentLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) AllowPersistentLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPersistentLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CustomLicenseAcquisitionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) CustomLicenseAcquisitionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetAllowPersistentLicense(val interface{}) {
	_jsii_.Set(
		j,
		"allowPersistentLicense",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetCustomLicenseAcquisitionUrlTemplate(val *string) {
	_jsii_.Set(
		j,
		"customLicenseAcquisitionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ResetAllowPersistentLicense() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowPersistentLicense",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ResetCustomLicenseAcquisitionUrlTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLicenseAcquisitionUrlTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#dash MediaStreamingPolicy#dash}.
	Dash interface{} `field:"optional" json:"dash" yaml:"dash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#download MediaStreamingPolicy#download}.
	Download interface{} `field:"optional" json:"download" yaml:"download"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#hls MediaStreamingPolicy#hls}.
	Hls interface{} `field:"optional" json:"hls" yaml:"hls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#smooth_streaming MediaStreamingPolicy#smooth_streaming}.
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
}

type MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dash() interface{}
	SetDash(val interface{})
	DashInput() interface{}
	Download() interface{}
	SetDownload(val interface{})
	DownloadInput() interface{}
	// Experimental.
	Fqn() *string
	Hls() interface{}
	SetHls(val interface{})
	HlsInput() interface{}
	InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols)
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDash()
	ResetDownload()
	ResetHls()
	ResetSmoothStreaming()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) Dash() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) DashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) Download() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"download",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) DownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) Hls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) HlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetDash(val interface{}) {
	_jsii_.Set(
		j,
		"dash",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetDownload(val interface{}) {
	_jsii_.Set(
		j,
		"download",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetHls(val interface{}) {
	_jsii_.Set(
		j,
		"hls",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetSmoothStreaming(val interface{}) {
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ResetDash() {
	_jsii_.InvokeVoid(
		m,
		"resetDash",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ResetDownload() {
	_jsii_.InvokeVoid(
		m,
		"resetDownload",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ResetHls() {
	_jsii_.InvokeVoid(
		m,
		"resetHls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		m,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCbcsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultContentKey() MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference
	DefaultContentKeyInput() *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey
	DrmFairplay() MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference
	DrmFairplayInput() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	EnabledProtocols() MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference
	EnabledProtocolsInput() *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCbcs
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcs)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey)
	PutDrmFairplay(value *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay)
	PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols)
	ResetDefaultContentKey()
	ResetDrmFairplay()
	ResetEnabledProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCbcsOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) DefaultContentKey() MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKeyOutputReference
	_jsii_.Get(
		j,
		"defaultContentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) DefaultContentKeyInput() *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey
	_jsii_.Get(
		j,
		"defaultContentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) DrmFairplay() MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCbcsDrmFairplayOutputReference
	_jsii_.Get(
		j,
		"drmFairplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) DrmFairplayInput() *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay
	_jsii_.Get(
		j,
		"drmFairplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) EnabledProtocols() MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocolsOutputReference
	_jsii_.Get(
		j,
		"enabledProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) EnabledProtocolsInput() *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols
	_jsii_.Get(
		j,
		"enabledProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCbcs {
	var returns *MediaStreamingPolicyCommonEncryptionCbcs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCbcsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCbcsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCbcsOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCbcsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCbcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey) {
	_jsii_.InvokeVoid(
		m,
		"putDefaultContentKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutDrmFairplay(value *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay) {
	_jsii_.InvokeVoid(
		m,
		"putDrmFairplay",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols) {
	_jsii_.InvokeVoid(
		m,
		"putEnabledProtocols",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ResetDefaultContentKey() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultContentKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ResetDrmFairplay() {
	_jsii_.InvokeVoid(
		m,
		"resetDrmFairplay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ResetEnabledProtocols() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabledProtocols",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCenc struct {
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// drm_playready block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#drm_playready MediaStreamingPolicy#drm_playready}
	DrmPlayready *MediaStreamingPolicyCommonEncryptionCencDrmPlayready `field:"optional" json:"drmPlayready" yaml:"drmPlayready"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#drm_widevine_custom_license_acquisition_url_template MediaStreamingPolicy#drm_widevine_custom_license_acquisition_url_template}.
	DrmWidevineCustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"drmWidevineCustomLicenseAcquisitionUrlTemplate" yaml:"drmWidevineCustomLicenseAcquisitionUrlTemplate"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

type MediaStreamingPolicyCommonEncryptionCencDefaultContentKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#label MediaStreamingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#policy_name MediaStreamingPolicy#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

type MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetLabel()
	ResetPolicyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey {
	var returns *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ResetPolicyName() {
	_jsii_.InvokeVoid(
		m,
		"resetPolicyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCencDrmPlayready struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#custom_attributes MediaStreamingPolicy#custom_attributes}.
	CustomAttributes *string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#custom_license_acquisition_url_template MediaStreamingPolicy#custom_license_acquisition_url_template}.
	CustomLicenseAcquisitionUrlTemplate *string `field:"optional" json:"customLicenseAcquisitionUrlTemplate" yaml:"customLicenseAcquisitionUrlTemplate"`
}

type MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAttributes() *string
	SetCustomAttributes(val *string)
	CustomAttributesInput() *string
	CustomLicenseAcquisitionUrlTemplate() *string
	SetCustomLicenseAcquisitionUrlTemplate(val *string)
	CustomLicenseAcquisitionUrlTemplateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCencDrmPlayready
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencDrmPlayready)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCustomAttributes()
	ResetCustomLicenseAcquisitionUrlTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) CustomAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) CustomAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) CustomLicenseAcquisitionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) CustomLicenseAcquisitionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLicenseAcquisitionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCencDrmPlayready {
	var returns *MediaStreamingPolicyCommonEncryptionCencDrmPlayready
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetCustomAttributes(val *string) {
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetCustomLicenseAcquisitionUrlTemplate(val *string) {
	_jsii_.Set(
		j,
		"customLicenseAcquisitionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencDrmPlayready) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ResetCustomLicenseAcquisitionUrlTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLicenseAcquisitionUrlTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCencEnabledProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#dash MediaStreamingPolicy#dash}.
	Dash interface{} `field:"optional" json:"dash" yaml:"dash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#download MediaStreamingPolicy#download}.
	Download interface{} `field:"optional" json:"download" yaml:"download"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#hls MediaStreamingPolicy#hls}.
	Hls interface{} `field:"optional" json:"hls" yaml:"hls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#smooth_streaming MediaStreamingPolicy#smooth_streaming}.
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
}

type MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dash() interface{}
	SetDash(val interface{})
	DashInput() interface{}
	Download() interface{}
	SetDownload(val interface{})
	DownloadInput() interface{}
	// Experimental.
	Fqn() *string
	Hls() interface{}
	SetHls(val interface{})
	HlsInput() interface{}
	InternalValue() *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols)
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDash()
	ResetDownload()
	ResetHls()
	ResetSmoothStreaming()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) Dash() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) DashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) Download() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"download",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) DownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) Hls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) HlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols {
	var returns *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetDash(val interface{}) {
	_jsii_.Set(
		j,
		"dash",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetDownload(val interface{}) {
	_jsii_.Set(
		j,
		"download",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetHls(val interface{}) {
	_jsii_.Set(
		j,
		"hls",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetSmoothStreaming(val interface{}) {
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ResetDash() {
	_jsii_.InvokeVoid(
		m,
		"resetDash",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ResetDownload() {
	_jsii_.InvokeVoid(
		m,
		"resetDownload",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ResetHls() {
	_jsii_.InvokeVoid(
		m,
		"resetHls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		m,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyCommonEncryptionCencOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultContentKey() MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference
	DefaultContentKeyInput() *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey
	DrmPlayready() MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference
	DrmPlayreadyInput() *MediaStreamingPolicyCommonEncryptionCencDrmPlayready
	DrmWidevineCustomLicenseAcquisitionUrlTemplate() *string
	SetDrmWidevineCustomLicenseAcquisitionUrlTemplate(val *string)
	DrmWidevineCustomLicenseAcquisitionUrlTemplateInput() *string
	EnabledProtocols() MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference
	EnabledProtocolsInput() *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols
	// Experimental.
	Fqn() *string
	InternalValue() *MediaStreamingPolicyCommonEncryptionCenc
	SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCenc)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey)
	PutDrmPlayready(value *MediaStreamingPolicyCommonEncryptionCencDrmPlayready)
	PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols)
	ResetDefaultContentKey()
	ResetDrmPlayready()
	ResetDrmWidevineCustomLicenseAcquisitionUrlTemplate()
	ResetEnabledProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyCommonEncryptionCencOutputReference
type jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DefaultContentKey() MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCencDefaultContentKeyOutputReference
	_jsii_.Get(
		j,
		"defaultContentKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DefaultContentKeyInput() *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey {
	var returns *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey
	_jsii_.Get(
		j,
		"defaultContentKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DrmPlayready() MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCencDrmPlayreadyOutputReference
	_jsii_.Get(
		j,
		"drmPlayready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DrmPlayreadyInput() *MediaStreamingPolicyCommonEncryptionCencDrmPlayready {
	var returns *MediaStreamingPolicyCommonEncryptionCencDrmPlayready
	_jsii_.Get(
		j,
		"drmPlayreadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DrmWidevineCustomLicenseAcquisitionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmWidevineCustomLicenseAcquisitionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) DrmWidevineCustomLicenseAcquisitionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drmWidevineCustomLicenseAcquisitionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) EnabledProtocols() MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCencEnabledProtocolsOutputReference
	_jsii_.Get(
		j,
		"enabledProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) EnabledProtocolsInput() *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols {
	var returns *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols
	_jsii_.Get(
		j,
		"enabledProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) InternalValue() *MediaStreamingPolicyCommonEncryptionCenc {
	var returns *MediaStreamingPolicyCommonEncryptionCenc
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyCommonEncryptionCencOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyCommonEncryptionCencOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyCommonEncryptionCencOutputReference_Override(m MediaStreamingPolicyCommonEncryptionCencOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyCommonEncryptionCencOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetDrmWidevineCustomLicenseAcquisitionUrlTemplate(val *string) {
	_jsii_.Set(
		j,
		"drmWidevineCustomLicenseAcquisitionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCenc) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey) {
	_jsii_.InvokeVoid(
		m,
		"putDefaultContentKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutDrmPlayready(value *MediaStreamingPolicyCommonEncryptionCencDrmPlayready) {
	_jsii_.InvokeVoid(
		m,
		"putDrmPlayready",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols) {
	_jsii_.InvokeVoid(
		m,
		"putEnabledProtocols",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetDefaultContentKey() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultContentKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetDrmPlayready() {
	_jsii_.InvokeVoid(
		m,
		"resetDrmPlayready",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetDrmWidevineCustomLicenseAcquisitionUrlTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetDrmWidevineCustomLicenseAcquisitionUrlTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetEnabledProtocols() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabledProtocols",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#media_services_account_name MediaStreamingPolicy#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#name MediaStreamingPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#resource_group_name MediaStreamingPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// common_encryption_cbcs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#common_encryption_cbcs MediaStreamingPolicy#common_encryption_cbcs}
	CommonEncryptionCbcs *MediaStreamingPolicyCommonEncryptionCbcs `field:"optional" json:"commonEncryptionCbcs" yaml:"commonEncryptionCbcs"`
	// common_encryption_cenc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#common_encryption_cenc MediaStreamingPolicy#common_encryption_cenc}
	CommonEncryptionCenc *MediaStreamingPolicyCommonEncryptionCenc `field:"optional" json:"commonEncryptionCenc" yaml:"commonEncryptionCenc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#default_content_key_policy_name MediaStreamingPolicy#default_content_key_policy_name}.
	DefaultContentKeyPolicyName *string `field:"optional" json:"defaultContentKeyPolicyName" yaml:"defaultContentKeyPolicyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#id MediaStreamingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// no_encryption_enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#no_encryption_enabled_protocols MediaStreamingPolicy#no_encryption_enabled_protocols}
	NoEncryptionEnabledProtocols *MediaStreamingPolicyNoEncryptionEnabledProtocols `field:"optional" json:"noEncryptionEnabledProtocols" yaml:"noEncryptionEnabledProtocols"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#timeouts MediaStreamingPolicy#timeouts}
	Timeouts *MediaStreamingPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MediaStreamingPolicyNoEncryptionEnabledProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#dash MediaStreamingPolicy#dash}.
	Dash interface{} `field:"optional" json:"dash" yaml:"dash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#download MediaStreamingPolicy#download}.
	Download interface{} `field:"optional" json:"download" yaml:"download"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#hls MediaStreamingPolicy#hls}.
	Hls interface{} `field:"optional" json:"hls" yaml:"hls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#smooth_streaming MediaStreamingPolicy#smooth_streaming}.
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
}

type MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dash() interface{}
	SetDash(val interface{})
	DashInput() interface{}
	Download() interface{}
	SetDownload(val interface{})
	DownloadInput() interface{}
	// Experimental.
	Fqn() *string
	Hls() interface{}
	SetHls(val interface{})
	HlsInput() interface{}
	InternalValue() *MediaStreamingPolicyNoEncryptionEnabledProtocols
	SetInternalValue(val *MediaStreamingPolicyNoEncryptionEnabledProtocols)
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDash()
	ResetDownload()
	ResetHls()
	ResetSmoothStreaming()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference
type jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) Dash() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) DashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) Download() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"download",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) DownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) Hls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) HlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) InternalValue() *MediaStreamingPolicyNoEncryptionEnabledProtocols {
	var returns *MediaStreamingPolicyNoEncryptionEnabledProtocols
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference_Override(m MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetDash(val interface{}) {
	_jsii_.Set(
		j,
		"dash",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetDownload(val interface{}) {
	_jsii_.Set(
		j,
		"download",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetHls(val interface{}) {
	_jsii_.Set(
		j,
		"hls",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetInternalValue(val *MediaStreamingPolicyNoEncryptionEnabledProtocols) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetSmoothStreaming(val interface{}) {
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ResetDash() {
	_jsii_.InvokeVoid(
		m,
		"resetDash",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ResetDownload() {
	_jsii_.InvokeVoid(
		m,
		"resetDownload",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ResetHls() {
	_jsii_.InvokeVoid(
		m,
		"resetHls",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		m,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaStreamingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#create MediaStreamingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#delete MediaStreamingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_policy#read MediaStreamingPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

type MediaStreamingPolicyTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaStreamingPolicyTimeoutsOutputReference
type jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaStreamingPolicyTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaStreamingPolicyTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaStreamingPolicyTimeoutsOutputReference_Override(m MediaStreamingPolicyTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaStreamingPolicy.MediaStreamingPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

