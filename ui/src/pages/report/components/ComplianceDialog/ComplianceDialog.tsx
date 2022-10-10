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
      : null;

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
