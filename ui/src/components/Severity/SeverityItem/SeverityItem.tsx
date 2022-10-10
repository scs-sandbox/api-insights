import capitalize from '../../../utils/string';
import SeverityIcon, { Severity } from '../SeverityIcon/SeverityIcon';
import './SeverityItem.scss';

type Props = {
  severity: Severity;
  count?: number;
  showLabel?: boolean;
  label?: string;
};

export default function SeverityItem(props: Props) {
  const labelText = props.label || capitalize(props.severity);

  function renderCount() {
    const hasCount = props.count !== undefined && props.count !== null;

    if (!hasCount) return null;

    return <span className="severity-item-count">{props.count}</span>;
  }

  function renderLabel() {
    if (!props.showLabel) return null;

    return <span className="severity-item-label">{labelText}</span>;
  }

  const count = renderCount();
  const label = renderLabel();
  const className = `severity-item severity-item-${props.severity}`;

  return (
    <span className={className}>
      <SeverityIcon severity={props.severity} title={labelText} />
      {count}
      {label}
    </span>
  );
}
