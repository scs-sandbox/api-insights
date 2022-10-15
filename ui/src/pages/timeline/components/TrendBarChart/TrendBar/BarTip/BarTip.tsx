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
