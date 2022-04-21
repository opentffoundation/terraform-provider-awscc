// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iottwinmaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iottwinmaker_scene", sceneDataSourceType)
}

// sceneDataSourceType returns the Terraform awscc_iottwinmaker_scene data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTTwinMaker::Scene resource type.
func sceneDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the scene.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iottwinmaker:[a-z0-9-]+:[0-9]{12}:[\\/a-zA-Z0-9_\\-\\.:]+",
			//   "type": "string"
			// }
			Description: "The ARN of the scene.",
			Type:        types.StringType,
			Computed:    true,
		},
		"capabilities": {
			// Property: Capabilities
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of capabilities that the scene uses to render.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 256,
			//     "minLength": 0,
			//     "pattern": ".*",
			//     "type": "string"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of capabilities that the scene uses to render.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"content_location": {
			// Property: ContentLocation
			// CloudFormation resource type schema:
			// {
			//   "description": "The relative path that specifies the location of the content definition file.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "[sS]3://[A-Za-z0-9._/-]+",
			//   "type": "string"
			// }
			Description: "The relative path that specifies the location of the content definition file.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_date_time": {
			// Property: CreationDateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the scene was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date and time when the scene was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the scene.",
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the scene.",
			Type:        types.StringType,
			Computed:    true,
		},
		"scene_id": {
			// Property: SceneId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the scene.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
			//   "type": "string"
			// }
			Description: "The ID of the scene.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
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
			//   "description": "The ID of the scene.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
			//   "type": "string"
			// }
			Description: "The ID of the scene.",
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
		Description: "Data Source schema for AWS::IoTTwinMaker::Scene",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTTwinMaker::Scene").WithTerraformTypeName("awscc_iottwinmaker_scene")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"capabilities":       "Capabilities",
		"content_location":   "ContentLocation",
		"creation_date_time": "CreationDateTime",
		"description":        "Description",
		"scene_id":           "SceneId",
		"tags":               "Tags",
		"update_date_time":   "UpdateDateTime",
		"workspace_id":       "WorkspaceId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
