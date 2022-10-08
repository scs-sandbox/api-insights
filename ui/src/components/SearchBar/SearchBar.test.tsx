import { fireEvent, render, screen } from '@testing-library/react';
import SearchBar from './SearchBar';

describe('<SearchBar />', () => {
  test('Basic', () => {
    const onChange = jest.fn();
    const onClear = jest.fn();
    const { container } = render((
      <SearchBar searchKey="hello" onSearchKeyChanged={onChange} onSearchKeyCleared={onClear} />
    ));

    const value = screen.getByDisplayValue('hello');
    expect(value).toHaveClass('search-input');
    fireEvent.change(value, { target: { value: 'world' } });
    expect(onChange).toHaveBeenCalledTimes(1);

    const clear = container.querySelector('.search-clear');
    expect(clear).toBeInTheDocument();
    fireEvent.click(clear, {});
    expect(onClear).toHaveBeenCalledTimes(1);
  });
});
