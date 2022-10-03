// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_lakeformation_tag_association", tagAssociationResource)
}

// tagAssociationResource returns the Terraform awscc_lakeformation_tag_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::LakeFormation::TagAssociation resource.
func tagAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"lf_tags": {
			// Property: LFTags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of Lake Formation Tags to associate with the Lake Formation Resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CatalogId": {
			//         "maxLength": 12,
			//         "minLength": 12,
			//         "type": "string"
			//       },
			//       "TagKey": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "TagValues": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxLength": 256,
			//           "minLength": 0,
			//           "type": "string"
			//         },
			//         "maxItems": 50,
			//         "minItems": 1,
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "CatalogId",
			//       "TagKey",
			//       "TagValues"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "List of Lake Formation Tags to associate with the Lake Formation Resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog_id": {
						// Property: CatalogId
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(12, 12),
						},
					},
					"tag_key": {
						// Property: TagKey
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"tag_values": {
						// Property: TagValues
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 50),
							validate.ArrayForEach(validate.StringLenBetween(0, 256)),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.RequiresReplace(),
			},
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Resource to tag with the Lake Formation Tags",
			//   "properties": {
			//     "Catalog": {
			//       "additionalProperties": false,
			//       "type": "object"
			//     },
			//     "Database": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "Name"
			//       ],
			//       "type": "object"
			//     },
			//     "Table": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "DatabaseName": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "TableWildcard": {
			//           "additionalProperties": false,
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "DatabaseName"
			//       ],
			//       "type": "object"
			//     },
			//     "TableWithColumns": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "ColumnNames": {
			//           "insertionOrder": false,
			//           "items": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "DatabaseName": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "DatabaseName",
			//         "Name",
			//         "ColumnNames"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Resource to tag with the Lake Formation Tags",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog": {
						// Property: Catalog
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"database": {
						// Property: Database
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(12, 12),
									},
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
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
					"table": {
						// Property: Table
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(12, 12),
									},
								},
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"table_wildcard": {
									// Property: TableWildcard
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
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
					"table_with_columns": {
						// Property: TableWithColumns
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(12, 12),
									},
								},
								"column_names": {
									// Property: ColumnNames
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayForEach(validate.StringLenBetween(1, 255)),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
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
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resource_identifier": {
			// Property: ResourceIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
			//   "type": "string"
			// }
			Description: "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags_identifier": {
			// Property: TagsIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
			//   "type": "string"
			// }
			Description: "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
			Type:        types.StringType,
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
		Description: "A resource schema representing a Lake Formation Tag Association. While tag associations are not explicit Lake Formation resources, this CloudFormation resource can be used to associate tags with Lake Formation entities.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::TagAssociation").WithTerraformTypeName("awscc_lakeformation_tag_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":             "Catalog",
		"catalog_id":          "CatalogId",
		"column_names":        "ColumnNames",
		"database":            "Database",
		"database_name":       "DatabaseName",
		"lf_tags":             "LFTags",
		"name":                "Name",
		"resource":            "Resource",
		"resource_identifier": "ResourceIdentifier",
		"table":               "Table",
		"table_wildcard":      "TableWildcard",
		"table_with_columns":  "TableWithColumns",
		"tag_key":             "TagKey",
		"tag_values":          "TagValues",
		"tags_identifier":     "TagsIdentifier",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
