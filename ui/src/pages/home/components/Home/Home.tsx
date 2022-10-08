import { useState, MouseEvent } from 'react';
import AddServiceButton from '../AddServiceButton/AddServiceButton';
import { getScoreLevels } from '../../../../components/Frame/Service/ScoreLevel/ScoreLevel';
import HelpButton from '../../../../components/HelpButton/HelpButton';
import HelpDialog from '../../../../components/HelpDialog/HelpDialog';
import { ServiceData, usePatchService } from '../../../../query/service';
import ServiceDetail from '../ServiceDetail/ServiceDetail';
import { groupServiceByProdcut, groupServiceByOrg } from '../../../../components/Frame/Service/ServiceDropDown/ServiceDropDown';
import Searchbar from '../../../../components/SearchBar/SearchBar';
import { OrganizationData } from '../../../../query/organization';
import './Home.scss';
import EditServiceDialog, { EditingEvent } from '../EditServiceDialog/EditServiceDialog';

type Props = {
  orgName: string;
  authEnabled: boolean;
  newServiceId: string;
  searchKey: string;
  serviceList: ServiceData.Service[];
  authorizedOrganizationList: OrganizationData.Organization[];
  organizationList: OrganizationData.Organization[];
  onServiceCreated: (newService: ServiceData.Service) => void;
  onServiceUpdated: (updatedService: ServiceData.Service) => void;
  onClearSearchKey: () => void;
  onSearchKeyChanged: (searchKey: string) => void;
  onOrgChanged: (org: string) => void;
  onCancelNewServiceNotation: () => void;
}

export type UpdatingData = {
  title: string;
  productTag: string;
  organization: string;
  name: string;
  description: string;
  visibility: string;
  contactName: string;
  email: string;
  url: string;
};

type ServiceCollection = {
  name: string;
  services: ServiceData.Service[];
}

function keyFilterService(service: ServiceData.Service, keyword: string) {
  return service.title.includes(keyword) || service.name_id.includes(keyword);
}

function tagFilterService(service: ServiceData.Service, tag: string) {
  return service.product_tag === tag;
}

function orgFilterService(service: ServiceData.Service, orgName: string) {
  return service.organization_id === orgName;
}

function filterOrgsByName(
  groups: { [index: string]: ServiceData.Service[] },
  orgName: string,
  orgList: OrganizationData.Organization[],
) {
  if (orgName === 'all') return groups;
  if (!orgList) return {};

  if (!groups[orgName]) return {};

  return {
    [orgName]: groups[orgName],
  };
}

function filterGroupsByName(groups: {[index: string]: ServiceData.Service[]}, name: string) {
  if (name === 'all') return groups;

  if (!groups[name]) return {};

  return {
    [name]: groups[name],
  };
}

/**
 * Renders Home page showing available services
 * @returns JSX.Element
 */
