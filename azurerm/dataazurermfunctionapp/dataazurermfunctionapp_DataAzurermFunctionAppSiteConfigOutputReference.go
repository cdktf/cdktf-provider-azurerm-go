package dataazurermfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/dataazurermfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermFunctionAppSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AlwaysOn() cdktf.IResolvable
	AppScaleLimit() *float64
	AutoSwapSlotName() *string
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
	Cors() DataAzurermFunctionAppSiteConfigCorsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DotnetFrameworkVersion() *string
	ElasticInstanceMinimum() *float64
	// Experimental.
	Fqn() *string
	FtpsState() *string
	HealthCheckPath() *string
	Http2Enabled() cdktf.IResolvable
	InternalValue() *DataAzurermFunctionAppSiteConfig
	SetInternalValue(val *DataAzurermFunctionAppSiteConfig)
	IpRestriction() DataAzurermFunctionAppSiteConfigIpRestrictionList
	JavaVersion() *string
	LinuxFxVersion() *string
	MinTlsVersion() *string
	PreWarmedInstanceCount() *float64
	RuntimeScaleMonitoringEnabled() cdktf.IResolvable
	ScmIpRestriction() DataAzurermFunctionAppSiteConfigScmIpRestrictionList
	ScmType() *string
	ScmUseMainIpRestriction() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Use32BitWorkerProcess() cdktf.IResolvable
	VnetRouteAllEnabled() cdktf.IResolvable
	WebsocketsEnabled() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermFunctionAppSiteConfigOutputReference
type jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) AlwaysOn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) AppScaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) AutoSwapSlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) Cors() DataAzurermFunctionAppSiteConfigCorsList {
	var returns DataAzurermFunctionAppSiteConfigCorsList
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) DotnetFrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) Http2Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) InternalValue() *DataAzurermFunctionAppSiteConfig {
	var returns *DataAzurermFunctionAppSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) IpRestriction() DataAzurermFunctionAppSiteConfigIpRestrictionList {
	var returns DataAzurermFunctionAppSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) PreWarmedInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ScmIpRestriction() DataAzurermFunctionAppSiteConfigScmIpRestrictionList {
	var returns DataAzurermFunctionAppSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ScmUseMainIpRestriction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) Use32BitWorkerProcess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"use32BitWorkerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) VnetRouteAllEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) WebsocketsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}


func NewDataAzurermFunctionAppSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermFunctionAppSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermFunctionAppSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermFunctionApp.DataAzurermFunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermFunctionAppSiteConfigOutputReference_Override(d DataAzurermFunctionAppSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermFunctionApp.DataAzurermFunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference)SetInternalValue(val *DataAzurermFunctionAppSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermFunctionAppSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

