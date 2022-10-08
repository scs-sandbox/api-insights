import { render, screen, fireEvent } from '@testing-library/react';
import DropDown from './DropDown';

describe('<DropDown />', () => {
  test('options is null', () => {
    const { container } = render((
      <DropDown value="hello" />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);
  });

  test('Empty options', () => {
    const { container } = render((
      <DropDown value="hello" options={[]} />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);
  });

  test('Placeholder', () => {
    const { container } = render((
      <DropDown placeholder="select one" />
    ));

    const placeholder = container.querySelector('.dropdown-placeholder');
    expect(placeholder).toBeInTheDocument();
  });

  test('String Array & Basic Feature', () => {
    const onChange = jest.fn();

    const { container } = render((
      <DropDown value="hello" options={['hello', 'world']} onChange={onChange} />
    ));

    const value = screen.getByText('hello');
    expect(value).toBeInTheDocument();
    expect(value).toHaveClass('dropdown-value');

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = screen.getByText('world');
    expect(menuItem).toBeInTheDocument();
    fireEvent.click(menuItem);
    expect(menuItem).not.toBeInTheDocument();
    expect(onChange).toHaveBeenCalledTimes(1);
  });

  test('Object Array, value and options are given only', () => {
    const { container } = render((
      <DropDown value="" options={[{ text: 'hello' }, { text: 'world' }]} />
    ));

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = document.querySelector('.menu-item-label');
    expect(menuItem).toBeInTheDocument();
    expect(menuItem).toHaveTextContent('[unknown]');
    fireEvent.click(menuItem);
  });

  test('Object Array', () => {
    const options = [{ id: '1', text: 'hello' }, { id: '2', text: 'world' }];
    const requestOptionValue = (option: {id: string, text: string}) => option.id;
    const renderValue = (value: string) => (
      <div className="dropdown-value">{options.find((i) => i.id === value).text}</div>
    );
    const renderMenuItemLabel = (option: {id: string, text: string}) => (<div>{option.text}</div>);

    const { container } = render((
      <DropDown
        value="1"
        options={options}
        requestOptionValue={requestOptionValue}
        renderValue={renderValue}
        renderMenuItemLabel={renderMenuItemLabel}
      />
    ));

    const value = screen.getByText('hello');
    expect(value).toBeInTheDocument();
    expect(value).toHaveClass('dropdown-value');

    const valueBox = container.querySelector('.dropdown-value-box');
    expect(valueBox).toBeInTheDocument();
    fireEvent.click(valueBox);

    const menuItem = screen.getByText('world');
    expect(menuItem).toBeInTheDocument();
  });
});
