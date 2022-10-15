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

import { MouseEvent } from 'react';
import dayjs from 'dayjs';
import CircleScore from '../../CircleScore/CircleScore';
import EditGroupButton from '../../EditGroupButton/EditGroupButton';
import { ServiceData } from '../../../../../query/service';
import './SelectedService.scss';

type Props = {
  service: ServiceData.Service;
  placeholder?: string;
  specServiceSummary?: ServiceData.ServiceSummary;
  onServiceUpdated?: () => void;
  onClickService?: (event: MouseEvent<HTMLDivElement>) => void;
}

export default function SelectedService(props: Props) {
  if (!props.service) {
    return (
      <div className="selected-service" onClick={props.onClickService}>
        <div className="place-holder drop-block">
          {props.placeholder || 'Select a service'}
        </div>
      </div>
    );
  }

  const specServiceSummary = props.specServiceSummary || props.service.summary;

  const renderUpdateAt = () => {
    const updatedAt = specServiceSummary?.updated_at || '';

    if (!updatedAt) {
      return <div className="last-updated hidden" />;
    }

    const time = dayjs(updatedAt).format('MMM DD, HH:mm');

    return (
      <div className="last-updated">
        <span className="last-updated-title">Updated</span>
        {time}
      </div>
    );
  };

  const score = (
    <div className="score-col">
      <CircleScore size={64} thickness={4} value={specServiceSummary?.score}>
        <div className="info-text">
          <div className="value">{specServiceSummary?.score}</div>
        </div>
      </CircleScore>
    </div>
  );

  const updatedAt = renderUpdateAt();

  return (
    <div className="selected-service">
      {score}
      <div className="info-col">
        <div className="group-part">
          {props.service.product_tag}
          <EditGroupButton
            service={props.service}
            onServiceUpdated={props.onServiceUpdated}
          />
        </div>
        <div
          className="service-title drop-block"
          onClick={props.onClickService}
        >
          {props.service.title}
        </div>
        {updatedAt}
      </div>
    </div>
  );
}
