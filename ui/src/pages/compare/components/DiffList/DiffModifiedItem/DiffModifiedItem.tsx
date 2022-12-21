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

import { useState, useLayoutEffect, useEffect } from 'react';
import { MonacoDiffEditor } from 'react-monaco-editor';
import DifferenceIcon from '@mui/icons-material/Difference';
import { editor } from 'monaco-editor';
import { DiffData } from '../../../../../query/compare';
import { SpecData } from '../../../../../query/spec';
import Blue from './images/diff-modified.png';
import MarkdownViewer from '../../../../../components/MarkdownViewer/MarkdownViewer';
import waitFor from '../../../../../utils/waitFor';
import iterateObject from '../../../../../utils/iterateObject';
import { monacoMount } from '../../../../../utils/monacoInjection';
import DiffReference from '../DiffReference/DiffReference';
import ModifiedDetails from '../ModifiedDetails/ModifiedDetails';
import './DiffModifiedItem.scss';

type Props = {
  data: DiffData.DiffModifiedItem;
  leftSpec?: SpecData.Spec;
  rightSpec?: SpecData.Spec;
};

export default function DiffModifiedItem(props: Props) {
  const [show, setShow] = useState(true);
  const [refs, setRefs] = useState<string[]>([]);
  const [activeRefs, setActiveRefs] = useState<string[]>([]);
  const [showDiff, setShowDiff] = useState(false);
  const dropdownClass = (show) ? 'up-icon' : 'drop-icon';
  const removeActive = (ref: string) => {
    setActiveRefs(activeRefs.filter((item) => item !== ref));
  };
  const referenceButtons = activeRefs.map((ref, index) => {
    const key = `ref-${index}`;
    return (
      <DiffReference
        key={key}
        data={ref}
        leftSpec={props.leftSpec}
        rightSpec={props.rightSpec}
        removeActive={removeActive}
        setActiveRefs={setActiveRefs}
      />
    );
  });

  // const iterate = (obj: any, resultArray: any[]) => {
  //   Object.keys(obj).forEach((key) => {
  //     if (key === '$ref') {
  //       resultArray.push(obj[key]);
  //     }
  //     if (typeof obj[key] === 'object' && obj[key] !== null) {
  //       iterate(obj[key], resultArray);
  //     }
  //   });
  // };
  // const injectButtonToEditor = () => {
  //   Array.from(document.getElementsByClassName('mtk1')).forEach((element) => {
  //     refs.forEach((ref) => {
  //       if (element.innerHTML.includes(ref) && !(element.innerHTML.includes('injected'))) {
  //         // eslint-disable-next-line no-param-reassign
  //         const node = document.createElement('div');
  //         node.innerHTML = '<div class="reference-icon"></div>View Reference';
  //         node.className = 'injected';
  //         element.append(node);
  //         node.addEventListener('click', () => {
  //           setActiveRefs((actives) => {
  //             if (actives.includes(ref)) {
  //               // console.log('del');
  //               // console.log(actives.filter((item) => item !== ref));
  //               // return (actives.filter((item) => item !== ref));
  //               return actives;
  //             }
  //             console.log('adding');
  //             console.log([...actives, ref]);
  //             return ([...actives, ref]);
  //           });
  //         });
  //       }
  //     });
  //   });
  // };
  useEffect(() => {
    iterateObject(props.data.old, refs);
    iterateObject(props.data.new, refs);
  }, []);
  return (
    <div
      className="compare-row compare-row-modified"
      onClick={(e) => {
        e.stopPropagation();
        setShow(!show);
      }}
    >
      <div className="row-container">
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
        {props.data.breaking && <div className="row-item row-breaking">Breaking</div>}
        <div className={dropdownClass} />
      </div>
      {show && (
        <div onClick={(e) => { e.stopPropagation(); }}>
          <div className="detail">
            <ModifiedDetails data={props.data} />
            {/* <MarkdownViewer text={props.data.message} /> */}
          </div>
          <div className="">
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
              {showDiff && (
                <div className="up-icon" />
              )}
            </button>
          </div>
          {showDiff && (
            <div>
              <MonacoDiffEditor
                height="400px"
                width="100%"
                theme="vs-dark"
                original={JSON.stringify(props.data.old, null, '\t')}
                value={JSON.stringify(props.data.new, null, '\t')}
                options={{
                  minimap: {
                    enabled: false,
                  },
                  overviewRulerLanes: 0,
                }}
                editorDidMount={(diffEditor) => {
                  if (!refs) return;
                  monacoMount(diffEditor, refs, setActiveRefs);
                }}
              />
              {referenceButtons}
            </div>
          )}
        </div>
      )}
    </div>
  );
}
