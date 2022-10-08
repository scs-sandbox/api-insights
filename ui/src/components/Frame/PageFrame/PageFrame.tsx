import { HTMLAttributes, ReactNode } from 'react';
import classNames from '../../../utils/className';
import PageHeader from '../PageHeader/PageHeader';
import PageNavBar from '../PageNavBar/PageNavBar';
import './PageFrame.scss';

type Props = HTMLAttributes<HTMLElement> & {
  header?: ReactNode;
};

export default function PageFrame(props: Props) {
  const className = classNames('app-page', props.className);

  return (
    <div className={className}>
      <PageHeader>{props.header}</PageHeader>
      <div className="page-body">
        <PageNavBar />
        <div className="page-content">
          {props.children}
        </div>
      </div>
    </div>
  );
}
