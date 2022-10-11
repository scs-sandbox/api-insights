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

import { HTMLAttributes, ReactNode } from 'react';
import classNames from '../../../utils/className';
import PageHeader from '../PageHeader/PageHeader';
import PageNavBar from '../PageNavBar/PageNavBar';
import './PageFrame.scss';

type Props = HTMLAttributes<HTMLElement> & {
  header?: ReactNode;
};

export default function PageFrame(props: Props) {
  const className = classNames('app-page', props.className);

  return (
    <div className={className}>
      <PageHeader>{props.header}</PageHeader>
      <div className="page-body">
        <PageNavBar />
        <div className="page-content">
          {props.children}
        </div>
      </div>
    </div>
  );
}
