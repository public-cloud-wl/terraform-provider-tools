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
	_ function.Function = SlugFunction{}
)

func NewSlugFunction() function.Function {
	return SlugFunction{}
}

type SlugFunction struct{}

func (r SlugFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "slug"
}

func (r SlugFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Slug function",
		MarkdownDescription: "Return URL-friendly slugify string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "Input string to slugify.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r SlugFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	output := slug.Make(data)
	output = strings.ReplaceAll(output, "_", "-")
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}
