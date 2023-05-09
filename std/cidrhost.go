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

type Cidrhost struct{}
type CidrhostArgs struct {
	Input string `pulumi:"input"`
	Host  int    `pulumi:"host"`
}

type CidrhostResult struct {
	Result string `pulumi:"result"`
}

func (r *Cidrhost) Annotate(a infer.Annotator) {
	a.Describe(r, `Takes an IP address range in CIDR notation as input
and creates an IP address with the given host number.
If given host number is negative, the count starts from the end of the range.
For example, cidrhost("10.0.0.0/8", 2) returns 10.0.0.2 and cidrhost("10.0.0.0/8", -2) returns 10.255.255.254.`)
}

func cidrhost(ipaddress string, hostnum int) (string, error) {
	_, network, err := net.ParseCIDR(ipaddress)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR expression: %s", err)
	}

	ip, err := cidr.HostBig(network, big.NewInt(int64(hostnum)))
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}

func (*Cidrhost) Call(_ p.Context, args CidrhostArgs) (CidrhostResult, error) {
	result, err := cidrhost(args.Input, args.Host)
	if err != nil {
		return CidrhostResult{}, err
	}

	return CidrhostResult{result}, nil
}
