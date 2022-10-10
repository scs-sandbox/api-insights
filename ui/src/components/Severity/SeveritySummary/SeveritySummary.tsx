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
