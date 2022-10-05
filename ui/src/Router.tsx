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
