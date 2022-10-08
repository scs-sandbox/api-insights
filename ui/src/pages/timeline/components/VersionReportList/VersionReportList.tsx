import VersionReport, { VersionData } from '../VersionReport/VersionReport';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import SpecStateIcon, {
  SpecState,
} from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import { ServiceData } from '../../../../query/service';
import './VersionReportList.scss';

type Props = {
  service: ServiceData.Service;
  versionList?: VersionData[];
};

export default function VersionReportList(props: Props) {
  const className = 'spec-version-report-list';

  const renderVersion = (data: VersionData) => (
    <li
      key={data.version}
      data-version={data.version}
      className="version-block-item"
    >
      <div className="item-line">
        <SpecStateIcon value={data.latestRevision.state as SpecState} />
      </div>
      <div className="item-content">
        <div className="item-content-time">
          <SpecTime icon="dark" time={data.updated_at} />
        </div>
        <VersionReport service={props.service} data={data} />
      </div>
    </li>
  );

  const versionList = (props.versionList || []).map(renderVersion);

  return <div className={className}>{versionList}</div>;
}
