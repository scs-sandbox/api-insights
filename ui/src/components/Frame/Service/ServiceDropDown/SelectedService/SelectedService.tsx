import { MouseEvent } from 'react';
import dayjs from 'dayjs';
import CircleScore from '../../CircleScore/CircleScore';
import EditGroupButton from '../../EditGroupButton/EditGroupButton';
import { ServiceData } from '../../../../../query/service';
import './SelectedService.scss';

type Props = {
  service: ServiceData.Service;
  placeholder?: string;
  specServiceSummary?: ServiceData.ServiceSummary;
  onServiceUpdated?: () => void;
  onClickService?: (event: MouseEvent<HTMLDivElement>) => void;
}

export default function SelectedService(props: Props) {
  if (!props.service) {
    return (
      <div className="selected-service" onClick={props.onClickService}>
        <div className="place-holder drop-block">
          {props.placeholder || 'Select a service'}
        </div>
      </div>
    );
  }

  const specServiceSummary = props.specServiceSummary || props.service.summary;

  const renderUpdateAt = () => {
    const updatedAt = specServiceSummary?.updated_at || '';

    if (!updatedAt) {
      return <div className="last-updated hidden" />;
    }

    const time = dayjs(updatedAt).format('MMM DD, HH:mm');

    return (
      <div className="last-updated">
        <span className="last-updated-title">Updated</span>
        {time}
      </div>
    );
  };

  const score = (
    <div className="score-col">
      <CircleScore size={64} thickness={4} value={specServiceSummary?.score}>
        <div className="info-text">
          <div className="value">{specServiceSummary?.score}</div>
        </div>
      </CircleScore>
    </div>
  );

  const updatedAt = renderUpdateAt();

  return (
    <div className="selected-service">
      {score}
      <div className="info-col">
        <div className="group-part">
          {props.service.product_tag}
          <EditGroupButton
            service={props.service}
            onServiceUpdated={props.onServiceUpdated}
          />
        </div>
        <div
          className="service-title drop-block"
          onClick={props.onClickService}
        >
          {props.service.title}
        </div>
        {updatedAt}
      </div>
    </div>
  );
}
