package maintenanceconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/maintenanceconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaintenanceConfigurationInstallPatchesLinuxOutputReference interface {
	cdktf.ComplexObject
	ClassificationsToInclude() *[]*string
	SetClassificationsToInclude(val *[]*string)
	ClassificationsToIncludeInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PackageNamesMaskToExclude() *[]*string
	SetPackageNamesMaskToExclude(val *[]*string)
	PackageNamesMaskToExcludeInput() *[]*string
	PackageNamesMaskToInclude() *[]*string
	SetPackageNamesMaskToInclude(val *[]*string)
	PackageNamesMaskToIncludeInput() *[]*string
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
	ResetClassificationsToInclude()
	ResetPackageNamesMaskToExclude()
	ResetPackageNamesMaskToInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MaintenanceConfigurationInstallPatchesLinuxOutputReference
type jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ClassificationsToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ClassificationsToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) PackageNamesMaskToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamesMaskToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) PackageNamesMaskToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamesMaskToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) PackageNamesMaskToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamesMaskToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) PackageNamesMaskToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamesMaskToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMaintenanceConfigurationInstallPatchesLinuxOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MaintenanceConfigurationInstallPatchesLinuxOutputReference {
	_init_.Initialize()

	if err := validateNewMaintenanceConfigurationInstallPatchesLinuxOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMaintenanceConfigurationInstallPatchesLinuxOutputReference_Override(m MaintenanceConfigurationInstallPatchesLinuxOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetClassificationsToInclude(val *[]*string) {
	if err := j.validateSetClassificationsToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classificationsToInclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetPackageNamesMaskToExclude(val *[]*string) {
	if err := j.validateSetPackageNamesMaskToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageNamesMaskToExclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetPackageNamesMaskToInclude(val *[]*string) {
	if err := j.validateSetPackageNamesMaskToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageNamesMaskToInclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ResetClassificationsToInclude() {
	_jsii_.InvokeVoid(
		m,
		"resetClassificationsToInclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ResetPackageNamesMaskToExclude() {
	_jsii_.InvokeVoid(
		m,
		"resetPackageNamesMaskToExclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ResetPackageNamesMaskToInclude() {
	_jsii_.InvokeVoid(
		m,
		"resetPackageNamesMaskToInclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesLinuxOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

