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
import Green from './images/diff-added.png';
import { DiffData } from '../../../../../query/compare';

type Props = {
  data: DiffData.DiffAddedItem;
};

export default function DiffAddedItem(props: Props) {
  const [show, setShow] = useState(true);
  return (
    <div
      className="compare-row compare-row-added"
      onClick={(e) => {
        e.stopPropagation();
        setShow(!show);
      }}
    >
      <div className="row-item row-icon">
        {' '}
        <img className="icon" src={Green} alt="React Logo" />
      </div>
      <div className="row-item row-text">Added: </div>
      <div className="row-item row-code">
        {props.data?.method}
        {' '}
        {props.data?.path}
      </div>
      {props.data.breaking && <div className="row-breaking">Breaking</div>}
    </div>
  );
}
