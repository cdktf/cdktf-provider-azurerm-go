// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxwebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxWebAppSiteConfigApplicationStackOutputReference interface {
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
	DotnetVersion() *string
	SetDotnetVersion(val *string)
	DotnetVersionInput() *string
	// Experimental.
	Fqn() *string
	GoVersion() *string
	SetGoVersion(val *string)
	GoVersionInput() *string
	InternalValue() *LinuxWebAppSiteConfigApplicationStack
	SetInternalValue(val *LinuxWebAppSiteConfigApplicationStack)
	JavaServer() *string
	SetJavaServer(val *string)
	JavaServerInput() *string
	JavaServerVersion() *string
	SetJavaServerVersion(val *string)
	JavaServerVersionInput() *string
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	PhpVersion() *string
	SetPhpVersion(val *string)
	PhpVersionInput() *string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	RubyVersion() *string
	SetRubyVersion(val *string)
	RubyVersionInput() *string
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
	ResetDockerImageName()
	ResetDockerRegistryPassword()
	ResetDockerRegistryUrl()
	ResetDockerRegistryUsername()
	ResetDotnetVersion()
	ResetGoVersion()
	ResetJavaServer()
	ResetJavaServerVersion()
	ResetJavaVersion()
	ResetNodeVersion()
	ResetPhpVersion()
	ResetPythonVersion()
	ResetRubyVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxWebAppSiteConfigApplicationStackOutputReference
type jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DockerRegistryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DotnetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) DotnetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) InternalValue() *LinuxWebAppSiteConfigApplicationStack {
	var returns *LinuxWebAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaServerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaServerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) PhpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) PhpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) RubyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) RubyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rubyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLinuxWebAppSiteConfigApplicationStackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxWebAppSiteConfigApplicationStackOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSiteConfigApplicationStackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebApp.LinuxWebAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxWebAppSiteConfigApplicationStackOutputReference_Override(l LinuxWebAppSiteConfigApplicationStackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebApp.LinuxWebAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetDockerImageName(val *string) {
	if err := j.validateSetDockerImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerImageName",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetDockerRegistryPassword(val *string) {
	if err := j.validateSetDockerRegistryPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryPassword",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetDockerRegistryUrl(val *string) {
	if err := j.validateSetDockerRegistryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryUrl",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetDockerRegistryUsername(val *string) {
	if err := j.validateSetDockerRegistryUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistryUsername",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetDotnetVersion(val *string) {
	if err := j.validateSetDotnetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetGoVersion(val *string) {
	if err := j.validateSetGoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetInternalValue(val *LinuxWebAppSiteConfigApplicationStack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetJavaServer(val *string) {
	if err := j.validateSetJavaServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaServer",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetJavaServerVersion(val *string) {
	if err := j.validateSetJavaServerVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaServerVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetJavaVersion(val *string) {
	if err := j.validateSetJavaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetPhpVersion(val *string) {
	if err := j.validateSetPhpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phpVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetPythonVersion(val *string) {
	if err := j.validateSetPythonVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetRubyVersion(val *string) {
	if err := j.validateSetRubyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rubyVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetDockerImageName() {
	_jsii_.InvokeVoid(
		l,
		"resetDockerImageName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetDockerRegistryPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetDockerRegistryPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetDockerRegistryUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetDockerRegistryUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetDockerRegistryUsername() {
	_jsii_.InvokeVoid(
		l,
		"resetDockerRegistryUsername",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetDotnetVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetDotnetVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetGoVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetGoVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetJavaServer() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaServer",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetJavaServerVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaServerVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetPhpVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPhpVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ResetRubyVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetRubyVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigApplicationStackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

