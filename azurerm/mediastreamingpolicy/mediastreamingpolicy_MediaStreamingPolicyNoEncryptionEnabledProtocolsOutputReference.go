package mediastreamingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/mediastreamingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

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

	if err := validateNewMediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetDash(val interface{}) {
	if err := j.validateSetDashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dash",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetDownload(val interface{}) {
	if err := j.validateSetDownloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"download",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetHls(val interface{}) {
	if err := j.validateSetHlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hls",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetInternalValue(val *MediaStreamingPolicyNoEncryptionEnabledProtocols) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetSmoothStreaming(val interface{}) {
	if err := j.validateSetSmoothStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaStreamingPolicyNoEncryptionEnabledProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

