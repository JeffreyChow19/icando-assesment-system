import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

dayjs.extend(utc);
dayjs.extend(timezone);

export function formatDate(date: Date): string {
  return dayjs(date).tz('Asia/Jakarta').locale('id').format('dddd, D MMMM YYYY');
}

export function formatHour(date: Date): string {
  return dayjs(date).tz('Asia/Jakarta').locale('id').format('HH:mm') + ' WIB';
}