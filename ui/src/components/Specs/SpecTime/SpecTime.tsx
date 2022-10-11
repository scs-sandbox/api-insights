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

import { ReactNode } from 'react';
import dayjs from 'dayjs';
import './SpecTime.scss';

type Props = {
  icon?: ReactNode;
  beginText?: string;
  time: string;
};

export default function SpecTime(props: Props) {
  if (!props.time) return null;

  const value = dayjs(props.time).format('MMM DD, HH:mm');

  return (
    <span className="spec-time">
      {props.icon && <span className={`icon ${props.icon}`} />}
      {props.beginText && <span className="begin-text">{props.beginText}</span>}
      <span className="time-value">{value}</span>
    </span>
  );
}
