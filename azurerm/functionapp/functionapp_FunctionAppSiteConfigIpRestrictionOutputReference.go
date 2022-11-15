package functionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/functionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppSiteConfigIpRestrictionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Headers() FunctionAppSiteConfigIpRestrictionHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ServiceTag() *string
	SetServiceTag(val *string)
	ServiceTagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
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
	PutHeaders(value interface{})
	ResetAction()
	ResetHeaders()
	ResetIpAddress()
	ResetName()
	ResetPriority()
	ResetServiceTag()
	ResetVirtualNetworkSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppSiteConfigIpRestrictionOutputReference
type jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Headers() FunctionAppSiteConfigIpRestrictionHeadersList {
	var returns FunctionAppSiteConfigIpRestrictionHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ServiceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ServiceTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}


func NewFunctionAppSiteConfigIpRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FunctionAppSiteConfigIpRestrictionOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppSiteConfigIpRestrictionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFunctionAppSiteConfigIpRestrictionOutputReference_Override(f FunctionAppSiteConfigIpRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetServiceTag(val *string) {
	if err := j.validateSetServiceTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceTag",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference)SetVirtualNetworkSubnetId(val *string) {
	if err := j.validateSetVirtualNetworkSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) PutHeaders(value interface{}) {
	if err := f.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putHeaders",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		f,
		"resetAction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		f,
		"resetHeaders",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		f,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		f,
		"resetPriority",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetServiceTag() {
	_jsii_.InvokeVoid(
		f,
		"resetServiceTag",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		f,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigIpRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

