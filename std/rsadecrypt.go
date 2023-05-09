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
	"crypto/rsa"
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"golang.org/x/crypto/ssh"
)

type Rsadecrypt struct{}
type RsadecryptArgs struct {
	CipherText string `pulumi:"cipherText"`
	Key        string `pulumi:"key"`
}

type RsadecryptResult struct {
	Result string `pulumi:"result"`
}

func (r *Rsadecrypt) Annotate(a infer.Annotator) {
	a.Describe(r, `Decrypts an RSA-encrypted ciphertext.
The cipher text must be base64-encoded and the key must be in PEM format.`)
}

func formatError(err error) string {
	switch e := err.(type) {
	case asn1.SyntaxError:
		return strings.ReplaceAll(e.Error(), "asn1: syntax error", "invalid ASN1 data in the given private key")
	case asn1.StructuralError:
		return strings.ReplaceAll(e.Error(), "asn1: struture error", "invalid ASN1 data in the given private key")
	default:
		return fmt.Sprintf("invalid private key: %s", e)
	}
}

func (*Rsadecrypt) Call(_ p.Context, args RsadecryptArgs) (RsadecryptResult, error) {

	cipherTextBytes, err := base64.StdEncoding.DecodeString(args.CipherText)
	if err != nil {
		return RsadecryptResult{}, errors.New("failed to decode input cipher text. It must be base64-encoded")
	}

	// Parse the private key.
	key, err := ssh.ParseRawPrivateKey([]byte(args.Key))
	if err != nil {
		errorText := formatError(err)
		return RsadecryptResult{}, errors.New(errorText)
	}

	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return RsadecryptResult{}, fmt.Errorf("invalid private key type %t", key)
	}

	// Decrypt the ciphertext.
	decrypted, err := rsa.DecryptPKCS1v15(nil, privateKey, cipherTextBytes)
	if err != nil {
		return RsadecryptResult{}, fmt.Errorf("failed to decrypt: %s", err.Error())
	}

	return RsadecryptResult{string(decrypted)}, nil
}
