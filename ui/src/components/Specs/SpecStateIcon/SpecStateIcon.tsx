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

import { SpecData } from '../../../query/spec';
import './SpecStateIcon.scss';

type Props = {
  value?: SpecData.SpecState;
};

export default function EnvIcon(props: Props) {
  const value = (props.value || SpecData.SpecState.Development).toLowerCase();
  const className = `env-icon icon-${value}`;

  return <i className={className} />;
}
