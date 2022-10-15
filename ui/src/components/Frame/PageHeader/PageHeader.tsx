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

import { HTMLAttributes } from 'react';
import { Link, useOutletContext } from 'react-router-dom';
import { AppFrameContext } from '../AppFrame/AppFrame';
import ServiceDropDown from '../Service/ServiceDropDown/ServiceDropDown';
import './PageHeader.scss';

type Props = HTMLAttributes<HTMLElement>;

export default function PageHeader(props: Props) {
  const {
    serviceList,
    selectedService,
    specServiceSummary,
    onServiceSelected,
    refetchServiceList,
  } = useOutletContext() as AppFrameContext;

  return (
    <div className="page-header">
      <div className="back-col">
        <Link className="goback" to="/services">
          <div className="back-icon" />
          <div className="back-text">All Services</div>
        </Link>
      </div>
      <div className="service-col">
        <ServiceDropDown
          services={serviceList}
          selectedService={selectedService}
          specServiceSummary={specServiceSummary}
          onServiceUpdated={refetchServiceList}
          onChange={onServiceSelected}
        />
      </div>
      <div className="slot-col">{props.children}</div>
    </div>
  );
}
