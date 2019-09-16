// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package util

import (
	"encoding/binary"
	"time"
)


const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 

func Uint64ToInt64(u uint64) int64 {
	if u > MaxInt {
		zap.L().Panic("Height can't be converted to int64")
	}
	return int64(u)
}

func TimeToBytes(t time.Time) ([]byte, error) {
	return t.MarshalBinary()
}

func BytesToTime(b []byte) (time.Time, error) {
	var t time.Time
	if err := t.UnmarshalBinary(b); err != nil {
		return t, err
	}
	return t, nil
}

func Uint64ToBytes(u uint64) []byte {
	retval := make([]byte, 8)
	binary.LittleEndian.PutUint64(retval, u)

	return retval
}

func BytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func CopyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)

	return c
}

// IsAllZeros return true if all bytes are zero
func IsAllZeros(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
