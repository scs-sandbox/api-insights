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
        const issue = findings[severity] as ComplianceData.ComplianceIssue;
        if (!issue) return;

        const { rules } = issue;
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
