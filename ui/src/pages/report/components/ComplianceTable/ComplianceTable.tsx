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

import { useEffect, useState, MouseEvent } from 'react';
import MarkdownViewer from '../../../../components/MarkdownViewer/MarkdownViewer';
import Pagination from '../../../../components/Pagination/Pagination';
import SeverityIcon from '../../../../components/Severity/SeverityIcon/SeverityIcon';
import { AnalyserData } from '../../../../query/analyzer';
import { ComplianceData } from '../../../../query/compliance';
import capitalize from '../../../../utils/string';
import { AnalyseTableRowData, convertToTableData, sortRows } from './utils/table';
import './ComplianceTable.scss';

const PAGE_SIZE = 20;

export type ClickRowEventData = AnalyseTableRowData & {
  row: ComplianceData.ComplianceRuleDataItem;
};

export type ClickRowEvent = {
  data: ClickRowEventData;
};

type Props = {
  analyzerList: AnalyserData.Analyser[];
  specId: string;
  isLoading: boolean;
  data: ComplianceData.Compliance[];
  onClickItem: (e: ClickRowEvent) => void;
};

export default function ComplianceTable(props: Props) {
  const [expandedRow, setExpandedRow] = useState('');
  const [sortBy, setSortBy] = useState('severity');
  const [sortDesc, setSortDesc] = useState(true);
  const [pageNum, setPageNum] = useState(0);

  useEffect(() => {
    setPageNum(0);
  }, [props.specId]);

  const data = props.data && Array.isArray(props.data) ? props.data : [];
  const orignalRows = convertToTableData(data);
  const rows = !sortBy
    ? orignalRows
    : sortRows([...orignalRows], sortBy, sortDesc);
  const currentPageStartIndex = pageNum * PAGE_SIZE;
  const currentPageRows = rows.slice(
    currentPageStartIndex,
    currentPageStartIndex + PAGE_SIZE,
  );

  const onSortChange = (event: MouseEvent<HTMLElement>) => {
    const { sort } = event.currentTarget.dataset;
    if (sortBy === sort) {
      setSortDesc(!sortDesc);
    } else {
      setSortBy(sort);
    }
  };

  const onClickDetailItem = (event: MouseEvent<HTMLElement>) => {
    event.stopPropagation();

    const { id, index } = event.currentTarget.dataset;
    const item = currentPageRows.find((i) => i.id === id);

    const clickItemEventData = { ...item, row: item.detail[index] };

    if (props.onClickItem) {
      props.onClickItem({
        data: clickItemEventData,
      });
    }
  };

  const onClickRow = (event: MouseEvent<HTMLElement>) => {
    if ((event.target as HTMLElement).tagName === 'A') return;

    const { id } = event.currentTarget.dataset;
    const item = currentPageRows.find((i) => i.id === id);

    if (!item.detail || !item.detail.length) {
      return;
    }

    if (item.detail.length === 1) {
      const clickItemEventData = { ...item, row: item.detail[0] };

      if (props.onClickItem) {
        props.onClickItem({
          data: clickItemEventData,
        });
      }

      return;
    }

    setExpandedRow(id === expandedRow ? '' : id);
  };

  const onChangePageNum = (index: number) => {
    setPageNum(index - 1);
  };

  const renderLineRange = (complianceRuleData: ComplianceData.ComplianceRuleDataItem) => {
    if (!complianceRuleData) return null;

    const { range } = complianceRuleData as ComplianceData.ComplianceRangItem;

    if (!range) return null;

    const text = range.start.line === range.end.line
      ? `(Line ${range.start.line})`
      : `(Line ${range.start.line} - ${range.end.line})`;

    return <div className="line-range">{text}</div>;
  };

  const renderRowDetailContent = (complianceRuleData: ComplianceData.ComplianceRuleDataItem) => {
    const lineRange = renderLineRange(complianceRuleData);
    if (!complianceRuleData.path || !complianceRuleData.path.length) {
      return <div className="path">{lineRange}</div>;
    }

    if (complianceRuleData.path.length === 1) {
      return (
        <div className="path">
          <span className="url">{complianceRuleData.path[0]}</span>
          {lineRange}
        </div>
      );
    }

    if (complianceRuleData.path.length === 2) {
      const method = complianceRuleData.path[0];
      const url = complianceRuleData.path[1];

      return (
        <div className="path">
          <span className={`method ${method}`}>{method}</span>
          <span className="url">{url}</span>
          {lineRange}
        </div>
      );
    }

    const method = complianceRuleData.path[2].toLowerCase();
    const url = complianceRuleData.path[1];

    return (
      <div className="path">
        <span className={`method ${method}`}>{method}</span>
        <span className="url">{url}</span>
        {lineRange}
      </div>
    );
  };

  const renderRowDetail = (item: AnalyseTableRowData) => {
    if (!item.detail || !item.detail.length) return null;

    if (item.detail.length === 1) {
      return renderLineRange(item.detail[0]);
    }

    const summary = (
      <div className="summary">
        <div className="children">
          {item.detail.length}
          {' '}
          items affected
        </div>
      </div>
    );

    if (item.id !== expandedRow) {
      return <div className="detail">{summary}</div>;
    }

    const detailRows = item.detail.map((i, idx) => {
      const key = `${item.id}-${idx}`;
      const rowDetailConent = renderRowDetailContent(i);
      return (
        <div
          key={key}
          className="item"
          data-id={item.id}
          data-index={idx}
          onClick={onClickDetailItem}
        >
          {rowDetailConent}
        </div>
      );
    });

    return (
      <div className="detail open">
        {summary}
        <div className="list">{detailRows}</div>
      </div>
    );
  };

  const renderAnalyzerCell = (analyzer: string) => {
    if (!props.analyzerList) return null;

    const item = props.analyzerList.find((i) => i.name_id === analyzer) || {
      title: analyzer,
    };
    const analyzerTitle = item.title;

    return (
      <td className="center-cell analyzer-cell">
        <div>{analyzerTitle}</div>
      </td>
    );
  };

  const renderRow = (analyseTableRowData: AnalyseTableRowData) => {
    const severityTitle = capitalize(analyseTableRowData.severity);
    const analyzerCell = renderAnalyzerCell(analyseTableRowData.analyzer);
    const rowDetail = renderRowDetail(analyseTableRowData);

    return (
      <tr key={analyseTableRowData.id} data-id={analyseTableRowData.id} onClick={onClickRow}>
        {analyzerCell}
        <td className="center-cell">
          <SeverityIcon severity={analyseTableRowData.severity} />
          <div>{severityTitle}</div>
        </td>
        <td className="left-cell">
          <div className="violation-block">
            <div className="message">
              <MarkdownViewer text={analyseTableRowData.message} />
            </div>
            {rowDetail}
          </div>
        </td>
        <td className="recomm-cell">
          <MarkdownViewer text={analyseTableRowData.mitigation} />
        </td>
      </tr>
    );
  };

  const renderTableRows = () => {
    const tablRows = currentPageRows.map(renderRow);

    return <tbody>{tablRows}</tbody>;
  };

  const renderLoading = () => {
    if (!props.isLoading) return null;

    return <div className="loading">Loading...</div>;
  };

  const renderNoResult = () => {
    if (props.isLoading) return null;

    if (!props.data) return null;

    if (orignalRows.length) return null;

    return <div className="no-result">No results</div>;
  };

  const renderPagination = (className: string) => {
    if (rows.length <= PAGE_SIZE) return null;

    return (
      <div className={`compliance-pagination ${className}`}>
        <Pagination
          currentPage={pageNum + 1}
          pageSize={PAGE_SIZE}
          total={rows.length}
          onChange={onChangePageNum}
        />
      </div>
    );
  };

  const renderSortableHeader = (title: string, field: string) => {
    const indicator = props.data && !!props.data.length && sortBy === field && (
      <span className={`indicator ${sortDesc ? 'desc' : 'asc'}`} />
    );

    return (
      <div className="sortable-header" data-sort={field} onClick={onSortChange}>
        <span className="title">{title}</span>
        {indicator}
      </div>
    );
  };

  const analyzerHeader = props.analyzerList ? (
    <th className="analyzer-col center-cell">
      {renderSortableHeader('Analyzer', 'analyzer')}
    </th>
  ) : null;

  const topPagination = renderPagination('top');
  const tableRows = renderTableRows();
  const loading = renderLoading();
  const noResult = renderNoResult();
  const buttomPagination = renderPagination('bottom');

  return (
    <div className="compliance-table">
      {topPagination}
      {loading}
      <table className="result-table">
        <thead>
          <tr>
            {analyzerHeader}
            <th className="status-col center-cell">
              {renderSortableHeader('Severity', 'severity')}
            </th>
            <th className="violation-col center-cell">Findings</th>
            <th className="center-cell">Recommendation</th>
          </tr>
        </thead>
        {tableRows}
      </table>
      {noResult}
      {buttomPagination}
    </div>
  );
}
