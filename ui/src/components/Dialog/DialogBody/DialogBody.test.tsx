import { render, screen } from '@testing-library/react';
import DialogBody from './DialogBody';

describe('<DialogBody />', () => {
  test('Children', () => {
    render((
      <DialogBody>
        <button type="button">OK</button>
      </DialogBody>
    ));

    expect(screen.getByText('OK')).toBeInTheDocument();
  });

  test('ClassName', () => {
    const { container } = render((
      <DialogBody className="my-dialog-body">
        <button type="button">OK</button>
      </DialogBody>
    ));

    expect(container.querySelector('.dialog-body.my-dialog-body')).toBeInTheDocument();
  });
});
