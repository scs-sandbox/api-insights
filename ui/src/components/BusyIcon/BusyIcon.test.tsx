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
import BusyIcon, { BusyIconType } from './BusyIcon';

describe('<BusyIcon />', () => {
  test('Circle', () => {
    const { container } = render((
      <BusyIcon />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeNull();
    expect(container.querySelector('.icon-circle')).toBeInTheDocument();
  });

  test('Arrow Circle', () => {
    const { container } = render((
      <BusyIcon type={BusyIconType.ArrowCircle} />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeNull();
    expect(container.querySelector('.icon-arrowcircle')).toBeInTheDocument();
  });

  test('Busy', () => {
    const { container } = render((
      <BusyIcon busy />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeInTheDocument();
    expect(container.querySelector('.icon-circle')).toBeInTheDocument();
  });
});
