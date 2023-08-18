package mssqlvirtualmachineavailabilitygrouplistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/mssqlvirtualmachineavailabilitygrouplistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList
type jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachineAvailabilityGroupListener.MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList_Override(m MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachineAvailabilityGroupListener.MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) Get(index *float64) MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAvailabilityGroupListenerMultiSubnetIpConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

