/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Distribution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DistributionSpec   `json:"spec,omitempty"`
	Status            DistributionStatus `json:"status,omitempty"`
}

type DistributionSpec struct {
	DistributionSpec2 `json:",inline"`
	// +optional
	KubeformOutput DistributionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type DistributionSpecCustomErrorResponse struct {
	// +optional
	ErrorCachingMinTtl *int64 `json:"errorCachingMinTtl,omitempty" tf:"error_caching_min_ttl"`
	ErrorCode          *int64 `json:"errorCode" tf:"error_code"`
	// +optional
	ResponseCode *int64 `json:"responseCode,omitempty" tf:"response_code"`
	// +optional
	ResponsePagePath *string `json:"responsePagePath,omitempty" tf:"response_page_path"`
}

type DistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward *string `json:"forward" tf:"forward"`
	// +optional
	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names"`
}

type DistributionSpecDefaultCacheBehaviorForwardedValues struct {
	Cookies *DistributionSpecDefaultCacheBehaviorForwardedValuesCookies `json:"cookies" tf:"cookies"`
	// +optional
	Headers     []string `json:"headers,omitempty" tf:"headers"`
	QueryString *bool    `json:"queryString" tf:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys"`
}

type DistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType *string `json:"eventType" tf:"event_type"`
	// +optional
	IncludeBody *bool   `json:"includeBody,omitempty" tf:"include_body"`
	LambdaArn   *string `json:"lambdaArn" tf:"lambda_arn"`
}

type DistributionSpecDefaultCacheBehavior struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +optional
	CachePolicyID *string  `json:"cachePolicyID,omitempty" tf:"cache_policy_id"`
	CachedMethods []string `json:"cachedMethods" tf:"cached_methods"`
	// +optional
	Compress *bool `json:"compress,omitempty" tf:"compress"`
	// +optional
	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`
	// +optional
	FieldLevelEncryptionID *string `json:"fieldLevelEncryptionID,omitempty" tf:"field_level_encryption_id"`
	// +optional
	ForwardedValues *DistributionSpecDefaultCacheBehaviorForwardedValues `json:"forwardedValues,omitempty" tf:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	LambdaFunctionAssociation []DistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association"`
	// +optional
	MaxTtl *int64 `json:"maxTtl,omitempty" tf:"max_ttl"`
	// +optional
	MinTtl *int64 `json:"minTtl,omitempty" tf:"min_ttl"`
	// +optional
	OriginRequestPolicyID *string `json:"originRequestPolicyID,omitempty" tf:"origin_request_policy_id"`
	// +optional
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn"`
	// +optional
	SmoothStreaming *bool   `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`
	TargetOriginID  *string `json:"targetOriginID" tf:"target_origin_id"`
	// +optional
	TrustedKeyGroups []string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups"`
	// +optional
	TrustedSigners       []string `json:"trustedSigners,omitempty" tf:"trusted_signers"`
	ViewerProtocolPolicy *string  `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type DistributionSpecLoggingConfig struct {
	Bucket *string `json:"bucket" tf:"bucket"`
	// +optional
	IncludeCookies *bool `json:"includeCookies,omitempty" tf:"include_cookies"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type DistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward *string `json:"forward" tf:"forward"`
	// +optional
	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names"`
}

type DistributionSpecOrderedCacheBehaviorForwardedValues struct {
	Cookies *DistributionSpecOrderedCacheBehaviorForwardedValuesCookies `json:"cookies" tf:"cookies"`
	// +optional
	Headers     []string `json:"headers,omitempty" tf:"headers"`
	QueryString *bool    `json:"queryString" tf:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys"`
}

type DistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType *string `json:"eventType" tf:"event_type"`
	// +optional
	IncludeBody *bool   `json:"includeBody,omitempty" tf:"include_body"`
	LambdaArn   *string `json:"lambdaArn" tf:"lambda_arn"`
}

