package cdnfrontdoororigingroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/cdnfrontdoororigingroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorOriginGroupLoadBalancingOutputReference interface {
	cdktf.ComplexObject
	AdditionalLatencyInMilliseconds() *float64
	SetAdditionalLatencyInMilliseconds(val *float64)
	AdditionalLatencyInMillisecondsInput() *float64
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
	InternalValue() *CdnFrontdoorOriginGroupLoadBalancing
	SetInternalValue(val *CdnFrontdoorOriginGroupLoadBalancing)
	SampleSize() *float64
	SetSampleSize(val *float64)
	SampleSizeInput() *float64
	SuccessfulSamplesRequired() *float64
	SetSuccessfulSamplesRequired(val *float64)
	SuccessfulSamplesRequiredInput() *float64
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
	ResetAdditionalLatencyInMilliseconds()
	ResetSampleSize()
	ResetSuccessfulSamplesRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorOriginGroupLoadBalancingOutputReference
type jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) AdditionalLatencyInMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalLatencyInMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) AdditionalLatencyInMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalLatencyInMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) InternalValue() *CdnFrontdoorOriginGroupLoadBalancing {
	var returns *CdnFrontdoorOriginGroupLoadBalancing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) SuccessfulSamplesRequired() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulSamplesRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) SuccessfulSamplesRequiredInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulSamplesRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorOriginGroupLoadBalancingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorOriginGroupLoadBalancingOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorOriginGroupLoadBalancingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorOriginGroup.CdnFrontdoorOriginGroupLoadBalancingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorOriginGroupLoadBalancingOutputReference_Override(c CdnFrontdoorOriginGroupLoadBalancingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorOriginGroup.CdnFrontdoorOriginGroupLoadBalancingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetAdditionalLatencyInMilliseconds(val *float64) {
	if err := j.validateSetAdditionalLatencyInMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLatencyInMilliseconds",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetInternalValue(val *CdnFrontdoorOriginGroupLoadBalancing) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetSampleSize(val *float64) {
	if err := j.validateSetSampleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetSuccessfulSamplesRequired(val *float64) {
	if err := j.validateSetSuccessfulSamplesRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successfulSamplesRequired",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ResetAdditionalLatencyInMilliseconds() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalLatencyInMilliseconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ResetSuccessfulSamplesRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessfulSamplesRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorOriginGroupLoadBalancingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

