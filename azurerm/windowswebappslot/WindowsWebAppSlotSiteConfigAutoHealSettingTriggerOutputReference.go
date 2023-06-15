package windowswebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/windowswebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference interface {
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
	InternalValue() *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger
	SetInternalValue(val *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger)
	PrivateMemoryKb() *float64
	SetPrivateMemoryKb(val *float64)
	PrivateMemoryKbInput() *float64
	Requests() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference
	RequestsInput() *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests
	SlowRequest() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	SlowRequestInput() *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest
	StatusCode() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList
	StatusCodeInput() interface{}
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
	PutRequests(value *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests)
	PutSlowRequest(value *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest)
	PutStatusCode(value interface{})
	ResetPrivateMemoryKb()
	ResetRequests()
	ResetSlowRequest()
	ResetStatusCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference
type jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InternalValue() *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger {
	var returns *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PrivateMemoryKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMemoryKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PrivateMemoryKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMemoryKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Requests() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference {
	var returns WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) RequestsInput() *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests {
	var returns *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) SlowRequest() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference {
	var returns WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	_jsii_.Get(
		j,
		"slowRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) SlowRequestInput() *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest {
	var returns *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest
	_jsii_.Get(
		j,
		"slowRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) StatusCode() WindowsWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList {
	var returns WindowsWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) StatusCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference_Override(w WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetInternalValue(val *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetPrivateMemoryKb(val *float64) {
	if err := j.validateSetPrivateMemoryKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateMemoryKb",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutRequests(value *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests) {
	if err := w.validatePutRequestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRequests",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutSlowRequest(value *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest) {
	if err := w.validatePutSlowRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSlowRequest",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutStatusCode(value interface{}) {
	if err := w.validatePutStatusCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStatusCode",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetPrivateMemoryKb() {
	_jsii_.InvokeVoid(
		w,
		"resetPrivateMemoryKb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		w,
		"resetRequests",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetSlowRequest() {
	_jsii_.InvokeVoid(
		w,
		"resetSlowRequest",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		w,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

