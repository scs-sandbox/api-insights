import { useState, MouseEvent } from 'react';
import { MenuItem } from '@mui/material';
import { ServiceData } from '../../../../../query/service';
import './ServiceGroup.scss';

type ServiceGroupProps = {
  selectedService: ServiceData.Service;
  services: ServiceData.Service[];
  group: string;
  onSelectedMenuItem?: (event: MouseEvent<HTMLElement>) => void;
};

export default function ServiceGroup(props: ServiceGroupProps) {
  const [hide, setHide] = useState(false);
  const handleToggle = () => {
    setHide(!hide);
  };
  const isSelectedGroup = props.services.find(
    (service) => service.id === props.selectedService?.id,
  );

  const renderService = (service: ServiceData.Service) => {
    const isSelectedService = service.id === props.selectedService?.id;

    const sx = isSelectedService ? {
      width: '100%',
      boxSizing: 'border-box',
      backgroundColor: ' rgba(0, 188, 235, 0.14)',
    } : { width: '100%', boxSizing: 'border-box' };

    const selectedFlag = isSelectedService
      && <div className="drop group-selected"> &#10003; </div>;

    return (
      <MenuItem
        key={service.id}
        data-id={service.id}
        onClick={props.onSelectedMenuItem}
        sx={sx}
      >
        <div className="item">
          {service.title}
          {selectedFlag}
        </div>
      </MenuItem>
    );
  };

  const count = `(${props.services.length})`;

  return (
    <div className="service-group-menu-item">
      <div
        className={`group-title ${isSelectedGroup ? 'group-selected' : ''}`}
        onClick={handleToggle}
      >
        {props.group}
        {' '}
        <div className="count">
          {count}
        </div>
        <div className="drop">{hide ? '\u276F' : '\u2303 '}</div>
      </div>
      <div className={`group-list ${hide ? 'hide' : ''} `}>
        {props.services.map(renderService)}
      </div>
    </div>
  );
}
