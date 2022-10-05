package appserviceslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/appserviceslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSlotLogsOutputReference interface {
	cdktf.ComplexObject
	ApplicationLogs() AppServiceSlotLogsApplicationLogsOutputReference
	ApplicationLogsInput() *AppServiceSlotLogsApplicationLogs
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
	DetailedErrorMessagesEnabled() interface{}
	SetDetailedErrorMessagesEnabled(val interface{})
	DetailedErrorMessagesEnabledInput() interface{}
	FailedRequestTracingEnabled() interface{}
	SetFailedRequestTracingEnabled(val interface{})
	FailedRequestTracingEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HttpLogs() AppServiceSlotLogsHttpLogsOutputReference
	HttpLogsInput() *AppServiceSlotLogsHttpLogs
	InternalValue() *AppServiceSlotLogs
	SetInternalValue(val *AppServiceSlotLogs)
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
	PutApplicationLogs(value *AppServiceSlotLogsApplicationLogs)
	PutHttpLogs(value *AppServiceSlotLogsHttpLogs)
	ResetApplicationLogs()
	ResetDetailedErrorMessagesEnabled()
	ResetFailedRequestTracingEnabled()
	ResetHttpLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsOutputReference
type jsiiProxy_AppServiceSlotLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ApplicationLogs() AppServiceSlotLogsApplicationLogsOutputReference {
	var returns AppServiceSlotLogsApplicationLogsOutputReference
	_jsii_.Get(
		j,
		"applicationLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ApplicationLogsInput() *AppServiceSlotLogsApplicationLogs {
	var returns *AppServiceSlotLogsApplicationLogs
	_jsii_.Get(
		j,
		"applicationLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) DetailedErrorMessagesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessagesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) DetailedErrorMessagesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessagesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) FailedRequestTracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) FailedRequestTracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) HttpLogs() AppServiceSlotLogsHttpLogsOutputReference {
	var returns AppServiceSlotLogsHttpLogsOutputReference
	_jsii_.Get(
		j,
		"httpLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) HttpLogsInput() *AppServiceSlotLogsHttpLogs {
	var returns *AppServiceSlotLogsHttpLogs
	_jsii_.Get(
		j,
		"httpLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) InternalValue() *AppServiceSlotLogs {
	var returns *AppServiceSlotLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsOutputReference {
	_init_.Initialize()

	if err := validateNewAppServiceSlotLogsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceSlotLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsOutputReference_Override(a AppServiceSlotLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetDetailedErrorMessagesEnabled(val interface{}) {
	if err := j.validateSetDetailedErrorMessagesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detailedErrorMessagesEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetFailedRequestTracingEnabled(val interface{}) {
	if err := j.validateSetFailedRequestTracingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedRequestTracingEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetInternalValue(val *AppServiceSlotLogs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) PutApplicationLogs(value *AppServiceSlotLogsApplicationLogs) {
	if err := a.validatePutApplicationLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) PutHttpLogs(value *AppServiceSlotLogsHttpLogs) {
	if err := a.validatePutHttpLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetApplicationLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetDetailedErrorMessagesEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedErrorMessagesEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetFailedRequestTracingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFailedRequestTracingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetHttpLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

