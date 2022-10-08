import classNames from '../../utils/className';
import './BusyIcon.scss';

export enum BusyIconType {
  Circle = 'circle',
  ArrowCircle = 'arrowcircle',
}

type Props = {
  type?: BusyIconType;
  busy?: boolean;
};

export default function BusyIcon(props: Props) {
  const className = classNames('busy-icon', props.busy ? 'state-busy' : '');

  const renderSvg = () => {
    if (props.type === 'arrowcircle') {
      return (
        <svg className="icon-arrowcircle" viewBox="0 0 101 101">
          <circle
            cx="50"
            cy="50"
            r="35"
            strokeWidth="6"
            strokeDasharray="200 100"
          />
          <path d="M75 10l5 19-20 4" strokeWidth="6" />
        </svg>
      );
    }

    return (
      <svg className="icon-circle" viewBox="0 0 100 100">
        <circle
          cx="50"
          cy="50"
          r="45"
          strokeWidth="10"
          strokeDasharray="200 100"
        />
      </svg>
    );
  };

  return <i className={className}>{renderSvg()}</i>;
}
