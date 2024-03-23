// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/nginxdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NginxDeploymentConfigurationOutputReference interface {
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
	ConfigFile() NginxDeploymentConfigurationConfigFileList
	ConfigFileInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *NginxDeploymentConfiguration
	SetInternalValue(val *NginxDeploymentConfiguration)
	PackageData() *string
	SetPackageData(val *string)
	PackageDataInput() *string
	ProtectedFile() NginxDeploymentConfigurationProtectedFileList
	ProtectedFileInput() interface{}
	RootFile() *string
	SetRootFile(val *string)
	RootFileInput() *string
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
	PutConfigFile(value interface{})
	PutProtectedFile(value interface{})
	ResetConfigFile()
	ResetPackageData()
	ResetProtectedFile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NginxDeploymentConfigurationOutputReference
type jsiiProxy_NginxDeploymentConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ConfigFile() NginxDeploymentConfigurationConfigFileList {
	var returns NginxDeploymentConfigurationConfigFileList
	_jsii_.Get(
		j,
		"configFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ConfigFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) InternalValue() *NginxDeploymentConfiguration {
	var returns *NginxDeploymentConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) PackageData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) PackageDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ProtectedFile() NginxDeploymentConfigurationProtectedFileList {
	var returns NginxDeploymentConfigurationProtectedFileList
	_jsii_.Get(
		j,
		"protectedFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) ProtectedFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) RootFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) RootFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNginxDeploymentConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NginxDeploymentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNginxDeploymentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NginxDeploymentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNginxDeploymentConfigurationOutputReference_Override(n NginxDeploymentConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.nginxDeployment.NginxDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetInternalValue(val *NginxDeploymentConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetPackageData(val *string) {
	if err := j.validateSetPackageDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageData",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetRootFile(val *string) {
	if err := j.validateSetRootFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootFile",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NginxDeploymentConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) PutConfigFile(value interface{}) {
	if err := n.validatePutConfigFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putConfigFile",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) PutProtectedFile(value interface{}) {
	if err := n.validatePutProtectedFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putProtectedFile",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) ResetConfigFile() {
	_jsii_.InvokeVoid(
		n,
		"resetConfigFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) ResetPackageData() {
	_jsii_.InvokeVoid(
		n,
		"resetPackageData",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) ResetProtectedFile() {
	_jsii_.InvokeVoid(
		n,
		"resetProtectedFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NginxDeploymentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

