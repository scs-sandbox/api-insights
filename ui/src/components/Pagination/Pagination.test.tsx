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

import { fireEvent, render, screen } from '@testing-library/react';
import Pagination from './Pagination';

describe('<Pagination />', () => {
  test('Basic', () => {
    const { container } = render((
      <Pagination currentPage={2} total={100} pageSize={10} />
    ));

    const firstPage = container.querySelector('.first-page .page-btn-link');
    expect(firstPage).toHaveTextContent('<<');
    const prevPage = container.querySelector('.prev-page .page-btn-link');
    expect(prevPage).toHaveTextContent('<');
    const curPage = container.querySelector('.page-btn.active .page-btn-link');
    expect(curPage).toHaveTextContent('2');
    const nextPage = container.querySelector('.next-page .page-btn-link');
    expect(nextPage).toHaveTextContent('>');
    const lastPage = container.querySelector('.last-page .page-btn-link');
    expect(lastPage).toHaveTextContent('>>');
  });

  test('Specify text for first page, prev page...', () => {
    const { container } = render((
      <Pagination
        currentPage={2}
        total={100}
        pageSize={10}
        pageRangeDisplayed={5}
        firstPageText="First Page"
        prevPageText="Previous Page"
        nextPageText="Next Page"
        lastPageText="Last Page"
      />
    ));

    const firstPage = container.querySelector('.first-page .page-btn-link');
    expect(firstPage).toHaveTextContent('First Page');
    const prevPage = container.querySelector('.prev-page .page-btn-link');
    expect(prevPage).toHaveTextContent('Previous Page');
    const curPage = container.querySelector('.page-btn.active .page-btn-link');
    expect(curPage).toHaveTextContent('2');
    const nextPage = container.querySelector('.next-page .page-btn-link');
    expect(nextPage).toHaveTextContent('Next Page');
    const lastPage = container.querySelector('.last-page .page-btn-link');
    expect(lastPage).toHaveTextContent('Last Page');

    fireEvent.click(screen.getByText('First Page'), {});
    expect(curPage).toHaveTextContent('2');
  });

  test('onChange', () => {
    const onChange = jest.fn();

    render((
      <Pagination currentPage={2} total={100} pageSize={10} onChange={onChange} />
    ));

    fireEvent.click(screen.getByText('<<'), {});
    expect(onChange).toHaveBeenCalledTimes(1);
  });
});
