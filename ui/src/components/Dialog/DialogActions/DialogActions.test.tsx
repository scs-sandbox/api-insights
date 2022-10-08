import { render, screen } from '@testing-library/react';
import DialogActions from './DialogActions';

describe('<DialogActions />', () => {
  test('Children', () => {
    render((
      <DialogActions>
        <button type="button">OK</button>
      </DialogActions>
    ));

    expect(screen.getByText('OK')).toBeInTheDocument();
  });

  test('ClassName', () => {
    const { container } = render((
      <DialogActions className="my-dialog-actions">
        <button type="button">OK</button>
      </DialogActions>
    ));

    expect(container.querySelector('.dialog-actions.my-dialog-actions')).toBeInTheDocument();
  });
});
