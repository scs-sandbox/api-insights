import { useQuery } from 'react-query';
import Api from './api';

export namespace ComplianceData {
  // #region Summary
  export type ComplianceIssueRules = {
    [index: string]: number;
  };

  export type ComplianceIssueItem = {
    count: number;
    rules?: ComplianceIssueRules;
  };

  export type ComplianceIssues = {
    count: number;
    warning?: ComplianceIssueItem;
    info?: ComplianceIssueItem;
    error?: ComplianceIssueItem;
    hint?: ComplianceIssueItem;
  };

  export type ComplianceAnalyseSummary = {
    stats?: ComplianceIssues;
  };
  // #endregion

  // #region Findings
  export type CompliancePath = string[];

  export type ComplianceDiff = {
    old: string;
    new: string;
  };

  export type ComplianceDiffItem = {
    type: 'diff';
    path: CompliancePath;
    diff: ComplianceDiff;
  };

  export type ComplianceLoaction = {
    line: number;
    column: number;
  };

  export type ComplianceRange = {
    start: ComplianceLoaction;
    end: ComplianceLoaction;
  };

  export type ComplianceRangItem = {
    type: 'range';
    path: CompliancePath;
    range: ComplianceRange;
  };

  export type ComplianceRuleDataItem = ComplianceDiffItem | ComplianceRangItem;

  export type ComplianceRuleItem = {
    message?: string;
    mitigation?: string;
    data?: ComplianceRuleDataItem[];
  };

  export type ComplianceRules = {
    [index: string]: ComplianceRuleItem;
  };

  export type ComplianceIssue = {
    rules?: ComplianceRules;
  };

  export type ComplianceFindings = {
    unclassified?: ComplianceIssue;
    low?: ComplianceIssue;
    medium?: ComplianceIssue;
    high?: ComplianceIssue;
    critical?: ComplianceIssue;
    blocker?: ComplianceIssue;
    warning?: ComplianceIssueItem;
    info?: ComplianceIssueItem;
    error?: ComplianceIssueItem;
    hint?: ComplianceIssueItem;
  };
  // #endregion

  export type ComplianceAnalyseResult = {
    summary: ComplianceAnalyseSummary;
    findings: ComplianceFindings;
  };

  export type Compliance = {
    id: string;
    analyzer: string;
    result: ComplianceAnalyseResult;
    result_version: number;
    score: number;
    service_id: string;
    spec_id: string;
    status: string;
    created_at: string;
    updated_at: string;
  };
}

export function useFetchSpecCompliance(serviceId: string, specId: string) {
  return useQuery(['spec-compliance', specId], () => {
    if (!specId) return null;
    const url = `/services/${serviceId}/specs/${specId}/analyses`;
    return Api.get(url);
  });
}

export function useFetchServiceCompliance(serviceId: string) {
  return useQuery(['service-compliance', serviceId], () => {
    if (!serviceId) return null;
    const url = `/services/${serviceId}/specs/analyses`;
    return Api.get(url);
  });
}
