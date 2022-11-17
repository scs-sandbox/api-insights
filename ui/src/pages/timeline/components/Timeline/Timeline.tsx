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
import VersionReportList from '../VersionReportList/VersionReportList';
import UploadSpecButton, {
  UploadSpecParam,
} from '../../../../components/Specs/UploadSpecButton/UploadSpecButton';
import PageFrame from '../../../../components/Frame/PageFrame/PageFrame';
import Toggle from '../../../../components/Toggle/Toggle';
import ReconstructSnapshotButton, {
  ReconstructedEvent,
} from '../../../../components/Specs/ReconstructSnapshotButton/ReconstructSnapshotButton';
import HelpButton from '../../../../components/HelpButton/HelpButton';
import {
  buildVersionSortedList,
  filterNotEmptyVersionList,
} from './utils/VersionList';
import SnapshotGroup from '../SnapshotGroup/SnapshotGroup';
import { SpecData } from '../../../../query/spec';
import { ServiceData } from '../../../../query/service';
import { ComplianceData } from '../../../../query/compliance';
import './Timeline.scss';

type Props = {
  selectedService: ServiceData.Service;
  specList: SpecData.Spec[];
  complianceList: ComplianceData.Compliance[];
  isSpecDataLoading: boolean;
  isComplianceListLoading: boolean;
  snapshots: SpecData.Snapshot[];
  isSnapshotLoading: boolean;
  onSpecUploaded: (data: UploadSpecParam) => void;
  onReconstructed: (data: ReconstructedEvent) => void;
};

export default function Timeline(props: Props) {
  const [archivedHidden, setArchivedHidden] = useState(true);
  const versionList = buildVersionSortedList(
    props.specList,
    props.complianceList,
  );
  const filteredValidVersionList = archivedHidden
    ? filterNotEmptyVersionList(versionList)
    : versionList;
  const btnDisalbed = props.selectedService === undefined;

  const header = (
    <div className="page-header-content">
      <UploadSpecButton
        disabled={btnDisalbed}
        serviceId={props.selectedService?.id}
        onUploaded={props.onSpecUploaded}
      />
    </div>
  );

  const renderVersionList = () => {
    if (props.isSpecDataLoading || props.isComplianceListLoading) {
      return <div className="data-loading">Loading...</div>;
    }

    if (!filteredValidVersionList || !filteredValidVersionList.length) {
      return (
        <div className="no-result">
          <UploadSpecButton
            className="full-button blue-button"
            disabled={btnDisalbed}
            serviceId={props.selectedService?.id}
            onUploaded={props.onSpecUploaded}
          />
        </div>
      );
    }

    return (
      <VersionReportList
        service={props.selectedService}
        versionList={filteredValidVersionList}
      />
    );
  };

  const renderSnapshots = () => {
    if (!props.snapshots.length) {
      return (
        <div className="no-result">
          <ReconstructSnapshotButton
            className="full-button blue-button"
            disabled={props.isSpecDataLoading}
            serviceId={props.selectedService?.id}
            onReconstructed={props.onReconstructed}
          >
            Take Snapshot
          </ReconstructSnapshotButton>
        </div>
      );
    }

    return (
      <SnapshotGroup
        serviceTitle={props.selectedService.title}
        data={props.snapshots}
      />
    );
  };

  const renderReconstructedBlock = () => {
    if (props.isSnapshotLoading) return null;

    const buttonText = props.snapshots.length
      ? 'Reconstruct Snapshot'
      : 'Take Snapshot';
    const snapshots = renderSnapshots();

    return (
      <div className="resconstructed-block">
        <div className="group-title-bar reconstructed-bar">
          <div className="title">Reconstructed</div>
          <div className="action">
            <ReconstructSnapshotButton
              className="blue-button"
              disabled={props.isSpecDataLoading}
              serviceId={props.selectedService?.id}
              onReconstructed={props.onReconstructed}
            >
              {buttonText}
            </ReconstructSnapshotButton>
          </div>
        </div>
        {snapshots}
      </div>
    );
  };

  const renderHelp = () => {
    if (props.isSpecDataLoading) return null;

    return (
      <HelpButton
        show={!props.specList?.length}
        title="History of your versions, revision by revision"
        message={
          "Your revised specs group together by version, keeping your entire spec development history and offering reports of each. Revision history chart allows you to monitor for sudden changes and download each spec you've previously uploaded."
        }
      />
    );
  };

  const help = renderHelp();
  const versions = renderVersionList();
  const reconstructedBlock = renderReconstructedBlock();

  return (
    <PageFrame className="timeline-page" header={header}>
      {help}
      <div className="page-body-content">
        <div className="page-title-bar">
          <div className="title">Timeline</div>
          <div className="action">
            <Toggle
              label="Archived Hidden"
              checked={archivedHidden}
              onToggle={setArchivedHidden}
            />
          </div>
        </div>
        {versions}
        {reconstructedBlock}
      </div>
    </PageFrame>
  );
}
