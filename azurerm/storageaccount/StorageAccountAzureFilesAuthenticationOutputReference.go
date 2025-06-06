// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/storageaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountAzureFilesAuthenticationOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
	ActiveDirectoryInput() *StorageAccountAzureFilesAuthenticationActiveDirectory
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
	DefaultShareLevelPermission() *string
	SetDefaultShareLevelPermission(val *string)
	DefaultShareLevelPermissionInput() *string
	DirectoryType() *string
	SetDirectoryType(val *string)
	DirectoryTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageAccountAzureFilesAuthentication
	SetInternalValue(val *StorageAccountAzureFilesAuthentication)
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
	PutActiveDirectory(value *StorageAccountAzureFilesAuthenticationActiveDirectory)
	ResetActiveDirectory()
	ResetDefaultShareLevelPermission()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountAzureFilesAuthenticationOutputReference
type jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ActiveDirectory() StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference {
	var returns StorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ActiveDirectoryInput() *StorageAccountAzureFilesAuthenticationActiveDirectory {
	var returns *StorageAccountAzureFilesAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DefaultShareLevelPermission() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultShareLevelPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DefaultShareLevelPermissionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultShareLevelPermissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) DirectoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InternalValue() *StorageAccountAzureFilesAuthentication {
	var returns *StorageAccountAzureFilesAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageAccountAzureFilesAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountAzureFilesAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewStorageAccountAzureFilesAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountAzureFilesAuthenticationOutputReference_Override(s StorageAccountAzureFilesAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccount.StorageAccountAzureFilesAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetDefaultShareLevelPermission(val *string) {
	if err := j.validateSetDefaultShareLevelPermissionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultShareLevelPermission",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetDirectoryType(val *string) {
	if err := j.validateSetDirectoryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetInternalValue(val *StorageAccountAzureFilesAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) PutActiveDirectory(value *StorageAccountAzureFilesAuthenticationActiveDirectory) {
	if err := s.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ResetDefaultShareLevelPermission() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultShareLevelPermission",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountAzureFilesAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

