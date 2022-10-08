import { useOutletContext, useSearchParams } from 'react-router-dom';
import { useEffect } from 'react';
import {
  useFetchSpecList,
  useFetchSpecDetail,
  SpecData,
} from '../../query/spec';
import { useFetchSpecCompliance } from '../../query/compliance';
import { useFetchAnalyzerList } from '../../query/analyzer';
import { AppFrameContext } from '../../components/Frame/AppFrame/AppFrame';
import Report from './components/Report/Report';
import { ServiceData } from '../../query/service';

/**
 * Renders the compliance report for a given API spec
 * @returns JSX.element
 */
export default function ReportPage() {
  const {
    selectedService: service,
    updateSpecSummary,
  } = useOutletContext() as AppFrameContext;
  const { data: analyzerList } = useFetchAnalyzerList();
  const { refetch: refetchSpecList } = useFetchSpecList(service?.id);
  const [searchParams, setSearchParams] = useSearchParams();
  // Retrieve selected API spec id from url paremeter
  const searchParamSpec = searchParams.get('spec');
  const { data: selectedSpec, refetch: refetchSpecDetail } = useFetchSpecDetail(
    service?.id,
    searchParamSpec,
  );
  const {
    data: complianceList,
    isLoading: complianceListLoading,
  } = useFetchSpecCompliance(service?.id, selectedSpec ? selectedSpec.id : '');

  const resetSpecSummary = () => {
    const summary: ServiceData.ServiceSummary = selectedSpec ? {
      score: selectedSpec.score,
      version: selectedSpec.version,
      revision: selectedSpec.revision,
      updated_at: selectedSpec.updated_at,
    } : null;

    updateSpecSummary(summary);
  };

  const onSpecSelected = async (e: SpecData.Spec) => {
    const params = {
      service: searchParams.get('service'),
      spec: e.id,
    };
    await setSearchParams(params);
    refetchSpecDetail();

    resetSpecSummary();
  };

  const onSpecUploaded = () => {
    refetchSpecList();
  };

  useEffect(() => {
    resetSpecSummary();
  }, [selectedSpec]);

  useEffect(() => {
    refetchSpecDetail();
  }, [searchParamSpec]);

  const props = {
    service,
    complianceList,
    complianceListLoading,
    analyzerList,
    selectedSpec,
    onSpecSelected,
    onSpecUploaded,
    refetchSpecDetail,
  };

  return <Report {...props} />;
}
