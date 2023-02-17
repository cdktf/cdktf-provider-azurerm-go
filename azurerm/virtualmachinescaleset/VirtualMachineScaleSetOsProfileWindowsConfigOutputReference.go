package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetOsProfileWindowsConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalUnattendConfig() VirtualMachineScaleSetOsProfileWindowsConfigAdditionalUnattendConfigList
	AdditionalUnattendConfigInput() interface{}
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
	EnableAutomaticUpgrades() interface{}
	SetEnableAutomaticUpgrades(val interface{})
	EnableAutomaticUpgradesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VirtualMachineScaleSetOsProfileWindowsConfig
	SetInternalValue(val *VirtualMachineScaleSetOsProfileWindowsConfig)
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Winrm() VirtualMachineScaleSetOsProfileWindowsConfigWinrmList
	WinrmInput() interface{}
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
	PutAdditionalUnattendConfig(value interface{})
	PutWinrm(value interface{})
	ResetAdditionalUnattendConfig()
	ResetEnableAutomaticUpgrades()
	ResetProvisionVmAgent()
	ResetWinrm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetOsProfileWindowsConfigOutputReference
type jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) AdditionalUnattendConfig() VirtualMachineScaleSetOsProfileWindowsConfigAdditionalUnattendConfigList {
	var returns VirtualMachineScaleSetOsProfileWindowsConfigAdditionalUnattendConfigList
	_jsii_.Get(
		j,
		"additionalUnattendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) AdditionalUnattendConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) EnableAutomaticUpgrades() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgrades",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) EnableAutomaticUpgradesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgradesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) InternalValue() *VirtualMachineScaleSetOsProfileWindowsConfig {
	var returns *VirtualMachineScaleSetOsProfileWindowsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) Winrm() VirtualMachineScaleSetOsProfileWindowsConfigWinrmList {
	var returns VirtualMachineScaleSetOsProfileWindowsConfigWinrmList
	_jsii_.Get(
		j,
		"winrm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) WinrmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetOsProfileWindowsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineScaleSetOsProfileWindowsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetOsProfileWindowsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetOsProfileWindowsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetOsProfileWindowsConfigOutputReference_Override(v VirtualMachineScaleSetOsProfileWindowsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetOsProfileWindowsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetEnableAutomaticUpgrades(val interface{}) {
	if err := j.validateSetEnableAutomaticUpgradesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpgrades",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetInternalValue(val *VirtualMachineScaleSetOsProfileWindowsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) PutAdditionalUnattendConfig(value interface{}) {
	if err := v.validatePutAdditionalUnattendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAdditionalUnattendConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) PutWinrm(value interface{}) {
	if err := v.validatePutWinrmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putWinrm",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ResetAdditionalUnattendConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetAdditionalUnattendConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ResetEnableAutomaticUpgrades() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableAutomaticUpgrades",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		v,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ResetWinrm() {
	_jsii_.InvokeVoid(
		v,
		"resetWinrm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetOsProfileWindowsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

