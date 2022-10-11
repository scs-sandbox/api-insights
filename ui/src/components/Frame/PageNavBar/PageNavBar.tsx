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

import { Link, useOutletContext } from 'react-router-dom';
import { AppFrameContext } from '../AppFrame/AppFrame';
import './PageNavBar.scss';

type NavItemProps = {
  page: string;
  name: string;
  title: string;
};

function isMe(page: string) {
  const pageInUrl = window.location.pathname.substring(
    window.location.pathname.lastIndexOf('/') + 1,
  );
  return page === pageInUrl;
}

function NavItem(props: NavItemProps) {
  const { navStore } = useOutletContext() as AppFrameContext;
  const to = `/${props.page}?${navStore[props.page]}`;
  const className = `nav-item${isMe(props.page) ? ' active' : ''}`;

  return (
    <li className="nav-item-box">
      <Link className={className} to={to}>
        <div className={`nav-item-icon icon-${props.name}`} />
        <div className="nav-item-name">{props.title}</div>
      </Link>
    </li>
  );
}

export default function PageNavBar() {
  return (
    <div className="page-nav-bar">
      <ul className="nav-item-list">
        <NavItem page="timeline" name="timeline" title="Timeline" />
        <NavItem page="reports" name="reports" title="Reports" />
        <NavItem page="comparison" name="comparison" title="Comparison" />
      </ul>
    </div>
  );
}
