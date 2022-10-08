import { render } from '@testing-library/react';
import IssueIcon from './IssueIcon';

describe('<IssueIcon />', () => {
  test('Error', () => {
    const { container } = render((
      <IssueIcon severity="error" title="Error" />
    ));

    const icon = container.querySelector('.issue-icon-error');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Error');
  });

  test('Warning', () => {
    const { container } = render((
      <IssueIcon severity="warning" title="Warning" />
    ));

    const icon = container.querySelector('.issue-icon-warning');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Warning');
  });

  test('Hint', () => {
    const { container } = render((
      <IssueIcon severity="hint" title="Hint" />
    ));

    const icon = container.querySelector('.issue-icon-hint');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Hint');
  });

  test('Info', () => {
    const { container } = render((
      <IssueIcon severity="info" title="Info" />
    ));

    const icon = container.querySelector('.issue-icon-info');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Info');
  });

  test('Unknown', () => {
    const { container } = render((
      <IssueIcon severity="something" title="Something" />
    ));

    const icon = container.querySelector('.issue-icon-something');
    expect(icon).toBeInTheDocument();
    expect(icon).toHaveAttribute('title', 'Something');
  });
});
