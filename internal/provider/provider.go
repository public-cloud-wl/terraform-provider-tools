// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ToolsProvider satisfies various provider interfaces.
var _ provider.Provider = &ToolsProvider{}
var _ provider.ProviderWithFunctions = &ToolsProvider{}

// ToolsProvider defines the provider implementation.
type ToolsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *ToolsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tools"
	resp.Version = p.version
}

func (p *ToolsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *ToolsProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	// This provider has no configuration. It only exposes pure functions.
}

func (p *ToolsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *ToolsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *ToolsProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewSlugFunction,
		NewGetClosestInferiorGCPCloudNatDynamicPortFunction,
		NewNetworkIsInSubnetFunction,
		NewApplicationServiceHashFunction,
		NewRoleHashFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ToolsProvider{
			version: version,
		}
	}
}
