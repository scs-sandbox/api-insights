import { HTMLAttributes } from 'react';
import { Link, useOutletContext } from 'react-router-dom';
import { AppFrameContext } from '../AppFrame/AppFrame';
import ServiceDropDown from '../Service/ServiceDropDown/ServiceDropDown';
import './PageHeader.scss';

type Props = HTMLAttributes<HTMLElement>;

export default function PageHeader(props: Props) {
  const {
    serviceList,
    selectedService,
    specServiceSummary,
    onServiceSelected,
    refetchServiceList,
  } = useOutletContext() as AppFrameContext;

  return (
    <div className="page-header">
      <div className="back-col">
        <Link className="goback" to="/services">
          <div className="back-icon" />
          <div className="back-text">All Services</div>
        </Link>
      </div>
      <div className="service-col">
        <ServiceDropDown
          services={serviceList}
          selectedService={selectedService}
          specServiceSummary={specServiceSummary}
          onServiceUpdated={refetchServiceList}
          onChange={onServiceSelected}
        />
      </div>
      <div className="slot-col">{props.children}</div>
    </div>
  );
}
