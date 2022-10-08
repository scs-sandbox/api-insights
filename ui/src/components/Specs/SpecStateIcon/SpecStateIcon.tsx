import './SpecStateIcon.scss';

export enum SpecState {
  Archive = 'Archive',
  Release = 'Release',
  Development = 'Development',
  Latest = 'Latest',
}

type Props = {
  value?: SpecState;
};

export default function EnvIcon(props: Props) {
  const className = `env-icon icon-${(
    props.value || SpecState.Development
  ).toLowerCase()}`;

  return <i className={className} />;
}
