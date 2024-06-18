import type { FC } from "react";

type Props = {
  title: string;
  createdAt: string;
  updatedAt: string;
}

export const PostHeading: FC<Props> = ({ title, createdAt, updatedAt }) => {
  return (
    <div className="flex flex-col gap-2">
      <h1 className="text-xl font-bold">{title}</h1>
      <div className="flex flex-col gap-0.5 text-sm text-gray-200">
        <p>作成日時: {createdAt}</p>
        <p>更新日時: {updatedAt}</p>
      </div>
    </div>
  );
};