type DistributionSpecOrderedCacheBehavior struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +optional
	CachePolicyID *string  `json:"cachePolicyID,omitempty" tf:"cache_policy_id"`
	CachedMethods []string `json:"cachedMethods" tf:"cached_methods"`
	// +optional
	Compress *bool `json:"compress,omitempty" tf:"compress"`
	// +optional
	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`
	// +optional
	FieldLevelEncryptionID *string `json:"fieldLevelEncryptionID,omitempty" tf:"field_level_encryption_id"`
	// +optional
	ForwardedValues *DistributionSpecOrderedCacheBehaviorForwardedValues `json:"forwardedValues,omitempty" tf:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	LambdaFunctionAssociation []DistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association"`
	// +optional
	MaxTtl *int64 `json:"maxTtl,omitempty" tf:"max_ttl"`
	// +optional
	MinTtl *int64 `json:"minTtl,omitempty" tf:"min_ttl"`
	// +optional
	OriginRequestPolicyID *string `json:"originRequestPolicyID,omitempty" tf:"origin_request_policy_id"`
	PathPattern           *string `json:"pathPattern" tf:"path_pattern"`
	// +optional
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn"`
	// +optional
	SmoothStreaming *bool   `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`
	TargetOriginID  *string `json:"targetOriginID" tf:"target_origin_id"`
	// +optional
	TrustedKeyGroups []string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups"`
	// +optional
	TrustedSigners       []string `json:"trustedSigners,omitempty" tf:"trusted_signers"`
	ViewerProtocolPolicy *string  `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type DistributionSpecOriginCustomHeader struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type DistributionSpecOriginCustomOriginConfig struct {
	HttpPort  *int64 `json:"httpPort" tf:"http_port"`
	HttpsPort *int64 `json:"httpsPort" tf:"https_port"`
	// +optional
	OriginKeepaliveTimeout *int64  `json:"originKeepaliveTimeout,omitempty" tf:"origin_keepalive_timeout"`
	OriginProtocolPolicy   *string `json:"originProtocolPolicy" tf:"origin_protocol_policy"`
	// +optional
	OriginReadTimeout  *int64   `json:"originReadTimeout,omitempty" tf:"origin_read_timeout"`
	OriginSSLProtocols []string `json:"originSSLProtocols" tf:"origin_ssl_protocols"`
}

type DistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity *string `json:"originAccessIdentity" tf:"origin_access_identity"`
}

type DistributionSpecOrigin struct {
	// +optional
	CustomHeader []DistributionSpecOriginCustomHeader `json:"customHeader,omitempty" tf:"custom_header"`
	// +optional
	CustomOriginConfig *DistributionSpecOriginCustomOriginConfig `json:"customOriginConfig,omitempty" tf:"custom_origin_config"`
	DomainName         *string                                   `json:"domainName" tf:"domain_name"`
	OriginID           *string                                   `json:"originID" tf:"origin_id"`
	// +optional
	OriginPath *string `json:"originPath,omitempty" tf:"origin_path"`
	// +optional
	S3OriginConfig *DistributionSpecOriginS3OriginConfig `json:"s3OriginConfig,omitempty" tf:"s3_origin_config"`
}

type DistributionSpecOriginGroupFailoverCriteria struct {
	StatusCodes []int64 `json:"statusCodes" tf:"status_codes"`
}

type DistributionSpecOriginGroupMember struct {
	OriginID *string `json:"originID" tf:"origin_id"`
}

type DistributionSpecOriginGroup struct {
	FailoverCriteria *DistributionSpecOriginGroupFailoverCriteria `json:"failoverCriteria" tf:"failover_criteria"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	Member   []DistributionSpecOriginGroupMember `json:"member" tf:"member"`
	OriginID *string                             `json:"originID" tf:"origin_id"`
}

type DistributionSpecRestrictionsGeoRestriction struct {
	// +optional
	Locations       []string `json:"locations,omitempty" tf:"locations"`
	RestrictionType *string  `json:"restrictionType" tf:"restriction_type"`
}

type DistributionSpecRestrictions struct {
	GeoRestriction *DistributionSpecRestrictionsGeoRestriction `json:"geoRestriction" tf:"geo_restriction"`
}

