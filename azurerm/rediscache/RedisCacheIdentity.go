package rediscache


type RedisCacheIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#type RedisCache#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#identity_ids RedisCache#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}
