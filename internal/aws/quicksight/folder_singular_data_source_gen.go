// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_quicksight_folder", folderDataSource)
}

// folderDataSource returns the Terraform awscc_quicksight_folder data source.
// This Terraform data source corresponds to the CloudFormation AWS::QuickSight::Folder resource.
func folderDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Resource Name (ARN) for the folder.\u003c/p\u003e",
		//	  "pattern": "^arn:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Resource Name (ARN) for the folder.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe time that the folder was created.\u003c/p\u003e",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "<p>The time that the folder was created.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FolderId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w\\-]+$",
		//	  "type": "string"
		//	}
		"folder_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FolderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "SHARED",
		//	    "RESTRICTED"
		//	  ],
		//	  "type": "string"
		//	}
		"folder_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe time that the folder was last updated.\u003c/p\u003e",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "<p>The time that the folder was last updated.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ParentFolderArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"parent_folder_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Permissions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003ePermission for the resource.\u003c/p\u003e",
		//	    "properties": {
		//	      "Actions": {
		//	        "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "maxItems": 20,
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "Principal": {
		//	        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n         \u003cul\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^arn:.*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Actions",
		//	      "Principal"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 64,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"permissions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Actions
					"actions": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "<p>The IAM action to grant or revoke permissions on.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Principal
					"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:</p>\n         <ul>\n            <li>\n               <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>\n            </li>\n            <li>\n               <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>\n            </li>\n            <li>\n               <p>The ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.\n                    (This is less common.) </p>\n            </li>\n         </ul>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SharingModel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ACCOUNT",
		//	    "NAMESPACE"
		//	  ],
		//	  "type": "string"
		//	}
		"sharing_model": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "\u003cp\u003eTag key.\u003c/p\u003e",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "\u003cp\u003eTag value.\u003c/p\u003e",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag key.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag value.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::QuickSight::Folder",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::Folder").WithTerraformTypeName("awscc_quicksight_folder")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":           "Actions",
		"arn":               "Arn",
		"aws_account_id":    "AwsAccountId",
		"created_time":      "CreatedTime",
		"folder_id":         "FolderId",
		"folder_type":       "FolderType",
		"key":               "Key",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"parent_folder_arn": "ParentFolderArn",
		"permissions":       "Permissions",
		"principal":         "Principal",
		"sharing_model":     "SharingModel",
		"tags":              "Tags",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
