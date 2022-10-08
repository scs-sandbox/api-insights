import { SpecState } from '../../../../../components/Specs/SpecStateIcon/SpecStateIcon';
import { ComplianceData } from '../../../../../query/compliance';
import { SpecData } from '../../../../../query/spec';
import buildSortedListByStringField from '../../../../../utils/order';
import { RevisionData } from '../../Revision/Revision';
import { VersionData } from '../../VersionReport/VersionReport';

type VersionMap = {
  [index: string]: SpecData.Spec[];
};

type SpecMap = {
  [index: string]: RevisionData;
};

/**
 * Groups a list of spec data into lists of versions groups,
 * Specs of the same version group will be grouped together,
 * and will be sorted into a list of revisions for that version group.
 *
 * @param specList list Spec data objects
 * @returns returns list of version objects
 */
function groupVersionList(specList: SpecData.Spec[]) {
  if (!specList) return [];

  // group up by version
  const versionMap = specList.reduce(
    (pre: VersionMap, cur: SpecData.Spec) => ({
      ...pre,
      [cur.version]: [...(pre[cur.version] || []), cur],
    }),
    {},
  );

  return Object.keys(versionMap).map((key: string) => {
    const revisions = buildSortedListByStringField(
      versionMap[key],
      'updated_at',
      true,
    ) as SpecData.Spec[];
    const latestRevision = revisions[0];

    return {
      version: key,
      updated_at: latestRevision.updated_at,
      revisions,
      latestRevision,
    };
  });
}

/**
 * Appends Compliance Data to respective Spec data object.
 * @param specList list of specs for a given service
 * @param complianceList list of compliance results for specs
 */
function appendComplianceToSpec(
  specList: RevisionData[],
  complianceList: ComplianceData.Compliance[],
) {
  if (!specList) return;

  for (let i = 0; i < specList.length; i += 1) {
    const spec = specList[i];
    spec.complianceList = [];
  }

  const specMap = specList.reduce(
    (pre: SpecMap, cur: RevisionData) => ({
      ...pre,
      [cur.id]: cur,
    }),
    {},
  );

  (complianceList || []).forEach((i: ComplianceData.Compliance) => {
    const spec = specMap[i.spec_id];

    if (spec) {
      spec.complianceList = [...(spec.complianceList || []), i];
    }
  });
}

export function filterNotEmptyVersionList(versionList: VersionData[]) {
  if (!versionList) return null;

  return versionList
    .map((v) => ({
      ...v,
      revisions: v.revisions.filter((r) => r.state !== SpecState.Archive),
    }))
    .filter((v) => v.revisions.length);
}

export function buildVersionSortedList(
  specList: SpecData.Spec[],
  complianceList: ComplianceData.Compliance[],
): VersionData[] {
  const versionList = groupVersionList(specList);
  appendComplianceToSpec(specList as RevisionData[], complianceList);

  return buildSortedListByStringField(versionList, 'updated_at', true) as VersionData[];
}
