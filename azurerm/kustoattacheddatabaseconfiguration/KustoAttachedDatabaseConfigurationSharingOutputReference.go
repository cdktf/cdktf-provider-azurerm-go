// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoattacheddatabaseconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kustoattacheddatabaseconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoAttachedDatabaseConfigurationSharingOutputReference interface {
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
	ExternalTablesToExclude() *[]*string
	SetExternalTablesToExclude(val *[]*string)
	ExternalTablesToExcludeInput() *[]*string
	ExternalTablesToInclude() *[]*string
	SetExternalTablesToInclude(val *[]*string)
	ExternalTablesToIncludeInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KustoAttachedDatabaseConfigurationSharing
	SetInternalValue(val *KustoAttachedDatabaseConfigurationSharing)
	MaterializedViewsToExclude() *[]*string
	SetMaterializedViewsToExclude(val *[]*string)
	MaterializedViewsToExcludeInput() *[]*string
	MaterializedViewsToInclude() *[]*string
	SetMaterializedViewsToInclude(val *[]*string)
	MaterializedViewsToIncludeInput() *[]*string
	TablesToExclude() *[]*string
	SetTablesToExclude(val *[]*string)
	TablesToExcludeInput() *[]*string
	TablesToInclude() *[]*string
	SetTablesToInclude(val *[]*string)
	TablesToIncludeInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetExternalTablesToExclude()
	ResetExternalTablesToInclude()
	ResetMaterializedViewsToExclude()
	ResetMaterializedViewsToInclude()
	ResetTablesToExclude()
	ResetTablesToInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoAttachedDatabaseConfigurationSharingOutputReference
type jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InternalValue() *KustoAttachedDatabaseConfigurationSharing {
	var returns *KustoAttachedDatabaseConfigurationSharing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKustoAttachedDatabaseConfigurationSharingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoAttachedDatabaseConfigurationSharingOutputReference {
	_init_.Initialize()

	if err := validateNewKustoAttachedDatabaseConfigurationSharingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationSharingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoAttachedDatabaseConfigurationSharingOutputReference_Override(k KustoAttachedDatabaseConfigurationSharingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationSharingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetExternalTablesToExclude(val *[]*string) {
	if err := j.validateSetExternalTablesToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTablesToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetExternalTablesToInclude(val *[]*string) {
	if err := j.validateSetExternalTablesToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTablesToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetInternalValue(val *KustoAttachedDatabaseConfigurationSharing) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetMaterializedViewsToExclude(val *[]*string) {
	if err := j.validateSetMaterializedViewsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"materializedViewsToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetMaterializedViewsToInclude(val *[]*string) {
	if err := j.validateSetMaterializedViewsToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"materializedViewsToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetTablesToExclude(val *[]*string) {
	if err := j.validateSetTablesToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablesToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetTablesToInclude(val *[]*string) {
	if err := j.validateSetTablesToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablesToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetExternalTablesToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalTablesToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetExternalTablesToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalTablesToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetMaterializedViewsToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetMaterializedViewsToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetMaterializedViewsToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetMaterializedViewsToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetTablesToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetTablesToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetTablesToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetTablesToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

