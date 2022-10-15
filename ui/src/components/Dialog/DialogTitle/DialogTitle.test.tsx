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
import DialogTitle from './DialogTitle';

describe('<DialogTitle />', () => {
  test('Icon', () => {
    const { container } = render((
      <DialogTitle>My Dialog</DialogTitle>
    ));

    expect(container.querySelector('.title-icon')).toBeInTheDocument();
  });

  test('Title', () => {
    render((
      <DialogTitle>My Dialog</DialogTitle>
    ));

    expect(screen.getByText('My Dialog')).toBeInTheDocument();
  });

  test('Close Button', () => {
    const onClose = jest.fn();

    render((
      <DialogTitle onClose={onClose}>My Dialog</DialogTitle>
    ));

    expect(screen.getByRole('button')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button'), {});

    expect(onClose).toHaveBeenCalledTimes(1);
  });
});
