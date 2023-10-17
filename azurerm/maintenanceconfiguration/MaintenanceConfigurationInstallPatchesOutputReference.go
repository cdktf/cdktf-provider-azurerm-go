// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/maintenanceconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaintenanceConfigurationInstallPatchesOutputReference interface {
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
	InternalValue() *MaintenanceConfigurationInstallPatches
	SetInternalValue(val *MaintenanceConfigurationInstallPatches)
	Linux() MaintenanceConfigurationInstallPatchesLinuxList
	LinuxInput() interface{}
	Reboot() *string
	SetReboot(val *string)
	RebootInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Windows() MaintenanceConfigurationInstallPatchesWindowsList
	WindowsInput() interface{}
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
	PutLinux(value interface{})
	PutWindows(value interface{})
	ResetLinux()
	ResetReboot()
	ResetWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MaintenanceConfigurationInstallPatchesOutputReference
type jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) InternalValue() *MaintenanceConfigurationInstallPatches {
	var returns *MaintenanceConfigurationInstallPatches
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) Linux() MaintenanceConfigurationInstallPatchesLinuxList {
	var returns MaintenanceConfigurationInstallPatchesLinuxList
	_jsii_.Get(
		j,
		"linux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) LinuxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) Reboot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) RebootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) Windows() MaintenanceConfigurationInstallPatchesWindowsList {
	var returns MaintenanceConfigurationInstallPatchesWindowsList
	_jsii_.Get(
		j,
		"windows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) WindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsInput",
		&returns,
	)
	return returns
}


func NewMaintenanceConfigurationInstallPatchesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MaintenanceConfigurationInstallPatchesOutputReference {
	_init_.Initialize()

	if err := validateNewMaintenanceConfigurationInstallPatchesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMaintenanceConfigurationInstallPatchesOutputReference_Override(m MaintenanceConfigurationInstallPatchesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetInternalValue(val *MaintenanceConfigurationInstallPatches) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetReboot(val *string) {
	if err := j.validateSetRebootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reboot",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) PutLinux(value interface{}) {
	if err := m.validatePutLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLinux",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) PutWindows(value interface{}) {
	if err := m.validatePutWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWindows",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ResetLinux() {
	_jsii_.InvokeVoid(
		m,
		"resetLinux",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ResetReboot() {
	_jsii_.InvokeVoid(
		m,
		"resetReboot",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ResetWindows() {
	_jsii_.InvokeVoid(
		m,
		"resetWindows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

