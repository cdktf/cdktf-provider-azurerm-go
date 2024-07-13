// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot


type LinuxFunctionAppSlotAuthSettingsV2 struct {
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#login LinuxFunctionAppSlot#login}
	Login *LinuxFunctionAppSlotAuthSettingsV2Login `field:"required" json:"login" yaml:"login"`
	// active_directory_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#active_directory_v2 LinuxFunctionAppSlot#active_directory_v2}
	ActiveDirectoryV2 *LinuxFunctionAppSlotAuthSettingsV2ActiveDirectoryV2 `field:"optional" json:"activeDirectoryV2" yaml:"activeDirectoryV2"`
	// apple_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#apple_v2 LinuxFunctionAppSlot#apple_v2}
	AppleV2 *LinuxFunctionAppSlotAuthSettingsV2AppleV2 `field:"optional" json:"appleV2" yaml:"appleV2"`
	// Should the AuthV2 Settings be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#auth_enabled LinuxFunctionAppSlot#auth_enabled}
	AuthEnabled interface{} `field:"optional" json:"authEnabled" yaml:"authEnabled"`
	// azure_static_web_app_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#azure_static_web_app_v2 LinuxFunctionAppSlot#azure_static_web_app_v2}
	AzureStaticWebAppV2 *LinuxFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2 `field:"optional" json:"azureStaticWebAppV2" yaml:"azureStaticWebAppV2"`
	// The path to the App Auth settings. **Note:** Relative Paths are evaluated from the Site Root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#config_file_path LinuxFunctionAppSlot#config_file_path}
	ConfigFilePath *string `field:"optional" json:"configFilePath" yaml:"configFilePath"`
	// custom_oidc_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#custom_oidc_v2 LinuxFunctionAppSlot#custom_oidc_v2}
	CustomOidcV2 interface{} `field:"optional" json:"customOidcV2" yaml:"customOidcV2"`
	// The Default Authentication Provider to use when the `unauthenticated_action` is set to `RedirectToLoginPage`.
	//
	// Possible values include: `apple`, `azureactivedirectory`, `facebook`, `github`, `google`, `twitter` and the `name` of your `custom_oidc_v2` provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#default_provider LinuxFunctionAppSlot#default_provider}
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// The paths which should be excluded from the `unauthenticated_action` when it is set to `RedirectToLoginPage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#excluded_paths LinuxFunctionAppSlot#excluded_paths}
	ExcludedPaths *[]*string `field:"optional" json:"excludedPaths" yaml:"excludedPaths"`
	// facebook_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#facebook_v2 LinuxFunctionAppSlot#facebook_v2}
	FacebookV2 *LinuxFunctionAppSlotAuthSettingsV2FacebookV2 `field:"optional" json:"facebookV2" yaml:"facebookV2"`
	// The convention used to determine the url of the request made.
	//
	// Possible values include `ForwardProxyConventionNoProxy`, `ForwardProxyConventionStandard`, `ForwardProxyConventionCustom`. Defaults to `ForwardProxyConventionNoProxy`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#forward_proxy_convention LinuxFunctionAppSlot#forward_proxy_convention}
	ForwardProxyConvention *string `field:"optional" json:"forwardProxyConvention" yaml:"forwardProxyConvention"`
	// The name of the header containing the host of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#forward_proxy_custom_host_header_name LinuxFunctionAppSlot#forward_proxy_custom_host_header_name}
	ForwardProxyCustomHostHeaderName *string `field:"optional" json:"forwardProxyCustomHostHeaderName" yaml:"forwardProxyCustomHostHeaderName"`
	// The name of the header containing the scheme of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#forward_proxy_custom_scheme_header_name LinuxFunctionAppSlot#forward_proxy_custom_scheme_header_name}
	ForwardProxyCustomSchemeHeaderName *string `field:"optional" json:"forwardProxyCustomSchemeHeaderName" yaml:"forwardProxyCustomSchemeHeaderName"`
	// github_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#github_v2 LinuxFunctionAppSlot#github_v2}
	GithubV2 *LinuxFunctionAppSlotAuthSettingsV2GithubV2 `field:"optional" json:"githubV2" yaml:"githubV2"`
	// google_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#google_v2 LinuxFunctionAppSlot#google_v2}
	GoogleV2 *LinuxFunctionAppSlotAuthSettingsV2GoogleV2 `field:"optional" json:"googleV2" yaml:"googleV2"`
	// The prefix that should precede all the authentication and authorisation paths. Defaults to `/.auth`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#http_route_api_prefix LinuxFunctionAppSlot#http_route_api_prefix}
	HttpRouteApiPrefix *string `field:"optional" json:"httpRouteApiPrefix" yaml:"httpRouteApiPrefix"`
	// microsoft_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#microsoft_v2 LinuxFunctionAppSlot#microsoft_v2}
	MicrosoftV2 *LinuxFunctionAppSlotAuthSettingsV2MicrosoftV2 `field:"optional" json:"microsoftV2" yaml:"microsoftV2"`
	// Should the authentication flow be used for all requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#require_authentication LinuxFunctionAppSlot#require_authentication}
	RequireAuthentication interface{} `field:"optional" json:"requireAuthentication" yaml:"requireAuthentication"`
	// Should HTTPS be required on connections? Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#require_https LinuxFunctionAppSlot#require_https}
	RequireHttps interface{} `field:"optional" json:"requireHttps" yaml:"requireHttps"`
	// The Runtime Version of the Authentication and Authorisation feature of this App. Defaults to `~1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#runtime_version LinuxFunctionAppSlot#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// twitter_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#twitter_v2 LinuxFunctionAppSlot#twitter_v2}
	TwitterV2 *LinuxFunctionAppSlotAuthSettingsV2TwitterV2 `field:"optional" json:"twitterV2" yaml:"twitterV2"`
	// The action to take for requests made without authentication.
	//
	// Possible values include `RedirectToLoginPage`, `AllowAnonymous`, `Return401`, and `Return403`. Defaults to `RedirectToLoginPage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app_slot#unauthenticated_action LinuxFunctionAppSlot#unauthenticated_action}
	UnauthenticatedAction *string `field:"optional" json:"unauthenticatedAction" yaml:"unauthenticatedAction"`
}

