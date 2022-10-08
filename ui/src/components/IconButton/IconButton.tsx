import { ButtonHTMLAttributes, ReactNode } from 'react';
import classNames from '../../utils/className';
import './IconButton.scss';

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon?: ReactNode;
};

export default function IconButton(props: Props) {
  const {
    icon, children, className, ...other
  } = props;

  const fullClassName = classNames('button-primary', 'icon-button', className);

  const iconCol = icon ? (<div className="icon-col">{icon}</div>) : null;

  return (
    <button type="button" className={fullClassName} {...other}>
      <div className="button-row">
        {iconCol}
        <div className="text-col">{children}</div>
      </div>
    </button>
  );
}