type DistributionSpecTrustedKeyGroupsItems struct {
	// +optional
	KeyGroupID *string `json:"keyGroupID,omitempty" tf:"key_group_id"`
	// +optional
	KeyPairIDS []string `json:"keyPairIDS,omitempty" tf:"key_pair_ids"`
}

type DistributionSpecTrustedKeyGroups struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Items []DistributionSpecTrustedKeyGroupsItems `json:"items,omitempty" tf:"items"`
}

type DistributionSpecTrustedSignersItems struct {
	// +optional
	AwsAccountNumber *string `json:"awsAccountNumber,omitempty" tf:"aws_account_number"`
	// +optional
	KeyPairIDS []string `json:"keyPairIDS,omitempty" tf:"key_pair_ids"`
}

type DistributionSpecTrustedSigners struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Items []DistributionSpecTrustedSignersItems `json:"items,omitempty" tf:"items"`
}

type DistributionSpecViewerCertificate struct {
	// +optional
	AcmCertificateArn *string `json:"acmCertificateArn,omitempty" tf:"acm_certificate_arn"`
	// +optional
	CloudfrontDefaultCertificate *bool `json:"cloudfrontDefaultCertificate,omitempty" tf:"cloudfront_default_certificate"`
	// +optional
	IamCertificateID *string `json:"iamCertificateID,omitempty" tf:"iam_certificate_id"`
	// +optional
	MinimumProtocolVersion *string `json:"minimumProtocolVersion,omitempty" tf:"minimum_protocol_version"`
	// +optional
	SslSupportMethod *string `json:"sslSupportMethod,omitempty" tf:"ssl_support_method"`
}

type DistributionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Aliases []string `json:"aliases,omitempty" tf:"aliases"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference"`
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	// +optional
	CustomErrorResponse  []DistributionSpecCustomErrorResponse `json:"customErrorResponse,omitempty" tf:"custom_error_response"`
	DefaultCacheBehavior *DistributionSpecDefaultCacheBehavior `json:"defaultCacheBehavior" tf:"default_cache_behavior"`
	// +optional
	DefaultRootObject *string `json:"defaultRootObject,omitempty" tf:"default_root_object"`
	// +optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name"`
	Enabled    *bool   `json:"enabled" tf:"enabled"`
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// +optional
	HostedZoneID *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	// +optional
	HttpVersion *string `json:"httpVersion,omitempty" tf:"http_version"`
	// +optional
	InProgressValidationBatches *int64 `json:"inProgressValidationBatches,omitempty" tf:"in_progress_validation_batches"`
	// +optional
	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6_enabled"`
	// +optional
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time"`
	// +optional
	LoggingConfig *DistributionSpecLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config"`
	// +optional
	OrderedCacheBehavior []DistributionSpecOrderedCacheBehavior `json:"orderedCacheBehavior,omitempty" tf:"ordered_cache_behavior"`
	Origin               []DistributionSpecOrigin               `json:"origin" tf:"origin"`
	// +optional
	OriginGroup []DistributionSpecOriginGroup `json:"originGroup,omitempty" tf:"origin_group"`
	// +optional
	PriceClass   *string                       `json:"priceClass,omitempty" tf:"price_class"`
	Restrictions *DistributionSpecRestrictions `json:"restrictions" tf:"restrictions"`
	// +optional
	RetainOnDelete *bool `json:"retainOnDelete,omitempty" tf:"retain_on_delete"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TrustedKeyGroups []DistributionSpecTrustedKeyGroups `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups"`
	// +optional
	TrustedSigners    []DistributionSpecTrustedSigners   `json:"trustedSigners,omitempty" tf:"trusted_signers"`
	ViewerCertificate *DistributionSpecViewerCertificate `json:"viewerCertificate" tf:"viewer_certificate"`
	// +optional
	WaitForDeployment *bool `json:"waitForDeployment,omitempty" tf:"wait_for_deployment"`
	// +optional
	WebACLID *string `json:"webACLID,omitempty" tf:"web_acl_id"`
}

type DistributionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DistributionList is a list of Distributions
type DistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Distribution CRD objects
	Items []Distribution `json:"items,omitempty"`
}
