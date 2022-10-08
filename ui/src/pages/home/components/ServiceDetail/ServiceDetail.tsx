import { MouseEvent } from 'react';
import { Link } from 'react-router-dom';
import dayjs from 'dayjs';
import PersonIcon from '@mui/icons-material/Person';
import PublicIcon from '@mui/icons-material/Public';
import LinkIcon from '@mui/icons-material/Link';
import EditIcon from '@mui/icons-material/Edit';
import Timer from '../../../../components/Timer/Timer';
import classNames from '../../../../utils/className';
import CircleScore from '../../../../components/Frame/Service/CircleScore/CircleScore';
import { ServiceData } from '../../../../query/service';
import './ServiceDetail.scss';

type Props = {
  authEnabled: boolean;
  service: ServiceData.Service;
  isNewCreated: boolean;
  onClickService: (clickedService: ServiceData.Service, e: MouseEvent<HTMLElement>) => void;
  onClickEdit: (clickedService: ServiceData.Service, e: MouseEvent<HTMLElement>) => void;
  onCancelNewServiceNotation: () => void;
}

export default function ServiceDetail(props: Props) {
  const time = props.service?.summary?.updated_at
    ? dayjs(props.service.summary.updated_at).format('MMM DD, HH:mm')
    : '';
  const name = props.service.contact?.name || '';
  const url = props.service.contact?.url || '';
  const email = props.service.contact?.email || '';
  const contactInfo = url ? decodeURIComponent(url) : '';
  const className = classNames(
    'service-detail',
    props.isNewCreated ? 'new' : '',
  );

  const newMark = props.isNewCreated && (
    <Timer ms={60 * 1000} onClose={props.onCancelNewServiceNotation}>
      <span className="new-mark">NEW</span>
    </Timer>
  );

  const noSpec = (!props.service.summary) ? (
    <div className="no-spec">No Spec</div>
  ) : null;

  const renderVisibility = () => {
    if (!props.authEnabled) return null;

    const isPrivate = props.service.visibility === 'private';
    const visIcon = isPrivate ? <PersonIcon className="icon" /> : <PublicIcon className="icon" />;
    const visText = isPrivate ? 'Private' : 'Public';

    return (
      <div className="visibility">
        {visIcon}
        {visText}
      </div>
    );
  };

  const visibility = renderVisibility();

  const renderEditButton = () => (
    <div className="edit" onClick={(e) => { e.preventDefault(); e.stopPropagation(); props.onClickEdit(props.service, e); }}>
      <EditIcon className="icon" />
      Edit Service
    </div>
  );

  const editButton = renderEditButton();

  return (
    <Link
      id={`service-${props.service.id}`}
      onClick={(e) => {
        e.stopPropagation();
        props.onClickService(props.service, e);
      }}
      className={className}
      to={`/timeline?service=${props.service.name_id || props.service.id}`}
    >
      <div className="detail-header">
        <div className="circle-container">
          <CircleScore value={props.service.summary?.score} size={96} darkTrack>
            {noSpec}
          </CircleScore>
        </div>
        <div className="header-text">
          <div className="service-name">
            {newMark}
            {props.service.title}
            <span className="version-label">{props.service.summary?.version}</span>
            {/* <span className="revision-label">{props.service.summary?.revision}</span> */}
          </div>
          <div className={classNames('updated-time', time ? '' : 'no-time')}>
            &#x21bb; Updated
            <span className="updated-time-value">{time}</span>
          </div>
        </div>
        <div className="upper-right">
          {visibility}
          {editButton}
        </div>
      </div>
      <div className="service-content">
        {/* <div className="content-title">Description:</div> */}
        <div
          className="content-description mul-line-2"
        >
          {props.service.description}
        </div>
        <div className="content-connect">
          <div
            className={`connect-detail ${!name && 'no-info'}`}
            title={name}
            onClick={(e) => { e.preventDefault(); e.stopPropagation(); window.location.href = `mailto:${email}?subject=Subject&body=message%20goes%20here`; }}
          >
            <PersonIcon className="icon" />
            <span className="connect-text">{name || email || 'No Contact'}</span>
          </div>
          <div
            className={`connect-detail ${!contactInfo && 'no-info'}`}
            title={contactInfo}
            onClick={(e) => { e.preventDefault(); e.stopPropagation(); window.open(contactInfo, '_blank').focus(); }}
          >
            <LinkIcon className="icon" />
            <span className="connect-text">{contactInfo ? 'Reference' : 'No Reference'}</span>
          </div>
        </div>
      </div>
    </Link>
  );
}
