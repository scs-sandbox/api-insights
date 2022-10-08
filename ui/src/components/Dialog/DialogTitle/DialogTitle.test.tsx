import { fireEvent, render, screen } from '@testing-library/react';
import DialogTitle from './DialogTitle';

describe('<DialogTitle />', () => {
  test('Icon', () => {
    const { container } = render((
      <DialogTitle>My Dialog</DialogTitle>
    ));

    expect(container.querySelector('.title-icon')).toBeInTheDocument();
  });

  test('Title', () => {
    render((
      <DialogTitle>My Dialog</DialogTitle>
    ));

    expect(screen.getByText('My Dialog')).toBeInTheDocument();
  });

  test('Close Button', () => {
    const onClose = jest.fn();

    render((
      <DialogTitle onClose={onClose}>My Dialog</DialogTitle>
    ));

    expect(screen.getByRole('button')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button'), {});

    expect(onClose).toHaveBeenCalledTimes(1);
  });
});
