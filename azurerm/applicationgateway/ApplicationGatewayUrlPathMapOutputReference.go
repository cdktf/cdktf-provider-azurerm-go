package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayUrlPathMapOutputReference interface {
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
	DefaultBackendAddressPoolId() *string
	DefaultBackendAddressPoolName() *string
	SetDefaultBackendAddressPoolName(val *string)
	DefaultBackendAddressPoolNameInput() *string
	DefaultBackendHttpSettingsId() *string
	DefaultBackendHttpSettingsName() *string
	SetDefaultBackendHttpSettingsName(val *string)
	DefaultBackendHttpSettingsNameInput() *string
	DefaultRedirectConfigurationId() *string
	DefaultRedirectConfigurationName() *string
	SetDefaultRedirectConfigurationName(val *string)
	DefaultRedirectConfigurationNameInput() *string
	DefaultRewriteRuleSetId() *string
	DefaultRewriteRuleSetName() *string
	SetDefaultRewriteRuleSetName(val *string)
	DefaultRewriteRuleSetNameInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PathRule() ApplicationGatewayUrlPathMapPathRuleList
	PathRuleInput() interface{}
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
	PutPathRule(value interface{})
	ResetDefaultBackendAddressPoolName()
	ResetDefaultBackendHttpSettingsName()
	ResetDefaultRedirectConfigurationName()
	ResetDefaultRewriteRuleSetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayUrlPathMapOutputReference
type jsiiProxy_ApplicationGatewayUrlPathMapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendAddressPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendAddressPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendAddressPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendAddressPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendAddressPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendHttpSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendHttpSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendHttpSettingsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendHttpSettingsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultBackendHttpSettingsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBackendHttpSettingsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRedirectConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRedirectConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRedirectConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRewriteRuleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRewriteRuleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRewriteRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRewriteRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) DefaultRewriteRuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRewriteRuleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) PathRule() ApplicationGatewayUrlPathMapPathRuleList {
	var returns ApplicationGatewayUrlPathMapPathRuleList
	_jsii_.Get(
		j,
		"pathRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) PathRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayUrlPathMapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayUrlPathMapOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayUrlPathMapOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayUrlPathMapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayUrlPathMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayUrlPathMapOutputReference_Override(a ApplicationGatewayUrlPathMapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayUrlPathMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetDefaultBackendAddressPoolName(val *string) {
	if err := j.validateSetDefaultBackendAddressPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBackendAddressPoolName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetDefaultBackendHttpSettingsName(val *string) {
	if err := j.validateSetDefaultBackendHttpSettingsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBackendHttpSettingsName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetDefaultRedirectConfigurationName(val *string) {
	if err := j.validateSetDefaultRedirectConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRedirectConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetDefaultRewriteRuleSetName(val *string) {
	if err := j.validateSetDefaultRewriteRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRewriteRuleSetName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) PutPathRule(value interface{}) {
	if err := a.validatePutPathRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPathRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ResetDefaultBackendAddressPoolName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultBackendAddressPoolName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ResetDefaultBackendHttpSettingsName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultBackendHttpSettingsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ResetDefaultRedirectConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRedirectConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ResetDefaultRewriteRuleSetName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRewriteRuleSetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

