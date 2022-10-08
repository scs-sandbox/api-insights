import { render, screen } from '@testing-library/react';
import FieldItem from './FieldItem';

describe('<FieldItem />', () => {
  test('Input', () => {
    render((
      <FieldItem label="Name">
        <input />
      </FieldItem>
    ));

    expect(screen.getByText('Name')).toBeInTheDocument();
    expect(screen.getByRole('textbox')).toBeInTheDocument();
  });
});
