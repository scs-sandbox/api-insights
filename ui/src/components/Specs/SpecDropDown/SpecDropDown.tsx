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

import { SpecData, useFetchSpecList } from '../../../query/spec';
import DropDown from '../../DropDown/DropDown';
import './SepcDropDown.scss';

type Props = {
  serviceId: string;
  selectedSpec?: SpecData.Spec;
  onChange: (spec: SpecData.Spec) => void;
};

export default function SpecDropDown(props: Props) {
  const { data, isLoading } = useFetchSpecList(props.serviceId);
  const specList = data as SpecData.Spec[];

  const onChange = (value: string) => {
    const spec = data.find((i) => i.id === value);
    if (props.onChange) {
      props.onChange(spec);
    }
  };

  const requestOptionValue = (option: SpecData.Spec) => option.id;

  const renderValue = (value: string) => {
    const spec = (specList || []).find((i) => i.id === value);

    if (!spec) return null;

    return (
      <div className="dropdown-value">
        <span className="version-label">{spec.version}</span>
        <span className="revision-label">{spec.revision}</span>
      </div>
    );
  };

  const renderMenuItemLabel = (option: SpecData.Spec) => (
    <div className="menu-item-label">
      <span className="version-label">{option.version}</span>
      <span className="revision-label">{option.revision}</span>
    </div>
  );

  const placeholder = isLoading ? 'Loading...' : 'Please select';
  const value = props.selectedSpec ? props.selectedSpec.id : '';

  return (
    <DropDown
      className="spec-dropdown"
      menuItemClassName="spec-menu-item"
      placeholder={placeholder}
      value={value}
      options={data}
      requestOptionValue={requestOptionValue}
      renderValue={renderValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={onChange}
    />
  );
}
