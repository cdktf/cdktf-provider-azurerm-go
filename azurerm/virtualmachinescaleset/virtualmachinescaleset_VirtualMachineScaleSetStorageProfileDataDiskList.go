package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetStorageProfileDataDiskList interface {
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
	Get(index *float64) VirtualMachineScaleSetStorageProfileDataDiskOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetStorageProfileDataDiskList
type jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetStorageProfileDataDiskList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualMachineScaleSetStorageProfileDataDiskList {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetStorageProfileDataDiskListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetStorageProfileDataDiskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetStorageProfileDataDiskList_Override(v VirtualMachineScaleSetStorageProfileDataDiskList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetStorageProfileDataDiskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) Get(index *float64) VirtualMachineScaleSetStorageProfileDataDiskOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VirtualMachineScaleSetStorageProfileDataDiskOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileDataDiskList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

