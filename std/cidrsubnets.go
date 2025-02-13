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
	"fmt"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	p "github.com/pulumi/pulumi-go-provider"
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
	a.Describe(r, `Takes an IP address prefix in CIDR notation (like 10.0.0.0/8) and creates a series of 
consecutive IP address ranges within that CIDR prefix. See https://opentofu.org/docs/language/functions/cidrsubnets/ 
for additional information`)
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
			return []string{},
				fmt.Errorf("not enough remaining address space for a subnet of %d bits after address %s", newbits, prevSubnet)
		}
		subnets[i] = subnet.String()
	}

	return subnets, nil
}

func (*Cidrsubnets) Call(_ p.Context, args CidrsubnetsArgs) (CidrsubnetsResult, error) {
	result, err := cidrsubnets(args.Input, args.Newbits...)
	if err != nil {
		return CidrsubnetsResult{}, err
	}

	return CidrsubnetsResult{result}, nil
}
