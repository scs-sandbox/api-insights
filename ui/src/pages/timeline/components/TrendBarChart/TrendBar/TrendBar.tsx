import { useState, MouseEvent } from 'react';
import Bar, { ChartDataItem } from './Bar/Bar';
import BarTip from './BarTip/BarTip';

type Props = ChartDataItem & {
  onMouseEnter?: (e: MouseEvent<HTMLElement>) => void;
  onMouseLeave?: (e: MouseEvent<HTMLElement>) => void;
};

export default function TipTrendBar(props: Props) {
  const [mouseMoveTimer, setMouseMoveTimer] = useState(0);
  const [mouse, setMouse] = useState(null);

  const onMouseMove = (event: MouseEvent<HTMLElement>) => {
    clearTimeout(mouseMoveTimer);

    const timer = setTimeout(() => {
      setMouse({
        x: event.clientX,
        y: event.clientY,
      });
    }, 50);

    setMouseMoveTimer(timer as unknown as number);
  };

  const onMouseLeave = (event: MouseEvent<HTMLElement>) => {
    clearTimeout(mouseMoveTimer);
    setMouse(null);
    setMouseMoveTimer(0);
    if (props.onMouseLeave) {
      props.onMouseLeave(event);
    }
  };

  const newProps = { ...props, onMouseMove, onMouseLeave };

  return (
    <Bar {...newProps}>
      <BarTip mouse={mouse} score={props.score} label={props.label} />
    </Bar>
  );
}
