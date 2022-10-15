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

import { fireEvent, render, screen } from '@testing-library/react';
import Toggle from './Toggle';

describe('<Toggle />', () => {
  test('Unchecked', () => {
    const { container } = render((
      <Toggle />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).not.toHaveClass('checked');
  });

  test('Checked', () => {
    const { container } = render((
      <Toggle checked label="Using" />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).toHaveClass('checked');

    const label = screen.getByText('Using');
    expect(label).toBeInTheDocument();

    fireEvent.click(toggle);
  });

  test('onToggle', () => {
    const onToggle = jest.fn();
    const { container } = render((
      <Toggle checked label="Using" onToggle={onToggle} />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).toHaveClass('checked');

    fireEvent.click(toggle);
    expect(onToggle).toHaveBeenCalledTimes(1);
  });
});
