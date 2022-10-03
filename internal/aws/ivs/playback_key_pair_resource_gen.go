// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ivs_playback_key_pair", playbackKeyPairResource)
}

// playbackKeyPairResource returns the Terraform awscc_ivs_playback_key_pair resource.
// This Terraform resource corresponds to the CloudFormation AWS::IVS::PlaybackKeyPair resource.
func playbackKeyPairResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-pair identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:playback-key/[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Description: "Key-pair identifier.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"fingerprint": {
			// Property: Fingerprint
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-pair identifier.",
			//   "type": "string"
			// }
			Description: "Key-pair identifier.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "^[a-zA-Z0-9-_]*$",
			//   "type": "string"
			// }
			Description: "An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 128),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"public_key_material": {
			// Property: PublicKeyMaterial
			// CloudFormation resource type schema:
			// {
			//   "description": "The public portion of a customer-generated key pair.",
			//   "type": "string"
			// }
			Description: "The public portion of a customer-generated key pair.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset model.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
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
		Description: "Resource Type definition for AWS::IVS::PlaybackKeyPair",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::PlaybackKeyPair").WithTerraformTypeName("awscc_ivs_playback_key_pair")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"fingerprint":         "Fingerprint",
		"key":                 "Key",
		"name":                "Name",
		"public_key_material": "PublicKeyMaterial",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
