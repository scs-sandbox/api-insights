/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

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
