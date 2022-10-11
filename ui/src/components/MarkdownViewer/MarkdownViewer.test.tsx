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
import MarkdownViewer from './MarkdownViewer';

describe('<MarkdownViewer />', () => {
  test('Basic', () => {
    render((
      <MarkdownViewer text="See [more](/home.html)" />
    ));

    const link = screen.getByText('more');
    expect(link).toHaveAttribute('href', '/home.html');
  });

  test('Basic', () => {
    render((
      <MarkdownViewer />
    ));
  });
});
