package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/cdnfrontdoorrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference interface {
	cdktf.ComplexObject
	CacheBehavior() *string
	SetCacheBehavior(val *string)
	CacheBehaviorInput() *string
	CacheDuration() *string
	SetCacheDuration(val *string)
	CacheDurationInput() *string
	CdnFrontdoorOriginGroupId() *string
	SetCdnFrontdoorOriginGroupId(val *string)
	CdnFrontdoorOriginGroupIdInput() *string
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
	CompressionEnabled() interface{}
	SetCompressionEnabled(val interface{})
	CompressionEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForwardingProtocol() *string
	SetForwardingProtocol(val *string)
	ForwardingProtocolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction
	SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction)
	QueryStringCachingBehavior() *string
	SetQueryStringCachingBehavior(val *string)
	QueryStringCachingBehaviorInput() *string
	QueryStringParameters() *[]*string
	SetQueryStringParameters(val *[]*string)
	QueryStringParametersInput() *[]*string
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
	ResetCacheBehavior()
	ResetCompressionEnabled()
	ResetForwardingProtocol()
	ResetQueryStringCachingBehavior()
	ResetQueryStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CacheBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CacheBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CacheDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CacheDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CdnFrontdoorOriginGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CdnFrontdoorOriginGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ForwardingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ForwardingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) QueryStringCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) QueryStringCachingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) QueryStringParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) QueryStringParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference_Override(c CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetCacheBehavior(val *string) {
	if err := j.validateSetCacheBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBehavior",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetCacheDuration(val *string) {
	if err := j.validateSetCacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDuration",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetCdnFrontdoorOriginGroupId(val *string) {
	if err := j.validateSetCdnFrontdoorOriginGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginGroupId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetCompressionEnabled(val interface{}) {
	if err := j.validateSetCompressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetForwardingProtocol(val *string) {
	if err := j.validateSetForwardingProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardingProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetQueryStringCachingBehavior(val *string) {
	if err := j.validateSetQueryStringCachingBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetQueryStringParameters(val *[]*string) {
	if err := j.validateSetQueryStringParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringParameters",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ResetCacheBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ResetCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ResetForwardingProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardingProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ResetQueryStringCachingBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCachingBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ResetQueryStringParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

