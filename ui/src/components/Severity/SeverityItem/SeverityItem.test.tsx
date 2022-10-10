import { render, screen } from '@testing-library/react';
import SeverityItem from './SeverityItem';

describe('<SeverityItem />', () => {
  test('little props', () => {
    const { container } = render((
      <SeverityItem severity="error" />
    ));

    const item = container.querySelector('.severity-item');
    expect(item).toBeInTheDocument();
    expect(item).toHaveClass('severity-item-error');

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Full props', () => {
    const { container } = render((
      <SeverityItem severity="error" count={5} showLabel label="Error(s)" />
    ));

    const item = container.querySelector('.severity-item');
    expect(item).toBeInTheDocument();

    const icon = container.querySelector('.severity-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error(s)');

    const count = screen.getByText('5');
    expect(count).toBeInTheDocument();
    expect(count).toHaveClass('severity-item-count');

    const label = screen.getByText('Error(s)');
    expect(label).toBeInTheDocument();
    expect(label).toHaveClass('severity-item-label');
  });
});
