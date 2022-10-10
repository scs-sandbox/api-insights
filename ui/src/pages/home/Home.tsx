import { useState, useEffect } from 'react';
import { useOutletContext, useSearchParams } from 'react-router-dom';
import elementScrollTo from '../../utils/elementScrollTo';
import { AppFrameContext } from '../../components/Frame/AppFrame/AppFrame';
import { ServiceData } from '../../query/service';
import { useFetchOrganizationList } from '../../query/organization';
import Home from './components/Home/Home';

/**
 * Renders Home page showing available services
 */
export default function HomePage() {
  const [searchParams, setSearchParams] = useSearchParams();
  // Search key used to filter service names, retreieved from url parameter
  const searchKey = searchParams.get('search') ? searchParams.get('search') : '';
  const orgName = searchParams.get('org') ? searchParams.get('org') : 'all';
  const { sysInfo, serviceList, refetchServiceList } = useOutletContext() as AppFrameContext;
  const [newServiceId, setNewServiceId] = useState('');

  // List of organizations from the authorized endpoint
  const {
    data: authorizedOrganizationList,
  } = useFetchOrganizationList();

  // List of organizations derrived from the service list details
  // stringified so that set operation will return a unique set of organizations
  const dynamicOrganizationList = (serviceList !== undefined)
    ? Array.from(new Set(serviceList.map((service) => (JSON.stringify({
      name_id: service.organization_id, id: '', title: service.organization_id, description: '',
    }))))) : [];

  const organizationList = dynamicOrganizationList.map((orgString) => JSON.parse(orgString)) || [];

  const onServiceCreated = (e?: ServiceData.Service) => {
    setNewServiceId(e?.id);
    refetchServiceList();
  };

  const onServiceUpdated = () => {
    refetchServiceList();
  };

  const onCancelNewServiceNotation = () => {
    setNewServiceId('');
  };

  const onClearSearchKey = () => {
    setSearchParams({ search: '', org: orgName });
  };

  const onSearchKeyChanged = (keyword: string) => {
    setSearchParams({ search: keyword, org: orgName });
  };

  const onOrgChanged = (org: string) => {
    setSearchParams({ search: searchKey, org });
  };

  // Scroll to newly created service group upon addition into service list.
  useEffect(() => {
    if (!newServiceId) return;

    elementScrollTo(`#service-${newServiceId}`, 80);

    if (!orgName) onClearSearchKey();
  }, [(serviceList || []).length]);

  return (
    <Home
      orgName={orgName}
      authEnabled={sysInfo.auth?.enabled}
      newServiceId={newServiceId}
      searchKey={searchKey}
      serviceList={serviceList}
      authorizedOrganizationList={authorizedOrganizationList}
      organizationList={organizationList}
      onServiceCreated={onServiceCreated}
      onServiceUpdated={onServiceUpdated}
      onClearSearchKey={onClearSearchKey}
      onSearchKeyChanged={onSearchKeyChanged}
      onOrgChanged={onOrgChanged}
      onCancelNewServiceNotation={onCancelNewServiceNotation}
    />
  );
}
