import elementScrollTo from './elementScrollTo';

jest.useFakeTimers();
jest.spyOn(global, 'setTimeout');
global.scrollTo = jest.fn();

describe('elementScrollTo', () => {
  test('An element that can be found, not delay', () => {
    elementScrollTo('.something', 0, 0);
  });

  test('An element that can be found, not delay', () => {
    const div = document.createElement('div');
    div.className = 'something';
    document.body.appendChild(div);
    elementScrollTo('.something', 0, 0);
  });

  test('An element that can be found, and delay', () => {
    const div = document.createElement('div');
    div.className = 'something';
    document.body.appendChild(div);
    elementScrollTo('.something', 0, 10);
    expect(setTimeout).toHaveBeenCalledTimes(1);
  });
});
