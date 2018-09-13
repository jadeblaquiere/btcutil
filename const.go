// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil

const (
	// MystikoPerBitcent is the number of satoshi in one bitcoin cent.
	MystikoPerBitcent = 1e6

	// MystikoPerBitcoin is the number of satoshi in one bitcoin (1 CTT).
	MystikoPerBitcoin = 1e8

	// MaxMystiko is the maximum transaction amount allowed in satoshi.
	MaxMystiko = 21e6 * MystikoPerBitcoin
)
