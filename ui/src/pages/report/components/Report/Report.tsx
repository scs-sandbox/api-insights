import { useState } from 'react';
import { Snackbar, Alert } from '@mui/material';
import ComplianceTable, {
  ClickRowEvent,
} from '../ComplianceTable/ComplianceTable';
import ComplianceDialog from '../ComplianceDialog/ComplianceDialog';
import {
  AnalyzerFilter,
  filterIsSelected,
} from '../AnalyzerFilter/AnalyzerFilter';
import IssueSummary from '../../../../components/Issues/IssueSummary/IssueSummary';
import EnvIcon, {
  SpecState,
} from '../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import UploadSpecButton, {
  UploadSpecParam,
} from '../../../../components/Specs/UploadSpecButton/UploadSpecButton';
import SpecSelect from '../../../../components/Specs/SpecDropDown/SpecDropDown';
import PageFrame from '../../../../components/Frame/PageFrame/PageFrame';
import Tabs from '../../../../components/Tabs/Tabs';
import HelpButton from '../../../../components/HelpButton/HelpButton';
import { AnalyserData, TriggerReanalyze } from '../../../../query/analyzer';
import { SpecData } from '../../../../query/spec';
import { ServiceData } from '../../../../query/service';
import { ComplianceData } from '../../../../query/compliance';
import './Report.scss';

type Props = {
  service: ServiceData.Service;
  complianceList: ComplianceData.Compliance[];
  complianceListLoading: boolean;
  analyzerList: AnalyserData.Analyser[];
  selectedSpec: SpecData.Spec;
  onSpecSelected: (spec: SpecData.Spec) => void;
  onSpecUploaded: (data: UploadSpecParam) => void;
  refetchSpecDetail: () => void;
  specUploadingErrorText?: string;
  onClearSpecUploadingError?: () => void;
};

