import type { FC } from "react";

type Props = {
  body: string;
  username: string;
}

export const PostContent: FC<Props> = ({ body, username }) => {
  return (
    <div className="flex flex-col gap-3 text-lg">
      <p className="whitespace-pre-line leading-8">{body}</p>
      <p className="text-right">{username}</p>
    </div>
  );
};
