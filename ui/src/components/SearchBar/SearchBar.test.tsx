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
import SearchBar from './SearchBar';

describe('<SearchBar />', () => {
  test('Basic', () => {
    const onChange = jest.fn();
    const onClear = jest.fn();
    const { container } = render((
      <SearchBar searchKey="hello" onSearchKeyChanged={onChange} onSearchKeyCleared={onClear} />
    ));

    const value = screen.getByDisplayValue('hello');
    expect(value).toHaveClass('search-input');
    fireEvent.change(value, { target: { value: 'world' } });
    expect(onChange).toHaveBeenCalledTimes(1);

    const clear = container.querySelector('.search-clear');
    expect(clear).toBeInTheDocument();
    if (clear) fireEvent.click(clear, {});
    expect(onClear).toHaveBeenCalledTimes(1);
  });
});
