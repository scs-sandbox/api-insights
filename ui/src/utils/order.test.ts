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

import buildSortedListByStringField from './order';

describe('buildSortedListByStringField()', () => {
  test('null', () => {
    const r = buildSortedListByStringField(null, 'score');
    expect(r.length).toEqual(0);
  });

  test('undefined', () => {
    const r = buildSortedListByStringField(undefined, 'score');
    expect(r.length).toEqual(0);
  });

  test('not array', () => {
    const r = buildSortedListByStringField({} as unknown[], 'score');
    expect(r.length).toEqual(0);
  });

  test('normal array', () => {
    const a = [{ name: 'Team-1', score: 100 }, { name: 'Team-2', score: 99 }];
    const b = buildSortedListByStringField(a, 'score') as { name: string, score: number }[];
    expect(b[0].name).toEqual('Team-2');
    expect(b[0].score).toEqual(99);
    expect(b[1].name).toEqual('Team-1');
    expect(b[1].score).toEqual(100);
  });

  test('normal array, equal items', () => {
    const a = [{ name: 'Team-1', score: 100 }, { name: 'Team-2', score: 100 }];
    const b = buildSortedListByStringField(a, 'score') as { name: string, score: number }[];
    expect(b[0].name).toEqual('Team-1');
    expect(b[0].score).toEqual(100);
    expect(b[1].name).toEqual('Team-2');
    expect(b[1].score).toEqual(100);
  });

  test('normal array, desc', () => {
    const a = [{ name: 'Team-1', score: 100 }, { name: 'Team-2', score: 101 }];
    const b = buildSortedListByStringField(a, 'score', true) as { name: string, score: number }[];
    expect(b[0].name).toEqual('Team-2');
    expect(b[0].score).toEqual(101);
    expect(b[1].name).toEqual('Team-1');
    expect(b[1].score).toEqual(100);
  });
});
