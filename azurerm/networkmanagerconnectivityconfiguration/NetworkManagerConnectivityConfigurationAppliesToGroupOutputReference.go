package networkmanagerconnectivityconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/networkmanagerconnectivityconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference interface {
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
	GlobalMeshEnabled() interface{}
	SetGlobalMeshEnabled(val interface{})
	GlobalMeshEnabledInput() interface{}
	GroupConnectivity() *string
	SetGroupConnectivity(val *string)
	GroupConnectivityInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkGroupId() *string
	SetNetworkGroupId(val *string)
	NetworkGroupIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseHubGateway() interface{}
	SetUseHubGateway(val interface{})
	UseHubGatewayInput() interface{}
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
	ResetGlobalMeshEnabled()
	ResetUseHubGateway()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference
type jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GlobalMeshEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalMeshEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GlobalMeshEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalMeshEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GroupConnectivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GroupConnectivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) NetworkGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) NetworkGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) UseHubGateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useHubGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) UseHubGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useHubGatewayInput",
		&returns,
	)
	return returns
}


func NewNetworkManagerConnectivityConfigurationAppliesToGroupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkManagerConnectivityConfigurationAppliesToGroupOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkManagerConnectivityConfiguration.NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkManagerConnectivityConfigurationAppliesToGroupOutputReference_Override(n NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkManagerConnectivityConfiguration.NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetGlobalMeshEnabled(val interface{}) {
	if err := j.validateSetGlobalMeshEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalMeshEnabled",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetGroupConnectivity(val *string) {
	if err := j.validateSetGroupConnectivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupConnectivity",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetNetworkGroupId(val *string) {
	if err := j.validateSetNetworkGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkGroupId",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference)SetUseHubGateway(val interface{}) {
	if err := j.validateSetUseHubGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHubGateway",
		val,
	)
}

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ResetGlobalMeshEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetGlobalMeshEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ResetUseHubGateway() {
	_jsii_.InvokeVoid(
		n,
		"resetUseHubGateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkManagerConnectivityConfigurationAppliesToGroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

