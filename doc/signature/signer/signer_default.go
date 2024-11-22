/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signer

import (
	"time"

	afgotime "github.com/dellekappa/did-go/doc/util/time"
)

func wrapTime(t time.Time) *afgotime.TimeWrapper {
	return &afgotime.TimeWrapper{Time: t}
}
