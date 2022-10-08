import capitalize from './string';

describe('String', () => {
  test('capitalize - normal', () => {
    const r = capitalize('hello, world');
    expect(r).toEqual('Hello, world');
  });

  test('capitalize - empty string', () => {
    const r = capitalize('');
    expect(r).toEqual('');
  });
});
