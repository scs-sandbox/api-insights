import { HTMLAttributes } from 'react';
import { CircularProgress } from '@mui/material';
import ScoreLevel from '../ScoreLevel/ScoreLevel';
import './CircleScore.scss';

type Props = HTMLAttributes<HTMLElement> & {
  size?: number;
  thickness?: number;
  value?: number;
  darkTrack?: boolean;
};

export default function CircleScore(props: Props) {
  const size = props.size || 56;
  const thickness = props.thickness || 4;
  const value = props.value || 0;

  const style = {
    width: `${size}px`,
    height: `${size}px`,
  };

  return (
    <div className="circle-score" style={style}>
      <ScoreLevel score={value} className="circle-part">
        <div className={`score-track${props.darkTrack ? ' dark-track' : ''}`}>
          <CircularProgress
            color="inherit"
            variant="determinate"
            size={size}
            thickness={thickness}
            value={100}
          />
        </div>
        <div className="score-value">
          <CircularProgress
            color="inherit"
            variant="determinate"
            size={size}
            thickness={thickness}
            value={value}
          />
        </div>
      </ScoreLevel>
      <div className="value-label">{props.children || value}</div>
    </div>
  );
}
