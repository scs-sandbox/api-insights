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

import { useState, MouseEvent } from 'react';
import Bar, { ChartDataItem } from './Bar/Bar';
import BarTip, { Mouse } from './BarTip/BarTip';

type Props = ChartDataItem & {
  onMouseEnter?: (e: MouseEvent<HTMLElement>) => void;
  onMouseLeave?: (e: MouseEvent<HTMLElement>) => void;
};

export default function TipTrendBar(props: Props) {
  const [mouseMoveTimer, setMouseMoveTimer] = useState(0);
  const [mouse, setMouse] = useState<Mouse>();

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
    setMouse(undefined);
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
