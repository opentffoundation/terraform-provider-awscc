// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package b2bi

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_b2bi_partnership", partnershipResource)
}

// partnershipResource returns the Terraform awscc_b2bi_partnership resource.
// This Terraform resource corresponds to the CloudFormation AWS::B2BI::Partnership resource.
func partnershipResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Capabilities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 1,
		//	    "pattern": "^[a-zA-Z0-9_-]+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"capabilities": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: CapabilityOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "OutboundEdi": {
		//	      "properties": {
		//	        "X12": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Common": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Delimiters": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "ComponentSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    },
		//	                    "DataElementSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    },
		//	                    "SegmentTerminator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "FunctionalGroupHeaders": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "ApplicationReceiverCode": {
		//	                      "maxLength": 15,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ApplicationSenderCode": {
		//	                      "maxLength": 15,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ResponsibleAgencyCode": {
		//	                      "maxLength": 2,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "InterchangeControlHeaders": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "AcknowledgmentRequestedCode": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ReceiverId": {
		//	                      "maxLength": 15,
		//	                      "minLength": 15,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ReceiverIdQualifier": {
		//	                      "maxLength": 2,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "RepetitionSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "type": "string"
		//	                    },
		//	                    "SenderId": {
		//	                      "maxLength": 15,
		//	                      "minLength": 15,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "SenderIdQualifier": {
		//	                      "maxLength": 2,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "UsageIndicatorCode": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "ValidateEdi": {
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"capability_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OutboundEdi
				"outbound_edi": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: X12
						"x12": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Common
								"common": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Delimiters
										"delimiters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: ComponentSeparator
												"component_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
														stringvalidator.RegexMatches(regexp.MustCompile("^[!&'()*+,\\-./:;?=%@\\[\\]_{}|<>~^`\"]$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: DataElementSeparator
												"data_element_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
														stringvalidator.RegexMatches(regexp.MustCompile("^[!&'()*+,\\-./:;?=%@\\[\\]_{}|<>~^`\"]$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: SegmentTerminator
												"segment_terminator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
														stringvalidator.RegexMatches(regexp.MustCompile("^[!&'()*+,\\-./:;?=%@\\[\\]_{}|<>~^`\"]$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: FunctionalGroupHeaders
										"functional_group_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: ApplicationReceiverCode
												"application_receiver_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(2, 15),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: ApplicationSenderCode
												"application_sender_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(2, 15),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: ResponsibleAgencyCode
												"responsible_agency_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 2),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: InterchangeControlHeaders
										"interchange_control_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: AcknowledgmentRequestedCode
												"acknowledgment_requested_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: ReceiverId
												"receiver_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(15, 15),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: ReceiverIdQualifier
												"receiver_id_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(2, 2),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: RepetitionSeparator
												"repetition_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: SenderId
												"sender_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(15, 15),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: SenderIdQualifier
												"sender_id_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(2, 2),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: UsageIndicatorCode
												"usage_indicator_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													Validators: []validator.String{ /*START VALIDATORS*/
														stringvalidator.LengthBetween(1, 1),
														stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]*$"), ""),
													}, /*END VALIDATORS*/
													PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
														stringplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: ValidateEdi
										"validate_edi": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
												boolplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Email
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 5,
		//	  "pattern": "^[\\w\\.\\-]+@[\\w\\.\\-]+$",
		//	  "type": "string"
		//	}
		"email": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 254),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w\\.\\-]+@[\\w\\.\\-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 254),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PartnershipArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"partnership_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PartnershipId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"partnership_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Phone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 22,
		//	  "minLength": 7,
		//	  "pattern": "^\\+?([0-9 \\t\\-()\\/]{7,})(?:\\s*(?:#|x\\.?|ext\\.?|extension) \\t*(\\d+))?$",
		//	  "type": "string"
		//	}
		"phone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(7, 22),
				stringvalidator.RegexMatches(regexp.MustCompile("^\\+?([0-9 \\t\\-()\\/]{7,})(?:\\s*(?:#|x\\.?|ext\\.?|extension) \\t*(\\d+))?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TradingPartnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"trading_partner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::B2BI::Partnership Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::B2BI::Partnership").WithTerraformTypeName("awscc_b2bi_partnership")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acknowledgment_requested_code": "AcknowledgmentRequestedCode",
		"application_receiver_code":     "ApplicationReceiverCode",
		"application_sender_code":       "ApplicationSenderCode",
		"capabilities":                  "Capabilities",
		"capability_options":            "CapabilityOptions",
		"common":                        "Common",
		"component_separator":           "ComponentSeparator",
		"created_at":                    "CreatedAt",
		"data_element_separator":        "DataElementSeparator",
		"delimiters":                    "Delimiters",
		"email":                         "Email",
		"functional_group_headers":      "FunctionalGroupHeaders",
		"interchange_control_headers":   "InterchangeControlHeaders",
		"key":                           "Key",
		"modified_at":                   "ModifiedAt",
		"name":                          "Name",
		"outbound_edi":                  "OutboundEdi",
		"partnership_arn":               "PartnershipArn",
		"partnership_id":                "PartnershipId",
		"phone":                         "Phone",
		"profile_id":                    "ProfileId",
		"receiver_id":                   "ReceiverId",
		"receiver_id_qualifier":         "ReceiverIdQualifier",
		"repetition_separator":          "RepetitionSeparator",
		"responsible_agency_code":       "ResponsibleAgencyCode",
		"segment_terminator":            "SegmentTerminator",
		"sender_id":                     "SenderId",
		"sender_id_qualifier":           "SenderIdQualifier",
		"tags":                          "Tags",
		"trading_partner_id":            "TradingPartnerId",
		"usage_indicator_code":          "UsageIndicatorCode",
		"validate_edi":                  "ValidateEdi",
		"value":                         "Value",
		"x12":                           "X12",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
