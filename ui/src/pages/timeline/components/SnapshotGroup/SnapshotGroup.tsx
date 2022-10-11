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
import SpecTime from '../../../../components/Specs/SpecTime/SpecTime';
import buildSortedListByStringField from '../../../../utils/order';
import Snapshot from '../Snapshot/Snapshot';
import { SpecData } from '../../../../query/spec';
import './SnapshotGroup.scss';

type Props = {
  data: SpecData.Snapshot[];
  serviceTitle: string;
};

export default function SnapshotGroup(props: Props) {
  const [openDetail, setOpenDetail] = useState(false);
  const [showAllRevisions, setShowAllRevisions] = useState(false);

  const descOrderedList = buildSortedListByStringField(
    props.data || [],
    'updated_at',
    true,
  ) as SpecData.Snapshot[];
  const latestSnapshot = descOrderedList[0];

  const onViewMore = () => {
    setShowAllRevisions(!showAllRevisions);
  };

  const renderSnapshot = (i: SpecData.Snapshot) => (
    <div className="snapshot-col" key={i.id}>
      <Snapshot data={i} />
    </div>
  );

  const renderSnapshots = () => {
    const DEFAULT_SHOW_COUNT = 4;
    const moreCount = descOrderedList.length - DEFAULT_SHOW_COUNT;
    const showCount = Math.min(DEFAULT_SHOW_COUNT, descOrderedList.length);
    const list = showAllRevisions
      ? descOrderedList
      : descOrderedList.slice(0, showCount);
    const viewMoreText = showAllRevisions
      ? 'View Less'
      : `View ${moreCount} More`;

    const viewMore = moreCount > 0 && (
      <div className="view-more-block">
        <div className="view-more" onClick={onViewMore}>
          {viewMoreText}
        </div>
      </div>
    );

    return (
      <div className="snapshots-block">
        <div className="snapshot-list">{list.map(renderSnapshot)}</div>
        {viewMore}
      </div>
    );
  };

  const snapshots = renderSnapshots();

  return (
    <div className="snapshop-group">
      <div className="group-header">
        <div className="item-content-time">
          <SpecTime icon="dark" time={latestSnapshot.updated_at} />
        </div>
      </div>
      <div className="panel-block group-body">
        <div className="group-brief" onClick={() => setOpenDetail(!openDetail)}>
          <div className="info-col">
            <span className="service">{props.serviceTitle}</span>
            <span className="bold">Reconstructed</span>
          </div>
          <div className="action-col" />
        </div>
        <div className={`group-detail${openDetail ? ' open' : ''}`}>
          <div className="snapshot-list">{snapshots}</div>
        </div>
      </div>
    </div>
  );
}
