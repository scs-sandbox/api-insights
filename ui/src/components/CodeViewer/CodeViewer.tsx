/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import MonacoEditor, { EditorDidMount } from 'react-monaco-editor';
import './CodeViewer.scss';

type Selection = {
  start: number;
  end: number;
};

type Props = {
  language?: string;
  theme?: string;
  value?: string;
  selections?: Selection[];
};

export default function CodeViewer(props: Props) {
  const editorDidMount: EditorDidMount = (editor, monaco) => {
    if (!props.selections || !props.selections.length) return;

    const options = {
      isWholeLine: true,
      inlineClassName: 'selected-line',
    };

    const selections = props.selections.map((i: Selection) => ({
      range: new monaco.Range(i.start, 1, i.end, 1),
      options,
    }));

    editor.deltaDecorations([], selections);

    const firstSelection = props.selections[0];
    editor.revealPositionInCenter({
      lineNumber: firstSelection.start,
      column: 1,
    });
  };

  return (
    <div className="code-viewer">
      <MonacoEditor
        language={props.language || 'json'}
        theme={props.theme || 'vs-dark'}
        value={props.value}
        options={{
          readOnly: true,
          selectOnLineNumbers: true,
          minimap: { enabled: false },
        }}
        editorDidMount={editorDidMount}
      />
    </div>
  );
}
