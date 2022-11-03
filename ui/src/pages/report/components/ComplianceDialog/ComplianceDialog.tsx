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

import { Dialog } from '@mui/material';
import { MonacoDiffEditor } from 'react-monaco-editor';
import CodeViewer from '../../../../components/CodeViewer/CodeViewer';
import MarkdownViewer from '../../../../components/MarkdownViewer/MarkdownViewer';
import SeverityItem from '../../../../components/Severity/SeverityItem/SeverityItem';
import { ClickRowEventData } from '../ComplianceTable/ComplianceTable';
import { ComplianceData } from '../../../../query/compliance';
import { SpecData } from '../../../../query/spec';
import './ComplianceDialog.scss';

type Props = {
  open: boolean;
  env: SpecData.SpecState;
  doc: string;
  data: ClickRowEventData;
  onClose: () => void;
};

export default function ComplianceDialog(props: Props) {
  if (!props.data) return null;

  const renderRange = () => {
    const row = props.data.row as ComplianceData.ComplianceRangItem;
    const codeSelections = row.range.start.line
      ? [{ start: row.range.start.line, end: row.range.end.line }]
      : undefined;

    return (
      <div className="source-code-block range-block">
        <CodeViewer
          language="json"
          selections={codeSelections}
          value={props.doc}
        />
      </div>
    );
  };

  const renderDiff = () => {
    const row = props.data.row as ComplianceData.ComplianceDiffItem;

    return (
      <div className="source-code-block diff-block">
        <div className="diff-viewer">
          <MonacoDiffEditor
            theme="vs-dark"
            original={row.diff.old}
            value={row.diff.new}
          />
        </div>
      </div>
    );
  };

  const renderDetail = () => {
    if (props.data.row.type === 'range') return renderRange();

    if (props.data.row.type === 'diff') return renderDiff();

    return null;
  };

  const maxWidth = props.data.row.type === 'diff' ? 'xl' : 'md';
  const detail = renderDetail();

  return (
    <Dialog open={props.open} onClose={props.onClose} fullWidth maxWidth={maxWidth}>
      <div className="compliance-dialog">
        <div className="dialog-title">
          <div className="main-part">
            <div className="env">{props.env}</div>
            <SeverityItem severity={props.data.severity} showLabel />
          </div>
          <div className="action-part">
            <div className="close-btn" onClick={props.onClose} />
          </div>
        </div>
        <div className="dialog-body">
          <div className="message-block">
            <MarkdownViewer text={props.data.message} />
          </div>
          <div className="row-block">
            <div className="solution-block">
              <div className="block-title">Solution</div>
              <div className="solution-content">
                <MarkdownViewer text={props.data.mitigation} />
              </div>
            </div>
            {detail}
          </div>
        </div>
      </div>
    </Dialog>
  );
}
