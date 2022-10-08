import { useEffect } from 'react';
import { useOutletContext } from 'react-router-dom';
import Timeline from './components/Timeline/Timeline';
import { useFetchServiceCompliance } from '../../query/compliance';
import { AppFrameContext } from '../../components/Frame/AppFrame/AppFrame';
import {
  SpecData,
  useFetchSpecList,
} from '../../query/spec';

export default function TimelinePage() {
  const { selectedService } = useOutletContext() as AppFrameContext;
  const {
    data: fullSpecList,
    isLoading: isSpecDataLoading,
    refetch: refetchSpecList,
  } = useFetchSpecList(selectedService?.id);
  const {
    data: complianceList,
    isLoading: isComplianceListLoading,
    refetch: refetchComplianceList,
  } = useFetchServiceCompliance(selectedService?.id);
  const specList = fullSpecList ? fullSpecList.filter((spec) => spec.state !== 'Reconstructed') : [];
  const snapshots = fullSpecList ? fullSpecList.filter((spec) => spec.state === 'Reconstructed') : [];
  const isSnapshotLoading = isSpecDataLoading;

  const refetchSpecCompliance = () => {
    refetchSpecList();
    refetchComplianceList();
  };

  const onReconstructed = () => {
    refetchSpecList();
  };

  useEffect(() => {
    if (!specList) return;

    // If a spec doesn't have a score, await backend analysis and perioidically refetch
    const needRefetchItem = specList.find(
      (i: SpecData.Spec) => i.score === null || i.score === undefined,
    );
    if (!needRefetchItem) return;

    setTimeout(() => {
      refetchSpecCompliance();
    }, 2000);
  }, [specList]);

  const props = {
    selectedService,
    specList,
    isSpecDataLoading,
    complianceList,
    isComplianceListLoading,
    snapshots: snapshots as SpecData.Snapshot[],
    isSnapshotLoading,
    onSpecUploaded: refetchSpecCompliance,
    onReconstructed,
  };

  return <Timeline {...props} />;
}
