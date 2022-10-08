import { ReactNode } from 'react';
import dayjs from 'dayjs';
import './SpecTime.scss';

type Props = {
  icon?: ReactNode;
  beginText?: string;
  time: string;
};

export default function SpecTime(props: Props) {
  if (!props.time) return null;

  const value = dayjs(props.time).format('MMM DD, HH:mm');

  return (
    <span className="spec-time">
      {props.icon && <span className={`icon ${props.icon}`} />}
      {props.beginText && <span className="begin-text">{props.beginText}</span>}
      <span className="time-value">{value}</span>
    </span>
  );
}
