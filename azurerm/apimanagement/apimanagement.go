package apimanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/api_management azurerm_api_management}.
type ApiManagement interface {
	cdktf.TerraformResource
	AdditionalLocation() ApiManagementAdditionalLocationList
	AdditionalLocationInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() ApiManagementCertificateList
	CertificateInput() interface{}
	ClientCertificateEnabled() interface{}
	SetClientCertificateEnabled(val interface{})
	ClientCertificateEnabledInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeveloperPortalUrl() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayDisabled() interface{}
	SetGatewayDisabled(val interface{})
	GatewayDisabledInput() interface{}
	GatewayRegionalUrl() *string
	GatewayUrl() *string
	HostnameConfiguration() ApiManagementHostnameConfigurationOutputReference
	HostnameConfigurationInput() *ApiManagementHostnameConfiguration
	Id() *string
	SetId(val *string)
	Identity() ApiManagementIdentityOutputReference
	IdentityInput() *ApiManagementIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementApiUrl() *string
	MinApiVersion() *string
	SetMinApiVersion(val *string)
	MinApiVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationSenderEmail() *string
	SetNotificationSenderEmail(val *string)
	NotificationSenderEmailInput() *string
	Policy() ApiManagementPolicyList
	PolicyInput() interface{}
	PortalUrl() *string
	PrivateIpAddresses() *[]*string
	Protocols() ApiManagementProtocolsOutputReference
	ProtocolsInput() *ApiManagementProtocols
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpAddresses() *[]*string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PublisherEmail() *string
	SetPublisherEmail(val *string)
	PublisherEmailInput() *string
	PublisherName() *string
	SetPublisherName(val *string)
	PublisherNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScmUrl() *string
	Security() ApiManagementSecurityOutputReference
	SecurityInput() *ApiManagementSecurity
	SignIn() ApiManagementSignInOutputReference
	SignInInput() *ApiManagementSignIn
	SignUp() ApiManagementSignUpOutputReference
	SignUpInput() *ApiManagementSignUp
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantAccess() ApiManagementTenantAccessOutputReference
	TenantAccessInput() *ApiManagementTenantAccess
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkConfiguration() ApiManagementVirtualNetworkConfigurationOutputReference
	VirtualNetworkConfigurationInput() *ApiManagementVirtualNetworkConfiguration
	VirtualNetworkType() *string
	SetVirtualNetworkType(val *string)
	VirtualNetworkTypeInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdditionalLocation(value interface{})
	PutCertificate(value interface{})
	PutHostnameConfiguration(value *ApiManagementHostnameConfiguration)
	PutIdentity(value *ApiManagementIdentity)
	PutPolicy(value interface{})
	PutProtocols(value *ApiManagementProtocols)
	PutSecurity(value *ApiManagementSecurity)
	PutSignIn(value *ApiManagementSignIn)
	PutSignUp(value *ApiManagementSignUp)
	PutTenantAccess(value *ApiManagementTenantAccess)
	PutTimeouts(value *ApiManagementTimeouts)
	PutVirtualNetworkConfiguration(value *ApiManagementVirtualNetworkConfiguration)
	ResetAdditionalLocation()
	ResetCertificate()
	ResetClientCertificateEnabled()
	ResetGatewayDisabled()
	ResetHostnameConfiguration()
	ResetId()
	ResetIdentity()
	ResetMinApiVersion()
	ResetNotificationSenderEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetProtocols()
	ResetPublicIpAddressId()
	ResetPublicNetworkAccessEnabled()
	ResetSecurity()
	ResetSignIn()
	ResetSignUp()
	ResetTags()
	ResetTenantAccess()
	ResetTimeouts()
	ResetVirtualNetworkConfiguration()
	ResetVirtualNetworkType()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiManagement
type jsiiProxy_ApiManagement struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagement) AdditionalLocation() ApiManagementAdditionalLocationList {
	var returns ApiManagementAdditionalLocationList
	_jsii_.Get(
		j,
		"additionalLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) AdditionalLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Certificate() ApiManagementCertificateList {
	var returns ApiManagementCertificateList
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) DeveloperPortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerPortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayRegionalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRegionalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) HostnameConfiguration() ApiManagementHostnameConfigurationOutputReference {
	var returns ApiManagementHostnameConfigurationOutputReference
	_jsii_.Get(
		j,
		"hostnameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) HostnameConfigurationInput() *ApiManagementHostnameConfiguration {
	var returns *ApiManagementHostnameConfiguration
	_jsii_.Get(
		j,
		"hostnameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Identity() ApiManagementIdentityOutputReference {
	var returns ApiManagementIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) IdentityInput() *ApiManagementIdentity {
	var returns *ApiManagementIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ManagementApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) MinApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) MinApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NotificationSenderEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationSenderEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NotificationSenderEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationSenderEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Policy() ApiManagementPolicyList {
	var returns ApiManagementPolicyList
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Protocols() ApiManagementProtocolsOutputReference {
	var returns ApiManagementProtocolsOutputReference
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ProtocolsInput() *ApiManagementProtocols {
	var returns *ApiManagementProtocols
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ScmUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Security() ApiManagementSecurityOutputReference {
	var returns ApiManagementSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SecurityInput() *ApiManagementSecurity {
	var returns *ApiManagementSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignIn() ApiManagementSignInOutputReference {
	var returns ApiManagementSignInOutputReference
	_jsii_.Get(
		j,
		"signIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignInInput() *ApiManagementSignIn {
	var returns *ApiManagementSignIn
	_jsii_.Get(
		j,
		"signInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignUp() ApiManagementSignUpOutputReference {
	var returns ApiManagementSignUpOutputReference
	_jsii_.Get(
		j,
		"signUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignUpInput() *ApiManagementSignUp {
	var returns *ApiManagementSignUp
	_jsii_.Get(
		j,
		"signUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TenantAccess() ApiManagementTenantAccessOutputReference {
	var returns ApiManagementTenantAccessOutputReference
	_jsii_.Get(
		j,
		"tenantAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TenantAccessInput() *ApiManagementTenantAccess {
	var returns *ApiManagementTenantAccess
	_jsii_.Get(
		j,
		"tenantAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Timeouts() ApiManagementTimeoutsOutputReference {
	var returns ApiManagementTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkConfiguration() ApiManagementVirtualNetworkConfigurationOutputReference {
	var returns ApiManagementVirtualNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkConfigurationInput() *ApiManagementVirtualNetworkConfiguration {
	var returns *ApiManagementVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"virtualNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/api_management azurerm_api_management} Resource.
func NewApiManagement(scope constructs.Construct, id *string, config *ApiManagementConfig) ApiManagement {
	_init_.Initialize()

	j := jsiiProxy_ApiManagement{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/api_management azurerm_api_management} Resource.
func NewApiManagement_Override(a ApiManagement, scope constructs.Construct, id *string, config *ApiManagementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagement) SetClientCertificateEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetGatewayDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayDisabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetMinApiVersion(val *string) {
	_jsii_.Set(
		j,
		"minApiVersion",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetNotificationSenderEmail(val *string) {
	_jsii_.Set(
		j,
		"notificationSenderEmail",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetPublicIpAddressId(val *string) {
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetPublisherEmail(val *string) {
	_jsii_.Set(
		j,
		"publisherEmail",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetPublisherName(val *string) {
	_jsii_.Set(
		j,
		"publisherName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetVirtualNetworkType(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkType",
		val,
	)
}

func (j *jsiiProxy_ApiManagement) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApiManagement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagement) PutAdditionalLocation(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putAdditionalLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutHostnameConfiguration(value *ApiManagementHostnameConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putHostnameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutIdentity(value *ApiManagementIdentity) {
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutProtocols(value *ApiManagementProtocols) {
	_jsii_.InvokeVoid(
		a,
		"putProtocols",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSecurity(value *ApiManagementSecurity) {
	_jsii_.InvokeVoid(
		a,
		"putSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSignIn(value *ApiManagementSignIn) {
	_jsii_.InvokeVoid(
		a,
		"putSignIn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSignUp(value *ApiManagementSignUp) {
	_jsii_.InvokeVoid(
		a,
		"putSignUp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutTenantAccess(value *ApiManagementTenantAccess) {
	_jsii_.InvokeVoid(
		a,
		"putTenantAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutTimeouts(value *ApiManagementTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutVirtualNetworkConfiguration(value *ApiManagementVirtualNetworkConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putVirtualNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) ResetAdditionalLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetGatewayDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetHostnameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHostnameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetMinApiVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMinApiVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetNotificationSenderEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationSenderEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPublicIpAddressId() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIpAddressId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSignIn() {
	_jsii_.InvokeVoid(
		a,
		"resetSignIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSignUp() {
	_jsii_.InvokeVoid(
		a,
		"resetSignUp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTenantAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetVirtualNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetVirtualNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetZones() {
	_jsii_.InvokeVoid(
		a,
		"resetZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementAdditionalLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#location ApiManagement#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#capacity ApiManagement#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#gateway_disabled ApiManagement#gateway_disabled}.
	GatewayDisabled interface{} `field:"optional" json:"gatewayDisabled" yaml:"gatewayDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#public_ip_address_id ApiManagement#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#virtual_network_configuration ApiManagement#virtual_network_configuration}
	VirtualNetworkConfiguration *ApiManagementAdditionalLocationVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#zones ApiManagement#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type ApiManagementAdditionalLocationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementAdditionalLocationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementAdditionalLocationList
type jsiiProxy_ApiManagementAdditionalLocationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementAdditionalLocationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementAdditionalLocationList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementAdditionalLocationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementAdditionalLocationList_Override(a ApiManagementAdditionalLocationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) Get(index *float64) ApiManagementAdditionalLocationOutputReference {
	var returns ApiManagementAdditionalLocationOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementAdditionalLocationOutputReference interface {
	cdktf.ComplexObject
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
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
	GatewayDisabled() interface{}
	SetGatewayDisabled(val interface{})
	GatewayDisabledInput() interface{}
	GatewayRegionalUrl() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	PrivateIpAddresses() *[]*string
	PublicIpAddresses() *[]*string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkConfiguration() ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference
	VirtualNetworkConfigurationInput() *ApiManagementAdditionalLocationVirtualNetworkConfiguration
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutVirtualNetworkConfiguration(value *ApiManagementAdditionalLocationVirtualNetworkConfiguration)
	ResetCapacity()
	ResetGatewayDisabled()
	ResetPublicIpAddressId()
	ResetVirtualNetworkConfiguration()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementAdditionalLocationOutputReference
type jsiiProxy_ApiManagementAdditionalLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GatewayDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GatewayDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GatewayRegionalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRegionalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) VirtualNetworkConfiguration() ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference {
	var returns ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) VirtualNetworkConfigurationInput() *ApiManagementAdditionalLocationVirtualNetworkConfiguration {
	var returns *ApiManagementAdditionalLocationVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"virtualNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewApiManagementAdditionalLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementAdditionalLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementAdditionalLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementAdditionalLocationOutputReference_Override(a ApiManagementAdditionalLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetGatewayDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayDisabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetPublicIpAddressId(val *string) {
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationOutputReference) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) PutVirtualNetworkConfiguration(value *ApiManagementAdditionalLocationVirtualNetworkConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putVirtualNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ResetGatewayDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ResetPublicIpAddressId() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIpAddressId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ResetVirtualNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		a,
		"resetZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementAdditionalLocationVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#subnet_id ApiManagement#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

type ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference interface {
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
	InternalValue() *ApiManagementAdditionalLocationVirtualNetworkConfiguration
	SetInternalValue(val *ApiManagementAdditionalLocationVirtualNetworkConfiguration)
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference
type jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) InternalValue() *ApiManagementAdditionalLocationVirtualNetworkConfiguration {
	var returns *ApiManagementAdditionalLocationVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference_Override(a ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetInternalValue(val *ApiManagementAdditionalLocationVirtualNetworkConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAdditionalLocationVirtualNetworkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#encoded_certificate ApiManagement#encoded_certificate}.
	EncodedCertificate *string `field:"required" json:"encodedCertificate" yaml:"encodedCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#store_name ApiManagement#store_name}.
	StoreName *string `field:"required" json:"storeName" yaml:"storeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
}

type ApiManagementCertificateList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementCertificateList
type jsiiProxy_ApiManagementCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementCertificateList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementCertificateList_Override(a ApiManagementCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateList) Get(index *float64) ApiManagementCertificateOutputReference {
	var returns ApiManagementCertificateOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementCertificateOutputReference interface {
	cdktf.ComplexObject
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
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
	EncodedCertificate() *string
	SetEncodedCertificate(val *string)
	EncodedCertificateInput() *string
	Expiry() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StoreName() *string
	SetStoreName(val *string)
	StoreNameInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificatePassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementCertificateOutputReference
type jsiiProxy_ApiManagementCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) EncodedCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) EncodedCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementCertificateOutputReference_Override(a ApiManagementCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetEncodedCertificate(val *string) {
	_jsii_.Set(
		j,
		"encodedCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetStoreName(val *string) {
	_jsii_.Set(
		j,
		"storeName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#location ApiManagement#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#name ApiManagement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#publisher_email ApiManagement#publisher_email}.
	PublisherEmail *string `field:"required" json:"publisherEmail" yaml:"publisherEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#publisher_name ApiManagement#publisher_name}.
	PublisherName *string `field:"required" json:"publisherName" yaml:"publisherName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#resource_group_name ApiManagement#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#sku_name ApiManagement#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// additional_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#additional_location ApiManagement#additional_location}
	AdditionalLocation interface{} `field:"optional" json:"additionalLocation" yaml:"additionalLocation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#client_certificate_enabled ApiManagement#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#gateway_disabled ApiManagement#gateway_disabled}.
	GatewayDisabled interface{} `field:"optional" json:"gatewayDisabled" yaml:"gatewayDisabled"`
	// hostname_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#hostname_configuration ApiManagement#hostname_configuration}
	HostnameConfiguration *ApiManagementHostnameConfiguration `field:"optional" json:"hostnameConfiguration" yaml:"hostnameConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#id ApiManagement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#identity ApiManagement#identity}
	Identity *ApiManagementIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#min_api_version ApiManagement#min_api_version}.
	MinApiVersion *string `field:"optional" json:"minApiVersion" yaml:"minApiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#notification_sender_email ApiManagement#notification_sender_email}.
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#policy ApiManagement#policy}.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// protocols block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#protocols ApiManagement#protocols}
	Protocols *ApiManagementProtocols `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#public_ip_address_id ApiManagement#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#public_network_access_enabled ApiManagement#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#security ApiManagement#security}
	Security *ApiManagementSecurity `field:"optional" json:"security" yaml:"security"`
	// sign_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#sign_in ApiManagement#sign_in}
	SignIn *ApiManagementSignIn `field:"optional" json:"signIn" yaml:"signIn"`
	// sign_up block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#sign_up ApiManagement#sign_up}
	SignUp *ApiManagementSignUp `field:"optional" json:"signUp" yaml:"signUp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tags ApiManagement#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant_access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tenant_access ApiManagement#tenant_access}
	TenantAccess *ApiManagementTenantAccess `field:"optional" json:"tenantAccess" yaml:"tenantAccess"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#timeouts ApiManagement#timeouts}
	Timeouts *ApiManagementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#virtual_network_configuration ApiManagement#virtual_network_configuration}
	VirtualNetworkConfiguration *ApiManagementVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#virtual_network_type ApiManagement#virtual_network_type}.
	VirtualNetworkType *string `field:"optional" json:"virtualNetworkType" yaml:"virtualNetworkType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#zones ApiManagement#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type ApiManagementHostnameConfiguration struct {
	// developer_portal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#developer_portal ApiManagement#developer_portal}
	DeveloperPortal interface{} `field:"optional" json:"developerPortal" yaml:"developerPortal"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#management ApiManagement#management}
	Management interface{} `field:"optional" json:"management" yaml:"management"`
	// portal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#portal ApiManagement#portal}
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#proxy ApiManagement#proxy}
	Proxy interface{} `field:"optional" json:"proxy" yaml:"proxy"`
	// scm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#scm ApiManagement#scm}
	Scm interface{} `field:"optional" json:"scm" yaml:"scm"`
}

type ApiManagementHostnameConfigurationDeveloperPortal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

type ApiManagementHostnameConfigurationDeveloperPortalList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementHostnameConfigurationDeveloperPortalOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationDeveloperPortalList
type jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationDeveloperPortalList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementHostnameConfigurationDeveloperPortalList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationDeveloperPortalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationDeveloperPortalList_Override(a ApiManagementHostnameConfigurationDeveloperPortalList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationDeveloperPortalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) Get(index *float64) ApiManagementHostnameConfigurationDeveloperPortalOutputReference {
	var returns ApiManagementHostnameConfigurationDeveloperPortalOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationDeveloperPortalOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationDeveloperPortalOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationDeveloperPortalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationDeveloperPortalOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationDeveloperPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationDeveloperPortalOutputReference_Override(a ApiManagementHostnameConfigurationDeveloperPortalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationDeveloperPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetNegotiateClientCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetSslKeyvaultIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationDeveloperPortalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationManagement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

type ApiManagementHostnameConfigurationManagementList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementHostnameConfigurationManagementOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationManagementList
type jsiiProxy_ApiManagementHostnameConfigurationManagementList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationManagementList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementHostnameConfigurationManagementList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationManagementList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationManagementList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationManagementList_Override(a ApiManagementHostnameConfigurationManagementList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationManagementList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) Get(index *float64) ApiManagementHostnameConfigurationManagementOutputReference {
	var returns ApiManagementHostnameConfigurationManagementOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationManagementOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationManagementOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationManagementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationManagementOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationManagementOutputReference_Override(a ApiManagementHostnameConfigurationManagementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetNegotiateClientCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetSslKeyvaultIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationOutputReference interface {
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
	DeveloperPortal() ApiManagementHostnameConfigurationDeveloperPortalList
	DeveloperPortalInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementHostnameConfiguration
	SetInternalValue(val *ApiManagementHostnameConfiguration)
	Management() ApiManagementHostnameConfigurationManagementList
	ManagementInput() interface{}
	Portal() ApiManagementHostnameConfigurationPortalList
	PortalInput() interface{}
	Proxy() ApiManagementHostnameConfigurationProxyList
	ProxyInput() interface{}
	Scm() ApiManagementHostnameConfigurationScmList
	ScmInput() interface{}
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
	PutDeveloperPortal(value interface{})
	PutManagement(value interface{})
	PutPortal(value interface{})
	PutProxy(value interface{})
	PutScm(value interface{})
	ResetDeveloperPortal()
	ResetManagement()
	ResetPortal()
	ResetProxy()
	ResetScm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) DeveloperPortal() ApiManagementHostnameConfigurationDeveloperPortalList {
	var returns ApiManagementHostnameConfigurationDeveloperPortalList
	_jsii_.Get(
		j,
		"developerPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) DeveloperPortalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InternalValue() *ApiManagementHostnameConfiguration {
	var returns *ApiManagementHostnameConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Management() ApiManagementHostnameConfigurationManagementList {
	var returns ApiManagementHostnameConfigurationManagementList
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Portal() ApiManagementHostnameConfigurationPortalList {
	var returns ApiManagementHostnameConfigurationPortalList
	_jsii_.Get(
		j,
		"portal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PortalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Proxy() ApiManagementHostnameConfigurationProxyList {
	var returns ApiManagementHostnameConfigurationProxyList
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ProxyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Scm() ApiManagementHostnameConfigurationScmList {
	var returns ApiManagementHostnameConfigurationScmList
	_jsii_.Get(
		j,
		"scm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ScmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementHostnameConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationOutputReference_Override(a ApiManagementHostnameConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) SetInternalValue(val *ApiManagementHostnameConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutDeveloperPortal(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putDeveloperPortal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutManagement(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutPortal(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putPortal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutProxy(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putProxy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutScm(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putScm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetDeveloperPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetScm() {
	_jsii_.InvokeVoid(
		a,
		"resetScm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationPortal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

type ApiManagementHostnameConfigurationPortalList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementHostnameConfigurationPortalOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationPortalList
type jsiiProxy_ApiManagementHostnameConfigurationPortalList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationPortalList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementHostnameConfigurationPortalList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationPortalList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationPortalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationPortalList_Override(a ApiManagementHostnameConfigurationPortalList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationPortalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) Get(index *float64) ApiManagementHostnameConfigurationPortalOutputReference {
	var returns ApiManagementHostnameConfigurationPortalOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationPortalOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationPortalOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationPortalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationPortalOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationPortalOutputReference_Override(a ApiManagementHostnameConfigurationPortalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetNegotiateClientCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetSslKeyvaultIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationProxy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#default_ssl_binding ApiManagement#default_ssl_binding}.
	DefaultSslBinding interface{} `field:"optional" json:"defaultSslBinding" yaml:"defaultSslBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

type ApiManagementHostnameConfigurationProxyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementHostnameConfigurationProxyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationProxyList
type jsiiProxy_ApiManagementHostnameConfigurationProxyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationProxyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementHostnameConfigurationProxyList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationProxyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationProxyList_Override(a ApiManagementHostnameConfigurationProxyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyList) Get(index *float64) ApiManagementHostnameConfigurationProxyOutputReference {
	var returns ApiManagementHostnameConfigurationProxyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationProxyOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	DefaultSslBinding() interface{}
	SetDefaultSslBinding(val interface{})
	DefaultSslBindingInput() interface{}
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetDefaultSslBinding()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationProxyOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) DefaultSslBinding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSslBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) DefaultSslBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSslBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationProxyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationProxyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationProxyOutputReference_Override(a ApiManagementHostnameConfigurationProxyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetDefaultSslBinding(val interface{}) {
	_jsii_.Set(
		j,
		"defaultSslBinding",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetNegotiateClientCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetSslKeyvaultIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetDefaultSslBinding() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSslBinding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationScm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#host_name ApiManagement#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate ApiManagement#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#certificate_password ApiManagement#certificate_password}.
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#key_vault_id ApiManagement#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}.
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}.
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

type ApiManagementHostnameConfigurationScmList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementHostnameConfigurationScmOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationScmList
type jsiiProxy_ApiManagementHostnameConfigurationScmList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationScmList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementHostnameConfigurationScmList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationScmList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationScmList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationScmList_Override(a ApiManagementHostnameConfigurationScmList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationScmList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmList) Get(index *float64) ApiManagementHostnameConfigurationScmOutputReference {
	var returns ApiManagementHostnameConfigurationScmOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementHostnameConfigurationScmOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationScmOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationScmOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationScmOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationScmOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationScmOutputReference_Override(a ApiManagementHostnameConfigurationScmOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationScmOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetHostName(val *string) {
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetKeyVaultId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetNegotiateClientCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetSslKeyvaultIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationScmOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#type ApiManagement#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#identity_ids ApiManagement#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type ApiManagementIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *ApiManagementIdentity
	SetInternalValue(val *ApiManagementIdentity)
	PrincipalId() *string
	TenantId() *string
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementIdentityOutputReference
type jsiiProxy_ApiManagementIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) InternalValue() *ApiManagementIdentity {
	var returns *ApiManagementIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewApiManagementIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementIdentityOutputReference_Override(a ApiManagementIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetInternalValue(val *ApiManagementIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#xml_content ApiManagement#xml_content}.
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#xml_link ApiManagement#xml_link}.
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

type ApiManagementPolicyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApiManagementPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementPolicyList
type jsiiProxy_ApiManagementPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApiManagementPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiManagementPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiManagementPolicyList {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiManagementPolicyList_Override(a ApiManagementPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiManagementPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyList) Get(index *float64) ApiManagementPolicyOutputReference {
	var returns ApiManagementPolicyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementPolicyOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XmlContent() *string
	SetXmlContent(val *string)
	XmlContentInput() *string
	XmlLink() *string
	SetXmlLink(val *string)
	XmlLinkInput() *string
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
	ResetXmlContent()
	ResetXmlLink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementPolicyOutputReference
type jsiiProxy_ApiManagementPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) XmlContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xmlContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) XmlContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xmlContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) XmlLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xmlLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) XmlLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xmlLinkInput",
		&returns,
	)
	return returns
}


func NewApiManagementPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementPolicyOutputReference_Override(a ApiManagementPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetXmlContent(val *string) {
	_jsii_.Set(
		j,
		"xmlContent",
		val,
	)
}

func (j *jsiiProxy_ApiManagementPolicyOutputReference) SetXmlLink(val *string) {
	_jsii_.Set(
		j,
		"xmlLink",
		val,
	)
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) ResetXmlContent() {
	_jsii_.InvokeVoid(
		a,
		"resetXmlContent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) ResetXmlLink() {
	_jsii_.InvokeVoid(
		a,
		"resetXmlLink",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementProtocols struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_http2 ApiManagement#enable_http2}.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
}

type ApiManagementProtocolsOutputReference interface {
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
	EnableHttp2() interface{}
	SetEnableHttp2(val interface{})
	EnableHttp2Input() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementProtocols
	SetInternalValue(val *ApiManagementProtocols)
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
	ResetEnableHttp2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementProtocolsOutputReference
type jsiiProxy_ApiManagementProtocolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) InternalValue() *ApiManagementProtocols {
	var returns *ApiManagementProtocols
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementProtocolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementProtocolsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementProtocolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementProtocolsOutputReference_Override(a ApiManagementProtocolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementProtocolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetEnableHttp2(val interface{}) {
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetInternalValue(val *ApiManagementProtocols) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementProtocolsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementProtocolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_backend_ssl30 ApiManagement#enable_backend_ssl30}.
	EnableBackendSsl30 interface{} `field:"optional" json:"enableBackendSsl30" yaml:"enableBackendSsl30"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_backend_tls10 ApiManagement#enable_backend_tls10}.
	EnableBackendTls10 interface{} `field:"optional" json:"enableBackendTls10" yaml:"enableBackendTls10"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_backend_tls11 ApiManagement#enable_backend_tls11}.
	EnableBackendTls11 interface{} `field:"optional" json:"enableBackendTls11" yaml:"enableBackendTls11"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_frontend_ssl30 ApiManagement#enable_frontend_ssl30}.
	EnableFrontendSsl30 interface{} `field:"optional" json:"enableFrontendSsl30" yaml:"enableFrontendSsl30"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_frontend_tls10 ApiManagement#enable_frontend_tls10}.
	EnableFrontendTls10 interface{} `field:"optional" json:"enableFrontendTls10" yaml:"enableFrontendTls10"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enable_frontend_tls11 ApiManagement#enable_frontend_tls11}.
	EnableFrontendTls11 interface{} `field:"optional" json:"enableFrontendTls11" yaml:"enableFrontendTls11"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled}.
	TlsRsaWithAes128CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes128CbcSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsRsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled}.
	TlsRsaWithAes128GcmSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128GcmSha256CiphersEnabled" yaml:"tlsRsaWithAes128GcmSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled}.
	TlsRsaWithAes256CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes256CbcSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#tls_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsRsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#triple_des_ciphers_enabled ApiManagement#triple_des_ciphers_enabled}.
	TripleDesCiphersEnabled interface{} `field:"optional" json:"tripleDesCiphersEnabled" yaml:"tripleDesCiphersEnabled"`
}

type ApiManagementSecurityOutputReference interface {
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
	EnableBackendSsl30() interface{}
	SetEnableBackendSsl30(val interface{})
	EnableBackendSsl30Input() interface{}
	EnableBackendTls10() interface{}
	SetEnableBackendTls10(val interface{})
	EnableBackendTls10Input() interface{}
	EnableBackendTls11() interface{}
	SetEnableBackendTls11(val interface{})
	EnableBackendTls11Input() interface{}
	EnableFrontendSsl30() interface{}
	SetEnableFrontendSsl30(val interface{})
	EnableFrontendSsl30Input() interface{}
	EnableFrontendTls10() interface{}
	SetEnableFrontendTls10(val interface{})
	EnableFrontendTls10Input() interface{}
	EnableFrontendTls11() interface{}
	SetEnableFrontendTls11(val interface{})
	EnableFrontendTls11Input() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementSecurity
	SetInternalValue(val *ApiManagementSecurity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsEcdheRsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsEcdheRsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsEcdheRsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsEcdheRsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TlsRsaWithAes128CbcSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes128CbcSha256CiphersEnabled(val interface{})
	TlsRsaWithAes128CbcSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes128CbcShaCiphersEnabled() interface{}
	SetTlsRsaWithAes128CbcShaCiphersEnabled(val interface{})
	TlsRsaWithAes128CbcShaCiphersEnabledInput() interface{}
	TlsRsaWithAes128GcmSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes128GcmSha256CiphersEnabled(val interface{})
	TlsRsaWithAes128GcmSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes256CbcSha256CiphersEnabled() interface{}
	SetTlsRsaWithAes256CbcSha256CiphersEnabled(val interface{})
	TlsRsaWithAes256CbcSha256CiphersEnabledInput() interface{}
	TlsRsaWithAes256CbcShaCiphersEnabled() interface{}
	SetTlsRsaWithAes256CbcShaCiphersEnabled(val interface{})
	TlsRsaWithAes256CbcShaCiphersEnabledInput() interface{}
	TripleDesCiphersEnabled() interface{}
	SetTripleDesCiphersEnabled(val interface{})
	TripleDesCiphersEnabledInput() interface{}
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
	ResetEnableBackendSsl30()
	ResetEnableBackendTls10()
	ResetEnableBackendTls11()
	ResetEnableFrontendSsl30()
	ResetEnableFrontendTls10()
	ResetEnableFrontendTls11()
	ResetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled()
	ResetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled()
	ResetTlsEcdheRsaWithAes128CbcShaCiphersEnabled()
	ResetTlsEcdheRsaWithAes256CbcShaCiphersEnabled()
	ResetTlsRsaWithAes128CbcSha256CiphersEnabled()
	ResetTlsRsaWithAes128CbcShaCiphersEnabled()
	ResetTlsRsaWithAes128GcmSha256CiphersEnabled()
	ResetTlsRsaWithAes256CbcSha256CiphersEnabled()
	ResetTlsRsaWithAes256CbcShaCiphersEnabled()
	ResetTripleDesCiphersEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementSecurityOutputReference
type jsiiProxy_ApiManagementSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendSsl30() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendSsl30",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendSsl30Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendSsl30Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls11() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableBackendTls11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBackendTls11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendSsl30() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendSsl30",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendSsl30Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendSsl30Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls10() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls10Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls11() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) EnableFrontendTls11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFrontendTls11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) InternalValue() *ApiManagementSecurity {
	var returns *ApiManagementSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsEcdheRsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128GcmSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes128GcmSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcSha256CiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcSha256CiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcShaCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TlsRsaWithAes256CbcShaCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TripleDesCiphersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tripleDesCiphersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) TripleDesCiphersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tripleDesCiphersEnabledInput",
		&returns,
	)
	return returns
}


func NewApiManagementSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementSecurityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementSecurityOutputReference_Override(a ApiManagementSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableBackendSsl30(val interface{}) {
	_jsii_.Set(
		j,
		"enableBackendSsl30",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableBackendTls10(val interface{}) {
	_jsii_.Set(
		j,
		"enableBackendTls10",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableBackendTls11(val interface{}) {
	_jsii_.Set(
		j,
		"enableBackendTls11",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableFrontendSsl30(val interface{}) {
	_jsii_.Set(
		j,
		"enableFrontendSsl30",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableFrontendTls10(val interface{}) {
	_jsii_.Set(
		j,
		"enableFrontendTls10",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetEnableFrontendTls11(val interface{}) {
	_jsii_.Set(
		j,
		"enableFrontendTls11",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetInternalValue(val *ApiManagementSecurity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsEcdheRsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsEcdheRsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsRsaWithAes128CbcSha256CiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsRsaWithAes128CbcSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsRsaWithAes128CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsRsaWithAes128CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsRsaWithAes128GcmSha256CiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsRsaWithAes128GcmSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsRsaWithAes256CbcSha256CiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsRsaWithAes256CbcSha256CiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTlsRsaWithAes256CbcShaCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsRsaWithAes256CbcShaCiphersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSecurityOutputReference) SetTripleDesCiphersEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tripleDesCiphersEnabled",
		val,
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendSsl30() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendSsl30",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendTls10() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendTls10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableBackendTls11() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBackendTls11",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendSsl30() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendSsl30",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendTls10() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendTls10",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetEnableFrontendTls11() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFrontendTls11",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheEcdsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheEcdsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheRsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheRsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsEcdheRsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsEcdheRsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128CbcSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128CbcSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes128GcmSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes128GcmSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes256CbcSha256CiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes256CbcSha256CiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTlsRsaWithAes256CbcShaCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsRsaWithAes256CbcShaCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ResetTripleDesCiphersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTripleDesCiphersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementSignIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

type ApiManagementSignInOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementSignIn
	SetInternalValue(val *ApiManagementSignIn)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementSignInOutputReference
type jsiiProxy_ApiManagementSignInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) InternalValue() *ApiManagementSignIn {
	var returns *ApiManagementSignIn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementSignInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementSignInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementSignInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementSignInOutputReference_Override(a ApiManagementSignInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetInternalValue(val *ApiManagementSignIn) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementSignUp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// terms_of_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#terms_of_service ApiManagement#terms_of_service}
	TermsOfService *ApiManagementSignUpTermsOfService `field:"required" json:"termsOfService" yaml:"termsOfService"`
}

type ApiManagementSignUpOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementSignUp
	SetInternalValue(val *ApiManagementSignUp)
	TermsOfService() ApiManagementSignUpTermsOfServiceOutputReference
	TermsOfServiceInput() *ApiManagementSignUpTermsOfService
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
	PutTermsOfService(value *ApiManagementSignUpTermsOfService)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementSignUpOutputReference
type jsiiProxy_ApiManagementSignUpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) InternalValue() *ApiManagementSignUp {
	var returns *ApiManagementSignUp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) TermsOfService() ApiManagementSignUpTermsOfServiceOutputReference {
	var returns ApiManagementSignUpTermsOfServiceOutputReference
	_jsii_.Get(
		j,
		"termsOfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) TermsOfServiceInput() *ApiManagementSignUpTermsOfService {
	var returns *ApiManagementSignUpTermsOfService
	_jsii_.Get(
		j,
		"termsOfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementSignUpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementSignUpOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementSignUpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignUpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementSignUpOutputReference_Override(a ApiManagementSignUpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignUpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetInternalValue(val *ApiManagementSignUp) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) PutTermsOfService(value *ApiManagementSignUpTermsOfService) {
	_jsii_.InvokeVoid(
		a,
		"putTermsOfService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementSignUpTermsOfService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#consent_required ApiManagement#consent_required}.
	ConsentRequired interface{} `field:"required" json:"consentRequired" yaml:"consentRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#text ApiManagement#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

type ApiManagementSignUpTermsOfServiceOutputReference interface {
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
	ConsentRequired() interface{}
	SetConsentRequired(val interface{})
	ConsentRequiredInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementSignUpTermsOfService
	SetInternalValue(val *ApiManagementSignUpTermsOfService)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
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
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementSignUpTermsOfServiceOutputReference
type jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ConsentRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consentRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ConsentRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consentRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) InternalValue() *ApiManagementSignUpTermsOfService {
	var returns *ApiManagementSignUpTermsOfService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewApiManagementSignUpTermsOfServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementSignUpTermsOfServiceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignUpTermsOfServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementSignUpTermsOfServiceOutputReference_Override(a ApiManagementSignUpTermsOfServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementSignUpTermsOfServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetConsentRequired(val interface{}) {
	_jsii_.Set(
		j,
		"consentRequired",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetInternalValue(val *ApiManagementSignUpTermsOfService) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) SetText(val *string) {
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		a,
		"resetText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementSignUpTermsOfServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementTenantAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

type ApiManagementTenantAccessOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementTenantAccess
	SetInternalValue(val *ApiManagementTenantAccess)
	PrimaryKey() *string
	SecondaryKey() *string
	TenantId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementTenantAccessOutputReference
type jsiiProxy_ApiManagementTenantAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) InternalValue() *ApiManagementTenantAccess {
	var returns *ApiManagementTenantAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SecondaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementTenantAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementTenantAccessOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementTenantAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementTenantAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementTenantAccessOutputReference_Override(a ApiManagementTenantAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementTenantAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetInternalValue(val *ApiManagementTenantAccess) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTenantAccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTenantAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#create ApiManagement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#delete ApiManagement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#read ApiManagement#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#update ApiManagement#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ApiManagementTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementTimeoutsOutputReference
type jsiiProxy_ApiManagementTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewApiManagementTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementTimeoutsOutputReference_Override(a ApiManagementTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiManagementVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#subnet_id ApiManagement#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

type ApiManagementVirtualNetworkConfigurationOutputReference interface {
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
	InternalValue() *ApiManagementVirtualNetworkConfiguration
	SetInternalValue(val *ApiManagementVirtualNetworkConfiguration)
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementVirtualNetworkConfigurationOutputReference
type jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) InternalValue() *ApiManagementVirtualNetworkConfiguration {
	var returns *ApiManagementVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementVirtualNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementVirtualNetworkConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementVirtualNetworkConfigurationOutputReference_Override(a ApiManagementVirtualNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetInternalValue(val *ApiManagementVirtualNetworkConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementVirtualNetworkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

