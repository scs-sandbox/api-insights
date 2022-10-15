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
import DropDown from './DropDown';

describe('<DropDown />', () => {
  test('options is null', () => {
    const { container } = render((
      <DropDown value="hello" />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);
  });

  test('Empty options', () => {
    const { container } = render((
      <DropDown value="hello" options={[]} />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);
  });

  test('Placeholder', () => {
    const { container } = render((
      <DropDown placeholder="select one" />
    ));

    const placeholder = container.querySelector('.dropdown-placeholder');
    expect(placeholder).toBeInTheDocument();
  });

  test('String Array & Basic Feature', () => {
    const onChange = jest.fn();

    const { container } = render((
      <DropDown value="hello" options={['hello', 'world']} onChange={onChange} />
    ));

    const value = screen.getByText('hello');
    expect(value).toBeInTheDocument();
    expect(value).toHaveClass('dropdown-value');

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = screen.getByText('world');
    expect(menuItem).toBeInTheDocument();
    fireEvent.click(menuItem);
    expect(menuItem).not.toBeInTheDocument();
    expect(onChange).toHaveBeenCalledTimes(1);
  });

  test('Object Array, value and options are given only', () => {
    const { container } = render((
      <DropDown value="" options={[{ text: 'hello' }, { text: 'world' }]} />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = document.querySelector('.menu-item-label');
    expect(menuItem).toBeInTheDocument();
    expect(menuItem).toHaveTextContent('[unknown]');
    fireEvent.click(menuItem);
  });

  test('Object Array', () => {
    const options = [{ id: '1', text: 'hello' }, { id: '2', text: 'world' }];
    const requestOptionValue = (option: {id: string, text: string}) => option.id;
    const renderValue = (value: string) => (
      <div className="dropdown-value">{options.find((i) => i.id === value).text}</div>
    );
    const renderMenuItemLabel = (option: {id: string, text: string}) => (<div>{option.text}</div>);

    const { container } = render((
      <DropDown
        value="1"
        options={options}
        requestOptionValue={requestOptionValue}
        renderValue={renderValue}
        renderMenuItemLabel={renderMenuItemLabel}
      />
    ));

    const value = screen.getByText('hello');
    expect(value).toBeInTheDocument();
    expect(value).toHaveClass('dropdown-value');

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = screen.getByText('world');
    expect(menuItem).toBeInTheDocument();
  });
});
