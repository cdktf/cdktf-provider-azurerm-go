package networkconnectionmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/networkconnectionmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference interface {
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
	InternalValue() *NetworkConnectionMonitorTestConfigurationHttpConfiguration
	SetInternalValue(val *NetworkConnectionMonitorTestConfigurationHttpConfiguration)
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreferHttps() interface{}
	SetPreferHttps(val interface{})
	PreferHttpsInput() interface{}
	RequestHeader() NetworkConnectionMonitorTestConfigurationHttpConfigurationRequestHeaderList
	RequestHeaderInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidStatusCodeRanges() *[]*string
	SetValidStatusCodeRanges(val *[]*string)
	ValidStatusCodeRangesInput() *[]*string
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
	PutRequestHeader(value interface{})
	ResetMethod()
	ResetPath()
	ResetPort()
	ResetPreferHttps()
	ResetRequestHeader()
	ResetValidStatusCodeRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference
type jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) InternalValue() *NetworkConnectionMonitorTestConfigurationHttpConfiguration {
	var returns *NetworkConnectionMonitorTestConfigurationHttpConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) PreferHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) PreferHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) RequestHeader() NetworkConnectionMonitorTestConfigurationHttpConfigurationRequestHeaderList {
	var returns NetworkConnectionMonitorTestConfigurationHttpConfigurationRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) RequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ValidStatusCodeRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validStatusCodeRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ValidStatusCodeRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validStatusCodeRangesInput",
		&returns,
	)
	return returns
}


func NewNetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference_Override(n NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetInternalValue(val *NetworkConnectionMonitorTestConfigurationHttpConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetPreferHttps(val interface{}) {
	if err := j.validateSetPreferHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferHttps",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference)SetValidStatusCodeRanges(val *[]*string) {
	if err := j.validateSetValidStatusCodeRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validStatusCodeRanges",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) PutRequestHeader(value interface{}) {
	if err := n.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRequestHeader",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		n,
		"resetMethod",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		n,
		"resetPath",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		n,
		"resetPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetPreferHttps() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferHttps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		n,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ResetValidStatusCodeRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetValidStatusCodeRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

