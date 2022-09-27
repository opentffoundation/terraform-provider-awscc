// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

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
	registry.AddResourceTypeFactory("awscc_datasync_location_fsx_lustre", locationFSxLustreResourceType)
}

// locationFSxLustreResourceType returns the Terraform awscc_datasync_location_fsx_lustre resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationFSxLustre resource type.
func locationFSxLustreResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"fsx_filesystem_arn": {
			// Property: FsxFilesystemArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the FSx for Lustre file system.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the FSx for Lustre file system.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
				validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
			// FsxFilesystemArn is a write-only property.
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the FSx for Lustre location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "^(efs|nfs|s3|smb|fsxw|hdfs|fsxl)://[a-zA-Z0-9.:/\\-]+$",
			//   "type": "string"
			// }
			Description: "The URL of the FSx for Lustre location that was described.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"security_group_arns": {
			// Property: SecurityGroupArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARNs of the security groups that are to use to configure the FSx for Lustre file system.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/sg-[a-f0-9]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "The ARNs of the security groups that are to use to configure the FSx for Lustre file system.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 5),
				validate.ArrayForEach(validate.StringLenAtMost(128)),
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/sg-[a-f0-9]+$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.RequiresReplace(),
			},
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "A subdirectory in the location's path.",
			//   "maxLength": 4096,
			//   "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
			//   "type": "string"
			// }
			Description: "A subdirectory in the location's path.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(4096),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Subdirectory is a write-only property.
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
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
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
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
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
		Description: "Resource schema for AWS::DataSync::LocationFSxLustre.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationFSxLustre").WithTerraformTypeName("awscc_datasync_location_fsx_lustre")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"fsx_filesystem_arn":  "FsxFilesystemArn",
		"key":                 "Key",
		"location_arn":        "LocationArn",
		"location_uri":        "LocationUri",
		"security_group_arns": "SecurityGroupArns",
		"subdirectory":        "Subdirectory",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Subdirectory",
		"/properties/FsxFilesystemArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
