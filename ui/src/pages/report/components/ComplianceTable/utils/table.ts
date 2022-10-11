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

import { ComplianceData } from '../../../../../query/compliance';

export type AnalyseTableRowData = {
  id: string;
  analyzer: string;
  severity: string;
  range: ComplianceData.ComplianceRange;
  code: string;
  message: string;
  mitigation: string;
  detail: ComplianceData.ComplianceRuleDataItem[];
};

function getSeverityOrder(severity: string) {
  switch (severity) {
    case 'error': return 0;
    case 'warning': return 1;
    case 'info': return 2;
    case 'hint': return 3;
    default: return 100;
  }
}

function compareRows(
  a: AnalyseTableRowData,
  b: AnalyseTableRowData,
  sortBy: string,
) {
  if (sortBy === 'severity') {
    const orderA = getSeverityOrder(a[sortBy]);
    const orderB = getSeverityOrder(b[sortBy]);
    return orderB - orderA;
  }

  if (a[sortBy] === b[sortBy]) return 0;

  return a[sortBy] < b[sortBy] ? -1 : 1;
}

export function sortRows(
  rows: AnalyseTableRowData[],
  sortBy: string,
  sortDesc: boolean,
) {
  return rows.sort((a: AnalyseTableRowData, b: AnalyseTableRowData) => {
    const sortValue = compareRows(a, b, sortBy);
    if (sortValue) {
      return sortDesc ? 0 - sortValue : sortValue;
    }

    return compareRows(a, b, '');
  });
}

/**
 * converts compliance data into format for frontend display
 * @param list list of compliance details
 * @returns AnalyseTableRowData[]
 */
export function convertToTableData(
  list: ComplianceData.Compliance[],
): AnalyseTableRowData[] {
  if (!list) return null;

  return list.reduce(
    (pre: AnalyseTableRowData[], cur: ComplianceData.Compliance) => {
      const findings = cur.result?.findings;
      if (!findings) return pre;

      let rows = [];
      Object.keys(findings).forEach((severity) => {
        const findingItem = findings[severity] as ComplianceData.ComplianceFindingItem;
        if (!findingItem) return;

        const { rules } = findingItem;
        if (!rules) return;

        Object.keys(rules).forEach((code) => {
          const rule = rules[code];

          const row = {
            id: `${pre.length + rows.length}`,
            analyzer: cur.analyzer,
            severity,
            code,
            message: rule.message,
            mitigation: rule.mitigation,
            detail: rule.data,
          };

          rows = [...rows, row];
        });
      });

      if (!rows.length) return pre;

      return [...pre, ...rows];
    },
    [],
  );
}
