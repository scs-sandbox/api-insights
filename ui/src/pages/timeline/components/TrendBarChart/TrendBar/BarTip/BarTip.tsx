import PopLayer from '../../../../../../components/PopLayer/PopLayer';
import { ChartDataItem } from '../Bar/Bar';
import './BarTip.scss';

type Props = ChartDataItem & {
  mouse: {
    x: number;
    y: number;
  };
};

export default function BarTip(props: Props) {
  if (!props.mouse) return null;

  const style = {
    transform: `translate(${props.mouse.x}px, ${props.mouse.y}px)`,
  };

  const viewWidth = window.innerWidth || document.body.clientWidth;
  const inViewRightPart = props.mouse.x > viewWidth / 2;
  const className = inViewRightPart ? 'in-right' : 'in-left';

  return (
    <PopLayer>
      <div className={`trend-bar-tip ${className}`} style={style}>
        <div className="tip-body">
          <div className="label">{props.label}</div>
          <div className="score">
            Score:
            {props.score}
          </div>
        </div>
      </div>
    </PopLayer>
  );
}
