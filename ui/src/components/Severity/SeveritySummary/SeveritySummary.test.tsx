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
import { ComplianceData } from '../../../query/compliance';
import SeveritySummary from './SeveritySummary';

describe('<SeveritySummary />', () => {
  test('All severities', () => {
    const data: ComplianceData.ComplianceSeveritySummary[] = [{
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
      info: {
        count: 100,
      },
      hint: {
        count: 1000,
      },
    },
    {
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
      info: {
        count: 100,
      },
      hint: {
        count: 1000,
      },
    }];

    const { container } = render((
      <SeveritySummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.severity-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem?.querySelector('.severity-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('2');
    expect(errorItem?.querySelector('.severity-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.severity-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem?.querySelector('.severity-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('20');
    expect(warningItem?.querySelector('.severity-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.severity-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem?.querySelector('.severity-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('200');
    expect(infoItem?.querySelector('.severity-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.severity-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem?.querySelector('.severity-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('2000');
    expect(hintItem?.querySelector('.severity-item-label')).toHaveTextContent('Hint');
  });

  test('Data is Null', () => {
    const data: ComplianceData.ComplianceSeveritySummary[] = [];

    const { container } = render((
      <SeveritySummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.severity-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem?.querySelector('.severity-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('0');
    expect(errorItem?.querySelector('.severity-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.severity-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem?.querySelector('.severity-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('0');
    expect(warningItem?.querySelector('.severity-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.severity-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem?.querySelector('.severity-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('0');
    expect(infoItem?.querySelector('.severity-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.severity-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem?.querySelector('.severity-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('0');
    expect(hintItem?.querySelector('.severity-item-label')).toHaveTextContent('Hint');
  });

  test('Part of severities', () => {
    const data: ComplianceData.ComplianceSeveritySummary[] = [{
      count: 1,
      error: {
        count: 1,
      },
    },
    {
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
    }];

    const { container } = render((
      <SeveritySummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.severity-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem?.querySelector('.severity-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('2');
    expect(errorItem?.querySelector('.severity-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.severity-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem?.querySelector('.severity-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('10');
    expect(warningItem?.querySelector('.severity-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.severity-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem?.querySelector('.severity-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('0');
    expect(infoItem?.querySelector('.severity-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.severity-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem?.querySelector('.severity-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('0');
    expect(hintItem?.querySelector('.severity-item-label')).toHaveTextContent('Hint');
  });
});
