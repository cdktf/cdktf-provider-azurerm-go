package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineOsProfileWindowsConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalUnattendConfig() VirtualMachineOsProfileWindowsConfigAdditionalUnattendConfigList
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
	InternalValue() *VirtualMachineOsProfileWindowsConfig
	SetInternalValue(val *VirtualMachineOsProfileWindowsConfig)
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
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Winrm() VirtualMachineOsProfileWindowsConfigWinrmList
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
	ResetTimezone()
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

// The jsii proxy struct for VirtualMachineOsProfileWindowsConfigOutputReference
type jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) AdditionalUnattendConfig() VirtualMachineOsProfileWindowsConfigAdditionalUnattendConfigList {
	var returns VirtualMachineOsProfileWindowsConfigAdditionalUnattendConfigList
	_jsii_.Get(
		j,
		"additionalUnattendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) AdditionalUnattendConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) EnableAutomaticUpgrades() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgrades",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) EnableAutomaticUpgradesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpgradesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) InternalValue() *VirtualMachineOsProfileWindowsConfig {
	var returns *VirtualMachineOsProfileWindowsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) Winrm() VirtualMachineOsProfileWindowsConfigWinrmList {
	var returns VirtualMachineOsProfileWindowsConfigWinrmList
	_jsii_.Get(
		j,
		"winrm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) WinrmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineOsProfileWindowsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineOsProfileWindowsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineOsProfileWindowsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachineOsProfileWindowsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineOsProfileWindowsConfigOutputReference_Override(v VirtualMachineOsProfileWindowsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachine.VirtualMachineOsProfileWindowsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetEnableAutomaticUpgrades(val interface{}) {
	if err := j.validateSetEnableAutomaticUpgradesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpgrades",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetInternalValue(val *VirtualMachineOsProfileWindowsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) PutAdditionalUnattendConfig(value interface{}) {
	if err := v.validatePutAdditionalUnattendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAdditionalUnattendConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) PutWinrm(value interface{}) {
	if err := v.validatePutWinrmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putWinrm",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ResetAdditionalUnattendConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetAdditionalUnattendConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ResetEnableAutomaticUpgrades() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableAutomaticUpgrades",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		v,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		v,
		"resetTimezone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ResetWinrm() {
	_jsii_.InvokeVoid(
		v,
		"resetWinrm",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineOsProfileWindowsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

