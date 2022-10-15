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

import DownloadIcon from '../../../../components/DownloadIcon/DownloadIcon';
import SpecStateIcon from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import { SpecData } from '../../../../query/spec';
import './Snapshot.scss';

type Props = {
  data: SpecData.Snapshot;
};

export default function Snapshot(props: Props) {
  return (
    <div className="snapshot-item block-item-light">
      <SpecStateIcon value={SpecData.SpecState.Release} />
      <div className="snapshot-item-info">
        <div className="detail">
          <SpecTime time={props.data.updated_at} />
        </div>
        <div className="snapshot-name">{props.data.name}</div>
      </div>
      <div className="spec-download">
        <DownloadIcon />
      </div>
    </div>
  );
}
