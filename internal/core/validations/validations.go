// MIT License
//
// Copyright (c) 2021 TFG Co
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package validations

import (
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

//IsMaxSurgeValid check if MaxSurge is valid. A MaxSurge valid is a number greater than zero or a number greater than zero with suffix '%'
func IsMaxSurgeValid(maxSurge string) bool {
	if maxSurge == "" {
		return false
	}
	minMaxSurge := 1
	relativeSymbol := "%"

	isRelative := strings.HasSuffix(maxSurge, relativeSymbol)
	maxSurgeInt, err := strconv.Atoi(strings.TrimSuffix(maxSurge, relativeSymbol))
	if err != nil {
		return false
	}

	if !isRelative {
		if minMaxSurge > maxSurgeInt {
			return false
		}
	}
	return true
}

// IsImagePullPolicySupported check if received policy is supported by maestro
func IsImagePullPolicySupported(policy string) bool {
	policies := []string{"Always", "Never", "IfNotPresent"}
	for _, item := range policies {
		if item == policy {
			return true
		}
	}
	return false
}

// IsProtocolSupported check if received protocol is supported by kubernetes
func IsProtocolSupported(protocol string) bool {
	protocols := []string{"tcp", "udp", "sctp"}
	for _, item := range protocols {
		if item == protocol {
			return true
		}
	}
	return false
}

// IsVersionValid check if version value is according to semantic definitions
func IsVersionValid(version string) bool {
	_, err := semver.NewVersion(version)
	return err == nil
}
