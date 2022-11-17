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

import { useState, MouseEvent, ChangeEvent } from 'react';
import { CircularProgress, Menu } from '@mui/material';
import ServiceGroup from './ServiceGroup/ServiceGroup';
import SelectedService from './SelectedService/SelectedService';
import Searchbar from '../../../SearchBar/SearchBar';
import { ServiceData } from '../../../../query/service';
import './ServiceDropDown.scss';

type Props = {
  services: ServiceData.Service[];
  selectedService: ServiceData.Service;
  specServiceSummary?: ServiceData.ServiceSummary;
  onServiceUpdated: () => void;
  placeholder?: string;
  onChange: (service: ServiceData.Service) => void;
};

type ProductGroups = {
  [index: string]: ServiceData.Service[];
}

type OrgGroups = {
  [index: string]: ServiceData.Service[];
}

export function groupServiceByOrg(allServices: ServiceData.Service[]) {
  return (allServices || []).reduce((pre: OrgGroups, cur: ServiceData.Service) => {
    const groups = pre[cur.organization_id] || [];
    return {
      ...pre,
      [cur.organization_id]: [...groups, cur],
    };
  }, {});
}

export function groupServiceByProdcut(allServices: ServiceData.Service[]) {
  return (allServices || []).reduce((pre: ProductGroups, cur: ServiceData.Service) => {
    const groups = pre[cur.product_tag] || [];
    return {
      ...pre,
      [cur.product_tag]: [...groups, cur],
    };
  }, {});
}

export default function ServiceDropDown(props: Props) {
  const [anchorEl, setAnchorEl] = useState<HTMLElement>();
  const openMenu = Boolean(anchorEl);
  const [searchKey, setSearchKey] = useState('');

  const filteredService = (props.services || [])
    .filter((service) => service.title.toLowerCase().includes(searchKey.toLowerCase()));

  const productGroups = groupServiceByProdcut(filteredService);

  const onOpenMenu = (event: MouseEvent<HTMLDivElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const onCloseMenu = () => {
    setAnchorEl(undefined);
  };

  const onSelectedMenuItem = (event: MouseEvent<HTMLElement>) => {
    onCloseMenu();

    const service = props.services.find(
      (i: ServiceData.Service) => i.id === event.currentTarget.dataset.id,
    );

    if (!service) return;

    if (props.onChange) {
      props.onChange(service);
    }
  };

  const renderSelectedService = () => {
    if (!props.services) {
      return (
        <div className="services-loading">
          <CircularProgress color="inherit" />
        </div>
      );
    }

    if (!props.services.length) {
      return <div className="services-message">No services</div>;
    }

    return (
      <SelectedService
        service={props.selectedService}
        placeholder={props.placeholder}
        specServiceSummary={props.specServiceSummary}
        onServiceUpdated={props.onServiceUpdated}
        onClickService={onOpenMenu}
      />
    );
  };

  function renderServiceGroups() {
    return Object.keys(productGroups).map((product) => (
      <ServiceGroup
        selectedService={props.selectedService}
        services={productGroups[product]}
        group={product}
        key={product}
        onSelectedMenuItem={onSelectedMenuItem}
      />
    ));
  }

  const onSearchKeyChanged = (e: ChangeEvent<HTMLInputElement>) => {
    setSearchKey(e.target.value);
  };

  const onSearchKeyCleared = () => {
    setSearchKey('');
  };

  const serviceGroups = renderServiceGroups();

  return (
    <div className="service-dropdown">
      {renderSelectedService()}
      <Menu open={openMenu} anchorEl={anchorEl} onClose={onCloseMenu}>
        <div className="service-dropdown-search-bar">
          <Searchbar
            searchKey={searchKey}
            onSearchKeyChanged={onSearchKeyChanged}
            onSearchKeyCleared={onSearchKeyCleared}
          />
        </div>
        {serviceGroups}
      </Menu>
    </div>
  );
}
