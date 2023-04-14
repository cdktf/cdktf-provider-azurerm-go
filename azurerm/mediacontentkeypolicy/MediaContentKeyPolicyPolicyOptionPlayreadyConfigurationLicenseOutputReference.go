package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference interface {
	cdktf.ComplexObject
	AllowTestDevices() interface{}
	SetAllowTestDevices(val interface{})
	AllowTestDevicesInput() interface{}
	BeginDate() *string
	SetBeginDate(val *string)
	BeginDateInput() *string
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
	ContentKeyLocationFromHeaderEnabled() interface{}
	SetContentKeyLocationFromHeaderEnabled(val interface{})
	ContentKeyLocationFromHeaderEnabledInput() interface{}
	ContentKeyLocationFromKeyId() *string
	SetContentKeyLocationFromKeyId(val *string)
	ContentKeyLocationFromKeyIdInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	// Experimental.
	Fqn() *string
	GracePeriod() *string
	SetGracePeriod(val *string)
	GracePeriodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	PlayRight() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
	PlayRightInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	RelativeBeginDate() *string
	SetRelativeBeginDate(val *string)
	RelativeBeginDateInput() *string
	RelativeExpirationDate() *string
	SetRelativeExpirationDate(val *string)
	RelativeExpirationDateInput() *string
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
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
	PutPlayRight(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight)
	ResetAllowTestDevices()
	ResetBeginDate()
	ResetContentKeyLocationFromHeaderEnabled()
	ResetContentKeyLocationFromKeyId()
	ResetContentType()
	ResetExpirationDate()
	ResetGracePeriod()
	ResetLicenseType()
	ResetPlayRight()
	ResetRelativeBeginDate()
	ResetRelativeExpirationDate()
	ResetSecurityLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) AllowTestDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowTestDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) AllowTestDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowTestDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) BeginDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) BeginDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromHeaderEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyLocationFromHeaderEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromHeaderEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyLocationFromHeaderEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentKeyLocationFromKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentKeyLocationFromKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PlayRight() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
	_jsii_.Get(
		j,
		"playRight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PlayRightInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight {
	var returns *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	_jsii_.Get(
		j,
		"playRightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeBeginDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeBeginDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeBeginDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeBeginDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeExpirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeExpirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference_Override(m MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetAllowTestDevices(val interface{}) {
	if err := j.validateSetAllowTestDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowTestDevices",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetBeginDate(val *string) {
	if err := j.validateSetBeginDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beginDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetContentKeyLocationFromHeaderEnabled(val interface{}) {
	if err := j.validateSetContentKeyLocationFromHeaderEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentKeyLocationFromHeaderEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetContentKeyLocationFromKeyId(val *string) {
	if err := j.validateSetContentKeyLocationFromKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentKeyLocationFromKeyId",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetExpirationDate(val *string) {
	if err := j.validateSetExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetGracePeriod(val *string) {
	if err := j.validateSetGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetRelativeBeginDate(val *string) {
	if err := j.validateSetRelativeBeginDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relativeBeginDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetRelativeExpirationDate(val *string) {
	if err := j.validateSetRelativeExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relativeExpirationDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PutPlayRight(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight) {
	if err := m.validatePutPlayRightParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPlayRight",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetAllowTestDevices() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowTestDevices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetBeginDate() {
	_jsii_.InvokeVoid(
		m,
		"resetBeginDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentKeyLocationFromHeaderEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKeyLocationFromHeaderEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentKeyLocationFromKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKeyLocationFromKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		m,
		"resetContentType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		m,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetLicenseType() {
	_jsii_.InvokeVoid(
		m,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetPlayRight() {
	_jsii_.InvokeVoid(
		m,
		"resetPlayRight",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetRelativeBeginDate() {
	_jsii_.InvokeVoid(
		m,
		"resetRelativeBeginDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetRelativeExpirationDate() {
	_jsii_.InvokeVoid(
		m,
		"resetRelativeExpirationDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

