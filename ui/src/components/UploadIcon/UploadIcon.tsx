import BusyIcon, { BusyIconType } from '../BusyIcon/BusyIcon';
import './UploadIcon.scss';

type Props = {
  busy?: boolean;
  busyType?: BusyIconType;
};

export default function UploadIcon(props: Props) {
  const className = `upload-icon${props.busy ? ' state-busy' : ''}`;
  const busyIcon = props.busy && <BusyIcon type={props.busyType} busy />;

  return <i className={className}>{busyIcon}</i>;
}
