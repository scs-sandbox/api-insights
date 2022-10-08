import { render, screen } from '@testing-library/react';
import PopLayer from './PopLayer';

describe('document.body has not had #poplayer yet', () => {
  test('Basic', () => {
    render((
      <PopLayer>
        <div>hello</div>
      </PopLayer>
    ));

    const div = screen.getByText('hello');
    expect(div).toBeInTheDocument();
  });

  test('document.body has had #poplayer already', () => {
    const domPopLayer = document.createElement('div');
    domPopLayer.id = 'pop-layer';
    document.body.appendChild(domPopLayer);

    render((
      <PopLayer>
        <div>hello</div>
      </PopLayer>
    ));

    const div = screen.getByText('hello');
    expect(div).toBeInTheDocument();
  });
});
