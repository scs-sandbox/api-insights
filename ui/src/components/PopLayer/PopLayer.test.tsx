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

import { render, screen } from '@testing-library/react';
import PopLayer from './PopLayer';

describe('document.body has not had #poplayer yet', () => {
  test('Basic', () => {
    render((
      <PopLayer>
        <div>hello</div>
      </PopLayer>
    ));

    const div = screen.getByText('hello');
    expect(div).toBeInTheDocument();
  });

  test('document.body has had #poplayer already', () => {
    const domPopLayer = document.createElement('div');
    domPopLayer.id = 'pop-layer';
    document.body.appendChild(domPopLayer);

    render((
      <PopLayer>
        <div>hello</div>
      </PopLayer>
    ));

    const div = screen.getByText('hello');
    expect(div).toBeInTheDocument();
  });
});
