/*
example:
    const a = [{name: "Team-1", score: 100}, {name: "Team-2", score: 99}]
    const b = buildSortedListByStringField(a, "score");

result:
    then b is [{name: "Team-2", score: 99}, {name: "Team-1", score: 100}]
 */

export default function buildSortedListByStringField(
  list: unknown[],
  stringField: string,
  desc?: boolean,
): unknown[] {
  if (!list) return [];
  if (!Array.isArray(list)) return [];

  return [...list].sort((a: unknown, b: unknown) => {
    const va = a[stringField];
    const vb = b[stringField];
    if (va === vb) return 0;

    const ascResult = va > vb ? 1 : -1;
    return desc ? 0 - ascResult : ascResult;
  });
}
