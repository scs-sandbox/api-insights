import { editor } from 'monaco-editor';
import waitFor from './waitFor';
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
export function injectSpaceToEditor(
  item: string,
  targetEditor: editor.IStandaloneCodeEditor,
  targetModel: editor.ITextModel | null,
  refName: string,
) {
  if (!targetModel) return;
  const modMatches = (targetModel)
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
export function injectButtonToEditor(refs: string[], setActiveRefs: any) {
  Array.from(document.getElementsByClassName('mtk1')).forEach((element) => {
    refs.forEach((ref) => {
      if (element.innerHTML.includes(ref) && !(element.innerHTML.includes('injected'))) {
        // eslint-disable-next-line no-param-reassign
        const node = document.createElement('div');
        node.innerHTML = '<div class="reference-icon"></div>View Reference';
        node.className = 'injected';
        element.append(node);
        node.addEventListener('click', () => {
          setActiveRefs((actives: string[]) => {
            if (actives.includes(ref)) {
              return actives;
            }
            return ([...actives, ref]);
          });
        });
      }
    });
  });
}
export function monacoMount(
  diffEditor: editor.IStandaloneDiffEditor,
  refs: string[],
  setActiveRefs: (refs: string[]) => void,
  refName?: string,
) {
  const ogEditor = diffEditor.getOriginalEditor();
  const modEditor = diffEditor.getModifiedEditor();
  const ogMdoel = ogEditor.getModel();
  const modMdoel = modEditor.getModel();
  let itemKey;
  const loadedKey = (refName) ? `.${refName}` : '.mtk1';
  new Set(refs).forEach((item: string) => {
    itemKey = `${item}"`;
    injectSpaceToEditor(itemKey, ogEditor, ogMdoel, refName || 'mtk1');
    injectSpaceToEditor(itemKey, modEditor, modMdoel, refName || 'mtk1');
  });
  waitFor(loadedKey).then(() => {
    injectButtonToEditor(refs, setActiveRefs);
    ogEditor.onDidScrollChange(() => {
      injectButtonToEditor(refs, setActiveRefs);
    });
    ogEditor.onMouseDown(() => {
      injectButtonToEditor(refs, setActiveRefs);
    });
    ogEditor.onMouseUp(() => {
      injectButtonToEditor(refs, setActiveRefs);
    });
    ogEditor.onDidChangeCursorSelection(() => {
      injectButtonToEditor(refs, setActiveRefs);
    });
  });
}
