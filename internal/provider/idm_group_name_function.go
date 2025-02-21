// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"regexp"
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
				Name:                "member",
				MarkdownDescription: "The member to calculate group from. Only member prefix with 'idmrole:' are consider. '*' refer to any people with a role in the application service.",
			},
			function.StringParameter{
				Name:                "as",
				MarkdownDescription: "The application service.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r IdmGroupName) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var member string
	var as string
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &member, &as))
	// Check if member is prefixed by 'idmrole:' and just return as is if not
	if !strings.HasPrefix(member, "idmrole:") {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, member))
		return
	}
	// Check if member is in format idmrole:[IDM_ROLE]@[DOMAIN]
	regexMember := "idmrole:[a-zA-Z0-9-_ ]+@[a-zA-Z0-9-_ \\.]+"
	m, err := regexp.MatchString(regexMember, member)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(function.NewFuncError("Unexpected error with regex \\" + regexMember + "\\ : " + err.Error()))
		return
	} else if !m {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError("Member prefixed by 'idmrole:' should be in format 'idmrole:[IDM_ROLE]@[DOMAIN]'\nMember: '" + member + "' doesn't match regex: " + regexMember + " "),
		)
		return
	}
	role_domain := strings.Split(strings.TrimPrefix(member, "idmrole:"), "@")
	role, domain := role_domain[0], role_domain[1]
	if role != "*" {
		output = "group:" + slug.Make("dl-idm-"+as+"-"+role) + "@" + domain
	} else {
		output = "group:" + slug.Make("dl-idm-"+as) + "@" + domain
	}
	output = strings.ReplaceAll(output, "_", "-")
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}
