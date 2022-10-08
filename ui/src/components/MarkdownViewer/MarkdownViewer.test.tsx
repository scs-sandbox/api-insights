import { render, screen } from '@testing-library/react';
import MarkdownViewer from './MarkdownViewer';

describe('<MarkdownViewer />', () => {
  test('Basic', () => {
    render((
      <MarkdownViewer text="See [more](/home.html)" />
    ));

    const link = screen.getByText('more');
    expect(link).toHaveAttribute('href', '/home.html');
  });

  test('Basic', () => {
    render((
      <MarkdownViewer />
    ));
  });
});
