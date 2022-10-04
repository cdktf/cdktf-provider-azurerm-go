package networkconnectionmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/networkconnectionmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectionMonitorTestConfigurationOutputReference interface {
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
	HttpConfiguration() NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference
	HttpConfigurationInput() *NetworkConnectionMonitorTestConfigurationHttpConfiguration
	IcmpConfiguration() NetworkConnectionMonitorTestConfigurationIcmpConfigurationOutputReference
	IcmpConfigurationInput() *NetworkConnectionMonitorTestConfigurationIcmpConfiguration
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PreferredIpVersion() *string
	SetPreferredIpVersion(val *string)
	PreferredIpVersionInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SuccessThreshold() NetworkConnectionMonitorTestConfigurationSuccessThresholdOutputReference
	SuccessThresholdInput() *NetworkConnectionMonitorTestConfigurationSuccessThreshold
	TcpConfiguration() NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference
	TcpConfigurationInput() *NetworkConnectionMonitorTestConfigurationTcpConfiguration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TestFrequencyInSeconds() *float64
	SetTestFrequencyInSeconds(val *float64)
	TestFrequencyInSecondsInput() *float64
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
	PutHttpConfiguration(value *NetworkConnectionMonitorTestConfigurationHttpConfiguration)
	PutIcmpConfiguration(value *NetworkConnectionMonitorTestConfigurationIcmpConfiguration)
	PutSuccessThreshold(value *NetworkConnectionMonitorTestConfigurationSuccessThreshold)
	PutTcpConfiguration(value *NetworkConnectionMonitorTestConfigurationTcpConfiguration)
	ResetHttpConfiguration()
	ResetIcmpConfiguration()
	ResetPreferredIpVersion()
	ResetSuccessThreshold()
	ResetTcpConfiguration()
	ResetTestFrequencyInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectionMonitorTestConfigurationOutputReference
type jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) HttpConfiguration() NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference {
	var returns NetworkConnectionMonitorTestConfigurationHttpConfigurationOutputReference
	_jsii_.Get(
		j,
		"httpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) HttpConfigurationInput() *NetworkConnectionMonitorTestConfigurationHttpConfiguration {
	var returns *NetworkConnectionMonitorTestConfigurationHttpConfiguration
	_jsii_.Get(
		j,
		"httpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) IcmpConfiguration() NetworkConnectionMonitorTestConfigurationIcmpConfigurationOutputReference {
	var returns NetworkConnectionMonitorTestConfigurationIcmpConfigurationOutputReference
	_jsii_.Get(
		j,
		"icmpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) IcmpConfigurationInput() *NetworkConnectionMonitorTestConfigurationIcmpConfiguration {
	var returns *NetworkConnectionMonitorTestConfigurationIcmpConfiguration
	_jsii_.Get(
		j,
		"icmpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PreferredIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PreferredIpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredIpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) SuccessThreshold() NetworkConnectionMonitorTestConfigurationSuccessThresholdOutputReference {
	var returns NetworkConnectionMonitorTestConfigurationSuccessThresholdOutputReference
	_jsii_.Get(
		j,
		"successThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) SuccessThresholdInput() *NetworkConnectionMonitorTestConfigurationSuccessThreshold {
	var returns *NetworkConnectionMonitorTestConfigurationSuccessThreshold
	_jsii_.Get(
		j,
		"successThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TcpConfiguration() NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference {
	var returns NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference
	_jsii_.Get(
		j,
		"tcpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TcpConfigurationInput() *NetworkConnectionMonitorTestConfigurationTcpConfiguration {
	var returns *NetworkConnectionMonitorTestConfigurationTcpConfiguration
	_jsii_.Get(
		j,
		"tcpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TestFrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"testFrequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) TestFrequencyInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"testFrequencyInSecondsInput",
		&returns,
	)
	return returns
}


func NewNetworkConnectionMonitorTestConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkConnectionMonitorTestConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectionMonitorTestConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkConnectionMonitorTestConfigurationOutputReference_Override(n NetworkConnectionMonitorTestConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetPreferredIpVersion(val *string) {
	if err := j.validateSetPreferredIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredIpVersion",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference)SetTestFrequencyInSeconds(val *float64) {
	if err := j.validateSetTestFrequencyInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testFrequencyInSeconds",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PutHttpConfiguration(value *NetworkConnectionMonitorTestConfigurationHttpConfiguration) {
	if err := n.validatePutHttpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHttpConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PutIcmpConfiguration(value *NetworkConnectionMonitorTestConfigurationIcmpConfiguration) {
	if err := n.validatePutIcmpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putIcmpConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PutSuccessThreshold(value *NetworkConnectionMonitorTestConfigurationSuccessThreshold) {
	if err := n.validatePutSuccessThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSuccessThreshold",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) PutTcpConfiguration(value *NetworkConnectionMonitorTestConfigurationTcpConfiguration) {
	if err := n.validatePutTcpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTcpConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetHttpConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetHttpConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetIcmpConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetIcmpConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetPreferredIpVersion() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferredIpVersion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetSuccessThreshold() {
	_jsii_.InvokeVoid(
		n,
		"resetSuccessThreshold",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetTcpConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetTcpConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ResetTestFrequencyInSeconds() {
	_jsii_.InvokeVoid(
		n,
		"resetTestFrequencyInSeconds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

