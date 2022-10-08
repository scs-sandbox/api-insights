/*
example:
    const className = classNames("button", props.className, props.active ? "active" : "");

result:
    className will be "button start active" if the props.className
    is "start" and the props.active is true
*/
export default function classNames(...args: string[]): string {
  return args
    .filter((i) => (i || '').trim())
    .map((i) => i.trim())
    .join(' ');
}
