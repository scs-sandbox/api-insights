import { render } from '@testing-library/react';
import SeverityIcon from './SeverityIcon';

describe('<SeverityIcon />', () => {
  test('Error', () => {
    const { container } = render((
      <SeverityIcon severity="error" title="Error" />
    ));

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Warning', () => {
    const { container } = render((
      <SeverityIcon severity="warning" title="Warning" />
    ));

    const icon = container.querySelector('.severity-icon-warning');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Warning');
  });

  test('Hint', () => {
    const { container } = render((
      <SeverityIcon severity="hint" title="Hint" />
    ));

    const icon = container.querySelector('.severity-icon-hint');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Hint');
  });

  test('Info', () => {
    const { container } = render((
      <SeverityIcon severity="info" title="Info" />
    ));

    const icon = container.querySelector('.severity-icon-info');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Info');
  });

  test('Unknown', () => {
    const { container } = render((
      <SeverityIcon severity="something" title="Something" />
    ));

    const icon = container.querySelector('.severity-icon-something');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Something');
  });
});
