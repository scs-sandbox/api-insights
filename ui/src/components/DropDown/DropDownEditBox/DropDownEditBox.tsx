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

import { ChangeEvent, MouseEvent } from 'react';
import './DropDownEditBox.scss';

type Props = {
  value?: string;
  readonly?: boolean;
  onChange?: (value: string) => void;
};

export default function DropDownEditBox(props: Props) {
  const onChangeText = (e: ChangeEvent<HTMLInputElement>) => {
    if (props.onChange) {
      props.onChange(e.target.value);
    }
  };

  const onClick = (e: MouseEvent<HTMLInputElement>) => {
    if (props.readonly) return;

    e.stopPropagation();
  };

  return (
    <div className="dropdown-value">
      <input
        className="dropdown-value-input"
        readOnly={props.readonly}
        value={props.value}
        onChange={onChangeText}
        onClick={onClick}
      />
    </div>
  );
}
