// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_resolver_rule", resolverRuleDataSource)
}

// resolverRuleDataSource returns the Terraform awscc_route53resolver_resolver_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::ResolverRule resource.
func resolverRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the resolver rule.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the resolver rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the Resolver rule",
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the Resolver rule",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResolverEndpointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint that the rule is associated with.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resolver_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint that the rule is associated with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResolverRuleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint that the rule is associated with.",
		//	  "type": "string"
		//	}
		"resolver_rule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint that the rule is associated with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
		//	  "enum": [
		//	    "FORWARD",
		//	    "SYSTEM",
		//	    "RECURSIVE",
		//	    "DELEGATE"
		//	  ],
		//	  "type": "string"
		//	}
		"rule_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetIps
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Ip": {
		//	        "description": "One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. ",
		//	        "type": "string"
		//	      },
		//	      "Ipv6": {
		//	        "description": "One IPv6 address that you want to forward DNS queries to. You can specify only IPv6 addresses. ",
		//	        "type": "string"
		//	      },
		//	      "Port": {
		//	        "description": "The port at Ip that you want to forward DNS queries to. ",
		//	        "maxLength": 65535,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "Protocol": {
		//	        "description": "The protocol that you want to use to forward DNS queries. ",
		//	        "enum": [
		//	          "Do53",
		//	          "DoH"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ServerNameIndication": {
		//	        "description": "The SNI of the target name servers for DoH/DoH-FIPS outbound endpoints",
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"target_ips": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Ip
					"ip": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Ipv6
					"ipv_6": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "One IPv6 address that you want to forward DNS queries to. You can specify only IPv6 addresses. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Port
					"port": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The port at Ip that you want to forward DNS queries to. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Protocol
					"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The protocol that you want to use to forward DNS queries. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ServerNameIndication
					"server_name_indication": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The SNI of the target name servers for DoH/DoH-FIPS outbound endpoints",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::ResolverRule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverRule").WithTerraformTypeName("awscc_route53resolver_resolver_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"domain_name":            "DomainName",
		"ip":                     "Ip",
		"ipv_6":                  "Ipv6",
		"key":                    "Key",
		"name":                   "Name",
		"port":                   "Port",
		"protocol":               "Protocol",
		"resolver_endpoint_id":   "ResolverEndpointId",
		"resolver_rule_id":       "ResolverRuleId",
		"rule_type":              "RuleType",
		"server_name_indication": "ServerNameIndication",
		"tags":                   "Tags",
		"target_ips":             "TargetIps",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
