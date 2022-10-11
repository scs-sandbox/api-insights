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
import Tabs from './Tabs';

describe('<Tabs />', () => {
  test('Basic', () => {
    const onChangeIndex = jest.fn();

    const headers = [0, 1].map((i) => (
      <div>{`th-${i}`}</div>
    ));

    const { container } = render((
      <Tabs
        selectedTabIndex={1}
        onChangeIndex={onChangeIndex}
        headers={headers}
      >
        <div>tab-0</div>
        <div>tab-1</div>
      </Tabs>
    ));

    const th0 = screen.getByText('th-0');
    expect(th0).toBeInTheDocument();
    const th1 = screen.getByText('th-1');
    expect(th1).toBeInTheDocument();

    const tb0 = screen.getByText('tab-0');
    expect(tb0).toBeInTheDocument();
    const tb1 = screen.getByText('tab-1');
    expect(tb1).toBeInTheDocument();

    const tabHeader0 = container.querySelector('[data-index="0"]');
    const tabHeader1 = container.querySelector('[data-index="1"]');
    expect(tabHeader1).toHaveClass('tab-header', 'active');

    fireEvent.click(tabHeader0, {});
    expect(onChangeIndex).toHaveBeenCalledTimes(1);
  });

  test('No onChangeIndex', () => {
    const onChangeIndex = jest.fn();

    const headers = [0, 1].map((i) => (
      <div>{`th-${i}`}</div>
    ));

    const { container } = render((
      <Tabs
        selectedTabIndex={1}
        onChangeIndex={onChangeIndex}
        headers={headers}
      >
        <div>tab-0</div>
        <div>tab-1</div>
      </Tabs>
    ));

    const tabHeader1 = container.querySelector('[data-index="1"]');

    fireEvent.click(tabHeader1, {});
    expect(onChangeIndex).toHaveBeenCalledTimes(1);
  });

  test('No headers', () => {
    render((
      <Tabs
        selectedTabIndex={0}
        headers={null}
      >
        <div>tab-0</div>
        <div>tab-1</div>
      </Tabs>
    ));

    const tb0 = screen.getByText('tab-0');
    expect(tb0).toBeInTheDocument();
  });
});
