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

import DiffModifiedItem from './DiffModifiedItem/DiffModifiedItem';
import DiffAddedItem from './DiffAddItem/DiffAddItem';
import DiffDeletedItem from './DiffDeletedItem/DiffDeletedItem';
import { DiffData } from '../../../../query/compare';
import { SpecData } from '../../../../query/spec';
import './DiffList.scss';

type Props = {
  data: DiffData.JsonDiffResult;
  leftSpec?: SpecData.Spec;
  rightSpec?: SpecData.Spec;
};

export default function DiffList(props: Props) {
  if (!props.data) {
    return (
      <div className="result-table">
        Select a service and two specs to compare
      </div>
    );
  }

  const addedRows = (props.data?.added || [])
    .map((addedItem, index) => {
      const key = `added-${index}`;
      return (
        <DiffAddedItem data={addedItem} key={key} />
      );
    });

  const modifiedRows = (props.data?.modified || [])
    .map((modifiedItem, index) => {
      const key = `modified-${index}`;
      return (
        <DiffModifiedItem
          data={modifiedItem}
          leftSpec={props.leftSpec}
          rightSpec={props.rightSpec}
          key={key}
        />
      );
    });

  const deletedRows = (props.data?.deleted || [])
    .map((deletedItem, index) => {
      const key = `deleted-${index}`;
      return (
        <DiffDeletedItem data={deletedItem} key={key} />
      );
    });

  const rows = [
    ...addedRows,
    ...modifiedRows,
    ...deletedRows,
  ];
  if (rows.length === 0) {
    return (
      <div className="result-table">
        Identical Specs
      </div>
    );
  }
  return (
    <div className="result-table">
      {rows}
    </div>
  );
}
