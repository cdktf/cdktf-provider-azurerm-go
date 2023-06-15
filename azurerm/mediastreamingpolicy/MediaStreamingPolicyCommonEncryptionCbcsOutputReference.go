package mediastreamingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediastreamingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaStreamingPolicyCommonEncryptionCbcsOutputReference interface {
	cdktf.ComplexObject
	ClearKeyEncryption() MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryptionOutputReference
	ClearKeyEncryptionInput() *MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryption
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
	PutClearKeyEncryption(value *MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryption)
	PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey)
	PutDrmFairplay(value *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay)
	PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols)
	ResetClearKeyEncryption()
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

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ClearKeyEncryption() MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryptionOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryptionOutputReference
	_jsii_.Get(
		j,
		"clearKeyEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ClearKeyEncryptionInput() *MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryption {
	var returns *MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryption
	_jsii_.Get(
		j,
		"clearKeyEncryptionInput",
		&returns,
	)
	return returns
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

	if err := validateNewMediaStreamingPolicyCommonEncryptionCbcsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference)SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCbcs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutClearKeyEncryption(value *MediaStreamingPolicyCommonEncryptionCbcsClearKeyEncryption) {
	if err := m.validatePutClearKeyEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClearKeyEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey) {
	if err := m.validatePutDefaultContentKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDefaultContentKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutDrmFairplay(value *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay) {
	if err := m.validatePutDrmFairplayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDrmFairplay",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols) {
	if err := m.validatePutEnabledProtocolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEnabledProtocols",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCbcsOutputReference) ResetClearKeyEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetClearKeyEncryption",
		nil, // no parameters
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
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
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

