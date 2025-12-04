// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = NetworkIsInSubnetFunction{}
)

func NewNetworkIsInSubnetFunction() function.Function {
	return NetworkIsInSubnetFunction{}
}

type NetworkIsInSubnetFunction struct{}

func (r NetworkIsInSubnetFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "network_is_in_subnet"
}

func (r NetworkIsInSubnetFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check if a CIDR is within another CIDR",
		MarkdownDescription: "Returns true if the first CIDR block is fully contained within the second CIDR block.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "child_cidr",
				MarkdownDescription: "The CIDR block to check (potential subnet).",
			},
			function.StringParameter{
				Name:                "parent_cidr",
				MarkdownDescription: "The CIDR block that should contain the child.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r NetworkIsInSubnetFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var childCidr, parentCidr string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &childCidr, &parentCidr))

	if resp.Error != nil {
		return
	}

	// Parse the child CIDR
	_, childNetwork, err := net.ParseCIDR(childCidr)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("invalid child CIDR %q: %s", childCidr, err))
		return
	}

	// Parse the parent CIDR
	_, parentNetwork, err := net.ParseCIDR(parentCidr)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("invalid parent CIDR %q: %s", parentCidr, err))
		return
	}

	// Check if child CIDR is within parent CIDR
	result := networkContains(parentNetwork, childNetwork)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// networkContains checks if the child network is fully contained within the parent network.
func networkContains(parent, child *net.IPNet) bool {
	// The parent network must contain the first IP of the child network
	if !parent.Contains(child.IP) {
		return false
	}

	// Calculate the last IP of the child network
	lastIP := lastIPInNetwork(child)

	// The parent network must also contain the last IP of the child network
	return parent.Contains(lastIP)
}

// lastIPInNetwork calculates the last IP address in a given network.
func lastIPInNetwork(network *net.IPNet) net.IP {
	ip := network.IP
	mask := network.Mask

	// Create a copy to avoid modifying the original
	lastIP := make(net.IP, len(ip))
	copy(lastIP, ip)

	// Apply the inverse of the mask to get the broadcast/last address
	for i := range lastIP {
		lastIP[i] = ip[i] | ^mask[i]
	}

	return lastIP
}
