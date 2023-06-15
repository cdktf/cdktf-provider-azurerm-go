package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/cdnfrontdoorrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleActionsOutputReference interface {
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
	InternalValue() *CdnFrontdoorRuleActions
	SetInternalValue(val *CdnFrontdoorRuleActions)
	RequestHeaderAction() CdnFrontdoorRuleActionsRequestHeaderActionList
	RequestHeaderActionInput() interface{}
	ResponseHeaderAction() CdnFrontdoorRuleActionsResponseHeaderActionList
	ResponseHeaderActionInput() interface{}
	RouteConfigurationOverrideAction() CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference
	RouteConfigurationOverrideActionInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirectAction() CdnFrontdoorRuleActionsUrlRedirectActionOutputReference
	UrlRedirectActionInput() *CdnFrontdoorRuleActionsUrlRedirectAction
	UrlRewriteAction() CdnFrontdoorRuleActionsUrlRewriteActionOutputReference
	UrlRewriteActionInput() *CdnFrontdoorRuleActionsUrlRewriteAction
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
	PutRequestHeaderAction(value interface{})
	PutResponseHeaderAction(value interface{})
	PutRouteConfigurationOverrideAction(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction)
	PutUrlRedirectAction(value *CdnFrontdoorRuleActionsUrlRedirectAction)
	PutUrlRewriteAction(value *CdnFrontdoorRuleActionsUrlRewriteAction)
	ResetRequestHeaderAction()
	ResetResponseHeaderAction()
	ResetRouteConfigurationOverrideAction()
	ResetUrlRedirectAction()
	ResetUrlRewriteAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InternalValue() *CdnFrontdoorRuleActions {
	var returns *CdnFrontdoorRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RequestHeaderAction() CdnFrontdoorRuleActionsRequestHeaderActionList {
	var returns CdnFrontdoorRuleActionsRequestHeaderActionList
	_jsii_.Get(
		j,
		"requestHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RequestHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResponseHeaderAction() CdnFrontdoorRuleActionsResponseHeaderActionList {
	var returns CdnFrontdoorRuleActionsResponseHeaderActionList
	_jsii_.Get(
		j,
		"responseHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResponseHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RouteConfigurationOverrideAction() CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference {
	var returns CdnFrontdoorRuleActionsRouteConfigurationOverrideActionOutputReference
	_jsii_.Get(
		j,
		"routeConfigurationOverrideAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RouteConfigurationOverrideActionInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction
	_jsii_.Get(
		j,
		"routeConfigurationOverrideActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRedirectAction() CdnFrontdoorRuleActionsUrlRedirectActionOutputReference {
	var returns CdnFrontdoorRuleActionsUrlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"urlRedirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRedirectActionInput() *CdnFrontdoorRuleActionsUrlRedirectAction {
	var returns *CdnFrontdoorRuleActionsUrlRedirectAction
	_jsii_.Get(
		j,
		"urlRedirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRewriteAction() CdnFrontdoorRuleActionsUrlRewriteActionOutputReference {
	var returns CdnFrontdoorRuleActionsUrlRewriteActionOutputReference
	_jsii_.Get(
		j,
		"urlRewriteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRewriteActionInput() *CdnFrontdoorRuleActionsUrlRewriteAction {
	var returns *CdnFrontdoorRuleActionsUrlRewriteAction
	_jsii_.Get(
		j,
		"urlRewriteActionInput",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsOutputReference_Override(c CdnFrontdoorRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetInternalValue(val *CdnFrontdoorRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutRequestHeaderAction(value interface{}) {
	if err := c.validatePutRequestHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutResponseHeaderAction(value interface{}) {
	if err := c.validatePutResponseHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResponseHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutRouteConfigurationOverrideAction(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction) {
	if err := c.validatePutRouteConfigurationOverrideActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouteConfigurationOverrideAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutUrlRedirectAction(value *CdnFrontdoorRuleActionsUrlRedirectAction) {
	if err := c.validatePutUrlRedirectActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirectAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutUrlRewriteAction(value *CdnFrontdoorRuleActionsUrlRewriteAction) {
	if err := c.validatePutUrlRewriteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewriteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetRequestHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetResponseHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetRouteConfigurationOverrideAction() {
	_jsii_.InvokeVoid(
		c,
		"resetRouteConfigurationOverrideAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetUrlRedirectAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirectAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetUrlRewriteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewriteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

