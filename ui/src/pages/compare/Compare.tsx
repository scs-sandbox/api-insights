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

import { useOutletContext, useSearchParams } from 'react-router-dom';
import Compare from './components/Compare/Compare';
import { SpecData, useFetchSpecList } from '../../query/spec';
import { useFetchCompare, useFetchMarkdown } from '../../query/compare';
import { useFetchSpecCompliance } from '../../query/compliance';
import { AppFrameContext } from '../../components/Frame/AppFrame/AppFrame';
import { buildApiAbsoluteUrl } from '../../query/api';

/**
 * This file sets up the data fetching functions for the compare page
 *
 */
export default function ComparePage() {
  const { selectedService } = useOutletContext() as AppFrameContext;
  const { data: specData } = useFetchSpecList(selectedService?.id);
  const [searchParams, setSearchParams] = useSearchParams();
  const searchParamSpec = searchParams.get('spec');
  const searchParamSpec2 = searchParams.get('spec2');
  const leftSpec = (specData || []).find((x) => x.id === searchParamSpec);
  const rightSpec = (specData || []).find((x) => x.id === searchParamSpec2);
  const {
    refetch: fetchCompare,
    data: compareData,
    isFetching: compareDataFetching,
  } = useFetchCompare(selectedService?.id, leftSpec?.id, rightSpec?.id);
  const {
    refetch: fetchMarkdown,
    data: markDownData,
  } = useFetchMarkdown(selectedService?.id, leftSpec?.id, rightSpec?.id);
  const { data: leftComplianceList } = useFetchSpecCompliance(
    selectedService?.id,
    leftSpec?.id,
  );
  const { data: rightComplianceList } = useFetchSpecCompliance(
    selectedService?.id,
    rightSpec?.id,
  );

  /**
   * Updating paremeters in the url that are used to keep track of
   * the latest selected serviceId and specIds on the page.
   * Loads existing parameters and overrides them with wahtever is updated.
   * @param {object} changingParam object that contains any possible url
   * paremeter updates, for example:
   * { service: string ,spec: 111, spec2: 222  }
   */
  const changeSearchParams = (changingParam: object) => {
    const paramSpec = searchParamSpec ? { spec: searchParamSpec } : {};
    const paramSpec2 = searchParamSpec2 ? { spec2: searchParamSpec2 } : {};
    const params = {
      service: searchParams.get('service'),
      ...paramSpec,
      ...paramSpec2,
      ...changingParam, // inputing the updated params last to override any existing parameters
    };

    setSearchParams(params);
  };

  /**
   * one of the two specs to be compared
   * @param e {SpecData.Spec}
   */
  const onSpecSelected = (e: SpecData.Spec) => {
    changeSearchParams({
      spec: e?.id,
    });
  };
  /**
   * one of the two specs to be compared
   * @param e {SpecData.Spec}
   */
  const onAltSpecSelected = (e: SpecData.Spec) => {
    changeSearchParams({
      spec2: e?.id,
    });
  };

  const onCompare = () => {
    fetchCompare(); // fetches results in json, rendered with helpful icons
    fetchMarkdown(); // fetches results in pure markdown, for markdown preview before download
  };

  const handleDownload = (fileName: string, content = '') => {
    const eleLink = document.createElement('a');
    eleLink.download = fileName;
    eleLink.style.display = 'none';
    const blob = new Blob([content]);
    eleLink.href = URL.createObjectURL(blob);
    document.body.appendChild(eleLink);
    eleLink.click();
    document.body.removeChild(eleLink);
  };

  const props = {
    selectedService,
    leftSpec,
    rightSpec,
    leftComplianceList,
    rightComplianceList,
    compareData,
    compareDataFetching,
    markDownData,
    onSpecSelected,
    onAltSpecSelected,
    onCompare,
    handleDownload,
  };

  return <Compare {...props} />;
}