export default function Home(props: Props) {
  const { orgName } = props;
  const [tagFilter, setTagFilter] = useState('all');
  const [showDialog, setShowDialog] = useState(false);
  const [serviceToEdit, setServiceToEdit] = useState(null);
  const [showEditService, setShowEditService] = useState(false);
  const {
    isLoading: isServicePatching,
    isError: isServicePatchingError,
    isSuccess: isServicePatchingSuccess,
    mutate: patchService,
    reset: resetPatchServiceResult,
  } = usePatchService();
  // Filter by organization
  const orgFilteredServiceList = (orgName !== 'all' && props.serviceList)
    ? props.serviceList.filter((service) => orgFilterService(service, orgName))
    : props.serviceList;
  // Filter services by search key
  const keyfilteredServiceList = (props.searchKey !== '' && props.serviceList)
    ? orgFilteredServiceList.filter((service) => keyFilterService(service, props.searchKey))
    : orgFilteredServiceList;
  // Filter by product group tag
  const filteredServiceList = (tagFilter !== 'all' && props.serviceList)
    ? keyfilteredServiceList.filter((service) => tagFilterService(service, tagFilter))
    : keyfilteredServiceList;

  const filteredOrgGroups = groupServiceByOrg(keyfilteredServiceList || []);
  const filteredProductGroups = groupServiceByProdcut(keyfilteredServiceList || []);

  const onServiceUpdated = (e: ServiceData.Service) => {
    if (props.onServiceUpdated) {
      props.onServiceUpdated(e);
    }
  };

  const onCancelNewServiceNotation = () => {
    if (props.onCancelNewServiceNotation) {
      props.onCancelNewServiceNotation();
    }
  };

  /**
   * To show warning dialog for any services that aren't available yet.
   * @param clickedService service seelcted
   * @param e click event
   */
  const onClickService = (clickedService: ServiceData.Service, e: MouseEvent<HTMLElement>) => {
    if (clickedService.product_tag === 'CX Cloud') {
      e.stopPropagation();
      e.preventDefault();
      setShowDialog(true);
    }
  };

  const onClickEdit = (clickedService: ServiceData.Service) => {
    setServiceToEdit(clickedService);
    setShowEditService(true);
  };

  const onUpdatingService = (e: EditingEvent) => {
    const visibilityField = props.authEnabled ? {
      visibility: e.visibility,
    } : {};
    const data = {
      contact: {
        email: e.email,
        name: e.contactName,
        url: e.url,
      },
      description: e.description,
      organization_id: e.organizationId,
      product_tag: e.productTag,
      title: e.title,
      id: serviceToEdit.id,
      ...visibilityField,
    };

    patchService(data, {
      onSuccess: (patchServiceData: ServiceData.Service) => {
        setShowEditService(false);
        onServiceUpdated(patchServiceData);
      },
    });
  };
  /**
   * Set tag filers that will filter by product group
   * @param value the string value of product group names
   */
  const handleFilterClick = (value) => {
    setTagFilter(value);
    if (props.onClearSearchKey) {
      props.onClearSearchKey();
    }
  };

  const onSearchKeyChanged = (e) => {
    if (props.onSearchKeyChanged) {
      props.onSearchKeyChanged(e.target.value);
    }
  };

  function renderOrgButton(org: string) {
    return (
      <div
        key={org}
        className={`tag-option ${orgName === org ? 'active' : ''}`}
        onClick={() => {
          props.onOrgChanged(org);
        }}
      >
        {org}
      </div>
    );
  }

  function renderTag(tag) {
    return (
      <div
        key={tag}
        className={`tag-option ${tagFilter === tag ? 'active' : ''}`}
        onClick={() => {
          handleFilterClick(tag);
        }}
      >
        {tag}
      </div>
    );
  }

  function renderLegend() {
    const list = getScoreLevels()
      .reverse()
      .map((i) => (
        <span key={i.title} className="legend-item">
          <span className={`item-icon ${i.className}`} />
          <span className="item-title">{i.title}</span>
        </span>
      ));

    return (
      <div className="legend">
        <span className="legend-item greyed">
          <span className="item-icon" />
          <span className="item-title">Legend</span>
        </span>
        {list}
      </div>
    );
  }

  function renderAddServiceButton() {
    return (
      <AddServiceButton
        authEnabled={props.authEnabled}
        organizationList={props.authorizedOrganizationList}
        tags={Object.keys(filteredProductGroups)}
        onServiceCreated={props.onServiceCreated}
      />
    );
  }

  /**
   * Renders a collection of services
   * @param collection product group that contains multiple services
   * @param index the index of the service group in the list of service groups
   * @returns  JSX.element
   */
  function renderGroup(collection: ServiceCollection, index: number) {
    // Render legend shows compliance score of the first service in the product group
    const legend = index === 0 ? renderLegend() : null;
    return (
      <div key={collection.name} className="group-container">
        <div className="group-item">
          <div className="group-title">{collection.name}</div>
          {legend}
        </div>
        <div className="group-list">
          {
            (collection.services || [])
              .map((service) => (
                <div key={service.id} className="group-col">
                  <ServiceDetail
                    authEnabled={props.authEnabled}
                    isNewCreated={props.newServiceId === service.id}
                    service={service}
                    onClickService={onClickService}
                    onClickEdit={onClickEdit}
                    onCancelNewServiceNotation={onCancelNewServiceNotation}
                  />
                </div>
              ))
          }
        </div>
      </div>
    );
  }

  // const renderGroups = filterGroupsByName(filteredProductGroups, tagFilter);
  const renderGroups = filterOrgsByName(filteredOrgGroups, orgName, props.organizationList);
  const groups = Object.keys(renderGroups)
    .map((key) => ({
      name: key,
      services: renderGroups[key],
    } as ServiceCollection))
    .map(renderGroup);

  const tags = Object.keys(filteredProductGroups).map(renderTag);
  const orgButtons = props.organizationList.map((org) => renderOrgButton(org.name_id));
  // Object.keys(filteredOrgGroups).map(renderOrgButton);

  function renderGroupContent() {
    if (!props.serviceList) {
      return (
        <div className="service-loading">Loading...</div>
      );
    }

    if (props.serviceList.length === 0) {
      const addServiceButton = renderAddServiceButton();

      return (
        <div className="add-block">
          {addServiceButton}
        </div>
      );
    }

    return (
      <>
        <div className="group-bar">
          <div className="group-bar-left">
            <div
              className={`tag-option ${orgName === 'all' ? 'active' : ''}`}
              onClick={() => props.onOrgChanged('all')}
            >
              View all
            </div>
            {/* {tags} */}
            {orgButtons}
          </div>
          <div className="group-bar-right">
            <Searchbar
              searchKey={props.searchKey}
              onSearchKeyChanged={onSearchKeyChanged}
              onSearchKeyCleared={props.onClearSearchKey}
            />
          </div>
        </div>
        <div className="groups-container">{groups}</div>
      </>
    );
  }

  const addServiceButton = renderAddServiceButton();
  const groupContent = renderGroupContent();

  return (
    <div className="app-page home-page">
      <div className="home-nav-bar">
        <div className="welcome">
          <div className="welcome-title">Organization Dashboard</div>
        </div>
        {addServiceButton}
      </div>
      <HelpButton
        show={!Array.isArray(props.serviceList) || !props.serviceList.length}
        title={"Welcome to your Org's Dashboard"}
        message="Add your services and sort them under different products. Once you have those services set up, the health of your services will show on each card to give you a convenient cross-section of your environment for servicing priority."
      />
      <HelpDialog
        open={showDialog}
        handleClose={() => {
          setShowDialog(false);
        }}
        message="This service is not acessible at this moment."
      />
      {showEditService && (
        <EditServiceDialog
          service={serviceToEdit}
          open={showEditService}
          busy={isServicePatching}
          productTagList={Object.keys(filteredProductGroups)}
          authEnabled={props.authEnabled}
          onEditing={onUpdatingService}
          organizationList={props.authorizedOrganizationList}
          handleClose={() => setShowEditService(false)}
        />
      )}
      <div className="group-content">
        {groupContent}
      </div>
    </div>
  );
}
