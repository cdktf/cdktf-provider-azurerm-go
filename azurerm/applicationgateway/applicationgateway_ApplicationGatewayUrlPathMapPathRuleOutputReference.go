package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayUrlPathMapPathRuleOutputReference interface {
	cdktf.ComplexObject
	BackendAddressPoolId() *string
	BackendAddressPoolName() *string
	SetBackendAddressPoolName(val *string)
	BackendAddressPoolNameInput() *string
	BackendHttpSettingsId() *string
	BackendHttpSettingsName() *string
	SetBackendHttpSettingsName(val *string)
	BackendHttpSettingsNameInput() *string
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
	FirewallPolicyId() *string
	SetFirewallPolicyId(val *string)
	FirewallPolicyIdInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Paths() *[]*string
	SetPaths(val *[]*string)
	PathsInput() *[]*string
	RedirectConfigurationId() *string
	RedirectConfigurationName() *string
	SetRedirectConfigurationName(val *string)
	RedirectConfigurationNameInput() *string
	RewriteRuleSetId() *string
	RewriteRuleSetName() *string
	SetRewriteRuleSetName(val *string)
	RewriteRuleSetNameInput() *string
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
	ResetBackendAddressPoolName()
	ResetBackendHttpSettingsName()
	ResetFirewallPolicyId()
	ResetRedirectConfigurationName()
	ResetRewriteRuleSetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayUrlPathMapPathRuleOutputReference
type jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendAddressPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendAddressPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendAddressPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendAddressPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendHttpSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendHttpSettingsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) BackendHttpSettingsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendHttpSettingsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) FirewallPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) FirewallPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) Paths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) PathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RedirectConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RedirectConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RedirectConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RewriteRuleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RewriteRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) RewriteRuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteRuleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayUrlPathMapPathRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayUrlPathMapPathRuleOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayUrlPathMapPathRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayUrlPathMapPathRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayUrlPathMapPathRuleOutputReference_Override(a ApplicationGatewayUrlPathMapPathRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayUrlPathMapPathRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetBackendAddressPoolName(val *string) {
	if err := j.validateSetBackendAddressPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendAddressPoolName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetBackendHttpSettingsName(val *string) {
	if err := j.validateSetBackendHttpSettingsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendHttpSettingsName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetFirewallPolicyId(val *string) {
	if err := j.validateSetFirewallPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyId",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetPaths(val *[]*string) {
	if err := j.validateSetPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paths",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetRedirectConfigurationName(val *string) {
	if err := j.validateSetRedirectConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetRewriteRuleSetName(val *string) {
	if err := j.validateSetRewriteRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rewriteRuleSetName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ResetBackendAddressPoolName() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendAddressPoolName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ResetBackendHttpSettingsName() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendHttpSettingsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ResetFirewallPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ResetRedirectConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ResetRewriteRuleSetName() {
	_jsii_.InvokeVoid(
		a,
		"resetRewriteRuleSetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewayUrlPathMapPathRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

