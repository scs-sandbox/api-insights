import { ComplianceData } from '../../../query/compliance';
import { KnownIssueType } from '../IssueIcon/IssueIcon';
import IssueItem from '../IssueItem/IssueItem';
import './IssueSummary.scss';

type Props = {
  data: ComplianceData.ComplianceIssues[];
  showLabel?: boolean;
};

type IssueItemData = {
  name: string;
  summary: {
    total: number;
  };
};

export const ISSUE_LIST: KnownIssueType[] = [
  'error',
  'warning',
  'info',
  'hint',
];

function calcSummary(issue: string, list: ComplianceData.ComplianceIssues[]) {
  const summary = {
    total: 0,
  };

  list.forEach((i) => {
    const stat: ComplianceData.ComplianceIssueItem = i[issue] || {};
    const count = stat.count || 0;
    summary.total += count;
  });

  return summary;
}

export default function IssueSummary(props: Props) {
  const data = props.data || [];

  const renderItem = (i: IssueItemData) => (
    <IssueItem
      key={i.name}
      count={i.summary.total}
      severity={i.name}
      showLabel={props.showLabel}
    />
  );

  const list = ISSUE_LIST.map((i) => ({
    name: i,
    summary: calcSummary(i, data),
  })).map(renderItem);

  return <div className="issue-summary">{list}</div>;
}
