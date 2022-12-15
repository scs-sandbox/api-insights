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

import { PropaneSharp } from '@mui/icons-material';
import { editor } from 'monaco-editor';
import { useState, useEffect } from 'react';
import { MonacoDiffEditor } from 'react-monaco-editor';
import { DiffData } from '../../../../../query/compare';
import { SpecData } from '../../../../../query/spec';
import iterateObject from '../../../../../utils/iterateObject';
import injectButtonToEditor from '../../../../../utils/injectButtonToEditor';
import waitFor from '../../../../../utils/waitFor';
import './DiffReference.scss';

type Props = {
    data: string;
    leftSpec?: SpecData.Spec;
    rightSpec?: SpecData.Spec;
    removeActive: (ref: string) => void;
    setActiveRefs: (refs: string[]) => void;
};

export default function DiffReference(props: Props) {
  const leftSchemaSection = (props.leftSpec)
    ? JSON.parse(props.leftSpec?.doc).components.schemas : {};
  const rightSchemaSection = (props.rightSpec)
    ? JSON.parse(props.rightSpec?.doc).components.schemas : {};
  const refName = props.data.split('/schemas/')[1];
  const [isOpen, toggleItemOpen] = useState<boolean | null>(null);
  const [refs, setRefs] = useState<string[]>([]);
  useEffect(() => {
    iterateObject(leftSchemaSection, refs);
    iterateObject(rightSchemaSection, refs);
  }, []);
  return (
    <div id={refName}>
      <div className="reference-button" onClick={() => props.removeActive(props.data)}>
        <div className="reference-label">Viewing reference</div>
        <div className="reference-name">
          {props.data}
          <div className="close-icon" />
        </div>
      </div>
      <MonacoDiffEditor
        height="400px"
        width="100%"
        original={JSON.stringify(leftSchemaSection[refName], null, '\t')}
        value={JSON.stringify(rightSchemaSection[refName], null, '\t')}
        options={{
          minimap: {
            enabled: false,
          },
          overviewRulerLanes: 0,
        }}
        editorDidMount={(diffEditor) => {
          console.log('ref did mount:', refs);
          if (!refs) return;
          const ogEditor = diffEditor.getOriginalEditor();
          const modEditor = diffEditor.getModifiedEditor();
          const ogMdoel = ogEditor.getModel();
          const modMdoel = modEditor.getModel();
          function injectSpaceToEditor(
            item: string,
            targetEditor: editor.IStandaloneCodeEditor,
            targetModel: editor.ITextModel | null,
          ) {
            if (!targetModel) return;
            const modMatches = (modMdoel)
              ? targetModel.findMatches(item, false, false, false, null, true, undefined)
              : [];
            modMatches.forEach((match) => {
              targetEditor.createDecorationsCollection([
                {
                  range: match.range,
                  options: {
                    isWholeLine: false,
                    after: {
                      attachedData: item,
                      content: '              ',
                      inlineClassName: refName,
                      cursorStops: 1,
                    },
                  },
                },
              ]);
            });
          }
          let itemKey;
          new Set(refs).forEach((item: any) => {
            itemKey = `${item}"`;
            injectSpaceToEditor(itemKey, ogEditor, ogMdoel);
            injectSpaceToEditor(itemKey, modEditor, modMdoel);
          });
          waitFor(`.${refName}`).then(() => {
            injectButtonToEditor(refs, props.setActiveRefs);
            // setRefs((updatedRefs) => {
            //   injectButtonToEditor(updatedRefs, props.setActiveRefs);
            //   return updatedRefs;
            // });
            ogEditor.onMouseDown(() => {
              injectButtonToEditor(refs, props.setActiveRefs);
              return undefined;
            });
            ogEditor.onMouseUp(() => {
              injectButtonToEditor(refs, props.setActiveRefs);
              return undefined;
            });
            ogEditor.onMouseLeave(() => {
              injectButtonToEditor(refs, props.setActiveRefs);
              return undefined;
            });
            ogEditor.onDidChangeCursorSelection(() => {
              injectButtonToEditor(refs, props.setActiveRefs);
              return undefined;
            });
            ogEditor.onDidScrollChange(() => {
              injectButtonToEditor(refs, props.setActiveRefs);
            });
          });
        }}
      />
    </div>
  );
}
