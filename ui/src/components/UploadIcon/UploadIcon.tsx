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

import BusyIcon, { BusyIconType } from '../BusyIcon/BusyIcon';
import './UploadIcon.scss';

type Props = {
  busy?: boolean;
  busyType?: BusyIconType;
};

export default function UploadIcon(props: Props) {
  const className = `upload-icon${props.busy ? ' state-busy' : ''}`;
  const busyIcon = props.busy && <BusyIcon type={props.busyType} busy />;

  return <i className={className}>{busyIcon}</i>;
}
