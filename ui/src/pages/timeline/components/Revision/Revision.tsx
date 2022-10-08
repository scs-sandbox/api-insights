import { useState } from 'react';
import { Link } from 'react-router-dom';
import DownloadIcon from '../../../../components/DownloadIcon/DownloadIcon';
import ScoreLevel from '../../../../components/Frame/Service/ScoreLevel/ScoreLevel';
import SpecStateIcon, {
  SpecState,
} from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import { buildApiAbsoluteUrl } from '../../../../query/api';
import { SpecData } from '../../../../query/spec';
import { ComplianceData } from '../../../../query/compliance';
import classNames from '../../../../utils/className';
import './Revision.scss';

export type RevisionData = SpecData.Spec & {
  complianceList: ComplianceData.Compliance[];
};

type Props = {
  data: RevisionData;
  to?: string;
  className?: string;
  onMouseEnter?: (item: RevisionData) => void;
  onMouseLeave?: (item: RevisionData) => void;
  onReleased?: (item: RevisionData) => void;
  onArchived?: (item: RevisionData) => void;
};

/**
 * TODO: comment
 * @param props props
 * @returns jsx
 */
export default function Revision(props: Props) {
  const [openMenu, setOpenMenu] = useState(false);

  const { data } = props;
  const className = classNames(
    'revision-item block-item-light',
    props.className,
  );

  const onMouseEnter = () => {
    if (props.onMouseEnter) {
      props.onMouseEnter(props.data);
    }
  };

  const onMouseLeave = () => {
    if (props.onMouseLeave) {
      props.onMouseLeave(props.data);
    }
  };

  const onOpenMenu = () => {
    setOpenMenu(true);
  };

  const onCloseMenu = () => {
    setOpenMenu(false);
  };

  const onRelease = () => {
    if (props.onReleased) {
      props.onReleased(null);
    }
  };

  const onArchive = () => {
    if (props.onArchived) {
      props.onArchived(null);
    }
  };

  const handleSpecDownload = () => {
    const url = `/services/${data.service_id}/specs/${data.id}/doc`;
    window.open(buildApiAbsoluteUrl(url));
  };
  const renderScore = () => {
    if (data.score === null || data.score === undefined) {
      return <div className="loader" />;
    }

    return (
      <ScoreLevel className="score block-item" score={data.score}>
        {data.score}
      </ScoreLevel>
    );
  };

  const renderMenu = () => {
    if (!openMenu) return null;

    return (
      <div className="action-menu" onClick={onCloseMenu}>
        <div className="menu-item" onClick={onRelease}>
          <SpecStateIcon value={SpecState.Release} />
          <div className="menu-item-label">Live</div>
        </div>
        <div className="menu-item" onClick={onArchive}>
          <SpecStateIcon value={SpecState.Archive} />
          <div className="menu-item-label">Archive</div>
        </div>
      </div>
    );
  };

  const score = renderScore();
  const menu = renderMenu();

  return (
    <div
      data-id={data.id}
      className={className}
      onMouseEnter={onMouseEnter}
      onMouseLeave={onMouseLeave}
    >
      <SpecStateIcon value={data.state as SpecState} />
      <div className="revision-item-info">
        <div className="detail">
          <SpecTime time={data.updated_at} />
          {score}
        </div>
        <span className="revision-label">{data.revision}</span>
      </div>
      <div
        className="revision-item-menu"
        onMouseEnter={onOpenMenu}
        onMouseLeave={onCloseMenu}
      >
        <div className="action-button" />
        {menu}
      </div>
      <div className="revision-item-action">
        <Link to={props.to} className="view-report button-rc">
          View Full Report
        </Link>
      </div>
      <div className="spec-download" onClick={handleSpecDownload}>
        <DownloadIcon />
      </div>
    </div>
  );
}
