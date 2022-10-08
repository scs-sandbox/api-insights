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

  /**
   * builds url to download the markdown results of the compared specs
   */
  const handleDownload = () => {
    const url = `/services/${selectedService.id}/specs/diff/${leftSpec.id}/${rightSpec.id}?format=markdown`;
    window.open(buildApiAbsoluteUrl(url));
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
