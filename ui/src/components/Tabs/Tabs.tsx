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

import { MouseEvent, ReactElement } from 'react';
import classNames from '../../utils/className';
import './Tabs.scss';

type Props = {
  selectedTabIndex: number;
  headers?: ReactElement[];
  children?: ReactElement[];
  onChangeIndex?: (index: number) => void;
};

export default function Tabs(props: Props) {
  const onClickHeader = (e: MouseEvent<HTMLElement>) => {
    const { index } = e.currentTarget.dataset;
    if (!index) return;

    const tabIndex = Number.parseInt(index, 10);
    if (props.onChangeIndex) {
      props.onChangeIndex(tabIndex);
    }
  };

  function renderHeaders() {
    const list = (props.headers || []).map((item, index) => {
      const className = classNames(
        'tab-header',
        index === props.selectedTabIndex ? 'active' : '',
      );

      const key = `key-${index}`;

      return (
        <div
          key={key}
          className={className}
          data-index={index}
          onClick={onClickHeader}
        >
          {item}
        </div>
      );
    });

    return <div className="tab-headers">{list}</div>;
  }

  const headers = renderHeaders();

  return (
    <div className="tabs">
      {headers}
      <div className="tab-bodys">{props.children}</div>
    </div>
  );
}
