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

import './SeverityIcon.scss';

export type KnownSeverity = 'error' | 'warning' | 'hint' | 'info';

export type Severity = KnownSeverity | string;

type Props = {
  severity: Severity;
  title?: string;
};

export default function SeverityIcon(props: Props) {
  const className = `severity-icon severity-icon-${props.severity}`;

  const allProps = {
    ...props,
    className,
  };

  return <span {...allProps} />;
}
