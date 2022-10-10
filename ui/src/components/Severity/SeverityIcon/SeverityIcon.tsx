import './SeverityIcon.scss';

export type KnownSeverity = 'error' | 'warning' | 'hint' | 'info';

export type Severity = KnownSeverity | string;

type Props = {
  severity: Severity;
  title?: string;
};

export default function SeverityIcon(props: Props) {
  const className = `severity-icon severity-icon-${props.severity}`;

  const allProps = {
    ...props,
    className,
  };

  return <span {...allProps} />;
}