export default function Report(props: Props) {
  const [selectedTabIndex, setSelectedTabIndex] = useState(0);
  const [analyzerFilterData, setAnalyzerFilterData] = useState(null);
  const [openComplianceRow, setOpenComplianceRow] = useState(null);
  const analyzerFailed = props.complianceList?.filter((i) => i.status !== 'Analyzed') || [];
  const btnDisalbed = props.service === undefined;
  const handleMessageClose = () => {
    if (props.onClearSpecUploadingError) {
      props.onClearSpecUploadingError();
    }
  };

  const onOpenComplianceDialog = (e: ClickRowEvent) => {
    setOpenComplianceRow(e.data);
  };

  const onCloseComplianceDialog = () => {
    setOpenComplianceRow(null);
  };

  const onChangeTabIndex = (index: number) => {
    setSelectedTabIndex(index);
  };
  const onReanalyze = async () => {
    await TriggerReanalyze(
      props.service.id,
      props.selectedSpec.id,
      props.analyzerList.map((analyzer) => analyzer.name_id),
    );
    props.refetchSpecDetail();
  };

  const renderComplianceDialog = () => {
    if (!openComplianceRow) return null;

    return (
      <ComplianceDialog
        open
        env={SpecState.Development}
        data={openComplianceRow}
        doc={props.selectedSpec.doc}
        onClose={onCloseComplianceDialog}
      />
    );
  };

  const renderStatistics = (data: ComplianceData.Compliance[]) => {
    if (!data || !data.length) return null;

    return (
      <div className="statistics-block">
        <IssueSummary data={data.map((i) => i.result.summary.stats)} showLabel />
      </div>
    );
  };

  const renderComplianceAnalyzerFilter = (forDrift: boolean) => {
    if (forDrift) return null;

    const analyzerList = props.analyzerList
      ? props.analyzerList
        .filter((i) => i.name_id !== 'drift') // Drift analyzer will be it's own tab on the page
        .map((i) => ({
          title: i.title,
          value: i.name_id,
          status:
              props.complianceList?.filter((x) => x.analyzer === i.name_id)[0]
                ?.status || 'none',
        }))
      : null;

    return (
      <AnalyzerFilter
        analyzerList={analyzerList}
        filterData={analyzerFilterData}
        onChange={setAnalyzerFilterData}
      />
    );
  };

  const renderComplianceAnalyzerBlock = (forDrift: boolean) => {
    const textPart = (
      <div className="text-part">
        <EnvIcon value={forDrift ? SpecState.Release : SpecState.Development} />
        <span className="description">
          {forDrift ? 'Spec drift' : 'Analyzing spec'}
        </span>
        <span className="version-label">{props.selectedSpec?.version}</span>
        <span className="revision-label">{props.selectedSpec?.revision}</span>
        {analyzerFailed.length > 0 ? (
          <span className="failed-warning">Warning: Analyzer Failure</span>
        ) : (
          ''
        )}
      </div>
    );

    const filterPart = (
      <div className="filter-part">
        <span className="text-part">
          {forDrift
            ? 'compared to the live deployed environment'
            : 'for compliance with'}
        </span>
        {renderComplianceAnalyzerFilter(forDrift)}
      </div>
    );

    return (
      <div className="analyzer-block">
        {textPart}
        {filterPart}
      </div>
    );
  };

  const renderComplianceTab = (forDrift: boolean, index: number) => {
    const list = props.complianceList
      ? props.complianceList.filter((i) => (forDrift
        ? i.analyzer === 'drift'
        : filterIsSelected(i.analyzer, analyzerFilterData)
              && i.analyzer !== 'drift'))
      : null;

    const statistics = renderStatistics(list);

    return (
      <div className={`tab-body${selectedTabIndex === index ? ' active' : ''}`}>
        <div className={`compliance-tab${forDrift ? ' for-drift' : ''}`}>
          {renderComplianceAnalyzerBlock(forDrift)}

          {statistics}
          <div className="reanalyze-button" onClick={onReanalyze}>
            <span className="reanalyze-icon" />
            Recalculate
          </div>
          <div className="table-block">
            <ComplianceTable
              key={props.selectedSpec ? props.selectedSpec.id : ''}
              analyzerList={forDrift ? null : props.analyzerList}
              isLoading={props.complianceListLoading}
              specId={props.selectedSpec ? props.selectedSpec.id : ''}
              data={list}
              onClickItem={onOpenComplianceDialog}
            />
          </div>
        </div>
      </div>
    );
  };

  const renderResult = () => {
    const headers = [
      <div key="Compliance" className="report-tab-header">
        <EnvIcon value={SpecState.Development} />
        <div className="title">Compliance</div>
      </div>,
      <div key="Drift" className="report-tab-header">
        <EnvIcon value={SpecState.Release} />
        <div className="title">Drift</div>
      </div>,
    ];

    return (
      <Tabs
        headers={headers}
        selectedTabIndex={selectedTabIndex}
        onChangeIndex={onChangeTabIndex}
      >
        {renderComplianceTab(false, 0)}
        {renderComplianceTab(true, 1)}
      </Tabs>
    );
  };

  const renderErrorMessage = () => {
    if (!props.specUploadingErrorText) return null;

    return (
      <Snackbar
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        open
        autoHideDuration={3000}
        onClose={handleMessageClose}
      >
        <Alert
          onClose={handleMessageClose}
          severity="error"
          sx={{ width: '100%' }}
        >
          {props.specUploadingErrorText}
        </Alert>
      </Snackbar>
    );
  };

  const header = (
    <div className="page-header-content">
      <div className="spec-col">
        <SpecSelect
          serviceId={props.service?.id}
          selectedSpec={props.selectedSpec}
          onChange={props.onSpecSelected}
        />
      </div>
      <div className="upload-col">
        <UploadSpecButton
          disabled={btnDisalbed}
          serviceId={props.service?.id}
          onUploaded={props.onSpecUploaded}
        />
      </div>
    </div>
  );

  const reportResult = renderResult();
  const dialog = renderComplianceDialog();
  const errorMessage = renderErrorMessage();

  return (
    <PageFrame className="report-page" header={header}>
      <HelpButton
        show={!props.selectedSpec}
        title="Reports for compliance and version drift"
        message={
          "Manage your service' compliance speedily with comprehensive priority ratings, grouped issues and easy to navigate line-numbered issues. View the issue in context by clicking on the instance!"
        }
      />
      <div className="page-body-content">{reportResult}</div>
      {dialog}
      {errorMessage}
    </PageFrame>
  );
}
