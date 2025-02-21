// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/gosimple/slug"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = GroupNameFunction{}
)

func NewGroupNameFunction() function.Function {
	return GroupNameFunction{}
}

type GroupNameFunction struct{}

func (r GroupNameFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "group_name"
}

func (r GroupNameFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Given an application service and an IDM role, return the associated group name.",
		MarkdownDescription: "Given an application service and an IDM role, return the associated group name.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "as",
				MarkdownDescription: "The application service.",
			},
			function.StringParameter{
				Name:                "role",
				MarkdownDescription: "The IDM role. '*' refer to any people with a role in the application service.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r GroupNameFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var as string
	var role string
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &as, &role))

	if role != "*" {
		output = slug.Make("dl-idm-" + strings.ToLower(as) + "-" + strings.ToLower(role))
	} else {
		output = slug.Make("dl-idm-" + strings.ToLower(as))
	}
	output = strings.ReplaceAll(output, "_", "-")
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}
