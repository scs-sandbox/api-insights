import { HTMLAttributes } from 'react';
import classNames from '../../../../utils/className';
import './ScoreLevel.scss';

type Props = HTMLAttributes<HTMLElement> & {
  score: number;
};

export function getScoreLevels() {
  return [
    {
      max: 49,
      title: '1-49 Alert',
      className: 'score-level level-alert',
    },
    {
      max: 59,
      title: '50-59 Warning',
      className: 'score-level level-warning',
    },
    {
      max: 69,
      title: '60-69 At Risk',
      className: 'score-level level-atrisk',
    },
    {
      max: 79,
      title: '70-79 Good',
      className: 'score-level level-good',
    },
    {
      max: 89,
      title: '80-89 Very Good',
      className: 'score-level level-verygood',
    },
    {
      max: 100,
      title: '> 90 Excellent',
      className: 'score-level level-excellent',
    },
  ];
}

export function calcScoreLevel(score: number) {
  const levels = getScoreLevels();

  const item = levels.find((i) => score <= i.max);
  if (item) return item;

  throw new Error('score is above 100');
}

export default function ScoreLevel(props: Props) {
  const { className, score, ...other } = props;
  const fullClassName = classNames(
    calcScoreLevel(score || 0).className,
    className,
  );

  return <div {...other} className={fullClassName} />;
}
