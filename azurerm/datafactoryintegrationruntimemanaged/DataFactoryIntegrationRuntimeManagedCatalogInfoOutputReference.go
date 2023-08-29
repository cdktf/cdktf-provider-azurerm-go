// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimemanaged

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/datafactoryintegrationruntimemanaged/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference interface {
	cdktf.ComplexObject
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorPassword() *string
	SetAdministratorPassword(val *string)
	AdministratorPasswordInput() *string
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
	InternalValue() *DataFactoryIntegrationRuntimeManagedCatalogInfo
	SetInternalValue(val *DataFactoryIntegrationRuntimeManagedCatalogInfo)
	PricingTier() *string
	SetPricingTier(val *string)
	PricingTierInput() *string
	ServerEndpoint() *string
	SetServerEndpoint(val *string)
	ServerEndpointInput() *string
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
	ResetAdministratorLogin()
	ResetAdministratorPassword()
	ResetPricingTier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference
type jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) AdministratorPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) AdministratorPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) InternalValue() *DataFactoryIntegrationRuntimeManagedCatalogInfo {
	var returns *DataFactoryIntegrationRuntimeManagedCatalogInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) PricingTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) PricingTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ServerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ServerEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataFactoryIntegrationRuntimeManagedCatalogInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference_Override(d DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetAdministratorPassword(val *string) {
	if err := j.validateSetAdministratorPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPassword",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetInternalValue(val *DataFactoryIntegrationRuntimeManagedCatalogInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetPricingTier(val *string) {
	if err := j.validateSetPricingTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pricingTier",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetServerEndpoint(val *string) {
	if err := j.validateSetServerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverEndpoint",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ResetAdministratorLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetAdministratorLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ResetAdministratorPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetAdministratorPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ResetPricingTier() {
	_jsii_.InvokeVoid(
		d,
		"resetPricingTier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

