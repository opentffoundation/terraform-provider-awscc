// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudfront_public_key", publicKeyResource)
}

// publicKeyResource returns the Terraform awscc_cloudfront_public_key resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFront::PublicKey resource.
func publicKeyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_time": {
			// Property: CreatedTime
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
		"id": {
			// Property: Id
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
		"public_key_config": {
			// Property: PublicKeyConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CallerReference": {
			//       "type": "string"
			//     },
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "EncodedKey": {
			//       "type": "string"
			//     },
			//     "Name": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "CallerReference",
			//     "Name",
			//     "EncodedKey"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"caller_reference": {
						// Property: CallerReference
						Type:     types.StringType,
						Required: true,
					},
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"encoded_key": {
						// Property: EncodedKey
						Type:     types.StringType,
						Required: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CloudFront::PublicKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::PublicKey").WithTerraformTypeName("awscc_cloudfront_public_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"caller_reference":  "CallerReference",
		"comment":           "Comment",
		"created_time":      "CreatedTime",
		"encoded_key":       "EncodedKey",
		"id":                "Id",
		"name":              "Name",
		"public_key_config": "PublicKeyConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
