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

import VersionReport, { VersionData } from '../VersionReport/VersionReport';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import SpecStateIcon from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import { ServiceData } from '../../../../query/service';
import { SpecData } from '../../../../query/spec';
import './VersionReportList.scss';

type Props = {
  service: ServiceData.Service;
  versionList?: VersionData[];
};

export default function VersionReportList(props: Props) {
  const className = 'spec-version-report-list';

  const renderVersion = (data: VersionData) => (
    <li
      key={data.version}
      data-version={data.version}
      className="version-block-item"
    >
      <div className="item-line">
        <SpecStateIcon value={data.latestRevision.state as SpecData.SpecState} />
      </div>
      <div className="item-content">
        <div className="item-content-time">
          <SpecTime icon="dark" time={data.updated_at} />
        </div>
        <VersionReport service={props.service} data={data} />
      </div>
    </li>
  );

  const versionList = (props.versionList || []).map(renderVersion);

  return <div className={className}>{versionList}</div>;
}
