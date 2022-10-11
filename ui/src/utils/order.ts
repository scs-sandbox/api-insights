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

/*
example:
    const a = [{name: "Team-1", score: 100}, {name: "Team-2", score: 99}]
    const b = buildSortedListByStringField(a, "score");

result:
    then b is [{name: "Team-2", score: 99}, {name: "Team-1", score: 100}]
 */

export default function buildSortedListByStringField(
  list: unknown[],
  stringField: string,
  desc?: boolean,
): unknown[] {
  if (!list) return [];
  if (!Array.isArray(list)) return [];

  return [...list].sort((a: unknown, b: unknown) => {
    const va = a[stringField];
    const vb = b[stringField];
    if (va === vb) return 0;

    const ascResult = va > vb ? 1 : -1;
    return desc ? 0 - ascResult : ascResult;
  });
}
