import { render, screen } from '@testing-library/react';
import FileInput from './FileInput';

describe('<FileInput />', () => {
  test('Elements', () => {
    const onChange = jest.fn();
    const { container } = render((
      <FileInput accept="*.json" value="a.json" onChange={onChange} />
    ));

    const value = screen.getByText('a.json');
    expect(value).toHaveClass('input-value');

    const action = container.querySelector('.input-action');
    expect(action).toBeInTheDocument();

    const input = container.querySelector('.sys-input');
    expect(input).toBeInTheDocument();
    expect(input).toHaveAttribute('type', 'file');
    expect(input).toHaveAttribute('accept', '*.json');
  });
});
