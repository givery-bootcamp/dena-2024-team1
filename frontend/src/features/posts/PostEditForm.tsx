import { Button } from "~/shared/components/Button";

export const PostEditForm = () => {
  return (
    <form className="flex flex-col gap-8">
      <label className="flex flex-col gap-1">
        <span className="font-bold">タイトル</span>
        <input className="h-10 rounded border border-gray-200 p-2" type="text" name="title" />
      </label>
      <label className="flex flex-col gap-1">
        <span className="font-bold">本文</span>
        <textarea className="rounded border border-gray-200 p-2" name="body" />
      </label>

      <Button className="h-12" type="submit">更新する</Button>
    </form>
  );
};