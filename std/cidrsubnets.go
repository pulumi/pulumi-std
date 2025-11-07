// Copyright 2025, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package std

import (
	"context"
	"fmt"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	Cidrsubnets     struct{}
	CidrsubnetsArgs struct {
		Input   string `pulumi:"input"`
		Newbits []int  `pulumi:"newbits"`
	}
)

type CidrsubnetsResult struct {
	Result []string `pulumi:"result"`
}

func (r *Cidrsubnets) Annotate(a infer.Annotator) {
	a.Describe(r,
		"Takes an IP address prefix in CIDR notation (like 10.0.0.0/8) and creates a series of "+
			"consecutive IP address ranges within that CIDR prefix. "+
			"See https://opentofu.org/docs/language/functions/cidrsubnets/ for additional information",
	)
}

func cidrsubnets(ipaddress string, newbitsList ...int) ([]string, error) {
	_, network, err := net.ParseCIDR(ipaddress)
	if err != nil {
		return []string{}, fmt.Errorf("invalid CIDR expression: %s", err.Error())
	}

	if len(newbitsList) == 0 {
		return []string{}, nil
	}

	subnets := make([]string, len(newbitsList))
	for i, newbits := range newbitsList {
		subnet, exceeds := cidr.NextSubnet(network, newbits)
		if exceeds {
			prevSubnet := network.String()
			if i > 0 {
				prevSubnet = subnets[i-1]
			}
			return []string{}, fmt.Errorf(
				"not enough remaining address space for a subnet of %d bits after address %s",
				newbits, prevSubnet,
			)
		}
		subnets[i] = subnet.String()
	}

	return subnets, nil
}

func (*Cidrsubnets) Invoke(_ context.Context, input infer.FunctionRequest[CidrsubnetsArgs]) (infer.FunctionResponse[CidrsubnetsResult], error) {
	_, network, err := net.ParseCIDR(input.Input.Input)
	if err != nil {
		return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{}}, fmt.Errorf("invalid CIDR expression: %w", err)
	}

	startPrefixLen, _ := network.Mask.Size()

	if len(input.Input.Newbits) == 0 {
		return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{[]string{}}}, nil
	}

	results := make([]string, len(input.Input.Newbits))

	firstLength := input.Input.Newbits[0] + startPrefixLen

	current, _ := cidr.PreviousSubnet(network, firstLength)
	for i, length := range input.Input.Newbits {
		if length < 1 {
			return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{}}, fmt.Errorf(
				"argument %d (%d) must extend prefix by at least one bit",
				i+1, length,
			)
		}

		if length > 32 {
			return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{}}, fmt.Errorf(
				"argument %d (%d) may not extend prefix by more than 32 bits",
				i+1, length,
			)
		}

		length += startPrefixLen
		if length > (len(network.IP) * 8) {
			protocol := "IP"
			switch len(network.IP) * 8 {
			case 32:
				protocol = "IPv4"
			case 128:
				protocol = "IPv6"
			}

			return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{}}, fmt.Errorf(
				"argument %d would extend prefix to %d bits, which is too long for an %s address",
				i+1, length, protocol,
			)
		}

		next, overflowed := cidr.NextSubnet(current, length)
		if overflowed || !network.Contains(next.IP) {
			return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{}}, fmt.Errorf(
				"not enough remaining address space for a subnet with a prefix of %d bits after %s",
				length, current.String(),
			)
		}

		current = next
		results[i] = current.String()
	}

	return infer.FunctionResponse[CidrsubnetsResult]{Output: CidrsubnetsResult{results}}, nil
}
