import DownloadIcon from '../../../../components/DownloadIcon/DownloadIcon';
import SpecStateIcon from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import { SpecData } from '../../../../query/spec';
import './Snapshot.scss';

type Props = {
  data: SpecData.Snapshot;
};

export default function Snapshot(props: Props) {
  return (
    <div className="snapshot-item block-item-light">
      <SpecStateIcon value={SpecData.SpecState.Release} />
      <div className="snapshot-item-info">
        <div className="detail">
          <SpecTime time={props.data.updated_at} />
        </div>
        <div className="snapshot-name">{props.data.name}</div>
      </div>
      <div className="spec-download">
        <DownloadIcon />
      </div>
    </div>
  );
}
