import { HTMLAttributes } from 'react';
import classNames from '../../../utils/className';
import './DialogBody.scss';

export default function DialogBody(props: HTMLAttributes<HTMLElement>) {
  const className = classNames('dialog-body', props.className);

  return <div className={className}>{props.children}</div>;
}
