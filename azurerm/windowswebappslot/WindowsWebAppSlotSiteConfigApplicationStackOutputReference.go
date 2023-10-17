// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/windowswebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotSiteConfigApplicationStackOutputReference interface {
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
	CurrentStack() *string
	SetCurrentStack(val *string)
	CurrentStackInput() *string
	DockerContainerName() *string
	SetDockerContainerName(val *string)
	DockerContainerNameInput() *string
	DockerContainerRegistry() *string
	SetDockerContainerRegistry(val *string)
	DockerContainerRegistryInput() *string
	DockerContainerTag() *string
	SetDockerContainerTag(val *string)
	DockerContainerTagInput() *string
	DockerImageName() *string
	SetDockerImageName(val *string)
	DockerImageNameInput() *string
	DockerRegistryPassword() *string
	SetDockerRegistryPassword(val *string)
	DockerRegistryPasswordInput() *string
	DockerRegistryUrl() *string
	SetDockerRegistryUrl(val *string)
	DockerRegistryUrlInput() *string
	DockerRegistryUsername() *string
	SetDockerRegistryUsername(val *string)
	DockerRegistryUsernameInput() *string
	DotnetCoreVersion() *string
	SetDotnetCoreVersion(val *string)
	DotnetCoreVersionInput() *string
	DotnetVersion() *string
	SetDotnetVersion(val *string)
	DotnetVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSlotSiteConfigApplicationStack
	SetInternalValue(val *WindowsWebAppSlotSiteConfigApplicationStack)
	JavaContainer() *string
	SetJavaContainer(val *string)
	JavaContainerInput() *string
	JavaContainerVersion() *string
	SetJavaContainerVersion(val *string)
	JavaContainerVersionInput() *string
	JavaEmbeddedServerEnabled() interface{}
	SetJavaEmbeddedServerEnabled(val interface{})
	JavaEmbeddedServerEnabledInput() interface{}
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	PhpVersion() *string
	SetPhpVersion(val *string)
	PhpVersionInput() *string
	Python() interface{}
	SetPython(val interface{})
	PythonInput() interface{}
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TomcatVersion() *string
	SetTomcatVersion(val *string)
	TomcatVersionInput() *string
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
	ResetCurrentStack()
	ResetDockerContainerName()
	ResetDockerContainerRegistry()
	ResetDockerContainerTag()
	ResetDockerImageName()
	ResetDockerRegistryPassword()
	ResetDockerRegistryUrl()
	ResetDockerRegistryUsername()
	ResetDotnetCoreVersion()
	ResetDotnetVersion()
	ResetJavaContainer()
	ResetJavaContainerVersion()
	ResetJavaEmbeddedServerEnabled()
	ResetJavaVersion()
	ResetNodeVersion()
	ResetPhpVersion()
	ResetPython()
	ResetPythonVersion()
	ResetTomcatVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSlotSiteConfigApplicationStackOutputReference
type jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) CurrentStack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) CurrentStackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerContainerTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerContainerTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DockerRegistryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DotnetCoreVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetCoreVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DotnetCoreVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetCoreVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DotnetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) DotnetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) InternalValue() *WindowsWebAppSlotSiteConfigApplicationStack {
	var returns *WindowsWebAppSlotSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaContainerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaContainerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaEmbeddedServerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javaEmbeddedServerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaEmbeddedServerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javaEmbeddedServerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) PhpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) PhpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) Python() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"python",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) PythonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pythonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) TomcatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tomcatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) TomcatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tomcatVersionInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSlotSiteConfigApplicationStackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSlotSiteConfigApplicationStackOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppSlotSiteConfigApplicationStackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSlotSiteConfigApplicationStackOutputReference_Override(w WindowsWebAppSlotSiteConfigApplicationStackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetCurrentStack(val *string) {
	if err := j.validateSetCurrentStackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentStack",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerContainerName(val *string) {
	if err := j.validateSetDockerContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerContainerName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerContainerRegistry(val *string) {
	if err := j.validateSetDockerContainerRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerContainerRegistry",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerContainerTag(val *string) {
	if err := j.validateSetDockerContainerTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerContainerTag",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerImageName(val *string) {
	if err := j.validateSetDockerImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerImageName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerRegistryPassword(val *string) {
	if err := j.validateSetDockerRegistryPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryPassword",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerRegistryUrl(val *string) {
	if err := j.validateSetDockerRegistryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDockerRegistryUsername(val *string) {
	if err := j.validateSetDockerRegistryUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryUsername",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDotnetCoreVersion(val *string) {
	if err := j.validateSetDotnetCoreVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetCoreVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetDotnetVersion(val *string) {
	if err := j.validateSetDotnetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetInternalValue(val *WindowsWebAppSlotSiteConfigApplicationStack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetJavaContainer(val *string) {
	if err := j.validateSetJavaContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaContainer",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetJavaContainerVersion(val *string) {
	if err := j.validateSetJavaContainerVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaContainerVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetJavaEmbeddedServerEnabled(val interface{}) {
	if err := j.validateSetJavaEmbeddedServerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaEmbeddedServerEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetJavaVersion(val *string) {
	if err := j.validateSetJavaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetPhpVersion(val *string) {
	if err := j.validateSetPhpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phpVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetPython(val interface{}) {
	if err := j.validateSetPythonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"python",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetPythonVersion(val *string) {
	if err := j.validateSetPythonVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference)SetTomcatVersion(val *string) {
	if err := j.validateSetTomcatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tomcatVersion",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetCurrentStack() {
	_jsii_.InvokeVoid(
		w,
		"resetCurrentStack",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerContainerName() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerContainerRegistry() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerRegistry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerContainerTag() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerContainerTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerImageName() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerImageName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerRegistryPassword() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerRegistryPassword",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerRegistryUrl() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerRegistryUrl",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDockerRegistryUsername() {
	_jsii_.InvokeVoid(
		w,
		"resetDockerRegistryUsername",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDotnetCoreVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetDotnetCoreVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetDotnetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetDotnetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetJavaContainer() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaContainer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetJavaContainerVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaContainerVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetJavaEmbeddedServerEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaEmbeddedServerEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetPhpVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetPhpVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetPython() {
	_jsii_.InvokeVoid(
		w,
		"resetPython",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ResetTomcatVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetTomcatVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigApplicationStackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

