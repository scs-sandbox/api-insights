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

import { ContactSupportOutlined } from '@mui/icons-material';
import { useQuery } from 'react-query';
import Api from './api';

export namespace ComplianceData {
  // #region Summary
  export type ComplianceSeveritySummaryItem = {
    count: number;
  };

  export type ComplianceSeveritySummary = {
    count: number;
    warning?: ComplianceSeveritySummaryItem;
    info?: ComplianceSeveritySummaryItem;
    error?: ComplianceSeveritySummaryItem;
    hint?: ComplianceSeveritySummaryItem;
  };

  export type ComplianceAnalyseSummary = {
    stats?: ComplianceSeveritySummary;
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

  export type ComplianceFindingItem = {
    rules?: ComplianceRules;
  };

  export type ComplianceFindings = {
    warning?: ComplianceFindingItem;
    info?: ComplianceFindingItem;
    error?: ComplianceFindingItem;
    hint?: ComplianceFindingItem;
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
