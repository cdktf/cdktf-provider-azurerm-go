package medialiveevent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/medialiveevent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventInputOutputReference interface {
	cdktf.ComplexObject
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
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
	Endpoint() MediaLiveEventInputEndpointList
	// Experimental.
	Fqn() *string
	InternalValue() *MediaLiveEventInput
	SetInternalValue(val *MediaLiveEventInput)
	IpAccessControlAllow() MediaLiveEventInputIpAccessControlAllowList
	IpAccessControlAllowInput() interface{}
	KeyFrameIntervalDuration() *string
	SetKeyFrameIntervalDuration(val *string)
	KeyFrameIntervalDurationInput() *string
	StreamingProtocol() *string
	SetStreamingProtocol(val *string)
	StreamingProtocolInput() *string
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
	PutIpAccessControlAllow(value interface{})
	ResetAccessToken()
	ResetIpAccessControlAllow()
	ResetKeyFrameIntervalDuration()
	ResetStreamingProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaLiveEventInputOutputReference
type jsiiProxy_MediaLiveEventInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) Endpoint() MediaLiveEventInputEndpointList {
	var returns MediaLiveEventInputEndpointList
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) InternalValue() *MediaLiveEventInput {
	var returns *MediaLiveEventInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) IpAccessControlAllow() MediaLiveEventInputIpAccessControlAllowList {
	var returns MediaLiveEventInputIpAccessControlAllowList
	_jsii_.Get(
		j,
		"ipAccessControlAllow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) IpAccessControlAllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAccessControlAllowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) KeyFrameIntervalDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameIntervalDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) KeyFrameIntervalDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFrameIntervalDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) StreamingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) StreamingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaLiveEventInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaLiveEventInputOutputReference {
	_init_.Initialize()

	if err := validateNewMediaLiveEventInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEventInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaLiveEventInputOutputReference_Override(m MediaLiveEventInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetInternalValue(val *MediaLiveEventInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetKeyFrameIntervalDuration(val *string) {
	if err := j.validateSetKeyFrameIntervalDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFrameIntervalDuration",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetStreamingProtocol(val *string) {
	if err := j.validateSetStreamingProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingProtocol",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) PutIpAccessControlAllow(value interface{}) {
	if err := m.validatePutIpAccessControlAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIpAccessControlAllow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		m,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ResetIpAccessControlAllow() {
	_jsii_.InvokeVoid(
		m,
		"resetIpAccessControlAllow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ResetKeyFrameIntervalDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFrameIntervalDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ResetStreamingProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamingProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaLiveEventInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

