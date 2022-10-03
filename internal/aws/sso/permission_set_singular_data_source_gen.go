// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sso

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sso_permission_set", permissionSetDataSource)
}

// permissionSetDataSource returns the Terraform awscc_sso_permission_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSO::PermissionSet resource.
func permissionSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"customer_managed_policy_references": {
			// Property: CustomerManagedPolicyReferences
			// CloudFormation resource type schema:
			// {
			//   "default": [],
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "[\\w+=,.@-]+",
			//         "type": "string"
			//       },
			//       "Path": {
			//         "maxLength": 512,
			//         "minLength": 1,
			//         "pattern": "((/[A-Za-z0-9\\.,\\+@=_-]+)*)/",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"path": {
						// Property: Path
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The permission set description.",
			//   "maxLength": 700,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The permission set description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"inline_policy": {
			// Property: InlinePolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The inline policy to put in permission set.",
			//   "type": "string"
			// }
			Description: "The inline policy to put in permission set.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The sso instance arn that the permission set is owned.",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}",
			//   "type": "string"
			// }
			Description: "The sso instance arn that the permission set is owned.",
			Type:        types.StringType,
			Computed:    true,
		},
		"managed_policies": {
			// Property: ManagedPolicies
			// CloudFormation resource type schema:
			// {
			//   "default": [],
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "The managed policy to attach.",
			//     "maxLength": 2048,
			//     "minLength": 20,
			//     "type": "string"
			//   },
			//   "maxItems": 20,
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this permission set.",
			//   "maxLength": 32,
			//   "minLength": 1,
			//   "pattern": "[\\w+=,.@-]+",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this permission set.",
			Type:        types.StringType,
			Computed:    true,
		},
		"permission_set_arn": {
			// Property: PermissionSetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The permission set that the policy will be attached to",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}",
			//   "type": "string"
			// }
			Description: "The permission set that the policy will be attached to",
			Type:        types.StringType,
			Computed:    true,
		},
		"permissions_boundary": {
			// Property: PermissionsBoundary
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CustomerManagedPolicyReference": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Name": {
			//           "maxLength": 128,
			//           "minLength": 1,
			//           "pattern": "[\\w+=,.@-]+",
			//           "type": "string"
			//         },
			//         "Path": {
			//           "maxLength": 512,
			//           "minLength": 1,
			//           "pattern": "((/[A-Za-z0-9\\.,\\+@=_-]+)*)/",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Name"
			//       ],
			//       "type": "object"
			//     },
			//     "ManagedPolicyArn": {
			//       "description": "The managed policy to attach.",
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"customer_managed_policy_reference": {
						// Property: CustomerManagedPolicyReference
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"path": {
									// Property: Path
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"managed_policy_arn": {
						// Property: ManagedPolicyArn
						Description: "The managed policy to attach.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"relay_state_type": {
			// Property: RelayStateType
			// CloudFormation resource type schema:
			// {
			//   "description": "The relay state URL that redirect links to any service in the AWS Management Console.",
			//   "maxLength": 240,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9\u0026amp;$@#\\/%?=~\\-_'\u0026quot;|!:,.;*+\\[\\]\\ \\(\\)\\{\\}]+",
			//   "type": "string"
			// }
			Description: "The relay state URL that redirect links to any service in the AWS Management Console.",
			Type:        types.StringType,
			Computed:    true,
		},
		"session_duration": {
			// Property: SessionDuration
			// CloudFormation resource type schema:
			// {
			//   "description": "The length of time that a user can be signed in to an AWS account.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The length of time that a user can be signed in to an AWS account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The metadata that you apply to the permission set to help you categorize and organize them.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "[\\w+=,.@-]+",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "[\\w+=,.@-]+",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SSO::PermissionSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::PermissionSet").WithTerraformTypeName("awscc_sso_permission_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_managed_policy_reference":  "CustomerManagedPolicyReference",
		"customer_managed_policy_references": "CustomerManagedPolicyReferences",
		"description":                        "Description",
		"inline_policy":                      "InlinePolicy",
		"instance_arn":                       "InstanceArn",
		"key":                                "Key",
		"managed_policies":                   "ManagedPolicies",
		"managed_policy_arn":                 "ManagedPolicyArn",
		"name":                               "Name",
		"path":                               "Path",
		"permission_set_arn":                 "PermissionSetArn",
		"permissions_boundary":               "PermissionsBoundary",
		"relay_state_type":                   "RelayStateType",
		"session_duration":                   "SessionDuration",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
