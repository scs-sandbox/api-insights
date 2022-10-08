import { render } from '@testing-library/react';
import DownloadIcon from './DownloadIcon';

describe('<DownloadIcon />', () => {
  test('Children', () => {
    const { container } = render((
      <DownloadIcon />
    ));

    expect(container.querySelector('.download-icon svg')).toBeInTheDocument();
  });
});
