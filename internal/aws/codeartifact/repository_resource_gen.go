// Code generated by generators/resource/main.go; DO NOT EDIT.

package codeartifact

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
	registry.AddResourceFactory("awscc_codeartifact_repository", repositoryResource)
}

// repositoryResource returns the Terraform awscc_codeartifact_repository resource.
// This Terraform resource corresponds to the CloudFormation AWS::CodeArtifact::Repository resource.
func repositoryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the repository.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ARN of the repository.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A text description of the repository.",
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Description: "A text description of the repository.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the domain that contains the repository.",
			//   "maxLength": 50,
			//   "minLength": 2,
			//   "pattern": "^([a-z][a-z0-9\\-]{0,48}[a-z0-9])$",
			//   "type": "string"
			// }
			Description: "The name of the domain that contains the repository.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(2, 50),
				validate.StringMatch(regexp.MustCompile("^([a-z][a-z0-9\\-]{0,48}[a-z0-9])$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"domain_owner": {
			// Property: DomainOwner
			// CloudFormation resource type schema:
			// {
			//   "description": "The 12-digit account ID of the AWS account that owns the domain.",
			//   "pattern": "[0-9]{12}",
			//   "type": "string"
			// }
			Description: "The 12-digit account ID of the AWS account that owns the domain.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"external_connections": {
			// Property: ExternalConnections
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of external connections associated with the repository.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of external connections associated with the repository.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the repository. This is used for GetAtt",
			//   "maxLength": 100,
			//   "minLength": 2,
			//   "pattern": "^([A-Za-z0-9][A-Za-z0-9._\\-]{1,99})$",
			//   "type": "string"
			// }
			Description: "The name of the repository. This is used for GetAtt",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"permissions_policy_document": {
			// Property: PermissionsPolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "The access control resource policy on the provided repository.",
			//   "maxLength": 5120,
			//   "minLength": 2,
			//   "type": "object"
			// }
			Description: "The access control resource policy on the provided repository.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"repository_name": {
			// Property: RepositoryName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the repository.",
			//   "maxLength": 100,
			//   "minLength": 2,
			//   "pattern": "^([A-Za-z0-9][A-Za-z0-9._\\-]{1,99})$",
			//   "type": "string"
			// }
			Description: "The name of the repository.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(2, 100),
				validate.StringMatch(regexp.MustCompile("^([A-Za-z0-9][A-Za-z0-9._\\-]{1,99})$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		"upstreams": {
			// Property: Upstreams
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of upstream repositories associated with the repository.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of upstream repositories associated with the repository.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
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
		Description: "The resource schema to create a CodeArtifact repository.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeArtifact::Repository").WithTerraformTypeName("awscc_codeartifact_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"description":                 "Description",
		"domain_name":                 "DomainName",
		"domain_owner":                "DomainOwner",
		"external_connections":        "ExternalConnections",
		"key":                         "Key",
		"name":                        "Name",
		"permissions_policy_document": "PermissionsPolicyDocument",
		"repository_name":             "RepositoryName",
		"tags":                        "Tags",
		"upstreams":                   "Upstreams",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
