import { HTMLAttributes } from 'react';
import classNames from '../../../utils/className';
import './DialogActions.scss';

type Props = HTMLAttributes<HTMLElement>;

export default function DialogActions(
  props: Props,
) {
  const className = classNames('dialog-actions', props.className);

  return <div className={className}>{props.children}</div>;
}
