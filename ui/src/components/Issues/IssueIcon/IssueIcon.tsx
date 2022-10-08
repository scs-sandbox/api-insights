import './IssueIcon.scss';

export type KnownIssueType = 'error' | 'warning' | 'hint' | 'info';

export type IssueType = KnownIssueType | string;

type Props = {
  severity: IssueType;
  title?: string;
};

export default function IssueIcon(props: Props) {
  const className = `issue-icon issue-icon-${props.severity}`;

  const allProps = {
    ...props,
    className,
  };

  return <span {...allProps} />;
}
