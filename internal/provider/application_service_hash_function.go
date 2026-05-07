package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

const projectHashLength = 16

var _ function.Function = &ApplicationServiceHashFunction{}

type ApplicationServiceHashFunction struct{}

func NewApplicationServiceHashFunction() function.Function {
	return &ApplicationServiceHashFunction{}
}

func (f *ApplicationServiceHashFunction) Metadata(
	_ context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	resp.Name = "application_service_hash"
}

func (f *ApplicationServiceHashFunction) Definition(
	_ context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	resp.Definition = function.Definition{
		Summary:     "Compute the deterministic hash for a project code",
		Description: "Returns the first 16 hexadecimal characters of the SHA-256 hash of the given project code. Used to build group aliases.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "project_code",
				Description: "The project code (e.g. 'nmbsip-gcp-bk01-vpn-production').",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ApplicationServiceHashFunction) Run(
	ctx context.Context,
	req function.RunRequest,
	resp *function.RunResponse,
) {
	var projectCode string

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		req.Arguments.Get(ctx, &projectCode),
	)
	if resp.Error != nil {
		return
	}

	hash := hashHexPrefix(projectCode, projectHashLength)

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, hash),
	)
}

// hashHexPrefix returns the first n hex characters of the SHA-256 hash of s.
func hashHexPrefix(s string, n int) string {
	sum := sha256.Sum256([]byte(s))
	full := hex.EncodeToString(sum[:])
	if n > len(full) {
		n = len(full)
	}
	return full[:n]
}