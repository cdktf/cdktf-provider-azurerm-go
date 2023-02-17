package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetStorageProfileOsDiskOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	CreateOption() *string
	SetCreateOption(val *string)
	CreateOptionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() *VirtualMachineScaleSetStorageProfileOsDisk
	SetInternalValue(val *VirtualMachineScaleSetStorageProfileOsDisk)
	ManagedDiskType() *string
	SetManagedDiskType(val *string)
	ManagedDiskTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VhdContainers() *[]*string
	SetVhdContainers(val *[]*string)
	VhdContainersInput() *[]*string
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
	ResetCaching()
	ResetImage()
	ResetManagedDiskType()
	ResetName()
	ResetOsType()
	ResetVhdContainers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetStorageProfileOsDiskOutputReference
type jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) CreateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) CreateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) InternalValue() *VirtualMachineScaleSetStorageProfileOsDisk {
	var returns *VirtualMachineScaleSetStorageProfileOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ManagedDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ManagedDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) VhdContainers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vhdContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) VhdContainersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vhdContainersInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetStorageProfileOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineScaleSetStorageProfileOsDiskOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetStorageProfileOsDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetStorageProfileOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetStorageProfileOsDiskOutputReference_Override(v VirtualMachineScaleSetStorageProfileOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetStorageProfileOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetCreateOption(val *string) {
	if err := j.validateSetCreateOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createOption",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetInternalValue(val *VirtualMachineScaleSetStorageProfileOsDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetManagedDiskType(val *string) {
	if err := j.validateSetManagedDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference)SetVhdContainers(val *[]*string) {
	if err := j.validateSetVhdContainersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vhdContainers",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetCaching() {
	_jsii_.InvokeVoid(
		v,
		"resetCaching",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		v,
		"resetImage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetManagedDiskType() {
	_jsii_.InvokeVoid(
		v,
		"resetManagedDiskType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetOsType() {
	_jsii_.InvokeVoid(
		v,
		"resetOsType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ResetVhdContainers() {
	_jsii_.InvokeVoid(
		v,
		"resetVhdContainers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetStorageProfileOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

