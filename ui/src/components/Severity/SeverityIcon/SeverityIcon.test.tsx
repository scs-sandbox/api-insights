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

import { render } from '@testing-library/react';
import SeverityIcon from './SeverityIcon';

describe('<SeverityIcon />', () => {
  test('Error', () => {
    const { container } = render((
      <SeverityIcon severity="error" title="Error" />
    ));

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Warning', () => {
    const { container } = render((
      <SeverityIcon severity="warning" title="Warning" />
    ));

    const icon = container.querySelector('.severity-icon-warning');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Warning');
  });

  test('Hint', () => {
    const { container } = render((
      <SeverityIcon severity="hint" title="Hint" />
    ));

    const icon = container.querySelector('.severity-icon-hint');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Hint');
  });

  test('Info', () => {
    const { container } = render((
      <SeverityIcon severity="info" title="Info" />
    ));

    const icon = container.querySelector('.severity-icon-info');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Info');
  });

  test('Unknown', () => {
    const { container } = render((
      <SeverityIcon severity="something" title="Something" />
    ));

    const icon = container.querySelector('.severity-icon-something');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Something');
  });
});
