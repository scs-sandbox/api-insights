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

import { render, screen, fireEvent } from '@testing-library/react';
import DropDownEditBox from './DropDownEditBox';

describe('<DropDownEditBox />', () => {
  test('Basic', () => {
    const onChange = jest.fn();

    render((
      <DropDownEditBox value="hello" onChange={onChange} />
    ));

    const input = screen.getByDisplayValue('hello');
    expect(input).toBeInTheDocument();
    expect(input).toHaveClass('dropdown-value-input');

    fireEvent.change(input, { target: { value: 'world' } });
    expect(onChange).toHaveBeenCalledTimes(1);

    fireEvent.click(input, {});
  });

  test('Readonly', () => {
    render((
      <DropDownEditBox value="hello" readonly />
    ));

    const input = screen.getByDisplayValue('hello');
    fireEvent.click(input, {});

    fireEvent.change(input, { target: { value: 'world' } });
  });
});
