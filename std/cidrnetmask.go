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
	"context"
	"fmt"
	"net"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Cidrnetmask struct{}
type CidrnetmaskArgs struct {
	Input string `pulumi:"input"`
}

type CidrnetmaskResult struct {
	Result string `pulumi:"result"`
}

func (r *Cidrnetmask) Annotate(a infer.Annotator) {
	a.Describe(r, `Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
that some systems expect for IPv4 interfaces.
For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.`)
}

func cidrnetmask(ipaddress string) (string, error) {
	_, network, err := net.ParseCIDR(ipaddress)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR expression: %s", err.Error())
	}

	if network.IP.To4() == nil {
		return "", fmt.Errorf("IPv6 addresses cannot have a netmask: %s", ipaddress)
	}

	return net.IP(network.Mask).String(), nil
}

func (*Cidrnetmask) Invoke(_ context.Context, input infer.FunctionRequest[CidrnetmaskArgs]) (infer.FunctionResponse[CidrnetmaskResult], error) {
	result, err := cidrnetmask(input.Input.Input)
	if err != nil {
		return infer.FunctionResponse[CidrnetmaskResult]{Output: CidrnetmaskResult{}}, err
	}

	return infer.FunctionResponse[CidrnetmaskResult]{Output: CidrnetmaskResult{result}}, nil
}
