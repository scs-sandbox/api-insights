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

import { ChangeEvent } from 'react';
import CircleScore from '../../../../components/Frame/Service/CircleScore/CircleScore';
import ScaleIcon from '../../../../components/Icons/ScaleIcon/ScaleIcon';
import classNames from '../../../../utils/className';
import './AnalyzerFilter.scss';

type AnalyserItem = {
  title: string;
  value: string;
  weight: number;
  status: string;
  score: number;
};

type AnalyserFilterItem = AnalyserItem & {
  selected: boolean;
};

export type AnalyserFilterData = {
  [index: string]: AnalyserFilterItem;
};

type Props = {
  analyzerList?: AnalyserItem[];
  filterData?: AnalyserFilterData;
  onChange: (data: AnalyserFilterData) => void;
};

export function filterIsSelected(
  filter: string,
  filterData?: AnalyserFilterData,
) {
  if (!filterData) return true;

  return filterData[filter] ? filterData[filter].selected : false;
}

function buildAllAnalyzerFilters(
  filterList: AnalyserItem[],
) {
  return filterList || [];
}

function buildNewFilterData(
  filterList?: AnalyserItem[],
  oldFilterData?: AnalyserFilterData,
) {
  if (!filterList) return oldFilterData;

  if (oldFilterData) return JSON.parse(JSON.stringify(oldFilterData));

  const allFilters = buildAllAnalyzerFilters(filterList);

  return allFilters.reduce(
    (pre: AnalyserFilterData, cur: AnalyserItem) => ({
      ...pre,
      [cur.value]: {
        ...cur,
        selected: true,
      },
    }),
    {},
  );
}

export function AnalyzerFilter(props: Props) {
  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const newData = buildNewFilterData(props.analyzerList, props.filterData);
    const itemAll = newData['*'];
    const itemTarget = newData[e.currentTarget.dataset.value || '*'];

    if (itemAll === itemTarget) {
      const newDataValues = Object.values(newData) as AnalyserFilterItem[];
      for (let i = 0; i < newDataValues.length; i += 1) {
        const filterItem = newDataValues[i];
        filterItem.selected = e.currentTarget.checked;
      }
    } else {
      itemTarget.selected = e.currentTarget.checked;
      const noUnselectedItem = !Object.values(newData).find(
        (i) => {
          const filterItem = i as AnalyserFilterItem;
          return (filterItem.value !== '*' && !filterItem.selected);
        },
      );
      if (itemAll) {
        itemAll.selected = noUnselectedItem;
      }
    }

    if (props.onChange) {
      props.onChange(newData);
    }
  };

  if (!props.analyzerList) {
    return <div className="analyzer-filter" />;
  }

  const filterListItems = buildAllAnalyzerFilters(
    props.analyzerList,
  ).map((i) => {
    const checked = filterIsSelected(i.value, props.filterData);
    const failed = i.status !== 'Analyzed';
    const failedText = failed ? (<div className="analyzer-failed">FAILED</div>) : null;
    const analyzerItemClassName = classNames('analyzer-item', failed ? 'status-failed' : '');

    return (
      <li key={i.value} className="filter-item">
        <label className="item-label">
          <input
            className="input-part"
            type="checkbox"
            checked={checked}
            data-value={i.value}
            onChange={onChange}
          />
          <div className={analyzerItemClassName}>
            <div className="analyzer-col-score">
              <CircleScore value={i.score} size={48} />
              {failedText}
            </div>
            <div className="analyzer-col-detail">
              <div className="analyzer-title">{i.title}</div>
              <div className="analyzer-item-weight">
                <ScaleIcon className="analyzer-weight-icon" />
                <span className="analyzer-weight-value">
                  {`${i.weight * 100}%`}
                </span>
              </div>
            </div>
          </div>
        </label>
      </li>
    );
  });

  return (
    <div className="analyzer-filter">
      <ul className="filter-list">{filterListItems}</ul>
    </div>
  );
}
