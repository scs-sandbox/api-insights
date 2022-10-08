import { render, screen } from '@testing-library/react';
import BusyIcon, { BusyIconType } from './BusyIcon';

describe('<BusyIcon />', () => {
  test('Circle', () => {
    const { container } = render((
      <BusyIcon />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeNull();
    expect(container.querySelector('.icon-circle')).toBeInTheDocument();
  });

  test('Arrow Circle', () => {
    const { container } = render((
      <BusyIcon type={BusyIconType.ArrowCircle} />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeNull();
    expect(container.querySelector('.icon-arrowcircle')).toBeInTheDocument();
  });

  test('Busy', () => {
    const { container } = render((
      <BusyIcon busy />
    ));

    expect(container.querySelector('.busy-icon.state-busy')).toBeInTheDocument();
    expect(container.querySelector('.icon-circle')).toBeInTheDocument();
  });
});
