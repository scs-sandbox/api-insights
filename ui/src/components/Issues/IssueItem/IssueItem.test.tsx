import { render, screen } from '@testing-library/react';
import IssueItem from './IssueItem';

describe('<IssueItem />', () => {
  test('little props', () => {
    const { container } = render((
      <IssueItem severity="error" />
    ));

    const item = container.querySelector('.issue-item');
    expect(item).toBeInTheDocument();
    expect(item).toHaveClass('issue-item-error');

    const icon = container.querySelector('.issue-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Full props', () => {
    const { container } = render((
      <IssueItem severity="error" count={5} showLabel label="Error(s)" />
    ));

    const item = container.querySelector('.issue-item');
    expect(item).toBeInTheDocument();

    const icon = container.querySelector('.issue-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error(s)');

    const count = screen.getByText('5');
    expect(count).toBeInTheDocument();
    expect(count).toHaveClass('issue-item-count');

    const label = screen.getByText('Error(s)');
    expect(label).toBeInTheDocument();
    expect(label).toHaveClass('issue-item-label');
  });
});
