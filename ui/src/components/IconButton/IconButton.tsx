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

import { ButtonHTMLAttributes, ReactNode } from 'react';
import classNames from '../../utils/className';
import './IconButton.scss';

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon?: ReactNode;
};

export default function IconButton(props: Props) {
  const {
    icon, children, className, ...other
  } = props;

  const fullClassName = classNames('button-primary', 'icon-button', className);

  const iconCol = icon ? (<div className="icon-col">{icon}</div>) : null;

  return (
    <button type="button" className={fullClassName} {...other}>
      <div className="button-row">
        {iconCol}
        <div className="text-col">{children}</div>
      </div>
    </button>
  );
}
