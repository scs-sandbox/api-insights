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

import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Navigate } from 'react-router';

import Home from './pages/home/Home';
import Compare from './pages/compare/Compare';
import Timeline from './pages/timeline/Timeline';
import Report from './pages/report/Report';
import AppFrame from './components/Frame/AppFrame/AppFrame';

const NavStore = {
  timeline: '',
  reports: '',
  comparison: '',
};

export default function Root() {
  return (
    <BrowserRouter basename={process.env.REACT_APP_ROOT_PATH || ''}>
      <Routes>
        <Route path="" element={<AppFrame navStore={NavStore} />}>
          <Route path="" element={<Navigate to="/services" replace />} />
          <Route path="services" element={<Home />} />
          <Route path="timeline" element={<Timeline />} />
          <Route path="reports" element={<Report />} />
          <Route path="comparison" element={<Compare />} />
        </Route>
        <Route path="*" element={<Navigate to="/services" replace />} />
      </Routes>
    </BrowserRouter>
  );
}
