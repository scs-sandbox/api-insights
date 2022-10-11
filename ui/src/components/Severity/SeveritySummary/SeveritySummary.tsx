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

import { ComplianceData } from '../../../query/compliance';
import { KnownSeverity } from '../SeverityIcon/SeverityIcon';
import SeverityItem from '../SeverityItem/SeverityItem';
import './SeveritySummary.scss';

type Props = {
  data: ComplianceData.ComplianceSeveritySummary[];
  showLabel?: boolean;
};

type SeverityItemData = {
  name: string;
  summary: {
    total: number;
  };
};

export const SEVERITY_LIST: KnownSeverity[] = [
  'error',
  'warning',
  'info',
  'hint',
];

function calcSummary(severity: string, list: ComplianceData.ComplianceSeveritySummary[]) {
  const summary = {
    total: 0,
  };

  list.forEach((i) => {
    const stat: ComplianceData.ComplianceSeveritySummaryItem = i[severity] || {};
    const count = stat.count || 0;
    summary.total += count;
  });

  return summary;
}

export default function SeveritySummary(props: Props) {
  const data = props.data || [];

  const renderItem = (i: SeverityItemData) => (
    <SeverityItem
      key={i.name}
      count={i.summary.total}
      severity={i.name}
      showLabel={props.showLabel}
    />
  );

  const list = SEVERITY_LIST.map((i) => ({
    name: i,
    summary: calcSummary(i, data),
  })).map(renderItem);

  return <div className="severity-summary">{list}</div>;
}
