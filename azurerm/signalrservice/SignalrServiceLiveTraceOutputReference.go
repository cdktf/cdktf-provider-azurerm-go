package signalrservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/signalrservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SignalrServiceLiveTraceOutputReference interface {
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
	ConnectivityLogsEnabled() interface{}
	SetConnectivityLogsEnabled(val interface{})
	ConnectivityLogsEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HttpRequestLogsEnabled() interface{}
	SetHttpRequestLogsEnabled(val interface{})
	HttpRequestLogsEnabledInput() interface{}
	InternalValue() *SignalrServiceLiveTrace
	SetInternalValue(val *SignalrServiceLiveTrace)
	MessagingLogsEnabled() interface{}
	SetMessagingLogsEnabled(val interface{})
	MessagingLogsEnabledInput() interface{}
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
	ResetConnectivityLogsEnabled()
	ResetEnabled()
	ResetHttpRequestLogsEnabled()
	ResetMessagingLogsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SignalrServiceLiveTraceOutputReference
type jsiiProxy_SignalrServiceLiveTraceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) ConnectivityLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) ConnectivityLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) HttpRequestLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) HttpRequestLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) InternalValue() *SignalrServiceLiveTrace {
	var returns *SignalrServiceLiveTrace
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) MessagingLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagingLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) MessagingLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagingLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSignalrServiceLiveTraceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SignalrServiceLiveTraceOutputReference {
	_init_.Initialize()

	if err := validateNewSignalrServiceLiveTraceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SignalrServiceLiveTraceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.signalrService.SignalrServiceLiveTraceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSignalrServiceLiveTraceOutputReference_Override(s SignalrServiceLiveTraceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.signalrService.SignalrServiceLiveTraceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetConnectivityLogsEnabled(val interface{}) {
	if err := j.validateSetConnectivityLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectivityLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetHttpRequestLogsEnabled(val interface{}) {
	if err := j.validateSetHttpRequestLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRequestLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetInternalValue(val *SignalrServiceLiveTrace) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetMessagingLogsEnabled(val interface{}) {
	if err := j.validateSetMessagingLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagingLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignalrServiceLiveTraceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ResetConnectivityLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetConnectivityLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ResetHttpRequestLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpRequestLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ResetMessagingLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMessagingLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrServiceLiveTraceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

