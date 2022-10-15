// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package speciterator

import (
	"bufio"
	"fmt"
	"github.com/buger/jsonparser"
	"strings"
)

type jsonParser struct {
	data    []byte
	locator *locator
}

func newJSONParser(data []byte) *jsonParser {
	locator := newLocator(data)
	return &jsonParser{data: data, locator: locator}
}

func (p *jsonParser) Iterate(cb Callback) error {
	return p.iterateObject(p.data, &Path{}, 0, cb, 0)
}

func (p *jsonParser) iterateObject(data []byte, path *Path, indent int, cb Callback, gOffset int) error {
	return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		path.Push(string(key))
		valueLen := lengthOfValue(value, dataType)
		bOffset := offset - valueLen + gOffset
		cb(path, p.locator.GetPos(bOffset))
		switch dataType {
		case jsonparser.Object:
			_ = p.iterateObject(value, path, indent+1, cb, bOffset)
		case jsonparser.Array:
			p.iterateArray(value, path, indent+1, cb, bOffset)
		default:
		}
		path.Pop()
		return nil
	})
}

func (p *jsonParser) iterateArray(data []byte, path *Path, indent int, cb Callback, gOffset int) {
	index := 0
	_, _ = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		key := fmt.Sprintf("%d", index)
		index++
		path.Push(key)
		bOffset := gOffset + offset + len(value) - lengthOfValue(value, dataType)
		switch dataType {
		case jsonparser.Object:
			_ = p.iterateObject(value, path, indent+1, cb, bOffset)
		case jsonparser.Array:
			p.iterateArray(value, path, indent+1, cb, bOffset)
		default:
			cb(path, p.locator.GetPos(bOffset))
		}
		path.Pop()
	})
}

type locator struct {
	lines []int
}

func newLocator(data []byte) *locator {
	lines := []int{0}
	last := 0
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		txt := scanner.Text()
		length := len([]byte(txt)) + 1
		lines = append(lines, length+last)
		last = length + last
	}
	return &locator{lines: lines}
}

func (l *locator) GetPos(offset int) *Pos {
	for i, v := range l.lines {
		if offset <= v {
			return &Pos{
				Line:   i,
				Column: offset - l.lines[i-1] + 1,
				Offset: offset,
			}
		}
	}
	return &Pos{
		Line:   0,
		Column: 0,
		Offset: offset,
	}
}

func lengthOfValue(value []byte, dataType jsonparser.ValueType) int {
	valueLen := len(value)
	switch dataType {
	case jsonparser.String:
		valueLen += 2
	}
	return valueLen
}
