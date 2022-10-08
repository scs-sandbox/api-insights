import capitalize from '../../../utils/string';
import IssueIcon, { IssueType } from '../IssueIcon/IssueIcon';
import './IssueItem.scss';

type Props = {
  severity: IssueType;
  count?: number;
  showLabel?: boolean;
  label?: string;
};

export default function IssueItem(props: Props) {
  const labelText = props.label || capitalize(props.severity);

  function renderCount() {
    const hasCount = props.count !== undefined && props.count !== null;

    if (!hasCount) return null;

    return <span className="issue-item-count">{props.count}</span>;
  }

  function renderLabel() {
    if (!props.showLabel) return null;

    return <span className="issue-item-label">{labelText}</span>;
  }

  const count = renderCount();
  const label = renderLabel();
  const className = `issue-item issue-item-${props.severity}`;

  return (
    <span className={className}>
      <IssueIcon severity={props.severity} title={labelText} />
      {count}
      {label}
    </span>
  );
}
