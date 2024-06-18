import dayjs from "dayjs";
// バックエンドから受け取った時間をフォーマットする関数
export function formatDateTime(dateTime: string): string {
    const date = new Date(dateTime);
    return dayjs(date).format("YYYY/MM/DD HH:mm");
}
