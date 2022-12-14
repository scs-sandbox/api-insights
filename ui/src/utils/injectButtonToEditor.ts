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

export default function injectButtonToEditor(refs: string[], setActiveRefs: any) {
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
            console.log('adding');
            console.log([...actives, ref]);
            return ([...actives, ref]);
          });
        });
      }
    });
  });
}
