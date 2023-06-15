package cosmosdbaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/cosmosdbaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbAccountBackupOutputReference interface {
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
	InternalValue() *CosmosdbAccountBackup
	SetInternalValue(val *CosmosdbAccountBackup)
	IntervalInMinutes() *float64
	SetIntervalInMinutes(val *float64)
	IntervalInMinutesInput() *float64
	RetentionInHours() *float64
	SetRetentionInHours(val *float64)
	RetentionInHoursInput() *float64
	StorageRedundancy() *string
	SetStorageRedundancy(val *string)
	StorageRedundancyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetIntervalInMinutes()
	ResetRetentionInHours()
	ResetStorageRedundancy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountBackupOutputReference
type jsiiProxy_CosmosdbAccountBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) InternalValue() *CosmosdbAccountBackup {
	var returns *CosmosdbAccountBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) IntervalInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) IntervalInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) RetentionInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) RetentionInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) StorageRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) StorageRedundancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRedundancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountBackupOutputReference {
	_init_.Initialize()

	if err := validateNewCosmosdbAccountBackupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbAccountBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountBackupOutputReference_Override(c CosmosdbAccountBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetInternalValue(val *CosmosdbAccountBackup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetIntervalInMinutes(val *float64) {
	if err := j.validateSetIntervalInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalInMinutes",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetRetentionInHours(val *float64) {
	if err := j.validateSetRetentionInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInHours",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetStorageRedundancy(val *string) {
	if err := j.validateSetStorageRedundancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRedundancy",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetIntervalInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetIntervalInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetRetentionInHours() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionInHours",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetStorageRedundancy() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageRedundancy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

