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
import SnackAlert from './SnackAlert';

describe('<SnackAlert />', () => {
  test('Success Message', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="success" message="Hello, world" onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });

  test('Error Message', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="error" message="Hello, world" onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });

  test('Duration', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="success" message="Hello, world" autoHideDuration={1000} onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });
});
