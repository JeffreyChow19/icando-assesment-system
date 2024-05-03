import { Dayjs } from "dayjs";

export function timeDiff(start: Dayjs, end: Dayjs): string {
  const secondDiff = end.diff(start, "second");

  const minute = Math.floor(secondDiff / 60);
  const second = secondDiff % 60;

  if (minute === 0) {
    return `${second} seconds`;
  }

  if (second === 0) {
    return `${minute} minutes`;
  }

  return `${minute} minute and ${second} seconds`;
}
