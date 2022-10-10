import { SpecData } from '../../../query/spec';
import './SpecStateIcon.scss';

type Props = {
  value?: SpecData.SpecState;
};

export default function EnvIcon(props: Props) {
  const value = (props.value || SpecData.SpecState.Development).toLowerCase();
  const className = `env-icon icon-${value}`;

  return <i className={className} />;
}
