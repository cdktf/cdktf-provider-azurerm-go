package apimanagementdiagnostic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/apimanagementdiagnostic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementDiagnosticBackendResponseOutputReference interface {
	cdktf.ComplexObject
	BodyBytes() *float64
	SetBodyBytes(val *float64)
	BodyBytesInput() *float64
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
	DataMasking() ApiManagementDiagnosticBackendResponseDataMaskingOutputReference
	DataMaskingInput() *ApiManagementDiagnosticBackendResponseDataMasking
	// Experimental.
	Fqn() *string
	HeadersToLog() *[]*string
	SetHeadersToLog(val *[]*string)
	HeadersToLogInput() *[]*string
	InternalValue() *ApiManagementDiagnosticBackendResponse
	SetInternalValue(val *ApiManagementDiagnosticBackendResponse)
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
	PutDataMasking(value *ApiManagementDiagnosticBackendResponseDataMasking)
	ResetBodyBytes()
	ResetDataMasking()
	ResetHeadersToLog()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementDiagnosticBackendResponseOutputReference
type jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) BodyBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bodyBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) BodyBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bodyBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) DataMasking() ApiManagementDiagnosticBackendResponseDataMaskingOutputReference {
	var returns ApiManagementDiagnosticBackendResponseDataMaskingOutputReference
	_jsii_.Get(
		j,
		"dataMasking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) DataMaskingInput() *ApiManagementDiagnosticBackendResponseDataMasking {
	var returns *ApiManagementDiagnosticBackendResponseDataMasking
	_jsii_.Get(
		j,
		"dataMaskingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) HeadersToLog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersToLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) HeadersToLogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersToLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) InternalValue() *ApiManagementDiagnosticBackendResponse {
	var returns *ApiManagementDiagnosticBackendResponse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementDiagnosticBackendResponseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementDiagnosticBackendResponseOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementDiagnosticBackendResponseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnosticBackendResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementDiagnosticBackendResponseOutputReference_Override(a ApiManagementDiagnosticBackendResponseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnosticBackendResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetBodyBytes(val *float64) {
	if err := j.validateSetBodyBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyBytes",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetHeadersToLog(val *[]*string) {
	if err := j.validateSetHeadersToLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headersToLog",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetInternalValue(val *ApiManagementDiagnosticBackendResponse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) PutDataMasking(value *ApiManagementDiagnosticBackendResponseDataMasking) {
	if err := a.validatePutDataMaskingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataMasking",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ResetBodyBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetBodyBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ResetDataMasking() {
	_jsii_.InvokeVoid(
		a,
		"resetDataMasking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ResetHeadersToLog() {
	_jsii_.InvokeVoid(
		a,
		"resetHeadersToLog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

