package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

const roleHashLength = 8

var _ function.Function = &RoleHashFunction{}

type RoleHashFunction struct{}

func NewRoleHashFunction() function.Function {
	return &RoleHashFunction{}
}

func (f *RoleHashFunction) Metadata(
	_ context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	resp.Name = "role_hash"
}

func (f *RoleHashFunction) Definition(
	_ context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	resp.Definition = function.Definition{
		Summary:     "Compute the deterministic hash for a role",
		Description: "Returns the first 8 hexadecimal characters of the SHA-256 hash of the given role name. Used to build group aliases.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "role",
				Description: "The role name (e.g. 'my-role-name').",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *RoleHashFunction) Run(
	ctx context.Context,
	req function.RunRequest,
	resp *function.RunResponse,
) {
	var role string

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		req.Arguments.Get(ctx, &role),
	)
	if resp.Error != nil {
		return
	}

	hash := hashHexPrefix(role, roleHashLength)

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, hash),
	)
}
