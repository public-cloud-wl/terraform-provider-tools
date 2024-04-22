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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure SlugifyProvider satisfies various provider interfaces.
var _ provider.Provider = &SlugifyProvider{}
var _ provider.ProviderWithFunctions = &SlugifyProvider{}

// SlugifyProvider defines the provider implementation.
type SlugifyProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SlugifyProviderModel describes the provider data model.
type SlugifyProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *SlugifyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scaffolding"
	resp.Version = p.version
}

func (p *SlugifyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *SlugifyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data SlugifyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *SlugifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *SlugifyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *SlugifyProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewSlugFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SlugifyProvider{
			version: version,
		}
	}
}
