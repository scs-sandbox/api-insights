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
import SeverityItem from './SeverityItem';

describe('<SeverityItem />', () => {
  test('little props', () => {
    const { container } = render((
      <SeverityItem severity="error" />
    ));

    const item = container.querySelector('.severity-item');
    expect(item).toBeInTheDocument();
    expect(item).toHaveClass('severity-item-error');

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Full props', () => {
    const { container } = render((
      <SeverityItem severity="error" count={5} showLabel label="Error(s)" />
    ));

    const item = container.querySelector('.severity-item');
    expect(item).toBeInTheDocument();

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error(s)');

    const count = screen.getByText('5');
    expect(count).toBeInTheDocument();
    expect(count).toHaveClass('severity-item-count');

    const label = screen.getByText('Error(s)');
    expect(label).toBeInTheDocument();
    expect(label).toHaveClass('severity-item-label');
  });
});
