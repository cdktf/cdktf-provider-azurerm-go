package maintenanceconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/maintenanceconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaintenanceConfigurationWindowOutputReference interface {
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	ExpirationDateTime() *string
	SetExpirationDateTime(val *string)
	ExpirationDateTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MaintenanceConfigurationWindow
	SetInternalValue(val *MaintenanceConfigurationWindow)
	RecurEvery() *string
	SetRecurEvery(val *string)
	RecurEveryInput() *string
	StartDateTime() *string
	SetStartDateTime(val *string)
	StartDateTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetDuration()
	ResetExpirationDateTime()
	ResetRecurEvery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MaintenanceConfigurationWindowOutputReference
type jsiiProxy_MaintenanceConfigurationWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ExpirationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ExpirationDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) InternalValue() *MaintenanceConfigurationWindow {
	var returns *MaintenanceConfigurationWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) RecurEvery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurEvery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) RecurEveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurEveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) StartDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) StartDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewMaintenanceConfigurationWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MaintenanceConfigurationWindowOutputReference {
	_init_.Initialize()

	if err := validateNewMaintenanceConfigurationWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceConfigurationWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMaintenanceConfigurationWindowOutputReference_Override(m MaintenanceConfigurationWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetExpirationDateTime(val *string) {
	if err := j.validateSetExpirationDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDateTime",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetInternalValue(val *MaintenanceConfigurationWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetRecurEvery(val *string) {
	if err := j.validateSetRecurEveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurEvery",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetStartDateTime(val *string) {
	if err := j.validateSetStartDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDateTime",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationWindowOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ResetExpirationDateTime() {
	_jsii_.InvokeVoid(
		m,
		"resetExpirationDateTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ResetRecurEvery() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurEvery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

