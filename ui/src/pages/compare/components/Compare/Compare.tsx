import { useState } from 'react';
import { Link } from 'react-router-dom';
import dayjs from 'dayjs';
import SpecSelect from '../../../../components/Specs/SpecDropDown/SpecDropDown';
import DiffList from '../DiffList/DiffList';
import CircleScore from '../../../../components/Frame/Service/CircleScore/CircleScore';
import IconButton from '../../../../components/IconButton/IconButton';
import PageFrame from '../../../../components/Frame/PageFrame/PageFrame';
import HelpButton from '../../../../components/HelpButton/HelpButton';
import SeveritySummary from '../../../../components/Severity/SeveritySummary/SeveritySummary';
import MarkdownViewer from '../../../../components/MarkdownViewer/MarkdownViewer';
import DownloadIcon from '../../../../components/DownloadIcon/DownloadIcon';
import { SpecData } from '../../../../query/spec';
import { ServiceData } from '../../../../query/service';
import { ComplianceData } from '../../../../query/compliance';
import { DiffData } from '../../../../query/compare';
import './Compare.scss';

type Props = {
  selectedService: ServiceData.Service;
  leftSpec?: SpecData.Spec;
  rightSpec?: SpecData.Spec;
  leftComplianceList?: ComplianceData.Compliance[];
  rightComplianceList?: ComplianceData.Compliance[];
  compareData: DiffData.JsonDiff;
  compareDataFetching: boolean;
  markDownData: DiffData.MarkdownDiff;
  onSpecSelected: (spec: SpecData.Spec) => void;
  onAltSpecSelected: (spec: SpecData.Spec) => void;
  onCompare: () => void;
  handleDownload: () => void;
};

function Compare(props: Props) {
  /**
   * boolean to control whether compare button is disabled
   * button is disabled when still in the middle of fetching previous compare,
   * or there aren't two specs selected
   */
  const btnDisalbed = !(props.leftSpec && props.rightSpec)
    || props.leftSpec.id === props.rightSpec.id
    || props.compareDataFetching;

  /**
   *  Boolean to toggle between compare result and markdown preview
   *  true by default to show compare result, false shows markdown preview
   */
  const [tab, setTab] = useState(true);

  /**
   *
   * @param spec API spec
   * @param complianceList  list of Compliance details
   * @returns JSX.element
   */
  function renderSpecDetail(
    spec: SpecData.Spec,
    complianceList: ComplianceData.Compliance[],
  ) {
    const time = dayjs(spec.updated_at).format('MMM DD, HH:mm');
    const link = `/reports?service=${spec.service_id}&spec=${spec.id}`;
    return (
      <div className="summary-table">
        <div className="summary-row">
          <div className="summary-name">
            <CircleScore size={38} value={spec.score} darkTrack />
            <div className="summary-name-text">
              <div className="summary-name-text-row">
                <span className="service">{props.selectedService.title}</span>
                <span className="version-label">{spec.version}</span>
                <span className="revision-label">{spec.revision}</span>
              </div>
              <div className="spec-time">{time}</div>
            </div>
          </div>
        </div>
        <div className="summary-row">
          <div className="summary-cell">
            <SeveritySummary data={complianceList.map((i) => i.result.summary.stats)} />
            <Link to={link} className="button-rc">
              View Report
            </Link>
          </div>
        </div>
      </div>
    );
  }

  /**
   * renders info on the two specs being compared
   * @returns JSX.element
   */
  function renderSpecSummaries() {
    const leftSpecDetail = renderSpecDetail(props.leftSpec, props.leftComplianceList);
    const rightSpecDetail = renderSpecDetail(props.rightSpec, props.rightComplianceList);
    return (
      <div>
        <div className="compare-summaries">
          {leftSpecDetail}
          <div className="compare-summaries-divider">to</div>
          {rightSpecDetail}
        </div>
      </div>
    );
  }
  /**
   * Renders the button for trigerring comparison, will show busy
   * if compare data fetching already in progress
   */
  const buttonIcon = (
    <i className={`compare-icon${props.compareDataFetching ? ' busy' : ''}`} />
  );

  const header = (
    <div className="page-header-content">
      <div className="compare-specs">
        <SpecSelect
          serviceId={props.selectedService?.id}
          selectedSpec={props.leftSpec}
          onChange={props.onSpecSelected}
        />
        <div className="compare-to">vs</div>
        <SpecSelect
          serviceId={props.selectedService?.id}
          selectedSpec={props.rightSpec}
          onChange={props.onAltSpecSelected}
        />
      </div>
      <div className="action">
        <IconButton
          icon={buttonIcon}
          onClick={props.onCompare}
          className={`${btnDisalbed ? 'disabled' : ''}`}
          disabled={btnDisalbed}
        >
          Compare Specs
        </IconButton>
      </div>
    </div>
  );

  return (
    <PageFrame className="compare-page" header={header}>
      <HelpButton
        show={!props.compareData}
        title="Compare any spec or snapshot"
        message="Select two versions to compare and get a detailed diff report with expanding issues. Try comparing now."
      />
      <div className="page-body-content">
        {props.compareData?.result?.json && renderSpecSummaries()}
        <div className="tab-row">
          <div
            className={`tab-button ${tab ? 'active-tab' : 'inactive-tab'}`}
            onClick={() => setTab(true)}
          >
            Comparing
          </div>
          <div
            className={`tab-button ${!tab ? 'active-tab' : 'inactive-tab'}`}
            onClick={() => setTab(false)}
          >
            Markdown
          </div>
        </div>

        <div className="result-container">
          {props.compareDataFetching && (
            <div className="loading-indicator">Loading...</div>
          )}
          {tab ? (
            <DiffList data={props.compareData?.result?.json} />
          ) : (
            <div className="markdown-container">
              <div className="button-rc download" onClick={props.handleDownload}>
                <DownloadIcon />
                Download changelog
              </div>
              <MarkdownViewer text={props.markDownData?.result?.markdown} />
            </div>
          )}
        </div>
      </div>
    </PageFrame>
  );
}

export default Compare;
