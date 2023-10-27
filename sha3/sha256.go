// Copyright © 2019 Weald Technology Trading
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

package sha3

import (
	"golang.org/x/crypto/sha3"
)

const _256hashlength = 32

// SHA3 is the Sha3 hashing method.
type SHA256 struct{}

// New creates a new Sha3 hashing method.
func New256() *SHA256 {
	return &SHA256{}
}

// HashLength returns the length of hashes generated by Hash() in bytes.
func (h *SHA256) HashLength() int {
	return _256hashlength
}

func (h *SHA256) HashName() string {
	return "sha256"
}

// Hash generates a SHA3 hash from input byte arrays.
func (h *SHA256) Hash(data ...[]byte) []byte {
	var hash [_256hashlength]byte
	if len(data) == 1 {
		hash = sha3.Sum256(data[0])
	} else {
		concatDataLen := 0
		for _, d := range data {
			concatDataLen += len(d)
		}
		concatData := make([]byte, concatDataLen)
		curOffset := 0
		for _, d := range data {
			copy(concatData[curOffset:], d)
			curOffset += len(d)
		}
		hash = sha3.Sum256(concatData)
	}

	return hash[:]
}
