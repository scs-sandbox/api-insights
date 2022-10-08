function scrollTo(
  querySelector: string,
  adjustOffset = 0,
) {
  const targetElement = document.querySelector(querySelector);
  if (!targetElement) return;

  const viewPortTop = document.body.getBoundingClientRect().top;
  const targetElementTop = targetElement.getBoundingClientRect().top;
  const offset = targetElementTop - viewPortTop - adjustOffset;
  window.scrollTo(0, offset);
}

export default function elementScrollTo(
  querySelector: string,
  adjustOffset = 0,
  delay = 100,
) {
  if (delay === 0) {
    scrollTo(querySelector, adjustOffset);
    return;
  }

  setTimeout(() => {
    scrollTo(querySelector, adjustOffset);
  }, delay);
}
