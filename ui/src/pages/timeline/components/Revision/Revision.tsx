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

import { useState } from 'react';
import { Link } from 'react-router-dom';
import DownloadIcon from '../../../../components/DownloadIcon/DownloadIcon';
import ScoreLevel from '../../../../components/Frame/Service/ScoreLevel/ScoreLevel';
import SpecStateIcon from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import { buildApiAbsoluteUrl } from '../../../../query/api';
import { SpecData, fetchSpecDetail } from '../../../../query/spec';
import { ServiceData } from '../../../../query/service';
import { ComplianceData } from '../../../../query/compliance';
import classNames from '../../../../utils/className';
import handleDownload from '../../../../utils/handleDownload';
import './Revision.scss';

export type RevisionData = SpecData.Spec & {
  complianceList: ComplianceData.Compliance[];
};

type Props = {
  data: RevisionData;
  linkTo: string;
  className?: string;
  service: ServiceData.Service;
  onMouseEnter?: (item: RevisionData) => void;
  onMouseLeave?: (item: RevisionData) => void;
  onReleased?: (item: RevisionData) => void;
  onArchived?: (item: RevisionData) => void;
};

/**
 * render revision according to props.data
 * @param props props
 * @returns JSX.element
 */
export default function Revision(props: Props) {
  const [openMenu, setOpenMenu] = useState(false);
  const { data, service } = props;
  const className = classNames(
    'revision-item block-item-light',
    props.className,
  );

  const onMouseEnter = () => {
    if (props.onMouseEnter) {
      props.onMouseEnter(props.data);
    }
  };

  const onMouseLeave = () => {
    if (props.onMouseLeave) {
      props.onMouseLeave(props.data);
    }
  };

  const onOpenMenu = () => {
    setOpenMenu(true);
  };

  const onCloseMenu = () => {
    setOpenMenu(false);
  };

  const onRelease = () => {
    if (props.onReleased) {
      // props.onReleased(null);
    }
  };

  const onArchive = () => {
    if (props.onArchived) {
      // props.onArchived(null);
    }
  };

  const handleSpecDownload = () => {
    const fileName = `${service.name_id}-${data.version}-${data.revision}.json`;
    fetchSpecDetail(service.id, data.id)?.then((result: any) => {
      handleDownload(fileName, result.doc);
    });
  };

  const renderScore = () => {
    if (data.score === null || data.score === undefined) {
      return <div className="loader" />;
    }

    return (
      <ScoreLevel className="score block-item" score={data.score}>
        {data.score}
      </ScoreLevel>
    );
  };

  const renderMenu = () => {
    if (!openMenu) return null;

    return (
      <div className="action-menu" onClick={onCloseMenu}>
        <div className="menu-item" onClick={onRelease}>
          <SpecStateIcon value={SpecData.SpecState.Release} />
          <div className="menu-item-label">Live</div>
        </div>
        <div className="menu-item" onClick={onArchive}>
          <SpecStateIcon value={SpecData.SpecState.Archive} />
          <div className="menu-item-label">Archive</div>
        </div>
      </div>
    );
  };

  const score = renderScore();
  const menu = renderMenu();

  return (
    <div
      data-id={data.id}
      className={className}
      onMouseEnter={onMouseEnter}
      onMouseLeave={onMouseLeave}
    >
      <SpecStateIcon value={data.state as SpecData.SpecState} />
      <div className="revision-item-info">
        <div className="detail">
          <SpecTime time={data.updated_at} />
          {score}
        </div>
        <span className="revision-label">{data.revision}</span>
      </div>
      <div
        className="revision-item-menu"
        onMouseEnter={onOpenMenu}
        onMouseLeave={onCloseMenu}
      >
        <div className="action-button" />
        {menu}
      </div>
      <div className="revision-item-action">
        <Link to={props.linkTo} className="view-report button-rc">
          View Full Report
        </Link>
      </div>
      <div className="spec-download" onClick={handleSpecDownload}>
        <DownloadIcon />
      </div>
    </div>
  );
}
