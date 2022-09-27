// Code generated by generators/resource/main.go; DO NOT EDIT.

package voiceid

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_voiceid_domain", domainResourceType)
}

// domainResourceType returns the Terraform awscc_voiceid_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::VoiceID::Domain resource type.
func domainResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-%@]*)$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1024),
				validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-%@]*)$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Description is a write-only property.
		},
		"domain_id": {
			// Property: DomainId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 22,
			//   "minLength": 22,
			//   "pattern": "^[a-zA-Z0-9]{22}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9_-]*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9_-]*$"), ""),
			},
			// Name is a write-only property.
		},
		"server_side_encryption_configuration": {
			// Property: ServerSideEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KmsKeyId": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "KmsKeyId"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
				},
			),
			Required: true,
			// ServerSideEncryptionConfiguration is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
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
		Description: "The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::VoiceID::Domain").WithTerraformTypeName("awscc_voiceid_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                          "Description",
		"domain_id":                            "DomainId",
		"key":                                  "Key",
		"kms_key_id":                           "KmsKeyId",
		"name":                                 "Name",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"tags":                                 "Tags",
		"value":                                "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Description",
		"/properties/Name",
		"/properties/ServerSideEncryptionConfiguration",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
