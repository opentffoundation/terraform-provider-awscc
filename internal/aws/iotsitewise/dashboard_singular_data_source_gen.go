// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_dashboard", dashboardDataSource)
}

// dashboardDataSource returns the Terraform awscc_iotsitewise_dashboard data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::Dashboard resource.
func dashboardDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DashboardArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dashboard definition specified in a JSON literal.",
		//	  "type": "string"
		//	}
		"dashboard_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The dashboard definition specified in a JSON literal.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A friendly name for the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A friendly name for the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the project in which to create the dashboard.",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the project in which to create the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the dashboard.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "A list of key-value pairs that contain metadata for the dashboard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Dashboard",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Dashboard").WithTerraformTypeName("awscc_iotsitewise_dashboard")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dashboard_arn":         "DashboardArn",
		"dashboard_definition":  "DashboardDefinition",
		"dashboard_description": "DashboardDescription",
		"dashboard_id":          "DashboardId",
		"dashboard_name":        "DashboardName",
		"key":                   "Key",
		"project_id":            "ProjectId",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
