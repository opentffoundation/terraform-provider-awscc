// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iottwinmaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iottwinmaker_workspace", workspaceDataSource)
}

// workspaceDataSource returns the Terraform awscc_iottwinmaker_workspace data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTTwinMaker::Workspace resource.
func workspaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the workspace.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iottwinmaker:[a-z0-9-]+:[0-9]{12}:[\\/a-zA-Z0-9_\\-\\.:]+",
			//   "type": "string"
			// }
			Description: "The ARN of the workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_date_time": {
			// Property: CreationDateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the workspace was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date and time when the workspace was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the workspace.",
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role": {
			// Property: Role
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the execution role associated with the workspace.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iam::[0-9]{12}:role/.*",
			//   "type": "string"
			// }
			Description: "The ARN of the execution role associated with the workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3_location": {
			// Property: S3Location
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the S3 bucket where resources associated with the workspace are stored.",
			//   "type": "string"
			// }
			Description: "The ARN of the S3 bucket where resources associated with the workspace are stored.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A map of key-value pairs to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A map of key-value pairs to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"update_date_time": {
			// Property: UpdateDateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time of the current update.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date and time of the current update.",
			Type:        types.StringType,
			Computed:    true,
		},
		"workspace_id": {
			// Property: WorkspaceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the workspace.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
			//   "type": "string"
			// }
			Description: "The ID of the workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTTwinMaker::Workspace",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTTwinMaker::Workspace").WithTerraformTypeName("awscc_iottwinmaker_workspace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"creation_date_time": "CreationDateTime",
		"description":        "Description",
		"role":               "Role",
		"s3_location":        "S3Location",
		"tags":               "Tags",
		"update_date_time":   "UpdateDateTime",
		"workspace_id":       "WorkspaceId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
