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

function scrollTo(
  querySelector: string,
  adjustOffset = 0,
) {
  const targetElement = document.querySelector(querySelector);
  if (!targetElement) return;

  const viewPortTop = document.body.getBoundingClientRect().top;
  const targetElementTop = targetElement.getBoundingClientRect().top;
  const offset = targetElementTop - viewPortTop - adjustOffset;
  window.scrollTo(0, offset);
}

export default function elementScrollTo(
  querySelector: string,
  adjustOffset = 0,
  delay = 100,
) {
  if (delay === 0) {
    scrollTo(querySelector, adjustOffset);
    return;
  }

  setTimeout(() => {
    scrollTo(querySelector, adjustOffset);
  }, delay);
}
