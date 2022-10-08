import classNames from './className';

describe('classNames', () => {
  test('a, b', () => {
    const r = classNames('a', 'b');
    expect(r).toEqual('a b');
  });

  test('a, null, b', () => {
    const r = classNames('a', null, 'b');
    expect(r).toEqual('a b');
  });

  test('a, space, b', () => {
    const r = classNames('a', ' ', 'b');
    expect(r).toEqual('a b');
  });
});
