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
	_ function.Function = IdmGroupName{}
)

func NewIdmGroupName() function.Function {
	return IdmGroupName{}
}

type IdmGroupName struct{}

func (r IdmGroupName) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "idm_group_name"
}

func (r IdmGroupName) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Give an application service and a member with format idmrole:[jobrole]@[domain], return the group name associated. If member is not prefixed by idmrole, just return it.",
		MarkdownDescription: "Give an application service and a member with format idmrole:[jobrole]@[domain], return the group name associated. If member is not prefixed by idmrole, just return it.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "as",
				MarkdownDescription: "The application service.",
			},
			function.StringParameter{
				Name:                "member",
				MarkdownDescription: "The member to calculate group from. Only member prefix with 'idmrole:' are consider. '*' refer to any people with a role in the application service.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r IdmGroupName) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var as string
	var member string
	var domain string
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &as, &member))
	if !strings.HasPrefix(member, "idmrole:") {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, member))
		return
	}
	role := strings.Split(strings.TrimPrefix(member, "idmrole:"), "@")[0]
	domain = strings.Split(strings.TrimPrefix(member, "idmrole:"), "@")[1]
	if role != "*" {
		output = slug.Make("dl-idm-"+strings.ToLower(as)+"-"+strings.ToLower(role)) + "@" + domain
	} else {
		output = slug.Make("dl-idm-"+strings.ToLower(as)) + "@" + domain
	}
	output = strings.ReplaceAll(output, "_", "-")
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}
