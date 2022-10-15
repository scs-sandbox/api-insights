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
import { MonacoDiffEditor } from 'react-monaco-editor';
import DifferenceIcon from '@mui/icons-material/Difference';
import { DiffData } from '../../../../../query/compare';
import Blue from './images/diff-modified.png';
import MarkdownViewer from '../../../../../components/MarkdownViewer/MarkdownViewer';

type Props = {
  data: DiffData.DiffModifiedItem;
};

export default function DiffModifiedItem(props: Props) {
  const [show, setShow] = useState(true);
  const [showDiff, setShowDiff] = useState(false);

  return (
    <div
      className="compare-row compare-row-modified"
      onClick={(e) => {
        e.stopPropagation();
        setShow(!show);
      }}
    >
      <div className="row-item row-icon">
        {props.data.breaking ? (
          <img src={Blue} className="icon" alt="React Logo" />
        ) : (
          <img className="icon" src={Blue} alt="React Logo" />
        )}
      </div>
      <div className="row-item row-text">Modified: </div>
      <div className="row-item row-code">
        {props.data?.method}
        {' '}
        {props.data?.path}
      </div>
      {props.data.breaking && <div className="row-breaking">Breaking</div>}
      {show && (
        <div>
          <div className="detail">
            <MarkdownViewer text={props.data.message} />
          </div>
          <div className="detail">
            <button
              type="button"
              className={`diff-button ${showDiff ? 'active' : ''}`}
              onClick={(e) => {
                e.stopPropagation();
                setShowDiff(!showDiff);
              }}
            >
              <DifferenceIcon className="compare-icon" />
              Code Diff
            </button>
          </div>
          {showDiff && (
            <div>
              <MonacoDiffEditor
                height="400px"
                width="75vw"
                original={JSON.stringify(props.data.old, null, '\t')}
                value={JSON.stringify(props.data.new, null, '\t')}
                options={{
                  minimap: {
                    enabled: false,
                  },
                  overviewRulerLanes: 0,
                }}
              />
            </div>
          )}
        </div>
      )}
    </div>
  );
}
