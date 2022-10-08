import { CSSProperties, ReactElement } from 'react';
import { Link } from 'react-router-dom';
import { calcScoreLevel } from '../../../../../../components/Frame/Service/ScoreLevel/ScoreLevel';
import classNames from '../../../../../../utils/className';
import './Bar.scss';

export type ChartDataItem = {
  score: number;
  label: string;
};

type Props = ChartDataItem & {
  'data-id'?: string;
  href?: string;
  to?: string;
  target?: string;
  highlight?: boolean;
  style?: CSSProperties;
  children: ReactElement;
};

export default function Bar(props: Props) {
  const height = props.score ? `${props.score}%` : '1px';
  const style = {
    ...props.style,
    height,
  };

  const { highlight, ...otherProps } = props;

  const className = classNames(
    'trend-bar',
    highlight ? 'highlight' : '',
    highlight ? calcScoreLevel(props.score).className : '',
  );

  const commonProps = {
    ...otherProps,
    className,
    style,
  };

  const axis = <div className="bar-axis">{props.label}</div>;

  if (props.href) {
    return (
      <a {...commonProps} href={props.href} target={props.target}>
        {axis}
        {props.children}
      </a>
    );
  }

  if (props.to) {
    return (
      <Link {...commonProps} to={props.to} target={props.target}>
        {axis}
        {props.children}
      </Link>
    );
  }

  return (
    <div {...commonProps}>
      {axis}
      {props.children}
    </div>
  );
}
