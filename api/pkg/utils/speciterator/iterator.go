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
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Callback func(path *Path, pos *Pos)

type SpecIterator interface {
	Iterate(cb Callback) error
}

type specIterator struct {
	data     []byte
	iterator SpecIterator
}

func NewSpecIterator(data []byte) SpecIterator {
	si := &specIterator{
		data: data,
	}

	if si.isJSONSpec() {
		si.iterator = newJSONParser(data)
	} else {
		si.iterator = newYAMLParser(data)
	}

	return si
}

func (si *specIterator) Iterate(cb Callback) error {
	return si.iterator.Iterate(cb)
}

func (si *specIterator) isJSONSpec() bool {
	trim := bytes.TrimLeftFunc(si.data, unicode.IsSpace)
	return bytes.HasPrefix(trim, []byte("{"))
}

type Pos struct {
	Line   int
	Column int
	Offset int
}

func (p *Pos) String() string {
	return fmt.Sprintf("%d:%d:%d", p.Line, p.Column, p.Offset)
}

type Path struct {
	path []string
}

func (p *Path) Push(pa string) {
	p.path = append(p.path, pa)
}

func (p *Path) Pop() {
	if len(p.path) == 0 {
		return
	}
	p.path = p.path[:len(p.path)-1]
}

func (p *Path) String() string {
	return strings.Join(p.path, "|")
}
