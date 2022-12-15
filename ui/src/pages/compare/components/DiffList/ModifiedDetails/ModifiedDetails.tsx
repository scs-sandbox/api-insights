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

import { useState, Fragment } from 'react';
import { DiffData } from '../../../../../query/compare';
import MarkdownViewer from '../../../../../components/MarkdownViewer/MarkdownViewer';
import SeverityIcon, { Severity } from '../../../../../components/Severity/SeverityIcon/SeverityIcon';

type Props = {
  data: DiffData.DiffModifiedItem;
};

export default function ModifiedDetails(props: Props) {
  const changeList = [
    props.data.parameters, props.data.requestBody,
    props.data.responses, props.data.security,
  ];
  const renderDetail = (data: DiffData.DetailItem) => {
    const isBreaking = data.breaking;
    let { message } = data;
    if (data.details) {
      data.details.forEach((detail) => {
        message = message.replaceAll(detail.message, '');
      });
    }
    return (
      <div>
        {isBreaking && <div className="breaking-container"><SeverityIcon severity="breaking" /></div>}
        <MarkdownViewer text={message} />
        <div className="markdown-children">
          {data.details && data.details.map((detail: DiffData.DetailItem) => (
            <Fragment key={detail.message}>
              {renderDetail(detail)}
            </Fragment>
          ))}
        </div>
      </div>
    );
  };
  const changes = changeList.map((change) => {
    if (change) {
      let { message } = change;
      if (change.details) {
        change.details.forEach((detail) => {
          message = message.replaceAll(detail.message, '');
        });
      }
      return (
        <div>
          <MarkdownViewer text={message} />
          {change.details && change.details.map((detail: DiffData.DetailItem) => (
            <Fragment key={detail.message}>
              {renderDetail(detail)}
            </Fragment>
          ))}
        </div>
      );
    }
    return null;
  });
  return (
    <div>
      {changes}
    </div>
  );
}
