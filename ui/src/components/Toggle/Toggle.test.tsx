import { fireEvent, render, screen } from '@testing-library/react';
import Toggle from './Toggle';

describe('<Toggle />', () => {
  test('Unchecked', () => {
    const { container } = render((
      <Toggle />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).not.toHaveClass('checked');
  });

  test('Checked', () => {
    const { container } = render((
      <Toggle checked label="Using" />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).toHaveClass('checked');

    const label = screen.getByText('Using');
    expect(label).toBeInTheDocument();

    fireEvent.click(toggle);
  });

  test('onToggle', () => {
    const onToggle = jest.fn();
    const { container } = render((
      <Toggle checked label="Using" onToggle={onToggle} />
    ));

    const toggle = container.querySelector('.toggle');
    expect(toggle).toBeInTheDocument();
    expect(toggle).toHaveClass('checked');

    fireEvent.click(toggle);
    expect(onToggle).toHaveBeenCalledTimes(1);
  });
});
