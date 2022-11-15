package networkconnectionmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/networkconnectionmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectionMonitorEndpointOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
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
	CoverageLevel() *string
	SetCoverageLevel(val *string)
	CoverageLevelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedIpAddresses() *[]*string
	SetExcludedIpAddresses(val *[]*string)
	ExcludedIpAddressesInput() *[]*string
	Filter() NetworkConnectionMonitorEndpointFilterOutputReference
	FilterInput() *NetworkConnectionMonitorEndpointFilter
	// Experimental.
	Fqn() *string
	IncludedIpAddresses() *[]*string
	SetIncludedIpAddresses(val *[]*string)
	IncludedIpAddressesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	TargetResourceId() *string
	SetTargetResourceId(val *string)
	TargetResourceIdInput() *string
	TargetResourceType() *string
	SetTargetResourceType(val *string)
	TargetResourceTypeInput() *string
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
	PutFilter(value *NetworkConnectionMonitorEndpointFilter)
	ResetAddress()
	ResetCoverageLevel()
	ResetExcludedIpAddresses()
	ResetFilter()
	ResetIncludedIpAddresses()
	ResetTargetResourceId()
	ResetTargetResourceType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectionMonitorEndpointOutputReference
type jsiiProxy_NetworkConnectionMonitorEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) CoverageLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coverageLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) CoverageLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coverageLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ExcludedIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ExcludedIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) Filter() NetworkConnectionMonitorEndpointFilterOutputReference {
	var returns NetworkConnectionMonitorEndpointFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) FilterInput() *NetworkConnectionMonitorEndpointFilter {
	var returns *NetworkConnectionMonitorEndpointFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) IncludedIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) IncludedIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TargetResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TargetResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TargetResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TargetResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkConnectionMonitorEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkConnectionMonitorEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectionMonitorEndpointOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectionMonitorEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkConnectionMonitorEndpointOutputReference_Override(n NetworkConnectionMonitorEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkConnectionMonitor.NetworkConnectionMonitorEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetCoverageLevel(val *string) {
	if err := j.validateSetCoverageLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coverageLevel",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetExcludedIpAddresses(val *[]*string) {
	if err := j.validateSetExcludedIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpAddresses",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetIncludedIpAddresses(val *[]*string) {
	if err := j.validateSetIncludedIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedIpAddresses",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetTargetResourceId(val *string) {
	if err := j.validateSetTargetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceId",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetTargetResourceType(val *string) {
	if err := j.validateSetTargetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceType",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) PutFilter(value *NetworkConnectionMonitorEndpointFilter) {
	if err := n.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetCoverageLevel() {
	_jsii_.InvokeVoid(
		n,
		"resetCoverageLevel",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetExcludedIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludedIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		n,
		"resetFilter",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetIncludedIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludedIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetTargetResourceId() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetResourceId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ResetTargetResourceType() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetResourceType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkConnectionMonitorEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

