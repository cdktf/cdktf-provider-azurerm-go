package labservicelab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/labservicelab/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LabServiceLabAutoShutdownOutputReference interface {
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
	DisconnectDelay() *string
	SetDisconnectDelay(val *string)
	DisconnectDelayInput() *string
	// Experimental.
	Fqn() *string
	IdleDelay() *string
	SetIdleDelay(val *string)
	IdleDelayInput() *string
	InternalValue() *LabServiceLabAutoShutdown
	SetInternalValue(val *LabServiceLabAutoShutdown)
	NoConnectDelay() *string
	SetNoConnectDelay(val *string)
	NoConnectDelayInput() *string
	ShutdownOnIdle() *string
	SetShutdownOnIdle(val *string)
	ShutdownOnIdleInput() *string
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
	ResetDisconnectDelay()
	ResetIdleDelay()
	ResetNoConnectDelay()
	ResetShutdownOnIdle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LabServiceLabAutoShutdownOutputReference
type jsiiProxy_LabServiceLabAutoShutdownOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) DisconnectDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disconnectDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) DisconnectDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disconnectDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) IdleDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) IdleDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) InternalValue() *LabServiceLabAutoShutdown {
	var returns *LabServiceLabAutoShutdown
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) NoConnectDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noConnectDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) NoConnectDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noConnectDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ShutdownOnIdle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownOnIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ShutdownOnIdleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shutdownOnIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLabServiceLabAutoShutdownOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LabServiceLabAutoShutdownOutputReference {
	_init_.Initialize()

	if err := validateNewLabServiceLabAutoShutdownOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LabServiceLabAutoShutdownOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLabAutoShutdownOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLabServiceLabAutoShutdownOutputReference_Override(l LabServiceLabAutoShutdownOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLabAutoShutdownOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetDisconnectDelay(val *string) {
	if err := j.validateSetDisconnectDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectDelay",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetIdleDelay(val *string) {
	if err := j.validateSetIdleDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDelay",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetInternalValue(val *LabServiceLabAutoShutdown) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetNoConnectDelay(val *string) {
	if err := j.validateSetNoConnectDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConnectDelay",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetShutdownOnIdle(val *string) {
	if err := j.validateSetShutdownOnIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownOnIdle",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabAutoShutdownOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ResetDisconnectDelay() {
	_jsii_.InvokeVoid(
		l,
		"resetDisconnectDelay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ResetIdleDelay() {
	_jsii_.InvokeVoid(
		l,
		"resetIdleDelay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ResetNoConnectDelay() {
	_jsii_.InvokeVoid(
		l,
		"resetNoConnectDelay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ResetShutdownOnIdle() {
	_jsii_.InvokeVoid(
		l,
		"resetShutdownOnIdle",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabAutoShutdownOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

