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

import DropDownEditBox from '../../../../components/DropDown/DropDownEditBox/DropDownEditBox';
import DropDown from '../../../../components/DropDown/DropDown';

type Props = {
  value: string;
  editBoxReadOnly?: boolean;
  list: string[];
  onChange: (e: string) => void;
};

export default function ProductTagDropDown(props: Props) {
  const renderValue = () => (
    <DropDownEditBox
      value={props.value}
      readonly={props.editBoxReadOnly}
      onChange={props.onChange}
    />
  );

  const renderMenuItemLabel = (option: string) => (
    <div className="menu-item-label">{option}</div>
  );

  return (
    <DropDown
      className="product-dropdown"
      value={props.value}
      options={props.list}
      renderValue={renderValue}
      renderMenuItemLabel={renderMenuItemLabel}
      onChange={props.onChange}
    />
  );
}
