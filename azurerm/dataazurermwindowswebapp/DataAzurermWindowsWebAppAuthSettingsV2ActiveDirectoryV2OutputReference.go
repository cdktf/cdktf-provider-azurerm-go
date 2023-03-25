package dataazurermwindowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/dataazurermwindowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference interface {
	cdktf.ComplexObject
	AllowedApplications() *[]*string
	AllowedAudiences() *[]*string
	AllowedGroups() *[]*string
	AllowedIdentities() *[]*string
	ClientId() *string
	ClientSecretCertificateThumbprint() *string
	ClientSecretSettingName() *string
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
	InternalValue() *DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2
	SetInternalValue(val *DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2)
	JwtAllowedClientApplications() *[]*string
	JwtAllowedGroups() *[]*string
	LoginParameters() cdktf.StringMap
	TenantAuthEndpoint() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WwwAuthenticationDisabled() cdktf.IResolvable
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

// The jsii proxy struct for DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference
type jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedApplications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedIdentities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretCertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InternalValue() *DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2 {
	var returns *DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedClientApplications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedClientApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) LoginParameters() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"loginParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TenantAuthEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantAuthEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) WwwAuthenticationDisabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"wwwAuthenticationDisabled",
		&returns,
	)
	return returns
}


func NewDataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference_Override(d DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetInternalValue(val *DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

