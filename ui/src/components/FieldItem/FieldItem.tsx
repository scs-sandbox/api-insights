import { HTMLAttributes, ReactNode } from 'react';
import classNames from '../../utils/className';
import './FieldItem.scss';

type Props = HTMLAttributes<HTMLElement> & {
  label?: ReactNode;
};

export default function FieldItem(props: Props) {
  const className = classNames('field-item', props.className);

  return (
    <div className={className}>
      <div className="field-name">{props.label}</div>
      <div className="field-input">{props.children}</div>
    </div>
  );
}
