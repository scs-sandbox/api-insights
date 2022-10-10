import { render } from '@testing-library/react';
import UploadIcon from './UploadIcon';

describe('<UploadIcon />', () => {
  test('Normal', () => {
    const { container } = render((
      <UploadIcon />
    ));

    expect(container.querySelector('.upload-icon.state-busy')).toBeNull();
    expect(container.querySelector('.upload-icon')).toBeInTheDocument();
  });

  test('Busy', () => {
    const { container } = render((
      <UploadIcon busy />
    ));

    expect(container.querySelector('.upload-icon.state-busy')).toBeInTheDocument();
    expect(container.querySelector('.upload-icon')).toBeInTheDocument();
  });
});
