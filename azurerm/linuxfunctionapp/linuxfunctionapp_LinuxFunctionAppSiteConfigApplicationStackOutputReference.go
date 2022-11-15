package linuxfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/linuxfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppSiteConfigApplicationStackOutputReference interface {
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
	Docker() LinuxFunctionAppSiteConfigApplicationStackDockerList
	DockerInput() interface{}
	DotnetVersion() *string
	SetDotnetVersion(val *string)
	DotnetVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *LinuxFunctionAppSiteConfigApplicationStack
	SetInternalValue(val *LinuxFunctionAppSiteConfigApplicationStack)
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	PowershellCoreVersion() *string
	SetPowershellCoreVersion(val *string)
	PowershellCoreVersionInput() *string
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
	UseCustomRuntime() interface{}
	SetUseCustomRuntime(val interface{})
	UseCustomRuntimeInput() interface{}
	UseDotnetIsolatedRuntime() interface{}
	SetUseDotnetIsolatedRuntime(val interface{})
	UseDotnetIsolatedRuntimeInput() interface{}
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
	PutDocker(value interface{})
	ResetDocker()
	ResetDotnetVersion()
	ResetJavaVersion()
	ResetNodeVersion()
	ResetPowershellCoreVersion()
	ResetPythonVersion()
	ResetUseCustomRuntime()
	ResetUseDotnetIsolatedRuntime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxFunctionAppSiteConfigApplicationStackOutputReference
type jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) Docker() LinuxFunctionAppSiteConfigApplicationStackDockerList {
	var returns LinuxFunctionAppSiteConfigApplicationStackDockerList
	_jsii_.Get(
		j,
		"docker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) DockerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) DotnetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) DotnetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) InternalValue() *LinuxFunctionAppSiteConfigApplicationStack {
	var returns *LinuxFunctionAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) PowershellCoreVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellCoreVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) PowershellCoreVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellCoreVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) UseCustomRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) UseCustomRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) UseDotnetIsolatedRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDotnetIsolatedRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) UseDotnetIsolatedRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDotnetIsolatedRuntimeInput",
		&returns,
	)
	return returns
}


func NewLinuxFunctionAppSiteConfigApplicationStackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxFunctionAppSiteConfigApplicationStackOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppSiteConfigApplicationStackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxFunctionAppSiteConfigApplicationStackOutputReference_Override(l LinuxFunctionAppSiteConfigApplicationStackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetDotnetVersion(val *string) {
	if err := j.validateSetDotnetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetInternalValue(val *LinuxFunctionAppSiteConfigApplicationStack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetJavaVersion(val *string) {
	if err := j.validateSetJavaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetPowershellCoreVersion(val *string) {
	if err := j.validateSetPowershellCoreVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powershellCoreVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetPythonVersion(val *string) {
	if err := j.validateSetPythonVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetUseCustomRuntime(val interface{}) {
	if err := j.validateSetUseCustomRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomRuntime",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference)SetUseDotnetIsolatedRuntime(val interface{}) {
	if err := j.validateSetUseDotnetIsolatedRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDotnetIsolatedRuntime",
		val,
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) PutDocker(value interface{}) {
	if err := l.validatePutDockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDocker",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetDocker() {
	_jsii_.InvokeVoid(
		l,
		"resetDocker",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetDotnetVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetDotnetVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetPowershellCoreVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPowershellCoreVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetUseCustomRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetUseCustomRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ResetUseDotnetIsolatedRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetUseDotnetIsolatedRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxFunctionAppSiteConfigApplicationStackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

