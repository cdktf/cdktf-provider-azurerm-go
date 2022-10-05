package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference interface {
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
	Condition() ApplicationGatewayRewriteRuleSetRewriteRuleConditionList
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	RequestHeaderConfiguration() ApplicationGatewayRewriteRuleSetRewriteRuleRequestHeaderConfigurationList
	RequestHeaderConfigurationInput() interface{}
	ResponseHeaderConfiguration() ApplicationGatewayRewriteRuleSetRewriteRuleResponseHeaderConfigurationList
	ResponseHeaderConfigurationInput() interface{}
	RuleSequence() *float64
	SetRuleSequence(val *float64)
	RuleSequenceInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference
	UrlInput() *ApplicationGatewayRewriteRuleSetRewriteRuleUrl
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
	PutCondition(value interface{})
	PutRequestHeaderConfiguration(value interface{})
	PutResponseHeaderConfiguration(value interface{})
	PutUrl(value *ApplicationGatewayRewriteRuleSetRewriteRuleUrl)
	ResetCondition()
	ResetRequestHeaderConfiguration()
	ResetResponseHeaderConfiguration()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference
type jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) Condition() ApplicationGatewayRewriteRuleSetRewriteRuleConditionList {
	var returns ApplicationGatewayRewriteRuleSetRewriteRuleConditionList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) RequestHeaderConfiguration() ApplicationGatewayRewriteRuleSetRewriteRuleRequestHeaderConfigurationList {
	var returns ApplicationGatewayRewriteRuleSetRewriteRuleRequestHeaderConfigurationList
	_jsii_.Get(
		j,
		"requestHeaderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) RequestHeaderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResponseHeaderConfiguration() ApplicationGatewayRewriteRuleSetRewriteRuleResponseHeaderConfigurationList {
	var returns ApplicationGatewayRewriteRuleSetRewriteRuleResponseHeaderConfigurationList
	_jsii_.Get(
		j,
		"responseHeaderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResponseHeaderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeaderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) RuleSequence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) RuleSequenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) Url() ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference {
	var returns ApplicationGatewayRewriteRuleSetRewriteRuleUrlOutputReference
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) UrlInput() *ApplicationGatewayRewriteRuleSetRewriteRuleUrl {
	var returns *ApplicationGatewayRewriteRuleSetRewriteRuleUrl
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewApplicationGatewayRewriteRuleSetRewriteRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayRewriteRuleSetRewriteRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayRewriteRuleSetRewriteRuleOutputReference_Override(a ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetRuleSequence(val *float64) {
	if err := j.validateSetRuleSequenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSequence",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) PutCondition(value interface{}) {
	if err := a.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) PutRequestHeaderConfiguration(value interface{}) {
	if err := a.validatePutRequestHeaderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestHeaderConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) PutResponseHeaderConfiguration(value interface{}) {
	if err := a.validatePutResponseHeaderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponseHeaderConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) PutUrl(value *ApplicationGatewayRewriteRuleSetRewriteRuleUrl) {
	if err := a.validatePutUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUrl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResetRequestHeaderConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestHeaderConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResetResponseHeaderConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseHeaderConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetRewriteRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

