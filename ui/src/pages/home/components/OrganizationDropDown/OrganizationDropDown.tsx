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

import { OrganizationData } from '../../../../query/organization';
import DropDownEditBox from '../../../../components/DropDown/DropDownEditBox/DropDownEditBox';
import DropDown from '../../../../components/DropDown/DropDown';

type Props = {
  value: string;
  editBoxReadOnly?: boolean;
  list: OrganizationData.Organization[];
  onChange: (e: string) => void;
};

export default function OrganizationDropDown(props: Props) {
  const requestOptionValue = (option: unknown) => (option as OrganizationData.Organization).name_id;

  const renderValue = () => {
    const foundOrganizationItem = (props.list || []).find(
      (i) => requestOptionValue(i) === props.value,
    );

    const value = foundOrganizationItem ? foundOrganizationItem.title : props.value;

    return (
      <DropDownEditBox value={value} readonly={props.editBoxReadOnly} onChange={props.onChange} />
    );
  };

  const renderMenuItemLabel = (option: unknown) => (
    <div className="menu-item-label">{(option as OrganizationData.Organization).title}</div>
  );

  return (
    <DropDown
      className="organization-dropdown"
      value={props.value}
      options={props.list}
      renderValue={renderValue}
      requestOptionValue={requestOptionValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={props.onChange}
    />
  );
}
