package mediastreamingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mediastreamingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaStreamingPolicyCommonEncryptionCencOutputReference interface {
	cdktf.ComplexObject
	ClearKeyEncryption() MediaStreamingPolicyCommonEncryptionCencClearKeyEncryptionOutputReference
	ClearKeyEncryptionInput() *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption
	ClearTrack() MediaStreamingPolicyCommonEncryptionCencClearTrackList
	ClearTrackInput() interface{}
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
	ContentKeyToTrackMapping() MediaStreamingPolicyCommonEncryptionCencContentKeyToTrackMappingList
	ContentKeyToTrackMappingInput() interface{}
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
	PutClearKeyEncryption(value *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption)
	PutClearTrack(value interface{})
	PutContentKeyToTrackMapping(value interface{})
	PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey)
	PutDrmPlayready(value *MediaStreamingPolicyCommonEncryptionCencDrmPlayready)
	PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols)
	ResetClearKeyEncryption()
	ResetClearTrack()
	ResetContentKeyToTrackMapping()
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

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ClearKeyEncryption() MediaStreamingPolicyCommonEncryptionCencClearKeyEncryptionOutputReference {
	var returns MediaStreamingPolicyCommonEncryptionCencClearKeyEncryptionOutputReference
	_jsii_.Get(
		j,
		"clearKeyEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ClearKeyEncryptionInput() *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption {
	var returns *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption
	_jsii_.Get(
		j,
		"clearKeyEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ClearTrack() MediaStreamingPolicyCommonEncryptionCencClearTrackList {
	var returns MediaStreamingPolicyCommonEncryptionCencClearTrackList
	_jsii_.Get(
		j,
		"clearTrack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ClearTrackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearTrackInput",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ContentKeyToTrackMapping() MediaStreamingPolicyCommonEncryptionCencContentKeyToTrackMappingList {
	var returns MediaStreamingPolicyCommonEncryptionCencContentKeyToTrackMappingList
	_jsii_.Get(
		j,
		"contentKeyToTrackMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ContentKeyToTrackMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyToTrackMappingInput",
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

	if err := validateNewMediaStreamingPolicyCommonEncryptionCencOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetDrmWidevineCustomLicenseAcquisitionUrlTemplate(val *string) {
	if err := j.validateSetDrmWidevineCustomLicenseAcquisitionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drmWidevineCustomLicenseAcquisitionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetInternalValue(val *MediaStreamingPolicyCommonEncryptionCenc) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutClearKeyEncryption(value *MediaStreamingPolicyCommonEncryptionCencClearKeyEncryption) {
	if err := m.validatePutClearKeyEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClearKeyEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutClearTrack(value interface{}) {
	if err := m.validatePutClearTrackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClearTrack",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutContentKeyToTrackMapping(value interface{}) {
	if err := m.validatePutContentKeyToTrackMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putContentKeyToTrackMapping",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutDefaultContentKey(value *MediaStreamingPolicyCommonEncryptionCencDefaultContentKey) {
	if err := m.validatePutDefaultContentKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDefaultContentKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutDrmPlayready(value *MediaStreamingPolicyCommonEncryptionCencDrmPlayready) {
	if err := m.validatePutDrmPlayreadyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDrmPlayready",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) PutEnabledProtocols(value *MediaStreamingPolicyCommonEncryptionCencEnabledProtocols) {
	if err := m.validatePutEnabledProtocolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEnabledProtocols",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetClearKeyEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetClearKeyEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetClearTrack() {
	_jsii_.InvokeVoid(
		m,
		"resetClearTrack",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaStreamingPolicyCommonEncryptionCencOutputReference) ResetContentKeyToTrackMapping() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKeyToTrackMapping",
		nil, // no parameters
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

