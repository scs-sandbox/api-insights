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
