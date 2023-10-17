// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorroute

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRoute",
		reflect.TypeOf((*CdnFrontdoorRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cache", GoGetter: "Cache"},
			_jsii_.MemberProperty{JsiiProperty: "cacheInput", GoGetter: "CacheInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorCustomDomainIds", GoGetter: "CdnFrontdoorCustomDomainIds"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorCustomDomainIdsInput", GoGetter: "CdnFrontdoorCustomDomainIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorEndpointId", GoGetter: "CdnFrontdoorEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorEndpointIdInput", GoGetter: "CdnFrontdoorEndpointIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginGroupId", GoGetter: "CdnFrontdoorOriginGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginGroupIdInput", GoGetter: "CdnFrontdoorOriginGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginIds", GoGetter: "CdnFrontdoorOriginIds"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginIdsInput", GoGetter: "CdnFrontdoorOriginIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginPath", GoGetter: "CdnFrontdoorOriginPath"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorOriginPathInput", GoGetter: "CdnFrontdoorOriginPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorRuleSetIds", GoGetter: "CdnFrontdoorRuleSetIds"},
			_jsii_.MemberProperty{JsiiProperty: "cdnFrontdoorRuleSetIdsInput", GoGetter: "CdnFrontdoorRuleSetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "forwardingProtocol", GoGetter: "ForwardingProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "forwardingProtocolInput", GoGetter: "ForwardingProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpsRedirectEnabled", GoGetter: "HttpsRedirectEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "httpsRedirectEnabledInput", GoGetter: "HttpsRedirectEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "linkToDefaultDomain", GoGetter: "LinkToDefaultDomain"},
			_jsii_.MemberProperty{JsiiProperty: "linkToDefaultDomainInput", GoGetter: "LinkToDefaultDomainInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "patternsToMatch", GoGetter: "PatternsToMatch"},
			_jsii_.MemberProperty{JsiiProperty: "patternsToMatchInput", GoGetter: "PatternsToMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCache", GoMethod: "PutCache"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCache", GoMethod: "ResetCache"},
			_jsii_.MemberMethod{JsiiMethod: "resetCdnFrontdoorCustomDomainIds", GoMethod: "ResetCdnFrontdoorCustomDomainIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetCdnFrontdoorOriginPath", GoMethod: "ResetCdnFrontdoorOriginPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetCdnFrontdoorRuleSetIds", GoMethod: "ResetCdnFrontdoorRuleSetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardingProtocol", GoMethod: "ResetForwardingProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirectEnabled", GoMethod: "ResetHttpsRedirectEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinkToDefaultDomain", GoMethod: "ResetLinkToDefaultDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "supportedProtocols", GoGetter: "SupportedProtocols"},
			_jsii_.MemberProperty{JsiiProperty: "supportedProtocolsInput", GoGetter: "SupportedProtocolsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CdnFrontdoorRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCache",
		reflect.TypeOf((*CdnFrontdoorRouteCache)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCacheOutputReference",
		reflect.TypeOf((*CdnFrontdoorRouteCacheOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compressionEnabled", GoGetter: "CompressionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "compressionEnabledInput", GoGetter: "CompressionEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypesToCompress", GoGetter: "ContentTypesToCompress"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypesToCompressInput", GoGetter: "ContentTypesToCompressInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCachingBehavior", GoGetter: "QueryStringCachingBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCachingBehaviorInput", GoGetter: "QueryStringCachingBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsInput", GoGetter: "QueryStringsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompressionEnabled", GoMethod: "ResetCompressionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentTypesToCompress", GoMethod: "ResetContentTypesToCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCachingBehavior", GoMethod: "ResetQueryStringCachingBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStrings", GoMethod: "ResetQueryStrings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CdnFrontdoorRouteCacheOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteConfig",
		reflect.TypeOf((*CdnFrontdoorRouteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteTimeouts",
		reflect.TypeOf((*CdnFrontdoorRouteTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteTimeoutsOutputReference",
		reflect.TypeOf((*CdnFrontdoorRouteTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_CdnFrontdoorRouteTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
