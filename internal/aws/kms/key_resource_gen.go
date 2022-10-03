// Code generated by generators/resource/main.go; DO NOT EDIT.

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_kms_key", keyResource)
}

// keyResource returns the Terraform awscc_kms_key resource.
// This Terraform resource corresponds to the CloudFormation AWS::KMS::Key resource.
func keyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.",
			//   "maxLength": 8192,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 8192),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enable_key_rotation": {
			// Property: EnableKeyRotation
			// CloudFormation resource type schema:
			// {
			//   "description": "Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.",
			//   "type": "boolean"
			// }
			Description: "Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"key_id": {
			// Property: KeyId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"key_policy": {
			// Property: KeyPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.",
			//   "type": "string"
			// }
			Description: "The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.",
			Type:        types.StringType,
			Required:    true,
		},
		"key_spec": {
			// Property: KeySpec
			// CloudFormation resource type schema:
			// {
			//   "default": "SYMMETRIC_DEFAULT",
			//   "description": "Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.",
			//   "enum": [
			//     "SYMMETRIC_DEFAULT",
			//     "RSA_2048",
			//     "RSA_3072",
			//     "RSA_4096",
			//     "ECC_NIST_P256",
			//     "ECC_NIST_P384",
			//     "ECC_NIST_P521",
			//     "ECC_SECG_P256K1",
			//     "HMAC_224",
			//     "HMAC_256",
			//     "HMAC_384",
			//     "HMAC_512",
			//     "SM2"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SYMMETRIC_DEFAULT",
					"RSA_2048",
					"RSA_3072",
					"RSA_4096",
					"ECC_NIST_P256",
					"ECC_NIST_P384",
					"ECC_NIST_P521",
					"ECC_SECG_P256K1",
					"HMAC_224",
					"HMAC_256",
					"HMAC_384",
					"HMAC_512",
					"SM2",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: "SYMMETRIC_DEFAULT"}),
				resource.UseStateForUnknown(),
			},
		},
		"key_usage": {
			// Property: KeyUsage
			// CloudFormation resource type schema:
			// {
			//   "default": "ENCRYPT_DECRYPT",
			//   "description": "Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.",
			//   "enum": [
			//     "ENCRYPT_DECRYPT",
			//     "SIGN_VERIFY",
			//     "GENERATE_VERIFY_MAC"
			//   ],
			//   "type": "string"
			// }
			Description: "Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ENCRYPT_DECRYPT",
					"SIGN_VERIFY",
					"GENERATE_VERIFY_MAC",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: "ENCRYPT_DECRYPT"}),
				resource.UseStateForUnknown(),
			},
		},
		"multi_region": {
			// Property: MultiRegion
			// CloudFormation resource type schema:
			// {
			//   "default": false,
			//   "description": "Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Bool{Value: false}),
				resource.UseStateForUnknown(),
			},
		},
		"pending_window_in_days": {
			// Property: PendingWindowInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			//   "maximum": 30,
			//   "minimum": 7,
			//   "type": "integer"
			// }
			Description: "Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(7, 30),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// PendingWindowInDays is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::Key").WithTerraformTypeName("awscc_kms_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"description":            "Description",
		"enable_key_rotation":    "EnableKeyRotation",
		"enabled":                "Enabled",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"key_policy":             "KeyPolicy",
		"key_spec":               "KeySpec",
		"key_usage":              "KeyUsage",
		"multi_region":           "MultiRegion",
		"pending_window_in_days": "PendingWindowInDays",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/PendingWindowInDays",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
