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
	"fmt"
	"gopkg.in/yaml.v3"
)

type yamlParser struct {
	data []byte
	root *yaml.Node
}

func newYAMLParser(data []byte) *yamlParser {
	p := &yamlParser{
		data: data,
	}

	return p
}

func (p *yamlParser) Iterate(cb Callback) error {
	var n yaml.Node
	err := yaml.Unmarshal(p.data, &n)
	if err != nil {
		return err
	}
	p.root = &n

	p.iterateDocumentNode(p.root, &Path{}, cb)
	return nil
}

func (p *yamlParser) iterateDocumentNode(n *yaml.Node, path *Path, cb Callback) {
	for _, node := range n.Content {
		switch node.Kind {
		case yaml.MappingNode:
			p.iterateMappingNode(node, path, cb)
		case yaml.SequenceNode:
			p.iterateSequenceNode(node, path, cb)
		default:
		}
	}
}

func (p *yamlParser) iterateMappingNode(n *yaml.Node, path *Path, cb Callback) {
	for i, node := range n.Content {
		switch node.Kind {
		case yaml.ScalarNode:
			if i%2 == 0 {
				path.Push(node.Value)
				cb(path, &Pos{Line: node.Line, Column: node.Column})
			} else {
				path.Pop()
			}
		case yaml.MappingNode:
			p.iterateMappingNode(node, path, cb)
		case yaml.SequenceNode:
			p.iterateSequenceNode(node, path, cb)
		default:
		}
	}

	path.Pop()
}

func (p *yamlParser) iterateSequenceNode(n *yaml.Node, path *Path, cb Callback) {
	for i, node := range n.Content {
		path.Push(fmt.Sprintf("%d", i))
		switch node.Kind {
		case yaml.ScalarNode:
			cb(path, &Pos{Line: node.Line, Column: node.Column})
			path.Pop()
		case yaml.MappingNode:
			cb(path, &Pos{Line: node.Line, Column: node.Column})
			p.iterateMappingNode(node, path, cb)
		case yaml.SequenceNode:
			p.iterateSequenceNode(node, path, cb)
		default:
		}
	}

	path.Pop()
}
