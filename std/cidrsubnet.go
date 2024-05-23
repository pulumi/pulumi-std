// Copyright 2022, Pulumi Corporation.
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
	"math/big"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Cidrsubnet struct{}
type CidrsubnetArgs struct {
	// The IP address range in CIDR notation.
	Input   string `pulumi:"input"`
	Newbits int    `pulumi:"newbits"`
	Netnum  int    `pulumi:"netnum"`
}

type CidrsubnetResult struct {
	Result string `pulumi:"result"`
}

func (r *Cidrsubnet) Annotate(a infer.Annotator) {
	a.Describe(r, `Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8) returns 10.2.0.0/16;
cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns 2607:f298:6051:516c:200::/72.`)
}

func cidrsubnet(ipaddress string, newbits int, netnum int) (string, error) {
	_, network, err := net.ParseCIDR(ipaddress)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR expression: %s", err.Error())
	}

	newNetwork, err := cidr.SubnetBig(network, newbits, big.NewInt(int64(netnum)))
	if err != nil {
		return "", err
	}

	return newNetwork.String(), nil
}

func (*Cidrsubnet) Call(_ p.Context, args CidrsubnetArgs) (CidrsubnetResult, error) {
	result, err := cidrsubnet(args.Input, args.Newbits, args.Netnum)
	if err != nil {
		return CidrsubnetResult{}, err
	}

	return CidrsubnetResult{result}, nil
}
