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

import Pages from 'react-js-pagination';
import './Pagination.scss';

type Props = {
  currentPage: number;
  pageSize: number;
  total: number;
  pageRangeDisplayed?: number;
  prevPageText?: string;
  nextPageText?: string;
  firstPageText?: string;
  lastPageText?: string;
  onChange?: (value: number) => void;
};

export default function Pagination(props: Props) {
  const onChange = (e: number) => {
    const index = e;
    if (props.onChange) {
      props.onChange(index);
    }
  };

  return (
    <Pages
      prevPageText={props.prevPageText || '<'}
      nextPageText={props.nextPageText || '>'}
      firstPageText={props.firstPageText || '<<'}
      lastPageText={props.lastPageText || '>>'}
      itemClassFirst="move-btn first-page"
      itemClassPrev="move-btn prev-page"
      itemClassNext="move-btn next-page"
      itemClassLast="move-btn last-page"
      itemClass="page-btn"
      linkClass="page-btn-link"
      activePage={props.currentPage}
      itemsCountPerPage={props.pageSize}
      totalItemsCount={props.total}
      pageRangeDisplayed={props.pageRangeDisplayed || 10}
      onChange={onChange}
    />
  );
}
