import { render, screen } from '@testing-library/react';
import SnackAlert from './SnackAlert';

describe('<SnackAlert />', () => {
  test('Success Message', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="success" message="Hello, world" onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });

  test('Error Message', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="error" message="Hello, world" onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });

  test('Duration', () => {
    const onClose = jest.fn();

    render((
      <SnackAlert severity="success" message="Hello, world" autoHideDuration={1000} onClose={onClose} />
    ));

    const message = screen.getByText('Hello, world');
    expect(message).toBeInTheDocument();
  });
});
