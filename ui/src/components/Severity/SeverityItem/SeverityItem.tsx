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

import capitalize from '../../../utils/string';
import SeverityIcon, { Severity } from '../SeverityIcon/SeverityIcon';
import './SeverityItem.scss';

type Props = {
  severity: Severity;
  count?: number;
  showLabel?: boolean;
  label?: string;
};

export default function SeverityItem(props: Props) {
  const labelText = props.label || capitalize(props.severity);

  function renderCount() {
    const hasCount = props.count !== undefined && props.count !== null;

    if (!hasCount) return null;

    return <span className="severity-item-count">{props.count}</span>;
  }

  function renderLabel() {
    if (!props.showLabel) return null;

    return <span className="severity-item-label">{labelText}</span>;
  }

  const count = renderCount();
  const label = renderLabel();
  const className = `severity-item severity-item-${props.severity}`;

  return (
    <span className={className}>
      <SeverityIcon severity={props.severity} title={labelText} />
      {count}
      {label}
    </span>
  );
}
