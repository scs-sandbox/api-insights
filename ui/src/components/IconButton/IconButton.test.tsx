import { render, screen } from '@testing-library/react';
import IconButton from './IconButton';

describe('<IconButton />', () => {
  test('No icon', () => {
    render((
      <IconButton>
        <div>Add</div>
      </IconButton>
    ));

    const children = screen.getByText('Add');
    expect(children).toBeInTheDocument();
  });

  test('Has icon', () => {
    const target = (
      <IconButton icon={<div>icon</div>}>
        <div>Add</div>
      </IconButton>
    );
    render(target);

    const icon = screen.getByText('icon');
    expect(icon).toBeInTheDocument();

    const children = screen.getByText('Add');
    expect(children).toBeInTheDocument();
  });
});
