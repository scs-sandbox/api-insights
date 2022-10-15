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
