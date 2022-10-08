import { render, screen, fireEvent } from '@testing-library/react';
import DropDownEditBox from './DropDownEditBox';

describe('<DropDownEditBox />', () => {
  test('Basic', () => {
    const onChange = jest.fn();

    render((
      <DropDownEditBox value="hello" onChange={onChange} />
    ));

    const input = screen.getByDisplayValue('hello');
    expect(input).toBeInTheDocument();
    expect(input).toHaveClass('dropdown-value-input');

    fireEvent.change(input, { target: { value: 'world' } });
    expect(onChange).toHaveBeenCalledTimes(1);

    fireEvent.click(input, {});
  });

  test('Readonly', () => {
    render((
      <DropDownEditBox value="hello" readonly />
    ));

    const input = screen.getByDisplayValue('hello');
    fireEvent.click(input, {});

    fireEvent.change(input, { target: { value: 'world' } });
  });
});
